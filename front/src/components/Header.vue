<template>
  <header class="app-header" :class="{ 'dark-mode': isDarkMode }">
    <div class="hamburger" @click="$emit('toggle-sidebar')">☰</div>
    <h1>業務管理</h1>
    <div class="header-buttons">
      <button @click="toggleDarkMode" class="mode-toggle">
        {{ isDarkMode ? '🌞' : '🌙' }}
      </button>
      <button @click="handleLogout" class="logout-button">ログアウト</button>
    </div>
  </header>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Header',
  data() {
    return {
      isDarkMode: false
    }
  },
  computed: {
    ...mapState(['user']),
    isLoggedIn() {
      return !!this.user
    }
  },
  methods: {
    ...mapActions(['logout']),
    async handleLogout() {
      const success = await this.logout()
      if (success) {
        this.$router.push('/login')
      } else {
        alert('ログアウトに失敗しました。もう一度お試しください。')
      }
    },
    toggleDarkMode() {
      this.isDarkMode = !this.isDarkMode
      document.body.classList.toggle('dark-mode', this.isDarkMode)
    }
  }
}
</script>

<style scoped>
.app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #2C3E50;
  color: white;
  height: 60px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}

.hamburger {
  font-size: 24px;
  cursor: pointer;
  margin-right: 15px;
}

.header-buttons {
  display: flex;
  align-items: center;
}

.mode-toggle {
  margin-right: 10px;
  background: none;
  border: none;
  font-size: 20px;
  cursor: pointer;
}

.logout-button {
  padding: 5px 10px;
  cursor: pointer;
  background-color: #34495E;
  color: white;
  border: none;
  border-radius: 4px;
}

.logout-button:hover {
  background-color: #41B883;
}

.dark-mode {
  background-color: #1a202c;
  color: #e2e8f0;
}

.dark-mode .logout-button {
  background-color: #2d3748;
}

.dark-mode .logout-button:hover {
  background-color: #4a5568;
}
</style>