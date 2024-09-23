<template>
  <v-container class="dashboard">
    <v-row>
      <v-col>
        <h1 class="text-h3 mb-4">ダッシュボード</h1>
        <p class="text-h5 mb-6">ようこそ、{{ username }}さん！</p>
      </v-col>
    </v-row>
    <v-row>
      <v-col cols="12" md="3">
        <v-card>
          <v-card-title>現在の売上</v-card-title>
          <v-card-text>
            <v-list>
              <v-list-item v-for="(activity, index) in recentActivities" :key="index">
                <v-list-item-content>
                  <v-list-item-title>{{ activity }}</v-list-item-title>
                </v-list-item-content>
              </v-list-item>
            </v-list>
          </v-card-text>
        </v-card>
      </v-col>
      <!-- 他のダッシュボード要素をここに追加 -->
    </v-row>
  </v-container>
</template>

<script>
export default {
  name: 'Dashboard',
  data() {
    return {
      username: 'ユーザー', // 実際にはストアやAPIから取得したユーザー名を使用
      recentActivities: ['活動1', '活動2', '活動3']
    }
  },
  mounted() {
    this.fetchDashboardData()
  },
  methods: {
    fetchDashboardData() {
      // APIからダッシュボードのデータを取得する処理を実装
    }
  }
}
</script>