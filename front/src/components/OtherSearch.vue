<template>
    <div class="other-search">
        <h1>一覧検索</h1>
        <div class="search-form">
            <div>
                <div class="input-group">
                    <input v-model="lastFourDigits" placeholder="電話番号（下4桁）" class="input" />
                    <select v-model="castID" class="input">
                        <option value="">キャストを選択</option>
                        <option v-for="cast in castList" :key="cast.cast_id" :value="cast.cast_id">
                            {{ cast.name }}
                        </option>
                    </select>
                    <select v-model="storeID" class="input">
                        <option value="">店舗を選択</option>
                        <option v-for="store in storeList" :key="store.store_id" :value="store.store_id">
                            {{ store.name }}
                        </option>
                    </select>
                </div>
                <div class="input-group">
                    <input v-model="startDate" type="date" placeholder="開始日" class="input" />
                    <input v-model="endDate" type="date" placeholder="終了日" class="input" />
                </div>
                <button @click="searchByOther" class="search-button">検索</button>
            </div>
        </div>
        <customer-detail v-if="selectedCustomer" :customer="selectedCustomer" @close="closeCustomerDetail" />
        <customer-list v-if="showList" :customers="searchResults" @select-customer="showDetails" />
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
            castList: [],
            storeList: []
        }
    },
    methods: {
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
        },
        async fetchStoreList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/store/dropdown`);
                this.storeList = response.data.data || [];
            } catch (error) {
                console.error('店舗リストの取得に失敗しました:', error);
            }
        }
    },
    mounted() {
        this.fetchCastList();
        this.fetchStoreList();
    }
}
</script>

<style scoped>
.other-search {
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
    background-color: #45a030;
}
</style>