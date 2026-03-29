package cmd

import (
	"fmt"
	output "gregops/pkg"
	"io/fs"
	"os"
	"path/filepath"
	"sort"
	"strings"

	"github.com/spf13/cobra"
)

const filesizesCmdName = "filesizes"

type dirSize struct {
	path string
	size int64
}

var filesizesCmd = &cobra.Command{
	Use:   filesizesCmdName + " [path]",
	Short: "Display directory sizes",
	Long: fmt.Sprintf(`Display directory sizes, similar to 'du -h -d 1'.
Calculates the total size of each subdirectory in the specified path.

Examples:
  %s %s ./src
  %s %s ./src --limit 5
  %s %s ./src --limit 20 --sort name`, CliName, filesizesCmdName, CliName, filesizesCmdName, CliName, filesizesCmdName),
	Args: cobra.MaximumNArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		limit, _ := cmd.Flags().GetInt("limit")
		sortMethod, _ := cmd.Flags().GetString("sort")
		formatStr, _ := cmd.Flags().GetString("output")

		format, err := output.ParseFormat(formatStr)
		if err != nil {
			cmd.Printf("Invalid format: %v\n", err)
			return
		}

		formatter := output.NewWithWriter(format, cmd.OutOrStdout())

		// Determine path - default to current directory
		path := "."
		if len(args) > 0 {
			path = args[0]
		}

		// Execute the native implementation
		if err := runFilesizes(path, sortMethod, limit, formatter); err != nil {
			formatter.PrintError(err)
			return
		}
	},
}

func init() {
	rootCmd.AddCommand(filesizesCmd)

	filesizesCmd.Flags().IntP("limit", "l", 10, "Number of items to display")
	filesizesCmd.Flags().StringP("sort", "s", "human", "Sort method: human, name, size")
	filesizesCmd.Flags().StringP("output", "o", "text", "Output format (text, json)")
}

func runFilesizes(rootPath, sortMethod string, limit int, formatter *output.Formatter) error {
	// Get directory sizes
	dirSizes, err := calculateDirectorySizes(rootPath)
	if err != nil {
		return fmt.Errorf("failed to calculate directory sizes: %w", err)
	}

	if len(dirSizes) == 0 {
		if err := formatter.PrintString("No directories found"); err != nil {
			return err
		}
		return nil
	}

	// Sort the results
	sortDirectorySizes(dirSizes, sortMethod)

	// Limit the results (take the largest ones for human/size sort)
	if limit > 0 && limit < len(dirSizes) {
		if sortMethod == "human" || sortMethod == "size" {
			dirSizes = dirSizes[len(dirSizes)-limit:]
		} else {
			dirSizes = dirSizes[:limit]
		}
	}

	// Prepare data for output
	var tableData []map[string]interface{}
	for _, ds := range dirSizes {
		// Format path to match du output (with ./ prefix for subdirs)
		displayPath := ds.path
		if ds.path != rootPath {
			relPath, _ := filepath.Rel(rootPath, ds.path)
			if relPath == "." {
				displayPath = rootPath
			} else {
				displayPath = "./" + relPath
			}
		} else {
			// For the root directory itself, use the original path
			if rootPath == "." {
				displayPath = "."
			} else {
				displayPath = rootPath
			}
		}

		tableData = append(tableData, map[string]interface{}{
			"Size":  formatSize(ds.size),
			"Path":  displayPath,
			"Bytes": ds.size,
		})
	}

	// Print the results using formatter
	return formatter.PrintTable(tableData)
}

func calculateDirectorySizes(rootPath string) ([]dirSize, error) {
	// Clean the path
	rootPath = filepath.Clean(rootPath)

	// Check if root path exists
	info, err := os.Stat(rootPath)
	if err != nil {
		return nil, fmt.Errorf("cannot access path %s: %w", rootPath, err)
	}

	var dirSizes []dirSize
	dirSizeMap := make(map[string]int64)
	dirCountMap := make(map[string]int64) // Track directory count for overhead

	if !info.IsDir() {
		// If it's a file, just return its disk usage
		return []dirSize{{path: rootPath, size: getDiskUsage(info.Size())}}, nil
	}

	// Walk the directory tree
	err = filepath.WalkDir(rootPath, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			// Skip directories we can't read
			return nil
		}

		// Get file info for size
		info, err := d.Info()
		if err != nil {
			return nil // Skip files we can't stat
		}

		// Determine which immediate subdirectory this file belongs to
		relPath, err := filepath.Rel(rootPath, path)
		if err != nil {
			return nil
		}

		// Find the top-level directory
		parts := strings.Split(filepath.ToSlash(relPath), "/")
		var targetDir string
		if len(parts) == 1 {
			// File/dir is directly in root directory
			targetDir = rootPath
		} else {
			// File/dir is in a subdirectory
			targetDir = filepath.Join(rootPath, parts[0])
		}

		if d.IsDir() {
			// Add directory overhead (directories consume disk space too)
			dirCountMap[targetDir]++
		} else {
			// Add file size (rounded up to block boundary to match du behavior)
			dirSizeMap[targetDir] += getDiskUsage(info.Size())
		}

		return nil
	})

	if err != nil {
		return nil, fmt.Errorf("error walking directory tree: %w", err)
	}

	// Add directory overhead (each directory typically takes 4KB)
	for dir, count := range dirCountMap {
		dirSizeMap[dir] += count * 4096 // 4KB per directory
	}

	// Convert map to slice
	for path, size := range dirSizeMap {
		if size > 0 { // Only include directories with content
			dirSizes = append(dirSizes, dirSize{path: path, size: size})
		}
	}

	return dirSizes, nil
}

// getDiskUsage approximates disk usage by rounding up to 4KB blocks
// This mimics how du calculates disk usage vs file size
func getDiskUsage(fileSize int64) int64 {
	const blockSize = 4096 // 4KB blocks
	if fileSize == 0 {
		return blockSize // Empty files still consume a block
	}
	// Round up to next block boundary
	return ((fileSize-1)/blockSize + 1) * blockSize
}

func sortDirectorySizes(dirSizes []dirSize, sortMethod string) {
	switch sortMethod {
	case "name":
		sort.Slice(dirSizes, func(i, j int) bool {
			return dirSizes[i].path < dirSizes[j].path
		})
	case "size", "human":
		sort.Slice(dirSizes, func(i, j int) bool {
			return dirSizes[i].size < dirSizes[j].size
		})
	}
}

func formatSize(bytes int64) string {
	const (
		B  = 1
		KB = 1024 * B
		MB = 1024 * KB
		GB = 1024 * MB
		TB = 1024 * GB
	)

	if bytes >= TB {
		return fmt.Sprintf("%.1fT", float64(bytes)/TB)
	}
	if bytes >= GB {
		return fmt.Sprintf("%.1fG", float64(bytes)/GB)
	}
	if bytes >= MB {
		return fmt.Sprintf("%.1fM", float64(bytes)/MB)
	}
	if bytes >= KB {
		return fmt.Sprintf("%.1fK", float64(bytes)/KB)
	}
	return fmt.Sprintf("%dB", bytes)
}
