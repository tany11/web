<template>
    <v-dialog v-model="dialog" fullscreen hide-overlay transition="dialog-bottom-transition">
        <v-card>
            <v-toolbar dark color="primary">
                <v-btn icon dark @click="closeModal">
                    <v-icon>mdi-close</v-icon>
                </v-btn>
                <v-toolbar-title>キャスト詳細</v-toolbar-title>
            </v-toolbar>
            <v-card-text>
                <v-container v-if="loading">
                    <v-row justify="center">
                        <v-progress-circular indeterminate color="primary"></v-progress-circular>
                    </v-row>
                </v-container>
                <v-container v-else-if="editedCast">
                    <v-row>
                        <v-col cols="12" sm="6">
                            <v-text-field v-model="editedCast.CastName" label="キャスト名"></v-text-field>
                            <v-text-field v-model="editedCast.CastID" label="キャストID" readonly></v-text-field>
                            <v-text-field v-model="editedCast.LineID" label="LINE ID"></v-text-field>
                            
                            <v-menu v-model="birthDateMenu" :close-on-content-click="false" transition="scale-transition" offset-y min-width="auto">
                                <template v-slot:activator="{ props }">
                                    <v-text-field :value="formattedBirthDate" label="生年月日" readonly v-bind="props"></v-text-field>
                                </template>
                                <v-date-picker v-model="editedCast.birthdate" @update:model-value="updateBirthDate"></v-date-picker>
                            </v-menu>
                            
                            <v-text-field v-model="editedCast.LastName" label="姓"></v-text-field>
                            <v-text-field v-model="editedCast.FirstName" label="名"></v-text-field>
                            
                            <v-menu v-model="effectiveDateMenu" :close-on-content-click="false" transition="scale-transition" offset-y min-width="auto">
                                <template v-slot:activator="{ props }">
                                    <v-text-field :value="formattedEffectiveDate" label="証明書有効期限" readonly v-bind="props"></v-text-field>
                                </template>
                                <v-date-picker v-model="editedCast.effective_date" @update:model-value="updateEffectiveDate"></v-date-picker>
                            </v-menu>
                        </v-col>
                        <v-col cols="12" sm="6">
                            <v-checkbox v-model="editedCast.TattooFlg" label="タトゥー" true-value="1" false-value="0"></v-checkbox>
                            <v-text-field v-if="editedCast.TattooFlg === '1'" v-model="editedCast.TattooArea" label="タトゥーの場所"></v-text-field>
                            <v-text-field v-model="editedCast.Allergy" label="アレルギー"></v-text-field>
                            <v-checkbox v-model="editedCast.StretchMarksFlg" label="妊娠線" true-value="1" false-value="0"></v-checkbox>
                            <v-checkbox v-model="editedCast.SmokingFlg" label="喫煙" true-value="1" false-value="0"></v-checkbox>
                            <v-checkbox v-model="editedCast.ForeignerFlg" label="外国人対応" true-value="1" false-value="0"></v-checkbox>
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col cols="12">
                            <h3>対応可能オプション</h3>
                        </v-col>
                    </v-row>
                    <v-row>
                        <v-col v-for="(value, key) in editedCast.option_flags" :key="key" cols="12" sm="2">
                            <v-checkbox v-model="editedCast.option_flags[key]" :label="key"></v-checkbox>
                        </v-col>
                    </v-row>
                </v-container>
                <v-container v-else>
                    <v-alert type="error" dense>
                        キャスト情報の取得に失敗しました。
                    </v-alert>
                </v-container>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="blue darken-1" text @click="closeModal">キャンセル</v-btn>
                <v-btn color="blue darken-1" text @click="save">保存</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import { ref, watch, onMounted, computed } from 'vue'
import axios from 'axios'
import { useStore } from 'vuex'

export default {
    props: {
        castId: {
            type: Number,
            required: true
        }
    },
    emits: ['close', 'update'],
    setup(props, { emit }) {
        const store = useStore()
        const dialog = ref(true)
        const editedCast = ref(null)
        const loading = ref(false)
        const birthDateMenu = ref(false)
        const effectiveDateMenu = ref(false)

        const formatDate = (date) => {
            if (!date) return null
            const d = new Date(date)
            if (isNaN(d.getTime())) return null
            console.log(d)
            // この関数は日付オブジェクトをJST（日本標準時）に調整して返します。
            // Vuetify3ではDate型で渡す必要があるため、文字列ではなくDate型を返します。
            return new Date(d.getTime() + (9 * 60 * 60 * 1000)); // JSTに調整したDate型を返す
        }

        const formattedBirthDate = computed(() => formatDate(editedCast.value?.birthdate))
        const formattedEffectiveDate = computed(() => formatDate(editedCast.value?.effective_date))

        const updateBirthDate = (date) => {
            if (editedCast.value) {
                editedCast.value.birthdate = date
            }
            birthDateMenu.value = false
        }

        const updateEffectiveDate = (date) => {
            if (editedCast.value) {
                editedCast.value.effective_date = date
            }
            effectiveDateMenu.value = false
        }

        const fetchCastDetails = async () => {
            if (!props.castId) return
            loading.value = true
            try {
                const response = await axios.get(`${store.state.apiBaseUrl}/cast/${props.castId}`)
                editedCast.value = response.data.data
                // 日付を適切な形式に変換
                if (editedCast.value.birthdate) {
                    editedCast.value.birthdate = new Date(editedCast.value.birthdate)
                }
                if (editedCast.value.effective_date) {
                    editedCast.value.effective_date = new Date(editedCast.value.effective_date)
                }
            } catch (error) {
                console.error('キャスト詳細の取得に失敗しました', error)
            } finally {
                loading.value = false
            }
        }

        const closeModal = () => {
            dialog.value = false
            emit('close')
        }

        const save = async () => {
            try {
                const castData = { ...editedCast.value }
                // 日付をISO形式に変換
                if (castData.birthdate) {
                    castData.birthdate = new Date(castData.birthdate).toISOString()
                }
                if (castData.effective_date) {
                    castData.effective_date = new Date(castData.effective_date).toISOString()
                }
                console.log("castData", castData)
                await axios.put(`${store.state.apiBaseUrl}/cast/${castData.ID}`, castData)
                emit('update')
                closeModal()
            } catch (error) {
                console.error('キャストの更新に失敗しました', error)
            }
        }

        onMounted(fetchCastDetails)

        watch(() => props.castId, (newCastId) => {
            if (newCastId) {
                fetchCastDetails()
            }
        })

        return {
            dialog,
            editedCast,
            loading,
            closeModal,
            save,
            formattedBirthDate,
            formattedEffectiveDate,
            updateBirthDate,
            updateEffectiveDate,
            birthDateMenu,
            effectiveDateMenu,
        }
    }
}
</script>