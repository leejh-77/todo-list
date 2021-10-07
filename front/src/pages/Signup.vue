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

import service from "../service/user";

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
      if (!this.validateParams()) {
        return
      }
      service.signup(this.email, this.password, this.username, res => {
        if (res.status === 201) {
          this.$router.push('/login')
        } else {
          alert("failed to signup")
        }
      })
    },
    validateParams : function() {
      if (this.email.length === 0) {
        alert("email required")
      } else if (this.password.length === 0) {
        alert("password required")
      } else if (this.username.length === 0) {
        alert("username required")
      } else if (this.confirmPassword.length === 0) {
        alert("confirm-password required")
      } else if (!this.validateEmail(this.email)) {
        alert("email format not right")
      } else if (this.password !== this.confirmPassword) {
        alert("passwords not matched")
      } else {
        return true
      }
      return false
    },
    validateEmail: function (email) {
      let re = /^(([^<>()[\]\\.,;:\s@"]+(\.[^<>()[\]\\.,;:\s@"]+)*)|(".+"))@((\[[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\.[0-9]{1,3}\])|(([a-zA-Z\-0-9]+\.)+[a-zA-Z]{2,}))$/;
      return re.test(email);
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