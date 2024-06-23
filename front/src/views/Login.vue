<template>
  <div class="login">
    <h2>ログイン</h2>
    <form @submit.prevent="login">
      <div>
        <label for="email">メールアドレス:</label>
        <input v-model="email" type="email" id="email" required>
      </div>
      <div>
        <label for="password">パスワード:</label>
        <input v-model="password" type="password" id="password" required>
      </div>
      <button type="submit">ログイン</button>
    </form>
    <p v-if="error" class="error">{{ error }}</p>
    <p>アカウントをお持ちでない方は <router-link to="/register">こちら</router-link> から新規登録</p>
  </div>
</template>

<script>
import axios from 'axios';

export default {
  name: 'Login',
  data() {
    return {
      email: '',
      password: '',
      error: null
    }
  },
  methods: {
    async login() {
      try {
        const response = await axios.get(`http://localhost:3000/api/v1/groups/${this.email}`, {
          // params: {
          //   password: this.password
          // }
        });
        
        const { token, userId } = response.data;
        
        // トークンとユーザーIDをストアに保存
        this.$store.commit('setToken', token);
        this.$store.commit('setUserId', userId);
        
        // ログイン状態を更新
        this.$store.commit('setLoggedIn', true);
        
        // ダッシュボードへリダイレクト
        this.$router.push('/dashboard');
      } catch (error) {
        console.error('ログインエラー:', error);
        this.error = error.response?.data?.message || 'ログインに失敗しました';
      }
    }
  }
}
</script>

<style scoped>
.error {
  color: red;
  margin-top: 10px;
}
</style>