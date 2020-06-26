<template>
<v-app class="UNKO">
  <v-app-bar app>
    <v-toolbar-title class="site-title"> <a :href="this.$site.base">{{ $siteTitle }}</a> </v-toolbar-title>
  </v-app-bar>

  <v-content>
  <v-container>
  <v-form
    ref="form"
  >
    <v-text-field
      v-model="address"
      :counter="50"
      :rules="[rule_address]"
      label="Mail Address"
    >
    </v-text-field>
    <v-text-field
      v-model="password"
      :counter="20"
      :rules="[rule_pass]"
      label="Password 8-20"
    >
    </v-text-field>
    <v-text-field
      v-model="repass"
      :counter="20"
      :rules="[rule_pass]"
      label="Password 8-20"
    >
    </v-text-field>
  </v-form>
  <v-btn class="white--text" color="indigo" @click="submit" dark>SUBMIT</v-btn>
  </v-container>
  </v-content>

</v-app>
</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
export default {
  name: 'LoginPage',
  vuetify: new Vuetify(),

  data () {
    return {
      address: '',
      password: '',
      repass: '',

      rule_address: value => {
        const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,50}))$/;
        return re.test(value) || "error address"
      },
      rule_pass: value => {
        const re = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{8,20}$/
        return re.test(value) || "error password"
      },
    }
  },
  methods: {
    submit () {
      if (this.$refs.form.validate()){
        console.log("SUBMIT")
      }
    }
  },
}
</script>

<style lang="stylus" scoped>
</style>