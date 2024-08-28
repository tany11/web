<template>
    <v-container class="cast-registration">
        <v-row justify="center">
            <v-col cols="12" sm="8" md="6">
                <h1 class="text-h4 text-center mb-6">キャスト登録</h1>
                <v-form @submit.prevent="submitForm">
                    <v-text-field v-model="castName" label="キャスト名 *" required></v-text-field>

                    <v-text-field v-model="password" label="パスワード *" type="password" required></v-text-field>

                    <v-text-field v-model="lineId" label="LINE ID *" required></v-text-field>

                    <v-menu v-model="menu" :close-on-content-click="false" transition="scale-transition" offset-y
                        min-width="auto">
                        <template v-slot:activator="{ props }">
                            <v-text-field v-model="formattedBirthDate" label="生年月日 *" readonly v-bind="props"
                                required></v-text-field>
                        </template>
                        <v-locale-provider locale="ja">
                            <v-date-picker v-model="birthDate" @update:model-value="updateBirthDate">
                            </v-date-picker>
                        </v-locale-provider>
                    </v-menu>

                    <v-text-field v-model="phoneNumber" label="電話番号" type="tel"></v-text-field>

                    <v-text-field v-model="email" label="メールアドレス" type="email"></v-text-field>

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
        const castName = ref('')
        const password = ref('')
        const lineId = ref('')
        const birthDate = ref(null)
        const formattedBirthDate = ref('')
        const phoneNumber = ref('')
        const email = ref('')
        const registrationResult = ref('')
        const resultClass = ref('')
        const menu = ref(false)
        const maxDate = new Date().toISOString().split('T')[0] // 今日の日付を最大値として設定
        const isFormValid = computed(() => {
            return castName.value.trim() !== '' &&
                password.value.trim() !== '' &&
                lineId.value.trim() !== '' &&
                birthDate.value !== null &&
                phoneNumber.value.trim() !== ''
        })

        const formatDate = (date) => {
            if (!date) return ''
            if (typeof date === 'string') {
                const [year, month, day] = date.split('-')
                return `${year}年${parseInt(month)}月${parseInt(day)}日`
            } else if (date instanceof Date) {
                const year = date.getFullYear()
                const month = date.getMonth() + 1 // getMonth()は0から始まるため、1を加算
                const day = date.getDate()
                return `${year}年${month}月${day}日`
            }
            return ''
        }

        const updateBirthDate = (date) => {
            birthDate.value = date
            formattedBirthDate.value = formatDate(date)
            menu.value = false
        }

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('すべての必須項目を入力してください。')
                return
            }

            const formData = {
                castName: castName.value,
                passwordHash: password.value,
                lineId: lineId.value,
                birthDate: birthDate.value,
                phoneNumber: phoneNumber.value,
                email: email.value
            }
            if (confirm('このキャストを登録してもよろしいですか？')) {
                try {
                    const response = await axios.post(`${store.state.apiBaseUrl}/cast`, formData)
                    console.log('キャストが登録されました', response.data)
                    registrationResult.value = 'キャストが正常に登録されました。'
                    resultClass.value = 'success'
                    resetForm()
                } catch (error) {
                    console.error('登録エラー', error)
                    registrationResult.value = 'キャストの登録に失敗しました。もう一度お試しください。'
                    resultClass.value = 'error'
                }
            }
        }

        const resetForm = () => {
            castName.value = ''
            password.value = ''
            lineId.value = ''
            birthDate.value = null
            phoneNumber.value = ''
            email.value = ''
        }

        return {
            castName,
            password,
            lineId,
            birthDate,
            phoneNumber,
            email,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass,
            menu,
            maxDate,
            formatDate,
            formattedBirthDate,
            updateBirthDate,
            menu,
        }
    }
}
</script>