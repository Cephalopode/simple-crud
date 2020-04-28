import '@mdi/font/css/materialdesignicons.css'
import Vue from 'vue'
import Vuetify from 'vuetify/lib'
import VNumberField from '../components/VNumberField.vue'

Vue.use(Vuetify)
Vue.component('v-number-field', VNumberField)

export default new Vuetify({
  icons: {
    iconfont: 'mdiSvg',
  },
})
