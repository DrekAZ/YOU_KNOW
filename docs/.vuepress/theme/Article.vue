<template>
<v-app>
  <v-app-bar app>
    <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
    <v-btn v-if="store.states.user" class="btn" dark color="indigo" @click="login">Log in</v-btn>
    <v-btn v-else-if="con_view" class="btn" dark color="indigo" @click="edit">編集<v-icon dark>mdi-pencil</v-icon></v-btn>
    <v-btn v-else class="btn" dark color="indigo" @click="save">保存<v-icon dark>mdi-folder-upload</v-icon></v-btn>
  </v-app-bar>

  <v-content>
    <v-container v-if="con_view">
      <Content />
    </v-container>
    <router-view ref="editor" /> <!-- v-else wo tuketara props send dekinai -->
  </v-content>
  <Footer />
</v-app>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
import Footer from '../components/Footer.vue'
import router from '../router'
import store from '../components/store'
//import Editor from './Editor.vue'
export default {
  name: 'Article',
  vuetify: new Vuetify(),
  router,
  components: {
    Footer,
  },
  data () {
    return {
      file_name: '',
      con_view: true,
    }
  },
  methods: {
    edit () {
      this.$router.push({
        name: 'edit',
        params: {file: this.$page.relativePath}
      }).catch(() => {})
      this.con_view =  false
    },
    save () {
      this.$refs.editor.save()
    },
    reset () {
    },
  },
  mounted () {
    this.$nextTick(function() {

    })
  }
}
</script>
<style lang="stylus" scoped>
body // later
  color #f5f5f5

.btn
  margin-left auto
  margin-right 50px
</style>