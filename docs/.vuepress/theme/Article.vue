<template>
<v-app>
  <v-app-bar app>
    <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
  </v-app-bar>
  <v-content>
    <v-container>
      <div class="content">
        <Content />
      </div>
    </v-container>
  </v-content>
  <Footer />
</v-app>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
import Footer from '../components/Footer.vue'
export default {
  name: 'Article',
  vuetify: new Vuetify(),
  
  components: {
    Footer
  },
}
</script>
<style lang="stylus" scoped>
.content
  border solid 1px #000
a
  color initial !important
  text-decoration none
</style>