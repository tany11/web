<template>
    <v-container class="media-registration">
        <v-row justify="center">
            <v-col cols="12" sm="8" md="6">
                <h1 class="text-h4 text-center mb-6">媒体登録</h1>
                <v-form @submit.prevent="submitForm">
                    <v-text-field v-model="mediaName" label="媒体名 *" required></v-text-field>

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
        const mediaName = ref('')
        const registrationResult = ref('')
        const resultClass = ref('')
        const isFormValid = computed(() => {
            return mediaName.value.trim() !== ''
        })

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('媒体名を入力してください。')
                return
            }

            const formData = {
                mediaName: mediaName.value,
            }
            if (confirm('この媒体を登録してもよろしいですか？')) {
                try {
                    const response = await axios.post(`${store.state.apiBaseUrl}/media`, formData)
                    console.log('媒体が登録されました', response.data)
                    registrationResult.value = '媒体が正常に登録されました。'
                    resultClass.value = 'success'
                    resetForm()
                } catch (error) {
                    console.error('登録エラー', error)
                    registrationResult.value = '媒体の登録に失敗しました。もう一度お試しください。'
                    resultClass.value = 'error'
                }
            }
        }

        const resetForm = () => {
            mediaName.value = ''
        }

        return {
            mediaName,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>