<template>
  <v-container class="fill-height" fluid>
    <v-row align="center" justify="center">
      <v-col cols="12" sm="8" md="6">
        <v-card class="elevation-12">
          <v-toolbar color="primary" dark flat>
            <v-toolbar-title>ログイン</v-toolbar-title>
          </v-toolbar>
          <v-card-text>
            <v-form @submit.prevent="handleLogin">
              <v-text-field v-model="staffId" label="スタッフID" name="staffId" prepend-icon="mdi-account" type="text"
                required></v-text-field>
              <v-text-field v-model="password" label="パスワード" name="password" prepend-icon="mdi-lock" type="password"
                required></v-text-field>
              <v-btn type="submit" color="primary" block class="mt-4">ログイン</v-btn>
            </v-form>
          </v-card-text>
          <v-card-actions>
            <v-spacer></v-spacer>
            <v-btn text to="/register">新規登録</v-btn>
          </v-card-actions>
        </v-card>
        <v-alert v-if="error" type="error" class="mt-4">
          {{ error }}
        </v-alert>
      </v-col>
    </v-row>
  </v-container>
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