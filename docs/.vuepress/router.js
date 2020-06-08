import Vue from '../../node_modules/vue'
import Router from '../../node_modules/vue-router'
import Editor from '../../Editor.vue'

Vue.use(Router)
export default new Router ({
  mode: 'history',
  routes: [{
    name: 'edit',
    path: '/edit',
    props: true,
    component: Editor
  }]
})