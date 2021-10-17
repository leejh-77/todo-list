<template>
  <Modal @close="actionClose">
    <div class="main">
      <img class="image" :src="imageSrc"/>
      <label class="image-input-label" for="image-input">Edit photo</label>
      <input id="image-input" type="file" accept="image/png, image/jpeg" @change="onChangeImage">
      <b-input class="username" v-model="username"></b-input>
    </div>
    <template slot="footer">
      <b-button @click="actionSave">Save</b-button>
    </template>
  </Modal>
</template>

<script>
import Modal from "./Modal";
import {mapGetters} from "vuex";
import userService from '../service/user'

export default {
  name: "UserInfoModal",
  components: {Modal},
  data() {
    return {
      username: '',
      imageSrc: '',
    }
  },
  computed: {
    ...mapGetters([
        'user'
    ]),
  },
  methods: {
    actionClose() {
      this.$emit('close')
    },
    onChangeImage(evt) {
      let vue = this
      let file = evt.target.files[0]
      if (file) {
        let reader = new FileReader();
        reader.onload = function (e) {
          vue.imageSrc = e.target.result;
        }
        reader.readAsDataURL(file);
      }
    },
    actionSave() {
      let base64 = this.imageSrc.replace(/^data:(.*,)?/, '');
      userService.update(base64, this.username)
      .then(res => {
        console.log(res)
        this.actionClose()
      })
      .catch(e => console.log(e))
    }
  },
  mounted() {
    this.username = this.user.username
    let url = this.user.imageUrl
    if (url == null) {
      url = require('../assets/user_icon.png')
    }
    this.imageSrc = url
  }
}
</script>

<style scoped>
.main {
  width: 400px;
  align-content: center;
}

.image {
  width: 150px;
  margin-top: 50px;
}

.image-input-label {
  width: 100%;
  margin-top: 20px;
}

#image-input {
  display: none;
}

.username {
  margin-top: 20px;
}

</style>