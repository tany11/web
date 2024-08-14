え<template>
    <div class="staff-registration">
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
                <div class="checkbox-group">
                    <label>
                        <input type="checkbox" v-model="officeFlg"> 内勤
                    </label>
                    <label>
                        <input type="checkbox" v-model="driverFlg"> ドライバー
                    </label>
                    <label>
                        <input type="checkbox" v-model="webFlg"> Webスタッフ
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