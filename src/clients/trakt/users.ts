import TraktClientConfig from './client'

export class Users {
  async addUserListItems(userId: string, listId: string, payload: string): Promise<any> {
    return TraktClientConfig.client.post(`/users/${userId}/lists/${listId}/items`, payload)
    .then(function (response) {
      return response.data
    })
    .catch(function (error) {
      return [error.response.status, error.response.data]
    })
  }

  async removeUserListItems(userId: string, listId: string, payload: string): Promise<any> {
    return TraktClientConfig.client.post(`/users/${userId}/lists/${listId}/items/remove`, payload)
    .then(function (response) {
      return response.data
    })
    .catch(function (error) {
      return [error.response.status, error.response.data]
    })
  }

  async userList(userId: string, listId: string): Promise<any> {
    return TraktClientConfig.client.get(`/users/${userId}/lists/${listId}`)
    .then(function (response) {
      return response.data
    })
    .catch(function (error) {
      return [error.response.status, error.response.data]
    })
  }

  userListItems(userId: string, listId: string): Promise<any> {
    return TraktClientConfig.client.get(`/users/${userId}/lists/${listId}/items`)
    .then(function (response) {
      return response.data
    })
    .catch(function (error) {
      return [error.response.status, error.response.data]
    })
  }

  userLists(clientId: string, userId: string): Promise<any> {
    // self.headers['trakt-api-version'] = '2'
    // self.headers['trakt-api-key'] = client_id
    return TraktClientConfig.client.get(`/users/${userId}/lists`)
    .then(function (response) {
      return response.data
    })
    .catch(function (error) {
      console.log(error)
    })
  }
}
