<template>
  <v-list>
    <v-list-group v-for="item in menuItems" :key="item.title" :value="item.title" :prepend-icon="item.icon">
      <template v-slot:activator="{ props }">
        <v-list-item v-bind="props">
          <v-list-item-title :class="getTitleClass(item.title)">{{ item.title }}</v-list-item-title>
        </v-list-item>
      </template>

      <v-list-item v-for="subItem in item.items" :key="subItem.title" :to="subItem.to" :title="subItem.title" />
    </v-list-group>
  </v-list>
</template>

<script>
export default {
  name: 'Sidebar',
  data() {
    return {
      menuItems: [
        {
          title: 'メイン',
          icon: 'mdi-view-dashboard',
          items: [
            { title: 'ダッシュボード', to: '/dashboard' },
            { title: '受注', to: '/order' },
            { title: 'タイムボード', to: '/timeboard' },
            { title: 'タイムスケジュール', to: '/timeschedule' },
          ],
        },
        {
          title: 'キャスト',
          icon: 'mdi-account',
          items: [
            { title: 'キャスト登録', to: '/cast-registration' },
            { title: 'キャスト編集', to: '/cast-edit' },
          ],
        },
        {
          title: 'スタッフ',
          icon: 'mdi-account',
          items: [
            { title: 'スタッフ登録', to: '/staff-registration' },
            { title: 'スタッフ編集', to: '/staff-edit' },
          ],
        },
        {
          title: '顧客',
          icon: 'mdi-account',
          items: [
            { title: '顧客検索', to: '/customer-registration' },
          ],
        },
        {
          title: '店舗',
          icon: 'mdi-store',
          items: [
            { title: '店舗設定', to: '/store-registration' },
            { title: '店舗編集', to: '/store-edit' },
            { title: '媒体登録', to: '/media-registration' },
          ],
        },
      ],
    }
  },
  methods: {
    getTitleClass(title) {
      switch (title) {
        case 'キャスト':
          return 'pink--text'
        case 'スタッフ':
          return 'blue--text'
        case '顧客':
          return 'green--text'
        default:
          return ''
      }
    },
  },
}
</script>

<style scoped>
.pink--text {
  color: rgb(248, 160, 175) !important;
}

.blue--text {
  color: rgb(108, 174, 250) !important;
}

.green--text {
  color: rgb(122, 240, 122) !important;
}
</style>