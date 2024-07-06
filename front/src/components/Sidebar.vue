<template>
  <div class="sidebar" :class="{ 'sidebar-open': isOpen }">
    <nav>
      <ul>
        <li><router-link to="/dashboard">ダッシュボード</router-link></li>
        <li><router-link to="/order">受注</router-link></li>
        <li><router-link to="/staff">スタッフ管理</router-link></li>
        <li><router-link to="/cast">キャスト管理</router-link></li>
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
</style>