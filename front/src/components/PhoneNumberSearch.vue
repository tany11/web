<template>
    <v-container class="customer-management">
        <v-row>
            <v-col>
                <h1 class="text-h4 mb-4">電話番号検索</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12" md="8" offset-md="2">
                <v-form @submit.prevent="searchByPhone">
                    <v-text-field v-model="phoneNumber" label="電話番号" placeholder="電話番号を入力してください"
                        @keyup.enter="searchByPhone"></v-text-field>
                    <v-btn color="primary" block @click="searchByPhone" :loading="loading">
                        検索
                    </v-btn>
                </v-form>
            </v-col>
        </v-row>
        <customer-detail v-if="selectedCustomer" :customer="selectedCustomer" @close="closeCustomerDetail" />
        <customer-list v-if="searchType === 'other' && showList" :customers="searchResults"
            @select-customer="showDetails" />
    </v-container>
</template>

<script>
import { ref, computed } from 'vue'
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
        const startDate = ref('')
        const endDate = ref('')
        const searchResults = ref([])
        const selectedCustomer = ref(null)
        const showList = ref(false)
        const castList = ref([])
        const loading = ref(false)

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

        fetchCastList()

        return {
            searchType,
            phoneNumber,
            lastFourDigits,
            castID,
            startDate,
            endDate,
            searchResults,
            selectedCustomer,
            showList,
            castList,
            loading,
            searchByPhone,
            searchByOther,
            showDetails,
            closeCustomerDetail
        }
    }
}
</script>