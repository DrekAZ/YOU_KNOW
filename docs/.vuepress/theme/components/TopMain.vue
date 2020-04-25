<template>
  <main class="top-main">
    TEST
    <List />
  </main>
</template>

<script>
import List from '@theme/components/customList.vue'

export default {
  name: 'TopMain',

  components: {
    List,
  }
}
</script>
<style lang="stylus">
  .top-main
    margin 0 auto
    min-width 80vw
</style>