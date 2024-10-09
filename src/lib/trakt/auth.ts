import {exit} from 'node:process'
import {AuthDevice} from '../../clients/trakt/auth-device'
import {TraktConfig, TraktConfigI} from './config'

export namespace TraktAuth {

  export function isTokenValid(config: TraktConfigI): boolean {
    // WTF: TypeScript Date takes milliseconds, epoch is in seconds
    const dateNow = new Date()
    const epoch3months = 2_629_743 * 3

    if (config.accessToken === undefined) {
      console.log('erro')
      exit(1)
    }

    const epochTokenExpiration = config.accessToken?.createdAt + epoch3months
    const tokenExpirationDate = new Date(epochTokenExpiration * 1000)
    const dateDifference = tokenExpirationDate.getTime() - dateNow.getTime()

    return dateDifference > 0
  }

  export async function getToken(config: TraktConfigI): Promise<void> {
    const auth = new AuthDevice()
    const code = await auth.generateDeviceCode(config.clientId)
    config.deviceCode = code.device_code

    console.log(`Go to ${code.verification_url}`)
    console.log(`Input code: ${code.user_code}`)

    let i = 0
    while (i < code.expires_in) {
      // eslint-disable-next-line no-await-in-loop
      const response = await auth.getToken(config)

      if (response !== undefined) {
        console.log('Wombat')
        config.accessToken = {
          accessToken: response.access_token,
          createdAt: response.created_at,
          expiresIn: response.expires_in,
          refreshToken: response.refresh_token,
          scope: response.scope,
          tokenType: response.token_type,
        }
        TraktConfig.createConfig(config)
        break
      }

      // eslint-disable-next-line no-await-in-loop
      await sleep(code.interval)

      i += code.interval
    }
  }

  // export async function handleConfig(dir: string): Promise<void> {
  //   const config = loadConfig(dir)

  //   // If config was null prompt and store
  //   if (config === null) {
  //     throw new Error('wtf')
  //   }

  //   if (config.token === undefined) {
  //     // getToken()
  //     console.log('The config.token is undefined')
  //   }
  // }

  export function sleep(sec: number): Promise<void> {
    // eslint-disable-next-line no-promise-executor-return
    return new Promise(resolve => setTimeout(resolve, sec * 1000))
  }
}
