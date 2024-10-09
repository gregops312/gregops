import axios from 'axios'

class TraktClientConfig {
  baseUrl = 'https://api.trakt.tv'
  client = axios.create({
    baseURL: `${this.baseUrl}`,
    headers: {
      'Content-type': 'application/json',
    },
  })

  clientId = null
  clientSecret = null
  userId = null
}

export default new TraktClientConfig()
