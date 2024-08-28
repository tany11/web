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
                    <template v-slot:item.actions="{ item }">
                        <v-btn small color="primary" @click="openModal(item)" class="mr-2">
                            詳細
                        </v-btn>
                        <v-btn small color="error" @click="confirmDelete(item)">
                            削除
                        </v-btn>
                    </template>
                </v-data-table>
            </v-col>
        </v-row>

        <cast-detail v-if="showModal" :cast="selectedCast" @close="closeModal" />

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

        const openModal = (cast) => {
            selectedCast.value = cast
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
                    await axios.delete(`${store.state.apiBaseUrl}/cast/${castToDelete.value.ID}`)
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
            deleteCast
        }
    }
}
</script>