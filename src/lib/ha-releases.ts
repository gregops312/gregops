import {compareVersions} from 'compare-versions'
import {config} from 'dotenv'
import {GitHubReleasesClient} from '../clients/github/releases'

interface GetLatest {
  current: string
  owner: string
  release: boolean
  repo: string
}

export namespace SmartHomeReleases {
  export async function envLoader(filePath: string): Promise<void> {
    config({path: filePath})
  }

  export async function getLatest(arg: GetLatest): Promise<void> {
    // if (arg.release) {
    // } else {
    // }
    const r = await GitHubReleasesClient.getReleases(arg.owner, arg.repo)

    r.sort((a, b) => compareVersions(a.tag_name, b.tag_name)).reverse()

    // Compare current to latest
    const res = compareVersions(arg.current, r[0].tag_name)

    switch (res) {
    case 1: {
      console.log('I have no idea how you have a later version than what is available')
      break
    }

    case 0: {
      console.log('up-to-date')
      break
    }

    case -1: {
      console.log(`\tCurrent: ${arg.current}`)
      console.log(`\tLatest: ${r[0].tag_name}`)
      console.log(`\tUrl: ${r[0].html_url}`)
    }
    }
  }
}
