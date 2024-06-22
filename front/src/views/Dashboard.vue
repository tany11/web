<template>
    <div class="dashboard">
      <h1>ダッシュボード</h1>
      <p>ようこそ、{{ username }}さん！</p>
      <div class="dashboard-content">
        <!-- ダッシュボードの内容をここに追加 -->
        <h2>最近の活動</h2>
        <ul>
          <li>活動1</li>
          <li>活動2</li>
          <li>活動3</li>
        </ul>
      </div>
    </div>
  </template>
  
  <script>
  export default {
    name: 'Dashboard',
    data() {
      return {
        username: 'ユーザー' // 実際にはストアやAPIから取得したユーザー名を使用
      }
    },
    mounted() {
      // コンポーネントがマウントされたときに必要なデータを取得
      this.fetchDashboardData()
    },
    methods: {
      fetchDashboardData() {
        // APIからダッシュボードのデータを取得する処理を実装
      }
    }
  }
  </script>