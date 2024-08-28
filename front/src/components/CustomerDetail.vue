<template>
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
            <v-toolbar dark color="primary">
                <v-btn icon dark @click="closeModal">
                    <v-icon>mdi-close</v-icon>
                </v-btn>
                <v-toolbar-title>顧客詳細</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                <v-container>
                    <v-row>
                        <v-col cols="12" sm="6">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>名前</v-list-item-title>
                                    <v-list-item-subtitle>{{ customer.CustomerName }}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>電話番号</v-list-item-title>
                                    <v-list-item-subtitle>
                                        <v-btn text color="primary" :href="`tel:${customer.PhoneNumber}`">
                                            {{ customer.PhoneNumber }}
                                        </v-btn>
                                    </v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="12">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>住所</v-list-item-title>
                                    <v-list-item-subtitle>{{ customer.Address }}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="12">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>メモ</v-list-item-title>
                                    <v-list-item-subtitle>{{ customer.Memo }}</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>合計金額</v-list-item-title>
                                    <v-list-item-subtitle>{{ customer.TotalPrice }}円</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-list-item>
                                <v-list-item-content>
                                    <v-list-item-title>利用回数</v-list-item-title>
                                    <v-list-item-subtitle>{{ customer.TotalUseTime }}回</v-list-item-subtitle>
                                </v-list-item-content>
                            </v-list-item>
                        </v-col>
                    </v-row>
                    <v-divider class="my-4"></v-divider>
                    <h3 class="text-h6 mb-2">注文履歴</h3>
                    <v-data-table v-if="customer.OrderList && customer.OrderList.length" :headers="headers"
                        :items="customer.OrderList" :items-per-page="5" class="elevation-1">
                        <template v-slot:item.CreatedAt="{ item }">
                            {{ formatDate(item.CreatedAt) }}
                        </template>
                        <template v-slot:item.StoreID="{ item }">
                            {{ getStoreName(item.StoreID) }}
                        </template>
                        <template v-slot:item.CourseMin="{ item }">
                            {{ item.CourseMin }}分
                        </template>
                        <template v-slot:item.Price="{ item }">
                            {{ item.Price }}円
                        </template>
                        <template v-slot:item.ActualModel="{ item }">
                            {{ getCastName(item.ActualModel) }}
                        </template>
                    </v-data-table>
                    <v-alert v-else type="info" class="mt-4">
                        注文履歴はありません。
                    </v-alert>
                </v-container>
            </v-card-text>
        </v-card>
    </v-dialog>
</template>

<script>
import { ref, onMounted, watch } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
    name: 'CustomerDetail',
    props: {
        customer: Object
    },
    setup(props, { emit }) {
        const store = useStore()
        const dialog = ref(true)
        const castList = ref([])
        const storeList = ref([])
        const headers = [
            { text: '日付', value: 'CreatedAt' },
            { text: '利用店舗', value: 'StoreID' },
            { text: 'コース時間', value: 'CourseMin' },
            { text: '料金', value: 'Price' },
            { text: 'キャスト', value: 'ActualModel' }
        ]

        const formatDate = (dateString) => {
            const date = new Date(dateString)
            return date.toLocaleDateString('ja-JP')
        }

        const closeModal = () => {
            dialog.value = false
            emit('close')
        }

        const getCastName = (castId) => {
            const cast = castList.value.find(c => c.cast_id === castId)
            return cast ? cast.name : 'Unknown'
        }

        const getStoreName = (storeId) => {
            const store = storeList.value.find(s => s.id === storeId)
            return store ? store.name : 'Unknown'
        }

        const fetchCastList = async () => {
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/cast/dropdown`)
                castList.value = response.data.data || []
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error)
            }
        }

        const fetchStoreList = async () => {
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/store/dropdown`)
                storeList.value = response.data.data || []
            } catch (error) {
                console.error('店舗リストの取得に失敗しました:', error)
            }
        }

        onMounted(() => {
            fetchCastList()
            fetchStoreList()
        })

        watch(() => props.customer, (newCustomer) => {
            if (newCustomer) {
                dialog.value = true
            }
        })

        return {
            dialog,
            headers,
            formatDate,
            closeModal,
            getCastName,
            getStoreName
        }
    }
}
</script>