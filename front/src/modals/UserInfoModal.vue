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
      imageChanged: false
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
          vue.imageChanged = true
        }
        reader.readAsDataURL(file);
      }
    },
    actionSave() {
      var image
      if (this.imageChanged) {
        image = this.imageFromSource()
      }
      userService.update(image, this.username)
          .then(res => {
            this.$store.commit('setUser', res.data)
            this.actionClose()
          })
    },
    imageFromSource() {
      let idx = this.imageSrc.indexOf(',')
      let base64 = this.imageSrc.substring(idx + 1, this.imageSrc.length);

      idx = this.imageSrc.indexOf(';')
      let type = this.imageSrc.substring(0, idx)
      idx = type.indexOf(':')
      type = type.substring(idx + 1, type.length)
      return {
        data: base64,
        type: type
      }
    }
  },
  mounted() {
    this.username = this.user.username
    let image = this.user.image
    if (image == null) {
      image = require('../assets/user_icon.png')
    } else {
      image = "data:" + image.type + ";base64," + image.data
    }
    this.imageSrc = image
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
  border-radius: 50%;
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