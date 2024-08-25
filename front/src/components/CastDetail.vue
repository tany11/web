<template>
    <div class="modal" v-if="castDetails" @click.self="closeModal">
        <div class="modal-content">
            <span class="close" @click="closeModal">&times;</span>
            <h2>キャスト詳細</h2>
            <div class="cast-detail">
                <p><strong>キャスト名：</strong>{{ castDetails.CastName }}</p>
                <p><strong>キャストID：</strong>{{ castDetails.CastID }}</p>
                <p><strong>LINE ID：</strong>{{ castDetails.LineID }}</p>
                <p><strong>生年月日：</strong>{{ formatDate(castDetails.birthdate) }}</p>
                <p><strong>登録日：</strong>{{ formatDate(castDetails.CreatedAt) }}</p>
                <p><strong>最終更新日：</strong>{{ formatDate(castDetails.UpdatedAt) }}</p>

                <h3>出勤履歴</h3>
                <table v-if="castDetails.WorkHistory && castDetails.WorkHistory.length">
                    <thead>
                        <tr>
                            <th>日付</th>
                            <th>勤務店舗</th>
                            <th>勤務時間</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="work in castDetails.WorkHistory" :key="work.ID">
                            <td>{{ formatDate(work.Date) }}</td>
                            <td>{{ getStoreName(work.StoreID) }}</td>
                            <td>{{ formatWorkTime(work.StartTime, work.EndTime) }}</td>
                        </tr>
                    </tbody>
                </table>
                <p v-else>出勤履歴はありません。</p>
            </div>
        </div>
    </div>
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

        const fetchCastDetail = async () => {
            if (props.cast) {
                try {
                    const response = await axios.get(`${store.state.apiBaseUrl}/cast/${props.cast.ID}`)
                    castDetails.value = response.data
                } catch (error) {
                    console.error('キャスト詳細の取得に失敗しました', error)
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
            emit('close')
        }

        onMounted(() => {
            fetchCastDetail()
            fetchStoreList()
        })

        watch(() => props.cast, fetchCastDetail)

        return {
            castDetails,
            formatDate,
            getStoreName,
            formatWorkTime,
            closeModal
        }
    }
}
</script>

<style scoped>
.modal {
    position: fixed;
    z-index: 1000;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-content {
    background-color: #292929;
    padding: 30px;
    border: 1px solid #888;
    width: 80%;
    max-width: 600px;
    position: relative;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.close {
    color: #4CAF50;
    position: absolute;
    top: 10px;
    right: 20px;
    font-size: 28px;
    font-weight: bold;
    cursor: pointer;
}

.close:hover,
.close:focus {
    color: #45a049;
    text-decoration: none;
}

.cast-detail {
    margin-top: 20px;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 15px;
}

th,
td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
}

th {
    background-color: #34495E;
}
</style>