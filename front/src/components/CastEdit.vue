<template>
    <v-container class="cast-edit">
        <v-row>
            <v-col>
                <h1 class="text-h4 text-center mb-6">キャスト一覧</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-data-table :headers="headers" :items="casts" :items-per-page="10" :loading="loading"
                    class="elevation-1">
                    <template v-slot:item="{ item }">
                        <tr>
                            <td class="working-btn-cell">
                                <v-btn small :color="item.WorkingFlg === '1' ? 'success' : 'primary'"
                                    :class="{ 'lighten-3': item.WorkingFlg === '1' }" @click="toggleWorking(item)"
                                    class="working-btn">
                                    {{ item.WorkingFlg === '1' ? '出勤中' : '出勤' }}
                                </v-btn>
                            </td>
                            <td>{{ item.CastName }}</td>
                            <td>
                                <v-btn small color="primary" @click="openModal(item.ID)" class="mr-2">
                                    詳細
                                </v-btn>
                                <v-btn small color="error" @click="confirmDelete(item)">
                                    削除
                                </v-btn>
                            </td>
                        </tr>
                    </template>
                </v-data-table>
            </v-col>
        </v-row>

        <cast-detail v-if="showModal" :castId="selectedCast" @close="closeModal" @update="fetchCasts" />

        <v-dialog v-model="showDeleteDialog" max-width="300px">
            <v-card>
                <v-card-title>削除の確認</v-card-title>
                <v-card-text>このキャストを削除してもよろしいですか？</v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="showDeleteDialog = false">キャンセル</v-btn>
                    <v-btn color="blue darken-1" text @click="deleteCast">削除</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>

        <v-snackbar v-model="showError" color="error" timeout="3000">
            {{ errorMessage }}
        </v-snackbar>
    </v-container>
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
        const loading = ref(false)
        const showDeleteDialog = ref(false)
        const castToDelete = ref(null)
        const showError = ref(false)
        const errorMessage = ref('')

        const headers = [
            { text: '出勤状態', value: 'WorkingFlg', sortable: false },
            { text: 'キャスト名', value: 'CastName' },
            { text: '操作', value: 'actions', sortable: false }
        ]

        const fetchCasts = async () => {
            loading.value = true
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/cast`)
                casts.value = response.data.data
            } catch (error) {
                console.error('キャスト一覧の取得に失敗しました', error)
                showError.value = true
                errorMessage.value = 'キャスト一覧の取得に失敗しました'
            } finally {
                loading.value = false
            }
        }

        const toggleWorking = async (cast) => {
            try {
                await axios.put(`${store.state.apiBaseUrl}/cast/${cast.ID}/working`)
                await fetchCasts() // キャスト一覧を再取得してWorkingFlgを更新
            } catch (error) {
                console.error('出勤状態の更新に失敗しました', error)
                showError.value = true
                errorMessage.value = '出勤状態の更新に失敗しました'
            }
        }

        const openModal = (castId) => {
            selectedCast.value = castId
            showModal.value = true
        }

        const closeModal = () => {
            showModal.value = false
            selectedCast.value = null
        }

        const confirmDelete = (cast) => {
            castToDelete.value = cast
            showDeleteDialog.value = true
        }

        const deleteCast = async () => {
            if (castToDelete.value) {
                try {
                    await axios.put(`${store.state.apiBaseUrl}/cast/${castToDelete.value.ID}`, {
                        IsDeleted: '1'
                    })
                    await fetchCasts() // キャスト一覧を再取得
                } catch (error) {
                    console.error('キャストの削除に失敗しました', error)
                    showError.value = true
                    errorMessage.value = 'キャストの削除に失敗しました'
                } finally {
                    showDeleteDialog.value = false
                    castToDelete.value = null
                }
            }
        }

        onMounted(fetchCasts)

        return {
            casts,
            showModal,
            selectedCast,
            headers,
            loading,
            showDeleteDialog,
            showError,
            errorMessage,
            openModal,
            closeModal,
            confirmDelete,
            deleteCast,
            toggleWorking,
            fetchCasts // ここでfetchCastsを返す
        }
    }
}
</script>

<style scoped>
.v-btn.lighten-3 {
    opacity: 0.7;
}

.working-btn-cell {
    width: 80px;
    padding-right: 8px !important;
}

.working-btn {
    width: 72px;
}
</style>