<template>
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
            <v-toolbar dark color="primary">
                <v-btn icon dark @click="closeModal">
                    <v-icon>mdi-close</v-icon>
                </v-btn>
                <v-toolbar-title>スタッフ詳細</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                <v-container v-if="loading">
                    <v-row justify="center">
                        <v-progress-circular indeterminate color="primary"></v-progress-circular>
                    </v-row>
                </v-container>
                <v-container v-else-if="editedStaff">
                    <v-row>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="editedStaff.StaffLastName" label="姓"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="editedStaff.StaffFirstName" label="名"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="editedStaff.PhoneNumber" label="電話番号"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="editedStaff.LineID" label="LINE ID"></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="formattedBirthDate" label="生年月日" readonly></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-checkbox v-model="editedStaff.OfficeFlg" label="内勤" true-value="1"
                                false-value="0"></v-checkbox>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-checkbox v-model="editedStaff.DriverFlg" label="ドライバー" true-value="1"
                                false-value="0"></v-checkbox>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-checkbox v-model="editedStaff.WebFlg" label="Webスタッフ" true-value="1"
                                false-value="0"></v-checkbox>
                        </v-col>
                    </v-row>
                </v-container>
                <v-container v-else>
                    <v-alert type="error" dense>
                        スタッフ情報の取得に失敗しました。
                    </v-alert>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="closeModal">キャンセル</v-btn>
                <v-btn color="blue darken-1" text @click="save">保存</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import { ref, watch, onMounted, computed } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
    props: {
        staffId: {
            type: Number,
            required: true
        }
    },
    emits: ['close', 'update'],
    setup(props, { emit }) {
        const store = useStore()
        const dialog = ref(true)
        const editedStaff = ref(null)
        const loading = ref(false)
        const formattedBirthDate = computed(() => {
            return formatDisplayDate(editedStaff.value?.birthdate)
        })

        const fetchStaffDetails = async () => {
            if (!props.staffId) return
            loading.value = true
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/staff/${props.staffId}`)
                editedStaff.value = response.data.data
            } catch (error) {
                console.error('スタッフ詳細の取得に失敗しました', error)
            } finally {
                loading.value = false
            }
        }

        const formatDate = (date) => {
            if (!date) return ''
            const d = new Date(date)
            return d.toISOString().split('T')[0]
        }

        const formatDisplayDate = (date) => {
            if (!date) return ''
            const d = new Date(date)
            if (isNaN(d.getTime())) return '' // 無効な日付の場合は空文字を返す
            return `${d.getFullYear()}年${d.getMonth() + 1}月${d.getDate()}日`
        }

        const closeModal = () => {
            dialog.value = false
            emit('close')
        }

        const save = async () => {
            try {
                const staffData = { ...editedStaff.value }
                // 日付を適切な形式に変換
                if (staffData.birthdate) {
                    const date = new Date(staffData.birthdate)
                    staffData.birthdate = date.toISOString().split('T')[0] + 'T00:00:00+09:00'
                }
                await axios.put(`${store.state.apiBaseUrl}/staff/${staffData.id}`, staffData)
                emit('update')
                closeModal()
            } catch (error) {
                console.error('スタッフの更新に失敗しました', error)
            }
        }

        onMounted(fetchStaffDetails)

        watch(() => props.staffId, (newStaffId) => {
            if (newStaffId) {
                fetchStaffDetails()
            }
        })

        return {
            dialog,
            editedStaff,
            loading,
            closeModal,
            save,
            formattedBirthDate
        }
    }
}
</script>