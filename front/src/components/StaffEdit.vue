<template>
    <v-container class="staff-edit">
        <v-row>
            <v-col>
                <h1 class="text-h4 text-center mb-6">スタッフ一覧</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col>
                <v-data-table :headers="headers" :items="staffs" :items-per-page="10" :loading="loading"
                    class="elevation-1">
                    <template v-slot:item="{ item }">
                        <tr>
                            <td>{{ item.StaffLastName }} {{ item.StaffFirstName }}</td>
                            <td>{{ item.PhoneNumber }}</td>
                            <td>{{ item.LineID }}</td>
                            <td>
                                <v-chip small :color="item.OfficeFlg === '1' ? 'primary' : 'grey'"
                                    text-color="white">内勤</v-chip>
                                <v-chip small :color="item.DriverFlg === '1' ? 'primary' : 'grey'"
                                    text-color="white">ドライバー</v-chip>
                                <v-chip small :color="item.WebFlg === '1' ? 'primary' : 'grey'"
                                    text-color="white">Webスタッフ</v-chip>
                            </td>
                            <td>
                                <v-btn small color="primary" @click="openModal(item.id)" class="mr-2">
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

        <staff-detail v-if="showModal" :staffId="selectedStaff" @close="closeModal" />

        <v-dialog v-model="showDeleteDialog" max-width="300px">
            <v-card>
                <v-card-title>削除の確認</v-card-title>
                <v-card-text>このスタッフを削除してもよろしいですか？</v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="blue darken-1" text @click="showDeleteDialog = false">キャンセル</v-btn>
                    <v-btn color="blue darken-1" text @click="deleteStaff">削除</v-btn>
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
import StaffDetail from './StaffDetail.vue'

export default {
    components: {
        StaffDetail
    },
    setup() {
        const store = useStore()
        const staffs = ref([])
        const showModal = ref(false)
        const selectedStaff = ref(null)
        const loading = ref(false)
        const showDeleteDialog = ref(false)
        const staffToDelete = ref(null)
        const showError = ref(false)
        const errorMessage = ref('')

        const headers = [
            { text: 'スタッフ名', value: 'staffName' },
            { text: '電話番号', value: 'phoneNumber' },
            { text: 'LINE ID', value: 'lineId' },
            { text: '役割', value: 'roles', sortable: false },
            { text: '操作', value: 'actions', sortable: false }
        ]

        const fetchStaffs = async () => {
            loading.value = true
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/staff`)
                staffs.value = response.data.data
            } catch (error) {
                console.error('スタッフ一覧の取得に失敗しました', error)
                showError.value = true
                errorMessage.value = 'スタッフ一覧の取得に失敗しました'
            } finally {
                loading.value = false
            }
        }

        const openModal = (staffId) => {
            selectedStaff.value = staffId
            showModal.value = true
        }

        const closeModal = () => {
            showModal.value = false
            selectedStaff.value = null
        }

        const confirmDelete = (staff) => {
            staffToDelete.value = staff
            showDeleteDialog.value = true
        }

        const deleteStaff = async () => {
            if (staffToDelete.value) {
                try {
                    await axios.put(`${store.state.apiBaseUrl}/staff/${staffToDelete.value.ID}`, {
                        IsDeleted: '1'
                    })
                    await fetchStaffs() // スタッフ一覧を再取得
                } catch (error) {
                    console.error('スタッフの削除に失敗しました', error)
                    showError.value = true
                    errorMessage.value = 'スタッフの削除に失敗しました'
                } finally {
                    showDeleteDialog.value = false
                    staffToDelete.value = null
                }
            }
        }

        onMounted(fetchStaffs)

        return {
            staffs,
            showModal,
            selectedStaff,
            headers,
            loading,
            showDeleteDialog,
            showError,
            errorMessage,
            openModal,
            closeModal,
            confirmDelete,
            deleteStaff
        }
    }
}
</script>

<style scoped>
.v-chip {
    margin-right: 4px;
}
</style>