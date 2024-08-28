<template>
    <v-container class="store-edit">
        <v-row>
            <v-col>
                <h1 class="text-h4 mb-4">店舗編集</h1>
            </v-col>
        </v-row>
        <v-row>
            <v-col cols="12" md="8" offset-md="2">
                <v-form @submit.prevent="submitForm">
                    <v-text-field v-model="storeName" label="店舗名 *" required></v-text-field>
                    <v-btn type="submit" color="primary" block :disabled="!isFormValid">
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
        const storeName = ref('')
        const registrationResult = ref('')
        const resultClass = ref('')

        const isFormValid = computed(() => {
            return storeName.value.trim() !== ''
        })

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('店舗名を入力してください。')
                return
            }

            const formData = {
                storeName: storeName.value,
            }
            if (confirm('この店舗を登録してもよろしいですか？')) {
                try {
                    const response = await axios.post(`${store.state.apiBaseUrl}/store`, formData)
                    console.log('店舗が登録されました', response.data)
                    registrationResult.value = '店舗が正常に登録されました。'
                    resultClass.value = 'success'
                    resetForm()
                } catch (error) {
                    console.error('登録エラー', error)
                    registrationResult.value = '店舗の登録に失敗しました。もう一度お試しください。'
                    resultClass.value = 'error'
                }
            }
        }

        const resetForm = () => {
            storeName.value = ''
        }

        return {
            storeName,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>