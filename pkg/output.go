package output

import (
	"encoding/json"
	"fmt"
	"io"
	"os"
	"sort"
	"strings"
)

// Format represents the output format type
type Format string

const (
	FormatText Format = "text"
	FormatJSON Format = "json"
)

// Formatter handles output formatting for different formats
type Formatter struct {
	format Format
	writer io.Writer
}

// New creates a new formatter with the specified format
func New(format Format) *Formatter {
	return &Formatter{
		format: format,
		writer: os.Stdout,
	}
}

// NewWithWriter creates a new formatter with a custom writer
func NewWithWriter(format Format, writer io.Writer) *Formatter {
	return &Formatter{
		format: format,
		writer: writer,
	}
}

// SetWriter changes the output writer
func (f *Formatter) SetWriter(w io.Writer) {
	f.writer = w
}

// GetFormat returns the current format
func (f *Formatter) GetFormat() Format {
	return f.format
}

// Print outputs data in the configured format
func (f *Formatter) Print(data interface{}) error {
	switch f.format {
	case FormatJSON:
		return f.printJSON(data)
	case FormatText:
		return f.printText(data)
	default:
		return fmt.Errorf("unsupported format: %s", f.format)
	}
}

// PrintString outputs a simple string
func (f *Formatter) PrintString(s string) error {
	if f.format == FormatJSON {
		return f.printJSON(map[string]string{"value": s})
	}
	_, err := fmt.Fprintln(f.writer, s)
	return err
}

// PrintKeyValue outputs key-value pairs
func (f *Formatter) PrintKeyValue(pairs map[string]interface{}) error {
	if f.format == FormatJSON {
		return f.printJSON(pairs)
	}

	// Find the longest key for alignment
	maxKeyLen := 0
	keys := make([]string, 0, len(pairs))
	for key := range pairs {
		keys = append(keys, key)
		if len(key) > maxKeyLen {
			maxKeyLen = len(key)
		}
	}
	sort.Strings(keys)

	for _, key := range keys {
		value := pairs[key]
		padding := strings.Repeat(" ", maxKeyLen-len(key))
		_, err := fmt.Fprintf(f.writer, "%s:%s %v\n", key, padding, value)
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintTable outputs tabular data with headers and rows
func (f *Formatter) PrintTable(data []map[string]interface{}) error {
	if len(data) == 0 {
		return nil
	}

	if f.format == FormatJSON {
		return f.printJSON(data)
	}

	// Collect the union of keys across all rows, then sort for stable output
	keySet := make(map[string]struct{})
	for _, row := range data {
		for key := range row {
			keySet[key] = struct{}{}
		}
	}
	headers := make([]string, 0, len(keySet))
	for key := range keySet {
		headers = append(headers, key)
	}
	sort.Strings(headers)

	// Calculate column widths for alignment
	colWidths := make(map[string]int)
	for _, header := range headers {
		colWidths[header] = len(header)
	}

	// Check data for max widths
	for _, row := range data {
		for _, header := range headers {
			if val, ok := row[header]; ok {
				valStr := fmt.Sprintf("%v", val)
				if len(valStr) > colWidths[header] {
					colWidths[header] = len(valStr)
				}
			}
		}
	}

	// Print header
	var headerParts []string
	for _, header := range headers {
		headerParts = append(headerParts, fmt.Sprintf("%-*s", colWidths[header], header))
	}
	_, err := fmt.Fprintf(f.writer, "%s\n", strings.Join(headerParts, "  "))
	if err != nil {
		return err
	}

	// Print separator
	var separatorParts []string
	for _, header := range headers {
		separatorParts = append(separatorParts, strings.Repeat("-", colWidths[header]))
	}
	_, err = fmt.Fprintf(f.writer, "%s\n", strings.Join(separatorParts, "  "))
	if err != nil {
		return err
	}

	// Print rows
	for _, row := range data {
		var rowParts []string
		for _, header := range headers {
			val := ""
			if v, ok := row[header]; ok {
				val = fmt.Sprintf("%v", v)
			}
			rowParts = append(rowParts, fmt.Sprintf("%-*s", colWidths[header], val))
		}
		_, err := fmt.Fprintf(f.writer, "%s\n", strings.Join(rowParts, "  "))
		if err != nil {
			return err
		}
	}

	return nil
}

// PrintList outputs a list of items
func (f *Formatter) PrintList(items []interface{}) error {
	if f.format == FormatJSON {
		return f.printJSON(items)
	}

	for _, item := range items {
		_, err := fmt.Fprintf(f.writer, "%v\n", item)
		if err != nil {
			return err
		}
	}
	return nil
}

// PrintError outputs an error message
func (f *Formatter) PrintError(err error) error {
	if f.format == FormatJSON {
		return f.printJSON(map[string]string{"error": err.Error()})
	}
	_, writeErr := fmt.Fprintf(f.writer, "Error: %v\n", err)
	return writeErr
}

// Internal helper methods

func (f *Formatter) printJSON(data interface{}) error {
	encoder := json.NewEncoder(f.writer)
	encoder.SetIndent("", "  ")
	return encoder.Encode(data)
}

func (f *Formatter) printText(data interface{}) error {
	switch v := data.(type) {
	case string:
		return f.PrintString(v)
	case map[string]interface{}:
		return f.PrintKeyValue(v)
	case []map[string]interface{}:
		return f.PrintTable(v)
	case []interface{}:
		return f.PrintList(v)
	default:
		// Fallback to basic string representation
		_, err := fmt.Fprintf(f.writer, "%v\n", v)
		return err
	}
}

// Utility functions

// ParseFormat converts a string to Format type
func ParseFormat(s string) (Format, error) {
	switch strings.ToLower(s) {
	case "text", "txt":
		return FormatText, nil
	case "json":
		return FormatJSON, nil
	default:
		return FormatText, fmt.Errorf("unsupported format: %s (supported: text, json)", s)
	}
}

// IsValidFormat checks if the format string is valid
func IsValidFormat(s string) bool {
	_, err := ParseFormat(s)
	return err == nil
}

// GetSupportedFormats returns a list of supported formats
func GetSupportedFormats() []string {
	return []string{"text", "json"}
}

// Helper structs for common data structures

// FileSize represents file size information
type FileSize struct {
	Path  string `json:"path"`
	Size  string `json:"size"`
	Bytes int64  `json:"bytes,omitempty"`
}

// IPInfo represents IP address information
type IPInfo struct {
	IP       string `json:"ip"`
	City     string `json:"city,omitempty"`
	Region   string `json:"region,omitempty"`
	Country  string `json:"country,omitempty"`
	Location string `json:"location,omitempty"`
	Org      string `json:"org,omitempty"`
	Postal   string `json:"postal,omitempty"`
	Timezone string `json:"timezone,omitempty"`
}

// VersionInfo represents version information
type VersionInfo struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}
