<template>
<header>
  <v-app-bar app>
  <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
    <h1><Content slot-key="catchCopy" /></h1>
    <form class="search-place">
      <!-- <input class="search" placeholder="検索ワード..."> -->
      <SearchBox />
    </form>
  </v-app-bar app>
</header>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
import SearchBox from './SearchBox.vue'
export default {
  name: 'Header',
  vuetify: new Vuetify(),
  components: {
    SearchBox,
  }
}
</script>

<style lang="stylus" scoped>
header
  position flex
  justify-content center
  aligh-items center
  min-height 20rem

  >>> site-title
    z-index 1

  >>> h1
    position absolute
    top 50px
    left 50%
    transform translateX(-50%)
    text-align center

  >>> .search-place
    position absolute
    top 16rem 
    left 50%
    transform translateX(-50%)
    //margin 0 auto
</style>