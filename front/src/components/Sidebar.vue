<template>
    <div class="sidebar">
      <nav>
        <ul>
          <li><router-link to="/dashboard">ダッシュボード</router-link></li>
          <li><router-link to="/profile">プロフィール</router-link></li>
          <li><a @click="logout">ログアウト</a></li>
        </ul>
      </nav>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Sidebar',
    methods: {
      logout() {
        this.$store.dispatch('logout')
        this.$router.push('/login')
      }
    }
  }
  </script>