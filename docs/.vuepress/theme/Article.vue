<template>
<v-app>
  <v-app-bar app>
    <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
    <v-btn v-if="con_view" class="btn" dark color="indigo" @click="edit">編集<v-icon dark>mdi-pencil</v-icon></v-btn>
    <v-btn v-else class="btn" dark color="indigo" @click="save">保存<v-icon dark>mdi-folder-upload</v-icon></v-btn>
  </v-app-bar>
  <v-content>
    <v-container v-if="con_view">
      <Content />
    </v-container>
    <router-view v-else />
  </v-content>
  <Footer />
</v-app>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
import Footer from '../components/Footer.vue'
import router from '../router'
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
      this.$router.push({path: '/edit', params:{file_name: this.$page.relativePath}}).catch(() => {})
      this.con_view =  false
    },
    save () {

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