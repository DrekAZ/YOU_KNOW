<template>
  <div>
    <Header />
    <div class="container">
      <Sideb class="sid-left" />
      <TopMain class="main" />
      <Sideb class="sid-right" />
    </div>
    <div>
      {{ info }}
    </div>
  </div>
</template>

<script>
const axios = require('axios')
import Header from '@theme/components/Header.vue'
import TopMain from '@theme/components/TopMain.vue'
import Sideb from '@theme/components/customSide.vue'

export default {
  name: 'Top',
  components: {
    Header,
    TopMain,
    Sideb,
  },

  data() {
    return { info: null }
  },

  mounted() {
    axios
      .get('https://api.coindesk.com/v1/bpi/currentprice.json')
      .then(response => (this.info = response))
  }
}
</script>
<style lang="stylus" scoped>
.container
  display flex
.main

.sid-left
  //position absolute
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