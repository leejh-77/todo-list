<template>
  <b-form>
    <b-form-input id="emailAddress" placeholder="email" type="email" v-model="email"></b-form-input>
    <b-form-input id="password" placeholder="password" type="password" v-model="password"></b-form-input>
    <div>
      <b-button id="button-login" v-on:click="actionLogin">Login</b-button>
      <b-button id="button-signup" v-on:click="actionMoveToSignup">Signup</b-button>
    </div>
  </b-form>
</template>

<script>

import service from "../service/user";

export default {
  name: "Login",
  data: function () {
    return {
      email: '',
      password: ''
    }
  },
  methods: {
    actionLogin: function () {
      service.login(this.email, this.password)
          .then(res => {
            console.log(res)
            this.$router.push('/')
          })
          .catch(e => {
            console.log(e)
            alert("email or password is incorrect")
          })
    },
    actionMoveToSignup: function() {
      this.$router.push("/signup")
    }
  }
}
</script>

<style scoped>

#emailAddress {
  margin: auto auto 10px;
  width: 20%;
}

#password {
  margin: auto auto 10px;
  width: 20%;
}

#button-login {
  width: 10%;
  margin-right: 5px;
}

#button-signup {
  width: 10%;
}

</style>