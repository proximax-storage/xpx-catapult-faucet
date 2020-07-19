// import env from '@/environment/environment'
import axios from 'axios'

class ApiService {
  constructor (environment, axios, headers) {
    // this.api = environment.API_URL
    this.axios = axios
    this.headers = headers
  }

  get (cb = null) {
    return this.axios.get(`${cb}`)
  }

  post () {
    console.log('post')
  }

  put () {
    console.log('put')
  }

  delete () {
    console.log('delete')
  }

  getConfig () {
    return axios.get('/api/faucet/config')
  }
}

const apiService = new ApiService('', axios, {
  key: 'Content-type',
  value: 'application/json'
})

export {
  apiService
}
