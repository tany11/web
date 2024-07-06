<template>
    <div class="staff-registration">
        <h1>スタッフ登録</h1>
        <form @submit.prevent="submitForm">
            <div class="form-group name-group">
                <div class="name-field">
                    <label for="staffLastName">スタッフ姓 *</label>
                    <input type="text" id="staffLastName" v-model="staffLastName" required>
                </div>
                <div class="name-field">
                    <label for="staffFirstName">スタッフ名 *</label>
                    <input type="text" id="staffFirstName" v-model="staffFirstName" required>
                </div>
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
                <label for="phoneNumber">電話番号 *</label>
                <input type="tel" id="phoneNumber" v-model="phoneNumber" required>
            </div>

            <div class="form-group">
                <label>役職 *</label>
                <div class="radio-group">
                    <label>
                        <input type="radio" v-model="position" value="office" required> 内勤
                    </label>
                    <label>
                        <input type="radio" v-model="position" value="driver" required> ドライバー
                    </label>
                    <label>
                        <input type="radio" v-model="position" value="web" required> Webスタッフ
                    </label>
                </div>
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
        const staffLastName = ref('')
        const staffFirstName = ref('')
        const password = ref('')
        const lineId = ref('')
        const phoneNumber = ref('')
        const position = ref('')
        const registrationResult = ref('')
        const resultClass = ref('')

        const isFormValid = computed(() => {
            return staffLastName.value.trim() !== '' &&
                staffFirstName.value.trim() !== '' &&
                password.value.trim() !== '' &&
                lineId.value.trim() !== '' &&
                phoneNumber.value.trim() !== '' &&
                position.value !== ''
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
                position: position.value
            }

            try {
                const response = await axios.post('http://localhost:3000/api/v1/staff', formData)
                console.log('スタッフが登録されました', response.data)
                registrationResult.value = 'スタッフが正常に登録されました。'
                resultClass.value = 'success'
                // フォームをリセット
                resetForm()
            } catch (error) {
                console.error('登録エラー', error)
                registrationResult.value = 'スタッフの登録に失敗しました。もう一度お試しください。'
                resultClass.value = 'error'
            }
        }
        const resetForm = () => {
            staffLastName.value = ''
            staffFirstName.value = ''
            password.value = ''
            lineId.value = ''
            phoneNumber.value = ''
            position.value = ''
        }
        return {
            staffLastName,
            staffFirstName,
            password,
            lineId,
            phoneNumber,
            position,
            submitForm,
            isFormValid,
            registrationResult,
            resultClass,
            resetForm
        }
    }
}
</script>

<style scoped>
.staff-registration {
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

.name-group {
    display: flex;
    justify-content: space-between;
}

.name-field {
    width: 48%;
    /* 少し余裕を持たせる */
}

label {
    display: block;
    margin-bottom: 5px;
}

input[type="text"],
input[type="password"],
input[type="tel"] {
    width: 100%;
    padding: 8px;
    border: 1px solid #ddd;
    border-radius: 4px;
}

.radio-group {
    display: flex;
    gap: 20px;
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
</style>