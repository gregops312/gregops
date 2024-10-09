import {Command} from '@oclif/core'
import TraktClientConfig from '../../clients/trakt/client'

export class Run extends Command {
  async run(): Promise<void> {
    console.log()
  }

  storeConfig(config: any): void {
    TraktClientConfig.clientId = config.clientId
    TraktClientConfig.clientSecret = config.clientSecret
    TraktClientConfig.userId = config.userId
  }
}
