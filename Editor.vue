<template>
<v-app>
  <div>
  <v-textarea
    v-model="text"
  >
  </v-textarea>
  </div>
  <div>
  <v-textarea
    v-model="md"
  >
  </v-textarea>
  </div>
</v-app>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
import markdown from '@vuepress/markdown'
export default {
  name: 'Editor',
  vuetify: new Vuetify(),
  data () {
    return {
      text: '',
      md: '# unko',
    }
  },
  computed: {
    compile: function() {
      return markdown(this.md)
    }
  }
}
</script>

<style lang="stylus" scoped>
div
  width 50vw
</style>