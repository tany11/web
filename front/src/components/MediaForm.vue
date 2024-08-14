<template>
    <div class="cast-registration">
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="mediaName">媒体名 *</label>
                <input type="text" id="mediaName" v-model="mediaName" required>
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