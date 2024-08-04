<template>
    <div class="cast-registration">
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="castName">キャスト名 *</label>
                <input type="text" id="castName" v-model="castName" required>
            </div>

            <div class="form-group">
                <label for="password">パスワード *</label>
                <input type="password" id="password" v-model="password" required>
            </div>

            <div class="form-group">
                <label for="lineId">LINE ID *</label>
                <input type="text" id="lineId" v-model="lineId" required>
            </div>

            <div class="form-group">
                <label for="birthDate">生年月日 *</label>
                <input type="date" id="birthDate" v-model="birthDate" required>
            </div>

            <button type="submit" :disabled="!isFormValid">登録</button>
        </form>
        <div v-if="registrationResult" class="registration-result" :class="resultClass">
            {{ registrationResult }}
        </div>
    </div>
</template>

<script>
import { ref, computed } from 'vue'
import axios from 'axios'

export default {
    setup() {
        const castName = ref('')
        const password = ref('')
        const lineId = ref('')
        const birthDate = ref('')
        const registrationResult = ref('')
        const resultClass = ref('')

        const isFormValid = computed(() => {
            return castName.value.trim() !== '' &&
                password.value.trim() !== '' &&
                lineId.value.trim() !== '' &&
                birthDate.value !== ''
        })

        const submitForm = async () => {
            if (!isFormValid.value) {
                alert('すべての必須項目を入力してください。')
                return
            }

            const formData = {
                castName: castName.value,
                passwordHash: password.value,
                lineId: lineId.value,
                birthDate: birthDate.value
            }
            if (confirm('このキャストを登録してもよろしいですか？')) {
                try {
                    const response = await axios.post('http://localhost:3000/api/v1/cast', formData)
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
            birthDate.value = ''
        }

        return {
            castName,
            password,
            lineId,
            birthDate,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>

<style scoped>
.cast-registration {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
}

.form-group {
    margin-bottom: 15px;
}

label {
    display: block;
    margin-bottom: 5px;
}

input[type="text"],
input[type="password"],
input[type="date"] {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
}

button {
    width: 100%;
    padding: 10px;
    background-color: #4CAF50;
    color: white;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

button:hover {
    background-color: #45a049;
}

.registration-result {
    margin-top: 15px;
    padding: 10px;
    border-radius: 4px;
}

.registration-result.success {
    background-color: #dff0d8;
    color: #3c763d;
}

.registration-result.error {
    background-color: #f2dede;
    color: #a94442;
}
</style>