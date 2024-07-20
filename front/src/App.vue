<template>
  <div id="app">
    <Header v-if="isLoggedIn" @toggle-sidebar="toggleSidebar" />
    <div class="content-wrapper">
      <Sidebar v-if="isLoggedIn" :isOpen="sidebarOpen" />
      <main class="main-content" :class="{ 'sidebar-open': sidebarOpen }">
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
  data() {
    return {
      sidebarOpen: false
    }
  },
  computed: {
    isLoggedIn() {
      return this.$store.state.isLoggedIn
    }
  },
  methods: {
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    },
    checkAuth() {
      const token = localStorage.getItem('token')
      const userId = localStorage.getItem('userId')
      if (token && userId) {
        this.$store.commit('setToken', token)
        this.$store.commit('setUserId', userId)
        this.$store.commit('setLoggedIn', true)
      }
    }
  },
  created() {
    this.checkAuth()
  },
  watch: {
    $route(to) {
      if (to.meta.requiresAuth && !this.isLoggedIn) {
        this.$router.push('/login')
      }
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
  margin-top: 60px;
}

.main-content {
  flex-grow: 1;
  padding: 20px;
  margin-left: 0;
  transition: margin-left 0.3s ease-in-out;
  overflow-y: auto;
  height: calc(100vh - 60px);
}

.main-content.sidebar-open {
  margin-left: 250px;
}

@media (max-width: 768px) {
  .main-content.sidebar-open {
    margin-left: 0;
  }
}
</style>