<template>
  <ul class="list-group">
    <li
      v-for="(item, index) in items"
      :key="index"
    >
      {{ item.title }} {{ item.author }}
    </li>
  </ul>
</template>
<script>
export default {
  name: 'CustomList',

  data() {
    return {
      items: [
        { title: 'gomi', author: 'takashi'},
        { title: 'kuso', author: 'nanako'},
      ]
    };
  }
}
</script>