<template>
    <div class="cast-edit">
        <h1>キャスト一覧</h1>
        <table class="cast-list">
            <thead>
                <tr>
                    <th>キャスト名</th>
                    <th>操作</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="cast in casts" :key="cast.ID">
                    <td>{{ cast.CastName }}</td>
                    <td>
                        <button @click="openModal(cast)" class="edit-btn">詳細</button>
                        <button @click="deleteCast(cast.ID)" class="delete-btn">削除</button>
                    </td>
                </tr>
            </tbody>
        </table>

        <CastDetail v-if="showModal" :cast="selectedCast" @close="closeModal" />
    </div>
</template>

<script>
import { ref, onMounted } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'
import CastDetail from './CastDetail.vue'

export default {
    components: {
        CastDetail
    },
    setup() {
        const store = useStore()
        const casts = ref([])
        const showModal = ref(false)
        const selectedCast = ref(null)

        const fetchCasts = async () => {
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/cast`)
                casts.value = response.data.data
            } catch (error) {
                console.error('キャスト一覧の取得に失敗しました', error)
            }
        }

        const openModal = (cast) => {
            selectedCast.value = cast
            showModal.value = true
        }

        const closeModal = () => {
            showModal.value = false
            selectedCast.value = null
        }

        const deleteCast = async (id) => {
            if (confirm('このキャストを削除してもよろしいですか？')) {
                try {
                    await axios.delete(`${store.state.apiBaseUrl}/cast/${id}`)
                    alert('キャストが削除されました')
                    fetchCasts() // キャスト一覧を再取得
                } catch (error) {
                    console.error('キャストの削除に失敗しました', error)
                    alert('キャストの削除に失敗しました')
                }
            }
        }

        onMounted(fetchCasts)

        return {
            casts,
            showModal,
            selectedCast,
            openModal,
            closeModal,
            deleteCast
        }
    }
}
</script>

<style scoped>
.cast-edit {
    max-width: 800px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
}

.cast-list {
    width: 100%;
    border-collapse: collapse;
    margin-top: 20px;
}

.cast-list th,
.cast-list td {
    border: 1px solid #ddd;
    padding: 12px;
    text-align: left;
}

.cast-list th {
    background-color: #f2f2f2;
    font-weight: bold;
}

.edit-btn,
.delete-btn {
    padding: 6px 12px;
    margin-right: 10px;
    border: none;
    border-radius: 4px;
    cursor: pointer;
    font-size: 14px;
}

.edit-btn {
    background-color: #4CAF50;
    color: white;
}

.delete-btn {
    background-color: #f44336;
    color: white;
}

.edit-btn:hover,
.delete-btn:hover {
    opacity: 0.8;
}
</style>