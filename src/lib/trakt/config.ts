import * as fs from 'node:fs'

export interface TraktConfigI {
  clientId: string
  clientSecret: string
  dir: string
  userId: string
  deviceCode?: string
  accessToken?: {
    accessToken: string
    createdAt: number
    expiresIn: number
    refreshToken: string
    scope: string
    tokenType: string
  }
}

export namespace TraktConfig {
  export const filename = 'trakt.conf'

  export function loadConfig(dir: string): TraktConfigI {
    return JSON.parse(fs.readFileSync(`${dir}/${filename}`, 'utf8'))
  }

  export function createConfig(config: TraktConfigI): void {
    // fs.writeFileSync(`${config.dir}/${filename}`, JSON.stringify({
    //   clientId: config.clientId,
    //   clientSecret: config.clientSecret,
    //   userId: config.userId,
    // }, null, 2))
    fs.writeFileSync(`${config.dir}/${filename}`, JSON.stringify(config, null, 2))
  }
}
