<template>
<v-app>
  <div class="editor">
  <div class="edit">
    <v-textarea
      :value="text"
      @input="compile"
      lazy>
    </v-textarea>
  </div>
  <div class="preview" v-html="md"></div>
  </div>
</v-app>
</template>

<script>
import Vuetify from './node_modules/vuetify'
import axios from './node_modules/axios'
import markdownit from './node_modules/markdown-it'

export default {
  name: 'Editor',
  vuetify: new Vuetify(),
  markdownit: new markdownit(),
  components: {
  },
  data () {
    return {
      text: '---\r\ntitle: 数学\r\ndate: 2020-05-05\r\nid: 1\r\ntags:\r\n  - 数学\r\n---\r\n## 数学です。\r\n#UNKO',
      md: '',
    }
  },

  props: {
    file_name: {
      type: String,
      default: '',
      required: true
    },
  },
  computed: {
  },

  methods :{
    compile (e) {
        this.md = markdownit().render(e)
    },
    Get_data () {
      const path = this.file_name.replace(/\/.*\/(.*\.md)/g, $1)
      this.axios.get('https://localhost:8081/get'+path).then((res) => {
        this.text = res.data.content
      })
      .catch((e) => {})
    },
  },
  mounted () {
    this.$nextTick(function() {
      //Get_data();
      this.md = markdownit().render(this.text)
    })
  },
}
</script>

<style lang="stylus" scoped>
.editor
  display flex

.edit
  width 50vw
  height 90vh
  display inline-block

.preview
  width 50vw
  height 90vh
  display inline-block
</style>