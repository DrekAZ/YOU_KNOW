<template>
<v-app>
  <div class="editor">
  <div class="edit">
    <v-textarea
      solo
      full-width
      auto-grow
      class="mr-2"
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
    file: {
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
      this.text = e
    },
    save () {
      //console.log(this.text)
      this.Update_data()
    },
    Get_data () {
      const path = this.file.replace(/\/.*\/(.*\.md)/g, '$1')
      axios.get('https://localhost:8081/get', {
        params: {
          name: path
        }
      }).then((res) => {
        this.text = res.data.content
      })
      .catch((e) => { if(e.response) console.error('get error') })
    },
    Update_data () {
      const path = this.file.replace(/\/.*\/(.*\.md)/g, '$1')
      axios.post('https://localhost:8081/update', {
        name: path,
        markdown: this.text
      }).then((res) => {
        console.log(res)
      })
      .catch((e) => { if(e.response) console.error('update error') })
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
  overflow-wrap break-word
  display inline-block

.preview
  width 50vw
  height 90vh
  overflow-wrap break-word
  display inline-block
</style>