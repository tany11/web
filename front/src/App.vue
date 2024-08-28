<template>
  <v-app>
    <v-app-bar v-if="isLoggedIn" app>
      <v-app-bar-nav-icon @click="toggleSidebar"></v-app-bar-nav-icon>
      <v-toolbar-title>アプリケーション名</v-toolbar-title>
      <v-spacer></v-spacer>
      <v-btn icon @click="toggleDarkMode">
        <v-icon>{{ isDarkMode ? 'mdi-weather-sunny' : 'mdi-weather-night' }}</v-icon>
      </v-btn>
      <v-btn text @click="handleLogout">ログアウト</v-btn>
    </v-app-bar>

    <v-navigation-drawer v-if="isLoggedIn" v-model="sidebarOpen" app>
      <Sidebar />
    </v-navigation-drawer>

    <v-main>
      <v-container fluid>
        <router-view></router-view>
      </v-container>
    </v-main>
  </v-app>
</template>

<script>
import Sidebar from './components/Sidebar.vue'
import { mapState, mapActions } from 'vuex'

export default {
  name: 'App',
  components: {
    Sidebar
  },
  data() {
    return {
      sidebarOpen: true,
      isDarkMode: false,
    }
  },
  computed: {
    ...mapState(['user']),
    isLoggedIn() {
      return this.$store.state.isLoggedIn
    }
  },
  methods: {
    ...mapActions(['logout']),
    toggleSidebar() {
      this.sidebarOpen = !this.sidebarOpen
    },
    toggleDarkMode() {
      this.isDarkMode = !this.isDarkMode
      this.$vuetify.theme.global.name = this.isDarkMode ? 'dark' : 'light'
      localStorage.setItem('darkMode', this.isDarkMode)
    },
    async handleLogout() {
      const success = await this.logout()
      if (success) {
        this.$router.push('/login')
      } else {
        // エラーメッセージの表示にはVuetifyのスナックバーを使用することをお勧めします
        this.$store.commit('setSnackbar', { show: true, text: 'ログアウトに失敗しました。もう一度お試しください。' })
      }
    },
    checkAuth() {
      const token = localStorage.getItem('token')
      const userId = localStorage.getItem('userId')
      if (token && userId) {
        this.$store.commit('setToken', token)
        this.$store.commit('setUserId', userId)
        this.$store.commit('setLoggedIn', true)
      } else {
        this.$store.commit('setLoggedIn', false)
        this.$store.commit('setToken', null)
        this.$store.commit('setUserId', null)
      }
    }
  },
  created() {
    this.checkAuth()
    const darkMode = localStorage.getItem('darkMode')
    if (darkMode !== null) {
      this.isDarkMode = darkMode === 'true'
      this.$vuetify.theme.global.name = this.isDarkMode ? 'dark' : 'light'
    }
  },
  watch: {
    isDarkMode(newVal) {
      localStorage.setItem('darkMode', newVal)
    }
  }
}
</script>

<style></style>