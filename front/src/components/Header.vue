<template>
  <header class="app-header">
    <div class="hamburger" @click="$emit('toggle-sidebar')">☰</div>
    <h1>業務管理</h1>
    <button @click="handleLogout">ログアウト</button>
  </header>
</template>

<script>
import { mapState, mapActions } from 'vuex'

export default {
  name: 'Header',
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
</style>