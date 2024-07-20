<template>
  <div class="login">
    <h2>ログイン</h2>
    <form @submit.prevent="handleLogin">
      <div>
        <label for="staffId">スタッフID:</label>
        <input v-model="staffId" type="text" id="staffId" required>
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
import { mapActions } from 'vuex'

export default {
  name: 'Login',
  data() {
    return {
      staffId: '',
      password: '',
      error: null
    }
  },
  methods: {
    ...mapActions(['login']),
    async handleLogin() {
      try {
        const success = await this.login({
          staff_id: this.staffId,
          password: this.password
        })
        if (success) {
          this.$router.push('/dashboard')
        } else {
          this.error = 'ログインに失敗しました。スタッフIDとパスワードを確認してください。'
        }
      } catch (error) {
        console.error('ログインエラー:', error)
        this.error = 'ログイン中にエラーが発生しました。後でもう一度お試しください。'
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