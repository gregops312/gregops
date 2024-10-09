import {Args, Command, Flags, ux} from '@oclif/core'
import * as fs from 'node:fs'
// import * as fsp from 'node:fs/promises'
// import * as notifier from 'node-notifier'

export default class FileSizes extends Command {
  static description = 'File size commands'

  static examples = [
    '$ <%= config.bin %> local file-sizes',
  ]

  static args = {
    dir: Args.string({
      default: process.cwd(),
    }),
  }

  static flags = {
    extended: ux.ux.table.flags().extended,
    filter: ux.ux.table.flags().filter,
    sort: ux.ux.table.flags().sort,
  }

  async run(): Promise<void> {
    const {args, flags} = await this.parse(FileSizes)
    const data = []
    const files = fs.readdirSync(args.dir)

    for (const file of files) {
      // console.log(file)
      // eslint-disable-next-line no-await-in-loop
      const something = await fs.statSync(`${args.dir}/${file}`)
      data.push({file: file, fmtSize: this.humanFileSize(something.size), size: something.size})
      // data.push({file: file, size: something.size})
    }

    ux.ux.table(
      data,
      {
        formattedSize: {get: row => row.fmtSize, header: 'Size'},
        name: {get: row => row.file, header: 'File'},
        size: {get: row => row.size, header: 'True Size', extended: true},
      },
      {
        extended: flags.extended,
        filter: flags.filter,
        sort: flags.sort,
      },
    )

    // notifier.notify({
    //   title: 'gregops',
    //   message: 'file sizes complete',
    //   sound: true,
    // })
  }

  // https://stackoverflow.com/questions/10420352/converting-file-size-in-bytes-to-human-readable-string
  humanFileSize(bytes: number, si = false, dp = 1): string {
    const thresh = si ? 1000 : 1024

    if (Math.abs(bytes) < thresh) {
      return bytes + ' B'
    }

    const units = si ?
      ['kB', 'MB', 'GB', 'TB', 'PB', 'EB', 'ZB', 'YB'] :
      ['K', 'M', 'G', 'T', 'P', 'E', 'Z', 'Y']
    let u = -1
    const r = 10 ** dp

    do {
      bytes /= thresh
      ++u
    } while (Math.round(Math.abs(bytes) * r) / r >= thresh && u < units.length - 1)

    return bytes.toFixed(dp) + units[u]
  }
}
