import {Args, Command, Flags, ux} from '@oclif/core'
import * as fs from 'node:fs'
import * as fsp from 'node:fs/promises'
import * as notifier from 'node-notifier'

export default class FileSizes extends Command {
  static description = 'File size commands'

  static examples = [
    '$ <%= config.bin %> local file-sizes',
  ]

  static args = {
    dir: Args.string(),
  }

  static flags = {
    filter: ux.ux.table.flags().filter,
    sort: ux.ux.table.flags().sort,
  }

  async run(): Promise<void> {
    const data = []
    const {args, flags} = await this.parse(FileSizes)
    const dir = args.dir ?? process.cwd()
    const files = fs.readdirSync(dir)

    const stats = await fs.statSync(`${process.cwd()}/package.json`)
    console.log(stats.size / 1024)

    for (const file of files) {
      // eslint-disable-next-line no-await-in-loop
      const something = await fsp.stat(`${dir}/${file}`)
      console.log(file)
      data.push({file: file, fmtSize: this.formatSize(something.size), size: something.size})
    }

    ux.ux.table(
      data,
      {
        formattedSize: {get: row => row.fmtSize, header: 'Size'},
        name: {get: row => row.file, header: 'File'},
        size: {get: row => row.size, header: 'True Size', extended: true},
      },
      {
        filter: flags.filter,
        sort: flags.sort,
      },
    )

    notifier.notify({
      title: 'gregops',
      message: 'file sizes complete',
      sound: true,
    })
  }

  formatSize(fileSize: number): string {
    const labels = ['', 'K', 'M', 'G', 'T', 'P']
    console.log(fileSize)
    let i = 0
    while (fileSize > 1000) {
      console.log(fileSize / 1024)
      fileSize /= (fileSize / 1024)
      console.log(fileSize)
      i++
    }

    return `${fileSize.toString()}${labels[i]}`
  }
}
