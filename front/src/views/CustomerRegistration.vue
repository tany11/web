<template>
    <div class="customer-management">
        <h1>顧客管理</h1>
        <div class="search-form">
            <div class="input-group">
                <input v-model="phoneNumber" placeholder="電話番号" class="input" />
                <input v-model="lastFourDigits" placeholder="電話番号（下4桁）" class="input" />
            </div>
            <div class="input-group">
                <input v-model="castName" placeholder="キャスト名" class="input" />
            </div>
            <div class="input-group">
                <input v-model="startDate" type="date" placeholder="開始日" class="input" />
                <input v-model="endDate" type="date" placeholder="終了日" class="input" />
            </div>
            <button @click="search" class="search-button">検索</button>
        </div>
        <customer-list v-if="showList" :customers="searchResults" @select-customer="showDetails" />
        <customer-detail v-if="selectedCustomer" :customer="selectedCustomer" />
    </div>
</template>

<script>
import CustomerList from '@/components/CustomerList.vue'
import CustomerDetail from '@/components/CustomerDetail.vue'

export default {
    name: 'CustomerManagement',
    components: {
        CustomerList,
        CustomerDetail
    },
    data() {
        return {
            phoneNumber: '',
            lastFourDigits: '',
            castName: '',
            startDate: '',
            endDate: '',
            searchResults: [],
            selectedCustomer: null,
            showList: false
        }
    },
    methods: {
        search() {
            // ここで検索ロジックを実装
            // API呼び出しなどを行い、結果をsearchResultsに格納
            this.showList = true
            this.selectedCustomer = null
        },
        showDetails(customer) {
            this.selectedCustomer = customer
            this.showList = false
        }
    }
}
</script>

<style scoped>
.customer-management {
    padding: 40px;
    max-width: 800px;
    margin: 0 auto;
}



.input-group {
    display: flex;
    gap: 15px;
    margin-bottom: 15px;
}

.input {
    flex: 1;
    padding: 12px;
    border: 1px solid #ddd;
    border-radius: 5px;
    font-size: 1rem;
}

.search-button {
    width: 100%;
    padding: 12px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 5px;
    font-size: 1rem;
    cursor: pointer;
    transition: background-color 0.3s;
}

.search-button:hover {
    background-color: #45a049;
}
</style>