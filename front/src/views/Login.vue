<template>
    <div class="login">
      <h2>ログイン</h2>
      <form @submit.prevent="login">
        <div>
          <label for="username">ユーザーID:</label>
          <input v-model="username" type="text" id="username" required>
        </div>
        <div>
          <label for="password">パスワード:</label>
          <input v-model="password" type="password" id="password" required>
        </div>
        <button type="submit">ログイン</button>
      </form>
      <p>アカウントをお持ちでない方は <router-link to="/register">こちら</router-link> から新規登録</p>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Login',
    data() {
      return {
        username: '',
        password: ''
      }
    },
    methods: {
      login() {
        // ログイン処理を実装
        this.$store.dispatch('login', { username: this.username, password: this.password })
          .then(() => {
            this.$router.push('/dashboard')
          })
          .catch(error => {
            console.error('ログインエラー:', error)
            // エラーメッセージを表示するなどの処理
          })
      }
    }
  }
  </script>