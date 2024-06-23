<template>
  <div id="app">
    <Header v-if="isLoggedIn" />
    <div class="content-wrapper">
      <Sidebar v-if="isLoggedIn" />
      <main class="main-content">
        <router-view></router-view>
      </main>
    </div>
  </div>
</template>

<script>
import Sidebar from './components/Sidebar.vue'
import Header from './components/Header.vue'

export default {
  name: 'App',
  components: {
    Sidebar,
    Header
  },
  computed: {
    isLoggedIn() {
      return this.$store.state.auth.isLoggedIn
    }
  }
}
</script>

<style>
#app {
  display: flex;
  flex-direction: column;
  min-height: 100vh;
}

.content-wrapper {
  display: flex;
  flex-grow: 1;
  margin-top: 60px; /* ヘッダーの高さ分のマージンを追加 */
}

.main-content {
  flex-grow: 1;
  padding: 20px;
  margin-left: 250px; /* サイドバーの幅分のマージンを追加 */
  overflow-y: auto;
  height: calc(100vh - 60px); /* ヘッダーの高さを引いた高さに設定 */
}
</style>