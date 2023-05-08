import {Command, Flags} from '@oclif/core'
import {SmartHomeReleases} from '../../lib/ha-releases'

export default class Releases extends Command {
  static description = 'TBD'

  static flags = {
    filePath: Flags.string({
      char: 'f',
      description: 'GitHub API token',
      required: true,
    }),
  }

  async run(): Promise<void> {
    const {flags} = await this.parse(Releases)
    await SmartHomeReleases.envLoader(flags.filePath)

    this.log('Grafana')
    if (process.env.GRAFANA_VERSION === undefined) {
      console.log('ENV GRAFANA_VERSION not found, can not process')
    } else {
      await SmartHomeReleases.getLatest({
        current: process.env.GRAFANA_VERSION,
        owner: 'grafana',
        release: true,
        repo: 'grafana',
      })
    }

    // this.log('Home Assistant')
    // if (process.env.HOME_ASSISTANT_VERSION === undefined) {
    //   console.log('ENV HOME_ASSISTANT_VERSION not found, can not process')
    // } else {
    //   await SmartHomeReleases.getLatest({
    //     current: process.env.HOME_ASSISTANT_VERSION,
    //     owner: 'home-assistant',
    //     release: true,
    //     repo: 'core',
    //   })
    // }

    this.log('Influx DB')
    if (process.env.INFLUXDB_VERSION === undefined) {
      console.log('ENV INFLUXDB_VERSION not found, can not process')
    } else {
      await SmartHomeReleases.getLatest({
        current: process.env.INFLUXDB_VERSION,
        owner: 'influxdata',
        release: true,
        repo: 'influxdb',
      })
    }

    // this.log('Nginx')
    // if (process.env.NGINX_VERSION === undefined) {
    //   console.log('ENV NGINX_VERSION not found, can not process')
    // } else {
    //   await SmartHomeReleases.getLatest({
    //     current: process.env.NGINX_VERSION,
    //     owner: 'nginx',
    //     release: false,
    //     repo: 'nginx',
    //   })
    // }

    // this.log('Node Red')
    // if (process.env.NODE_RED_VERSION === undefined) {
    //   console.log('ENV NODE_RED_VERSION not found, can not process')
    // } else {
    //   await SmartHomeReleases.getLatest({
    //     current: process.env.NODE_RED_VERSION,
    //     owner: 'node-red',
    //     release: false,
    //     repo: 'node-red',
    //   })
    // }

    this.log('Wyze Bridge')
    if (process.env.WYZE_VERSION === undefined) {
      console.log('ENV WYZE_VERSION not found, can not process')
    } else {
      await SmartHomeReleases.getLatest({
        current: process.env.WYZE_VERSION,
        owner: 'mrlt8',
        release: true,
        repo: 'docker-wyze-bridge',
      })
    }

    this.log('Zwave')
    if (process.env.ZWAVE_VERSION === undefined) {
      console.log('ENV ZWAVE_VERSION not found, can not process')
    } else {
      await SmartHomeReleases.getLatest({
        current: process.env.ZWAVE_VERSION,
        owner: 'zwave-js',
        release: true,
        repo: 'zwave-js-ui',
      })
    }
  }
}
