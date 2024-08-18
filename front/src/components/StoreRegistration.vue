<template>
    <div class="store-registration">
        <h1>店舗登録</h1>
        <form @submit.prevent="submitForm">
            <div class="form-group">
                <label for="storeCode">店舗コード *</label>
                <input type="number" id="storeCode" v-model.number="storeCode" required min="0">
            </div>
            <div class="form-group">
                <label for="storeName">店舗名 *</label>
                <input type="text" id="storeName" v-model="storeName" required>
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
        const storeCode = ref('')
        const storeName = ref('')
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
        }

        return {
            storeCode,
            storeName,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass
        }
    }
}
</script>

<style scoped>
.store-registration {
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

input[type="number"] {
    -moz-appearance: textfield;
    appearance: textfield;
}

input[type="number"]::-webkit-inner-spin-button,
input[type="number"]::-webkit-outer-spin-button {
    -webkit-appearance: none;
    margin: 0;
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