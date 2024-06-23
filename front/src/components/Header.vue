<template>
    <header class="app-header">
      <h1>業務管理</h1>
      <button @click="logout" class="logout-button">ログアウト</button>
    </header>
  </template>
  
  <script>
  // ... 既存のスクリプト ...
  </script>
  
  <style scoped>
  .app-header {
  display: flex;
  justify-content: space-between;
  align-items: center;
  padding: 0 20px;
  background-color: #2C3E50;
  color: white;
  height: 60px;
  position: fixed;
  top: 0;
  left: 0;
  right: 0;
  z-index: 1000;
}
  
  .logout-button {
    padding: 5px 10px;
    cursor: pointer;
    background-color: #34495E;
    color: white;
    border: none;
    border-radius: 4px;
  }
  
  .logout-button:hover {
    background-color: #41B883;
  }
  </style>