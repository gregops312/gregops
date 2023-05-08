/* eslint-disable camelcase */
import {Octokit} from 'octokit'
import {validate} from 'compare-versions'

interface Release {
  draft: boolean
  html_url: string
  prerelease: boolean
  tag_name: string
}

export namespace GitHubReleasesClient {
  export async function getReleases(owner: string, repo: string): Promise<Release[]> {
    const list = []
    const octokit = new Octokit()

    const iterator = octokit.paginate.iterator(`GET /repos/${owner}/${repo}/releases`, {
      owner: owner,
      repo: repo,
      per_page: 100,
      headers: {
        'X-GitHub-Api-Version': '2022-11-28',
      },
    })

    // return Promise.all(
    // )

    for await (const {data: releases} of iterator) {
      for (const release of releases) {
        const rel = release as Release
        if (!rel.draft && !rel.prerelease && validate(rel.tag_name)) {
          list.push(rel)
        }
      }
    }

    return list
  }
}
