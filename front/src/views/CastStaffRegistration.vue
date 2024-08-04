<template>
    <div class="registration">
        <h1>登録フォーム</h1>
        <div class="toggle-buttons">
            <button @click="setMode('cast')" :class="{ active: mode === 'cast' }">キャスト</button>
            <button @click="setMode('staff')" :class="{ active: mode === 'staff' }">スタッフ</button>
        </div>
        <component :is="currentComponent" @registration-complete="handleRegistrationComplete"></component>
        <div v-if="registrationResult" class="registration-result" :class="resultClass">
            {{ registrationResult }}
        </div>
    </div>
</template>

<script>
import { defineComponent, ref, computed } from 'vue'
import CastForm from '@/components/CastForm.vue'
import StaffForm from '@/components/StaffForm.vue'


export default defineComponent({
    components: {
        CastForm,
        StaffForm
    },
    setup() {
        const mode = ref('cast')
        const registrationResult = ref('')
        const resultClass = ref('')

        const currentComponent = computed(() => {
            return mode.value === 'cast' ? CastForm : StaffForm
        })

        const setMode = (newMode) => {
            mode.value = newMode
            registrationResult.value = ''
            resultClass.value = ''
        }

        const handleRegistrationComplete = (result) => {
            registrationResult.value = result.message
            resultClass.value = result.success ? 'success' : 'error'
        }

        return {
            mode,
            currentComponent,
            setMode,
            registrationResult,
            resultClass,
            handleRegistrationComplete
        }
    }
})
</script>

<style scoped>
.registration {
    max-width: 600px;
    margin: 0 auto;
    padding: 20px;
}

h1 {
    text-align: center;
}

.toggle-buttons {
    display: flex;
    justify-content: center;
    margin-bottom: 20px;
}

.toggle-buttons button {
    padding: 10px 20px;
    margin: 0 10px;
    background-color: #f0f0f0;
    border: none;
    border-radius: 4px;
    cursor: pointer;
}

.toggle-buttons button.active {
    background-color: #4CAF50;
    color: white;
}
</style>