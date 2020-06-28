<template>

  <v-dialog v-model="dialog" width="500">
    <template v-slot:activator="{on, attrs}">
      <v-btn class="mr-2 white--text" color="indigo" dark v-bind="attrs" v-on="on">Log in</v-btn>
    </template>
    <v-card>
      <v-card-title class="headline grey lighten-2" primary-title>Log in</v-card-title>
      <v-content>
        <v-card-text>
          choise
        </v-card-text>
        <v-divider></v-divider>
        <v-card-actions>
          <v-btn @click="goole_oauth">OAuth</v-btn>
        </v-card-actions>
      </v-content>
    </v-card>
  </v-dialog>

  <!--<v-form
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
  </v-form>-->


</template>

<script>
import Vuetify from '../../../node_modules/vuetify'
export default {
  name: 'LoginPage',
  vuetify: new Vuetify(),

  data () {
    return {
      dialog: false,
      result: null,
      /*address: '',
      password: '',
      repass: '',
      rule_address: value => {
        const re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,50}))$/;
        return re.test(value) || "error address"
      },
      rule_pass: value => {
        const re = /^(?=.*\d)(?=.*[a-z])(?=.*[A-Z])[0-9a-zA-Z]{8,20}$/
        return re.test(value) || "error password"
      },*/
    }
  },
  methods: {
    submit () {
      console.log("SUBMIT")
    },
    google_oauth () {
      OAuth.initialize('8q7_NoiupHfdTbb8eVDG75zDKzc')
      OAuth.popup('google').done( result => {
        // do some stuff with result
        //this.$emit('result', result)
        console.log(result)
      })
    },
  },
}
</script>

<style lang="stylus" scoped>
</style>