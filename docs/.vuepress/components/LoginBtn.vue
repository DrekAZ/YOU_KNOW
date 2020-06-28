<template>
<div class="button">
  <v-btn @click="goole_oauth">OAuth</v-btn>
</div>
</template>

<script>
export default {
  name: 'OAuthBtn',

  data () {
    return {
      data: '',
    }
  },
  methods: {
    google_oauth () {
      OAuth.initialize('8q7_NoiupHfdTbb8eVDG75zDKzc')
      OAuth.popup('google').done(function(result) {
        console.log(result)
        // do some stuff with result
        this.$emit('result', result)
      })
    },
  },
}
</script>

<style lang="stylus" scoped>
</style>