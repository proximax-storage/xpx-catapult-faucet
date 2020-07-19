import Vue from 'vue'
import Vuetify from 'vuetify/lib'

Vue.use(Vuetify)

export default new Vuetify({
  theme: {
    options: {
      customProperties: true
    },
    themes: {
      light: {
        primary: '#df4c48',
        error: '#ee6723',
        success: '#1eb3aa',
        warning: '#f4a400'
      }
    }
  }
})
