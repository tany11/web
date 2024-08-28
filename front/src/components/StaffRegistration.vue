<template>
    <v-container class="staff-registration">
        <v-row justify="center">
            <v-col cols="12" sm="8" md="6">
                <h1 class="text-h4 text-center mb-6">スタッフ登録</h1>
                <v-form @submit.prevent="submitForm">
                    <v-row>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="staffLastName" label="スタッフ姓 *" required></v-text-field>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="staffFirstName" label="スタッフ名 *" required></v-text-field>
                        </v-col>
                    </v-row>

                    <v-text-field v-model="password" label="パスワード *" type="password" required></v-text-field>

                    <v-text-field v-model="lineId" label="LINE ID *" required></v-text-field>

                    <v-text-field v-model="phoneNumber" label="電話番号 *" type="tel" required></v-text-field>

                    <v-checkbox v-model="officeFlg" label="内勤"></v-checkbox>

                    <v-checkbox v-model="driverFlg" label="ドライバー"></v-checkbox>

                    <v-checkbox v-model="webFlg" label="Webスタッフ"></v-checkbox>

                    <v-btn type="submit" color="primary" block :disabled="!isFormValid" class="mt-4">
                        登録
                    </v-btn>
                </v-form>
                <v-alert v-if="registrationResult" :type="resultClass" class="mt-4">
                    {{ registrationResult }}
                </v-alert>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
import { ref, computed } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
    setup() {
        const store = useStore()
        const staffLastName = ref('')
        const staffFirstName = ref('')
        const password = ref('')
        const lineId = ref('')
        const phoneNumber = ref('')
        const officeFlg = ref(false)
        const driverFlg = ref(false)
        const webFlg = ref(false)
        const registrationResult = ref('')
        const resultClass = ref('')

        const isFormValid = computed(() => {
            return staffLastName.value.trim() !== '' &&
                staffFirstName.value.trim() !== '' &&
                password.value.trim() !== '' &&
                lineId.value.trim() !== '' &&
                phoneNumber.value.trim() !== '' &&
                (officeFlg.value || driverFlg.value || webFlg.value)
        })

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('すべての必須項目を入力してください。')
                return
            }

            const formData = {
                staffLastName: staffLastName.value,
                staffFirstName: staffFirstName.value,
                passwordHash: password.value,
                lineId: lineId.value,
                phoneNumber: phoneNumber.value,
                officeFlg: officeFlg.value ? "1" : "0",
                driverFlg: driverFlg.value ? "1" : "0",
                webFlg: webFlg.value ? "1" : "0"
            }
            if (confirm('スタッフを登録してもよろしいですか？')) {
                try {
                    const response = await axios.post(`${store.state.apiBaseUrl}/staff`, formData)
                    console.log('スタッフが登録されました', response.data)
                    registrationResult.value = 'スタッフが正常に登録されました。'
                    resultClass.value = 'success'
                    resetForm()
                } catch (error) {
                    console.error('登録エラー', error)
                    registrationResult.value = 'スタッフの登録に失敗しました。もう一度お試しください。'
                    resultClass.value = 'error'
                }
            }
        }

        const resetForm = () => {
            staffLastName.value = ''
            staffFirstName.value = ''
            password.value = ''
            lineId.value = ''
            phoneNumber.value = ''
            officeFlg.value = false
            driverFlg.value = false
            webFlg.value = false
        }

        return {
            staffLastName,
            staffFirstName,
            password,
            lineId,
            phoneNumber,
            officeFlg,
            driverFlg,
            webFlg,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>