import * as fs from 'node:fs'

import {Command, Flags, ux} from '@oclif/core'
import {TraktConfig} from '../lib/trakt/config'
import {TraktAuth} from '../lib/trakt/auth'
import { Sources } from '../lib/trakt/sources'

export class Media extends Command {
  static description = `
  - Uses XML files of owner media in sources such as Movies Anywhere, Google, Apple, etc
  - Compares those sources to what you have in Trakt lists
  - Will prompt any diffs of missing items from your sources that were previously in Trakt

  The goal is to make sure anything you bought was never removed because licesning changes...
  `
  static summary = 'Diff owned media between sources and Trakt lists'

  static examples = [
    '$ <%= config.bin %>',
  ]

  static flags = {
    moviesAnywhereFile: Flags.string({
      default: 'movies_anywhere.xml',
    }),
    moviesAnywhereList: Flags.string({
      char: 'm',
      default: 'movies anywhere',
    }),
    userId: Flags.string({
      char: 'u',
      default: 'gkman',
    }),
  }

  // Run should check if it can run and has valid info stored
  //  if fails it should tell you to run auth or something
  //  if Success it should run down the lists
  //    This needs to be passed files for each thing to compare against
  async run(): Promise<void> {
    const {flags} = await this.parse(Media)

    // Check for config
    let config
    if (fs.existsSync(`${this.config.dataDir}/${TraktConfig.filename}`) === false) {
      this.log('You must create a Trakt app to obtain a Client ID & Secret.')
      this.log('Navigate to: https://trakt.tv/oauth/applications')
      const clientId = await ux.prompt('Trakt client ID?')
      const clientSecret = await ux.prompt('Trakt client secret?')
      const userId = await ux.prompt('Trakt user ID?')
      TraktConfig.createConfig({dir: this.config.dataDir, clientId: clientId, clientSecret: clientSecret, userId: userId})
    }

    config = await TraktConfig.loadConfig(this.config.dataDir)

    // Auth to Trakt
    if (!TraktAuth.isTokenValid(config)) {
      console.log('Token is not valid, attempting to refresh token.')
      await TraktAuth.getToken(config)
    }

    // Load sources



    //    Flags use default filename in current working directory path
    //    No source files = exit

    // FOR LOOP
    //    Check source type for matching or defined Trakt list
    //        Flag option to change default list name

    //    Add source objects to list
    //    Compare list to source
    //    Flag diffs
    // FOR LOOP END
  }
}
