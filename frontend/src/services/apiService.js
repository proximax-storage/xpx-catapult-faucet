// import env from '@/environment/environment'
import axios from 'axios'

class ApiService {
  constructor (environment, axios, headers) {
    // this.api = environment.API_URL
    this.axios = axios
    this.headers = headers
  }

  /**
   *
   *
   * @returns
   * @memberof ApiService
   */
  getConfig () {
    return axios.get('/api/faucet/config')
  }

  /**
   *
   * @param {*} address
   * @param {*} mosaic
   */
  getMosaic (address, mosaic) {
    return axios.get(`/api/faucet/GetXpx/${address}/${mosaic}`)
  }
}

const apiService = new ApiService('', axios, {
  key: 'Content-type',
  value: 'application/json'
})

export {
  apiService
}
