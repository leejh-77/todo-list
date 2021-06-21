<template>
  <b-form>
    <b-form-input id="input-field" placeholder="email" type="email" v-model="email"></b-form-input>
    <b-form-input id="input-field" placeholder="username" v-model="username"></b-form-input>
    <b-form-input id="input-field" placeholder="password" type="password" v-model="password"></b-form-input>
    <b-form-input id="input-field" placeholder="confirm password" type="password" v-model="confirmPassword"></b-form-input>
    <b-button id="button" v-on:click="actionSignup">SignUp</b-button>
  </b-form>
</template>

<script>

import service from "../network/service";

export default {
  name: "Signup",
  data: function() {
    return {
      email: '',
      password: '',
      confirmPassword: '',
      username: ''
    }
  },
  methods: {
    actionSignup: function() {
      service.signup(this.email, this.password, this.username, (res) => {
        console.log(res)
      })
    }
  }
}
</script>

<style scoped>

#input-field {
  margin: auto auto 10px;
  width: 20%;
}

#button {
  width: 20%;
}

</style>