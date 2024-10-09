import {AxiosError} from 'axios'
import TraktClientConfig from './client'
import {TraktConfigI} from '../../lib/trakt/config'

export class AuthDevice {
  async generateDeviceCode(clientId: string): Promise<any> {
    // eslint-disable-next-line camelcase
    const data = {client_id: clientId}

    try {
      return (await TraktClientConfig.client.post('/oauth/device/code', data)).data
    } catch (error) {
      console.log(error)
    }
  }

  async getToken(config: TraktConfigI): Promise<any> {
    const data = {
      code: config.deviceCode,
      // eslint-disable-next-line camelcase
      client_id: config.clientId,
      // eslint-disable-next-line camelcase
      client_secret: config.clientSecret,
    }

    try {
      return await (await TraktClientConfig.client.post('/oauth/device/token', data)).data
    } catch (error: any) {
      const aError = error as AxiosError
      switch (aError.response?.status) {
      case 400: {
        console.log('Pending - waiting for the user to authorize your app')
        break
      }

      case 404: {
        console.log('Not Found - invalid device_code')
        break
      }

      case 409: {
        console.log('Already Used - user already approved this code')
        break
      }

      case 410: {
        console.log('Expired - the tokens have expired, restart the process')
        break
      }

      case 418: {
        console.log('Denied - user explicitly denied this code')
        break
      }

      case 429: {
        console.log('Slow Down - your app is polling too quickly')
        break
      }

      default: {
        console.log(error)
      }
      }
    }
  }
}

