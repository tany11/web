<template>
  <div class="register">
    <h2>新規登録</h2>
    <form @submit.prevent="register">
      <div>
        <label for="username">ユーザーID:</label>
        <input v-model="username" type="text" id="username" required>
      </div>
      <div>
        <label for="email">メールアドレス:</label>
        <input v-model="email" type="email" id="email" required>
      </div>
      <div>
        <label for="password">パスワード:</label>
        <input v-model="password" type="password" id="password" required>
      </div>
      <div>
        <label for="confirmPassword">パスワード（確認）:</label>
        <input v-model="confirmPassword" type="password" id="confirmPassword" required>
      </div>
      <button type="submit">登録</button>
    </form>
    <p>すでにアカウントをお持ちの方は <router-link to="/login">こちら</router-link> からログイン</p>
  </div>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';

export default {
  name: 'Register',
  data() {
    return {
      username: '',
      email: '',
      password: '',
      confirmPassword: ''
    }
  },
  computed: {
    ...mapState(['apiBaseUrl'])
  },
  methods: {
    async register() {
      if (this.password !== this.confirmPassword) {
        alert('パスワードが一致しません');
        return;
      }

      try {
        const response = await axios.post(`${this.apiBaseUrl}/groups`, {
          username: this.username,
          email: this.email,
          password: this.password
        });

        console.log('ユーザー登録成功:', response.data);
        alert('ユーザー登録が完了しました。ログインページに移動します。');
        this.$router.push('/login');
      } catch (error) {
        console.error('ユーザー登録エラー:', error);
        alert('ユーザー登録に失敗しました。' + (error.response?.data?.message || ''));
      }
    }
  }
}
</script>