<template>
  <div>
    <Header class="header" />
    <div class="container">
      <Sideb class="sid-left" />
      <TopMain />
      <Sideb class="sid-right" />
    </div>
  </div>
</template>
<script>
import Header from '@theme/components/Header.vue'
import TopMain from '@theme/components/TopMain.vue'
import Sideb from '@theme/components/customSide.vue'
export default {
  name: 'SearchTop',
  components: {
    Header,
    TopMain,
    Sideb,
  },
}
</script>
<style lang="stylus" scoped>
.header
  height 10rem
</style>
<style lang="stylus" scoped>
.container
  display flex
.sid-left
  min-width 100px
  height 300px
  margin-left 1rem 
.sid-right
  min-width 100px
  margin-right 1rem

@media(max-width: $MQMobile)
  .sid-left 
    display none
  .sid-right
    display none
</style>