<template>
  <div class="sidebar" :class="{ 'sidebar-open': isOpen }">
    <nav>
      <ul>
        <li class="category">メイン</li>
        <li><router-link to="/dashboard">ダッシュボード</router-link></li>
        <li><router-link to="/order">受注</router-link></li>
        <li><router-link to="/timeboard">タイムボード</router-link></li>
        <li class="category">キャスト</li>
        <li><router-link to="/cast-registration">キャスト登録</router-link></li>
        <li><router-link to="/cast-edit">キャスト編集</router-link></li>
        <li class="category">スタッフ</li>
        <li><router-link to="/staff-registration">スタッフ登録</router-link></li>
        <li><router-link to="/staff-edit">スタッフ編集</router-link></li>
        <li class="category">顧客</li>
        <li><router-link to="/customer-registration">顧客検索</router-link></li>
        <li class="category">店舗</li>
        <li><router-link to="/store-registration">店舗設定</router-link></li>
        <li><router-link to="/store-edit">店舗編集</router-link></li>
        <li><router-link to="/media-registration">媒体登録</router-link></li>
      </ul>
    </nav>
  </div>
</template>

<script>
export default {
  name: 'Sidebar',
  props: {
    isOpen: {
      type: Boolean,
      default: false
    }
  }
}
</script>

<style scoped>
.sidebar {
  width: 250px;
  background-color: #34495E;
  height: 100vh;
  position: fixed;
  left: -250px;
  top: 0;
  overflow-y: auto;
  transition: left 0.3s ease-in-out;
  padding-top: 60px;
}

.sidebar-open {
  left: 0;
}

nav ul {
  list-style-type: none;
  padding: 0;
}

nav ul li {
  padding: 10px;
}

nav ul li a {
  text-decoration: none;
  color: white;
  display: block;
  padding: 10px;
  transition: background-color 0.3s;
}

nav ul li a:hover {
  background-color: #41B883;
}

nav ul li a.router-link-active {
  background-color: #41B883;
}

.category {
  font-weight: bold;
  color: #41B883;
  padding: 15px 10px 5px;
  text-transform: uppercase;
  font-size: 0.9em;
}

nav ul li:not(.category) {
  padding-left: 20px;
}
</style>