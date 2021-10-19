<template>
  <Modal @close="actionCloseModal">
    <p>Write name for folder</p>
    <div>
      <b-input v-model="inputString"/>
    </div>
    <template slot="footer">
      <b-button @click="actionAddFolder">Add</b-button>
    </template>
  </Modal>
</template>

<script>
import Modal from "./Modal";
import service from '../service/folder'
import {mapGetters} from "vuex";

export default {
  name: "AddFolderModal",
  components: {
    Modal
  },
  computed: {
    ...mapGetters([
        'workspace'
    ])
  },
  data() {
    return {
      inputString : ''
    }
  },
  methods: {
    actionAddFolder() {
      let wid = this.workspace.id
      service.addFolder(wid, this.inputString)
      .then(res => {
        this.$emit('created', res.data)
        this.actionCloseModal()
      })
      .catch(() => {
        this.actionCloseModal()
      })
    },
    actionCloseModal() {
      this.$emit('close')
    }
  }
}
</script>

<style scoped>

</style>