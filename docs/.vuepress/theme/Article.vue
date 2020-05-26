<template>
<v-app>
  <v-app-bar app>
    <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
    <v-btn class="ma-2" fab dark color="indigo" @click="edit"><v-icon dark>mdi-pencil</v-icon></v-btn>
  </v-app-bar>
  <v-content>
    <v-container>
      <div class="content">
        <Content v-if="con_view" />
        <router-view v-else />
      </div>
    </v-container>
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
      this.$router.push({name: 'editor', params:{file_name: this.$page.relativePath}}).catch(() => {})
      this.con_view =  false
    }
  },
}
</script>
<style lang="stylus" scoped>
body // later
  color #f5f5f5
</style>