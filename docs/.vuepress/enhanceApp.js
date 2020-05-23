import vuetify from '../../node_modules/vuetify'
//import vuetify from './plugin/vuetify'
import '../../node_modules/vuetify/dist/vuetify.min.css'

export default ({
  Vue, // the version of Vue being used in the VuePress app
  options, // the options for the root Vue instance
  router, // the router instance for the app
  siteData, // site metadata
  isServer // is this enhancement applied in server-rendering or client
}) => {
  Vue.use(vuetify)
}