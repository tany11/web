<template>
    <v-container class="customer-management">
        <v-row>
            <v-col>
                <h1 class="text-h4 mb-4">顧客管理</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-tabs v-model="searchType" grow>
                    <v-tab value="phone">電話番号検索</v-tab>
                    <v-tab value="other">その他の検索</v-tab>
                </v-tabs>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-card>
                    <v-card-text>
                        <v-form @submit.prevent="searchType === 'phone' ? searchByPhone() : searchByOther()">
                            <v-container>
                                <v-row v-if="searchType === 'phone'">
                                    <v-col cols="12">
                                        <v-text-field v-model="phoneNumber" label="電話番号" placeholder="電話番号を入力してください"
                                            @keyup.enter="searchByPhone"></v-text-field>
                                    </v-col>
                                </v-row>
                                <v-row v-else>
                                    <v-col cols="12" sm="6">
                                        <v-text-field v-model="lastFourDigits" label="電話番号（下4桁）"
                                            placeholder="下4桁を入力してください"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-select v-model="castID" :items="castList" item-text="name"
                                            item-value="cast_id" label="キャスト"></v-select>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-select v-model="storeID" :items="storeList" item-text="name" item-value="id"
                                            label="店舗"></v-select>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-menu v-model="startDateMenu" :close-on-content-click="false"
                                            :nudge-right="40" transition="scale-transition" offset-y min-width="auto">
                                            <template v-slot:activator="{ on, attrs }">
                                                <v-text-field v-model="startDate" label="開始日"
                                                    prepend-icon="mdi-calendar" readonly v-bind="attrs"
                                                    v-on="on"></v-text-field>
                                            </template>
                                            <v-date-picker v-model="startDate"
                                                @input="startDateMenu = false"></v-date-picker>
                                        </v-menu>
                                    </v-col>
                                    <v-col cols="12" sm="6">
                                        <v-menu v-model="endDateMenu" :close-on-content-click="false" :nudge-right="40"
                                            transition="scale-transition" offset-y min-width="auto">
                                            <template v-slot:activator="{ on, attrs }">
                                                <v-text-field v-model="endDate" label="終了日" prepend-icon="mdi-calendar"
                                                    readonly v-bind="attrs" v-on="on"></v-text-field>
                                            </template>
                                            <v-date-picker v-model="endDate"
                                                @input="endDateMenu = false"></v-date-picker>
                                        </v-menu>
                                    </v-col>
                                </v-row>
                                <v-row>
                                    <v-col>
                                        <v-btn color="primary" block
                                            @click="searchType === 'phone' ? searchByPhone() : searchByOther()"
                                            :loading="loading">
                                            検索
                                        </v-btn>
                                    </v-col>
                                </v-row>
                            </v-container>
                        </v-form>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>
        <customer-detail v-if="selectedCustomer" :customer="selectedCustomer" @close="closeCustomerDetail" />
        <customer-list v-if="searchType === 'other' && showList" :customers="searchResults"
            @select-customer="showDetails" />
    </v-container>
</template>

<script>
import { ref, computed, onMounted } from 'vue'
import { useStore } from 'vuex'
import axios from 'axios'
import CustomerList from '@/components/CustomerList.vue'
import CustomerDetail from '@/components/CustomerDetail.vue'

export default {
    name: 'CustomerManagement',
    components: {
        CustomerList,
        CustomerDetail
    },
    setup() {
        const store = useStore()
        const apiBaseUrl = computed(() => store.state.apiBaseUrl)

        const searchType = ref('phone')
        const phoneNumber = ref('')
        const lastFourDigits = ref('')
        const castID = ref('')
        const storeID = ref('')
        const startDate = ref('')
        const endDate = ref('')
        const searchResults = ref([])
        const selectedCustomer = ref(null)
        const showList = ref(false)
        const castList = ref([])
        const storeList = ref([])
        const loading = ref(false)
        const startDateMenu = ref(false)
        const endDateMenu = ref(false)

        const searchByPhone = async () => {
            loading.value = true
            try {
                const response = await axios.get(`${apiBaseUrl.value}/customers/detail/${phoneNumber.value}`)
                selectedCustomer.value = response.data.data
                showList.value = false
            } catch (error) {
                console.error('顧客の検索中にエラーが発生しました:', error)
                selectedCustomer.value = null
            } finally {
                loading.value = false
            }
        }

        const searchByOther = async () => {
            loading.value = true
            try {
                const params = {
                    phoneLast4: lastFourDigits.value,
                    castID: castID.value,
                    storeID: storeID.value,
                    createdFrom: startDate.value,
                    createdTo: endDate.value
                }
                const response = await axios.get(`${apiBaseUrl.value}/customers/search`, { params })
                searchResults.value = response.data.data
                showList.value = true
                selectedCustomer.value = null
            } catch (error) {
                console.error('顧客の検索中にエラーが発生しました:', error)
                searchResults.value = []
                showList.value = true
                selectedCustomer.value = null
            } finally {
                loading.value = false
            }
        }

        const showDetails = async (customer) => {
            loading.value = true
            try {
                const response = await axios.get(`${apiBaseUrl.value}/customers/detail/${customer.PhoneNumber}`)
                selectedCustomer.value = response.data.data
                showList.value = false
            } catch (error) {
                console.error('顧客の詳細取得中にエラーが発生しました:', error)
                selectedCustomer.value = null
            } finally {
                loading.value = false
            }
        }

        const closeCustomerDetail = () => {
            selectedCustomer.value = null
            if (searchType.value === 'other') {
                showList.value = true
            }
        }

        const fetchCastList = async () => {
            try {
                const response = await axios.get(`${apiBaseUrl.value}/cast/dropdown`)
                castList.value = response.data.data || []
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error)
            }
        }

        const fetchStoreList = async () => {
            try {
                const response = await axios.get(`${apiBaseUrl.value}/store/dropdown`)
                storeList.value = response.data.data || []
            } catch (error) {
                console.error('店舗リストの取得に失敗しました:', error)
            }
        }

        onMounted(() => {
            fetchCastList()
            fetchStoreList()
        })

        return {
            searchType,
            phoneNumber,
            lastFourDigits,
            castID,
            storeID,
            startDate,
            endDate,
            searchResults,
            selectedCustomer,
            showList,
            castList,
            storeList,
            loading,
            startDateMenu,
            endDateMenu,
            searchByPhone,
            searchByOther,
            showDetails,
            closeCustomerDetail
        }
    }
}
</script>