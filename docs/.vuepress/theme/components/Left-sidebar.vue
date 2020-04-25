<template>
  <aside class="left">
    anthour: 
  </aside>
</template>

<script>
export default {
  name: 'LeftSidebar',
}
</script>

<style lang="stylus">
.left 
  background-color #f5f5f5
@media (max-width: $MQMobile)
  .left 
    display none 
</style>