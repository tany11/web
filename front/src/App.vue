<template>
  <div id="app">
    <Sidebar v-if="isLoggedIn" />
    <router-view></router-view>
  </div>
</template>

<script>
import Sidebar from './components/Sidebar.vue'

export default {
  name: 'App',
  components: {
    Sidebar
  },
  computed: {
    isLoggedIn() {
      return this.$store.state.auth.isLoggedIn
    }
  }
}
</script>