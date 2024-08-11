<template>
    <div class="customer-management">
        <h1>顧客管理</h1>
        <div class="search-form">
            <div class="search-type-selector">
                <button @click="searchType = 'phone'" :class="{ active: searchType === 'phone' }">電話番号検索</button>
                <button @click="searchType = 'other'" :class="{ active: searchType === 'other' }">その他の検索</button>
            </div>
            <div v-if="searchType === 'phone'" class="input-group">
                <input v-model="phoneNumber" placeholder="電話番号" class="input" />
                <button @click="searchByPhone" class="search-button">検索</button>
            </div>
            <div v-else-if="searchType === 'other'">
                <div class="input-group">
                    <input v-model="lastFourDigits" placeholder="電話番号（下4桁）" class="input" />
                    <input v-model="castName" placeholder="キャスト名" class="input" />
                </div>
                <div class="input-group">
                    <input v-model="startDate" type="date" placeholder="開始日" class="input" />
                    <input v-model="endDate" type="date" placeholder="終了日" class="input" />
                </div>
                <button @click="searchByOther" class="search-button">検索</button>
            </div>
        </div>
        <customer-detail v-if="searchType === 'phone' && selectedCustomer" :customer="selectedCustomer" />
        <customer-list v-if="searchType === 'other' && showList" :customers="searchResults"
            @select-customer="showDetails" />
    </div>
</template>

<script>
import CustomerList from '@/components/CustomerList.vue'
import CustomerDetail from '@/components/CustomerDetail.vue'
import axios from 'axios'

export default {
    name: 'CustomerManagement',
    components: {
        CustomerList,
        CustomerDetail
    },
    data() {
        return {
            searchType: 'phone',
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
        async searchByPhone() {
            try {
                const response = await axios.get(`http://localhost:3000/api/v1/customers/detail/${this.phoneNumber}`);
                this.selectedCustomer = response.data.data;
                this.showList = false;
            } catch (error) {
                console.error('顧客の検索中にエラーが発生しました:', error);
                this.selectedCustomer = null;
            }
        },
        searchByOther() {
            // その他の条件での検索ロジックを実装
            // API呼び出しなどを行い、結果をsearchResultsに格納
            this.searchResults = [] // 仮の処理。実際はAPI結果を設定
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

.search-type-selector {
    display: flex;
    justify-content: space-between;
    margin-bottom: 15px;
}

.search-type-selector button {
    flex: 1;
    padding: 10px;
    background-color: #f0f0f0;
    border: 1px solid #ddd;
    cursor: pointer;
}

.search-type-selector button.active {
    background-color: #4CAF50;
    color: white;
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