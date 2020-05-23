<template>
<div>
<v-text-field
  v-model="search"
  @keyup.enter="submit"
  label="Search"
  outlined
  dense
  ></v-text-field>
</div>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
export default {
  name: 'SearchBox',
  vuetify: new Vuetify(),
  data () {
    return {
      search: ''
    }
  },
  methods: {
    submit (event) {
      if (event.keyCode === 13) {

      }
    },
  },
}
</script>

<style lang="stylus" scoped>
</style>