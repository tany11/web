<template>
    <div class="customer-management">
        <h1>電話番号検索</h1>
        <div class="search-form">
            <div>
                <input v-model="phoneNumber" placeholder="電話番号" class="input" @keyup.enter="searchByPhone" />
                <button @click="searchByPhone" class="search-button">検索</button>
            </div>
        </div>
        <customer-detail v-if="selectedCustomer" :customer="selectedCustomer" @close="closeCustomerDetail" />
        <customer-list v-if="searchType === 'other' && showList" :customers="searchResults"
            @select-customer="showDetails" />
    </div>
</template>

<script>
import CustomerList from '@/components/CustomerList.vue'
import CustomerDetail from '@/components/CustomerDetail.vue'
import axios from 'axios'
import { mapState } from 'vuex'

export default {
    name: 'CustomerManagement',
    components: {
        CustomerList,
        CustomerDetail
    },
    computed: {
        ...mapState(['apiBaseUrl'])
    },
    data() {
        return {
            searchType: 'phone',
            phoneNumber: '',
            lastFourDigits: '',
            castID: '',
            startDate: '',
            endDate: '',
            searchResults: [],
            selectedCustomer: null,
            showList: false,
            castList: []
        }
    },
    methods: {
        async searchByPhone() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/customers/detail/${this.phoneNumber}`);
                this.selectedCustomer = response.data.data;
                this.showList = false;
            } catch (error) {
                console.error('顧客の検索中にエラーが発生しました:', error);
                this.selectedCustomer = null;
            }
        },
        async searchByOther() {
            try {
                const params = {
                    phoneLast4: this.lastFourDigits,
                    castID: this.castID,
                    createdFrom: this.startDate,
                    createdTo: this.endDate
                }
                const response = await axios.get(`${this.apiBaseUrl}/customers/search`, { params })
                this.searchResults = response.data.data
                this.showList = true
                this.selectedCustomer = null
            } catch (error) {
                console.error('顧客の検索中にエラーが発生しました:', error)
                this.searchResults = []
                this.showList = true
                this.selectedCustomer = null
            }
        },
        async showDetails(customer) {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/customers/detail/${customer.PhoneNumber}`);
                this.selectedCustomer = response.data.data;
                this.showList = false;
            } catch (error) {
                console.error('顧客の詳細取得中にエラーが発生しました:', error);
                this.selectedCustomer = null;
            }
        },
        closeCustomerDetail() {
            this.selectedCustomer = null;
            if (this.searchType === 'other') {
                this.showList = true;
            }
        },
        async fetchCastList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.castList = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        }
    },
    mounted() {
        this.fetchCastList();
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