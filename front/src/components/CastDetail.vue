<template>
    <v-dialog v-model="dialogVisible" max-width="600px">
        <v-card>
            <v-card-title>
                キャスト詳細
                <v-spacer></v-spacer>
                <v-btn icon @click="closeModal">
                    <v-icon>mdi-close</v-icon>
                </v-btn>
            </v-card-title>
            <v-card-text>
                <v-container v-if="loading">
                    <v-row justify="center">
                        <v-progress-circular indeterminate color="primary"></v-progress-circular>
                    </v-row>
                </v-container>
                <v-container v-else-if="castDetails">
                    <v-row>
                        <v-col cols="12">
                            <div class="text-subtitle-1 font-weight-bold mb-2">基本情報</div>
                        </v-col>
                    </v-row>
                    <v-list dense>
                        <v-list-item>
                            <v-list-item-title>キャスト名：</v-list-item-title>
                            <v-list-item-subtitle>{{ castDetails.data.CastName }}</v-list-item-subtitle>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-title>キャストID：</v-list-item-title>
                            <v-list-item-subtitle>{{ castDetails.data.CastID }}</v-list-item-subtitle>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-title>LINE ID：</v-list-item-title>
                            <v-list-item-subtitle>{{ castDetails.data.LineID }}</v-list-item-subtitle>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-title>生年月日：</v-list-item-title>
                            <v-list-item-subtitle>{{ formatDate(castDetails.data.birthdate) }}</v-list-item-subtitle>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-title>登録日：</v-list-item-title>
                            <v-list-item-subtitle>{{ formatDate(castDetails.data.CreatedAt) }}</v-list-item-subtitle>
                        </v-list-item>
                        <v-list-item>
                            <v-list-item-title>最終更新日：</v-list-item-title>
                            <v-list-item-subtitle>{{ formatDate(castDetails.data.UpdatedAt) }}</v-list-item-subtitle>
                        </v-list-item>
                    </v-list>
                    <v-row>
                        <v-col cols="12">
                            <div class="text-subtitle-1 font-weight-bold mb-2">追加情報</div>
                        </v-col>
                    </v-row>
                    <v-data-table v-if="castDetails.WorkHistory && castDetails.WorkHistory.length" :headers="headers"
                        :items="castDetails.WorkHistory" dense>
                        <template v-slot:item.Date="{ item }">
                            {{ formatDate(item.Date) }}
                        </template>
                        <template v-slot:item.StoreID="{ item }">
                            {{ getStoreName(item.StoreID) }}
                        </template>
                        <template v-slot:item.WorkTime="{ item }">
                            {{ formatWorkTime(item.StartTime, item.EndTime) }}
                        </template>
                    </v-data-table>
                    <v-alert v-else type="info" dense>
                        出勤履歴はありません。
                    </v-alert>
                </v-container>
                <v-container v-else>
                    <v-alert type="error" dense>
                        キャスト情報の取得に失敗しました。
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
    name: 'CastDetail',
    props: ['cast'],
    emits: ['close'],
    setup(props, { emit }) {
        const store = useStore()
        const castDetails = ref(null)
        const storeList = ref([])
        const dialogVisible = ref(false)
        const loading = ref(false)

        const headers = [
            { text: '日付', value: 'Date' },
            { text: '勤務店舗', value: 'StoreID' },
            { text: '勤務時間', value: 'WorkTime' },
        ]

        const fetchCastDetail = async () => {
            if (props.cast) {
                loading.value = true
                try {
                    const response = await axios.get(`${store.state.apiBaseUrl}/cast/${props.cast.ID}`)
                    castDetails.value = response.data
                    dialogVisible.value = true
                } catch (error) {
                    console.error('キャスト詳細の取得に失敗しました', error)
                } finally {
                    loading.value = false
                }
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

        const formatDate = (dateString) => {
            if (!dateString) return 'N/A'
            const date = new Date(dateString)
            return date.toLocaleDateString('ja-JP')
        }

        const getStoreName = (storeId) => {
            const store = storeList.value.find(s => s.id === storeId)
            return store ? store.name : 'Unknown'
        }

        const formatWorkTime = (startTime, endTime) => {
            return `${startTime} - ${endTime}`
        }

        const closeModal = () => {
            dialogVisible.value = false
            emit('close')
        }

        onMounted(() => {
            fetchStoreList()
        })

        watch(() => props.cast, (newCast) => {
            if (newCast) {
                fetchCastDetail()
            } else {
                dialogVisible.value = false
            }
        }, { immediate: true })

        return {
            castDetails,
            dialogVisible,
            headers,
            formatDate,
            getStoreName,
            formatWorkTime,
            closeModal,
            loading
        }
    }
}
</script>