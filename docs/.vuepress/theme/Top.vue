<template>
<v-app>
  <Header />
  <v-container>
    <Main />

    <v-divider></v-divider>

    <v-card>
      <v-row class="ma-5" align="center" justify="center" dense>
        <v-col
          v-for="(card, index) in officials"
          :key="index"
          :cols=2.5
        >
          <v-card :href="card.path" class="ma-1">
            <v-icon large>{{ icons[index].icon }}</v-icon>
            <v-card-title v-text="card.title" class="justify-center"></v-card-title>
          </v-card>
        </v-col>
      </v-row>
      <v-row align="center" justify="end" class="mr-5">
        <a href="search">more...</a>
      </v-row>
    </v-card>

  </v-container>
  <Footer />
</v-app>
</template>

<script>
//import Vuetify from '../plugin/vuetify'
import Vuetify from '../../../node_modules/vuetify'
import Header from '../components/Header.vue'
import Main from '../components/Main.vue' 
import Footer from '../components/Footer.vue'
export default {
  name: 'Top',
  vuetify: new Vuetify(),
  
  components: {
    Header,
    Main,
    Footer,
  },
  data: function() {
    return {
      icons: [
        {icon: 'mdi-numeric'},
        {icon: 'mdi-laptop'},
        {icon: 'mdi-flask'},
        {icon: 'mdi-robot-industrial'},
        {icon: 'mdi-link-variant-plus'},
      ],
      search: {src: 'search'},
    }
  },
  computed: {
    officials: function() {
      return this.$site.pages.filter( page => page.id === 'official').sort((a, b) => a.frontmatter.id - b.frontmatter.id ).slice(0, 5)
    },
  },
  methods: {
    selectIcon: function(n) {
      return this.icons.filter( (icon, index) => index === n)
    }
  }
}
</script>
