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
                                    <v-col cols="12" sm="4">
                                        <v-text-field v-model="lastFourDigits" label="電話番号（下4桁）"
                                            placeholder="下4桁を入力してください"></v-text-field>
                                    </v-col>
                                    <v-col cols="12" sm="4">
                                        <v-select v-model="castID" :items="castList" item-title="name"
                                            item-value="cast_id" label="キャスト" :return-object="false" clearable
                                            @update:model-value="onCastChange">
                                            <template v-slot:selection="{ item }">
                                                {{ getCastName(item.value) }}
                                            </template>
                                        </v-select>
                                    </v-col>
                                    <v-col cols="12" sm="4">
                                        <v-select v-model="storeID" :items="storeList" item-title="name" item-value="id"
                                            label="店舗" :return-object="false" clearable
                                            @update:model-value="onStoreChange">
                                            <template v-slot:selection="{ item }">
                                                {{ getStoreName(item.value) }}
                                            </template>
                                        </v-select>
                                    </v-col>
                                    <v-col cols="12">
                                        <v-row>
                                            <v-col cols="12" sm="6">
                                                <v-menu v-model="startDateMenu" :close-on-content-click="false"
                                                    :nudge-right="40" transition="scale-transition" offset-y
                                                    min-width="auto">
                                                    <template v-slot:activator="{ props }">
                                                        <v-text-field v-model="startDateFormatted" label="開始日"
                                                            prepend-icon="mdi-calendar" readonly
                                                            v-bind="props"></v-text-field>
                                                    </template>
                                                    <v-date-picker v-model="startDate"
                                                        @update:model-value="updateStartDate"></v-date-picker>
                                                </v-menu>
                                            </v-col>
                                            <v-col cols="12" sm="6">
                                                <v-menu v-model="endDateMenu" :close-on-content-click="false"
                                                    :nudge-right="40" transition="scale-transition" offset-y
                                                    min-width="auto">
                                                    <template v-slot:activator="{ props }">
                                                        <v-text-field v-model="endDateFormatted" label="終了日"
                                                            prepend-icon="mdi-calendar" readonly
                                                            v-bind="props"></v-text-field>
                                                    </template>
                                                    <v-date-picker v-model="endDate"
                                                        @update:model-value="updateEndDate"></v-date-picker>
                                                </v-menu>
                                            </v-col>
                                        </v-row>
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
        const startDate = ref(null)
        const endDate = ref(null)
        const startDateFormatted = ref('')
        const endDateFormatted = ref('')
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

        const getCastName = (castId) => {
            const cast = castList.value.find(c => c.cast_id === castId)
            return cast ? cast.name : ''
        }

        const getStoreName = (storeId) => {
            const store = storeList.value.find(s => s.id === storeId)
            return store ? store.name : ''
        }

        const onCastChange = (value) => {
            if (value === null) {
                castID.value = null
            }
        }

        const onStoreChange = (value) => {
            if (value === null) {
                storeID.value = null
            }
        }

        const updateStartDate = (newDate) => {
            startDate.value = newDate
            startDateFormatted.value = formatDate(newDate)
            startDateMenu.value = false
        }

        const updateEndDate = (newDate) => {
            endDate.value = newDate
            endDateFormatted.value = formatDate(newDate)
            endDateMenu.value = false
        }

        const formatDate = (date) => {
            if (!date) return ''
            if (date instanceof Date) {
                return `${date.getFullYear()}/${String(date.getMonth() + 1).padStart(2, '0')}/${String(date.getDate()).padStart(2, '0')}`
            }
            if (typeof date === 'string') {
                const [year, month, day] = date.split('-')
                return `${year}/${month}/${day}`
            }
            return ''
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
            closeCustomerDetail,
            getCastName,
            getStoreName,
            onCastChange,
            onStoreChange,
            updateStartDate,
            updateEndDate,
            startDateFormatted,
            endDateFormatted
        }
    }
}
</script>