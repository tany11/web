<template>
    <v-container class="store-registration">
        <v-row justify="center">
            <v-col cols="12" sm="8" md="6">
                <h1 class="text-h4 text-center mb-6">店舗登録</h1>
                <v-form @submit.prevent="submitForm">
                    <v-radio-group v-model="selectedStore" row>
                        <v-radio label="高円寺店" value="0"></v-radio>
                        <v-radio label="立川店" value="1"></v-radio>
                    </v-radio-group>

                    <v-text-field v-model.number="storeCode" label="店舗コード *" type="number" required
                        min="0"></v-text-field>

                    <v-text-field v-model="storeName" label="店舗名 *" required></v-text-field>

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
        const storeCode = ref('')
        const storeName = ref('')
        const selectedStore = ref('0')
        const registrationResult = ref('')
        const resultClass = ref('')

        const isFormValid = computed(() => {
            return storeCode.value !== '' && !isNaN(storeCode.value) && storeName.value.trim() !== ''
        })

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('有効な店舗コード（整数）と店舗名を入力してください。')
                return
            }

            const formData = {
                storeCode: parseInt(storeCode.value, 10),
                storeName: storeName.value,
                DummyStoreFlg: selectedStore.value === '1' ? '1' : '0'
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
            storeCode.value = ''
            storeName.value = ''
            selectedStore.value = '0'
        }

        return {
            storeCode,
            storeName,
            selectedStore,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>