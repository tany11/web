<template>
    <v-card>
        <v-card-title>{{ isNewMemo ? 'メモ作成' : 'メモ詳細' }}</v-card-title>
        <v-card-text>
            <v-form @submit.prevent="saveMemo">
                <v-text-field v-model="editedMemo.Title" label="タイトル"></v-text-field>
                <v-textarea v-model="editedMemo.Notes" label="内容"></v-textarea>
                <v-select v-model="editedMemo.ActualModel" :items="castList" item-title="name" item-value="cast_id"
                    label="対象キャスト"></v-select>
                <v-text-field v-model="editedMemo.ScheduledBox" label="予定枠(30分単位)" type="text" :step="30"
                    @keydown.down.prevent="adjustValue('ScheduledBox', 30)"
                    @keydown.up.prevent="adjustValue('ScheduledBox', -30)"></v-text-field>
                <v-btn color="primary" block type="submit">{{ isNewMemo ? '作成' : '更新' }}</v-btn>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-btn @click="$emit('close')">閉じる</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="success" @click="completeMemo" v-if="!isNewMemo">完了</v-btn>
            <v-btn color="error" @click="confirmDelete" v-if="!isNewMemo">削除</v-btn>
        </v-card-actions>

        <v-dialog v-model="deleteConfirmDialog" max-width="300">
            <v-card>
                <v-card-title>削除の確認</v-card-title>
                <v-card-text>このメモを削除してもよろしいですか？</v-card-text>
                <v-card-actions>
                    <v-spacer></v-spacer>
                    <v-btn color="primary" text @click="deleteConfirmDialog = false">キャンセル</v-btn>
                    <v-btn color="error" text @click="deleteMemo">削除</v-btn>
                </v-card-actions>
            </v-card>
        </v-dialog>
    </v-card>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';

export default {
    props: {
        memo: Object,
    },
    data() {
        return {
            editedMemo: {},
            castList: [],
            deleteConfirmDialog: false,
        };
    },
    computed: {
        ...mapState(['apiBaseUrl']),
        isNewMemo() {
            return this.memo && this.memo.ID === null;
        },
        formattedScheduledTime() {
            if (!this.editedMemo.start_time) return '';
            const date = new Date(this.editedMemo.start_time);
            let hours = date.getHours();
            const minutes = date.getMinutes().toString().padStart(2, '0');

            if (hours >= 24) {
                hours = hours % 24;
            }

            const hoursStr = hours.toString().padStart(2, '0');

            return `${hoursStr}:${minutes}`;
        },
    },
    methods: {
        async fetchCastList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.castList = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        },
        async saveMemo() {
            try {
                let memoData = { ...this.editedMemo };
                memoData.ScheduledBox = memoData.ScheduledBox.toString();

                let response;
                if (this.isNewMemo) {
                    response = await axios.post(`${this.apiBaseUrl}/tips`, memoData);
                } else {
                    response = await axios.put(`${this.apiBaseUrl}/tips/${memoData.ID}`, memoData);
                }
                this.$emit('close');
                this.$emit('memo-updated', response.data);
            } catch (error) {
                console.error('メモの保存に失敗しました:', error);
            }
        },
        adjustValue(field, step) {
            const currentValue = parseInt(this.editedMemo[field]) || 0;
            this.editedMemo[field] = Math.max(0, currentValue + step).toString();
        },
        async completeMemo() {
            try {
                await axios.put(`${this.apiBaseUrl}/tips/${this.editedMemo.ID}/completion`);
                this.$emit('close');
                this.$emit('memo-updated');
            } catch (error) {
                console.error('メモの完了に失敗しました:', error);
            }
        },
        confirmDelete() {
            this.deleteConfirmDialog = true;
        },
        async deleteMemo() {
            try {
                await axios.put(`${this.apiBaseUrl}/tips/${this.editedMemo.ID}/delete`);
                this.deleteConfirmDialog = false;
                this.$emit('close');
                this.$emit('memo-deleted', this.editedMemo.ID);
            } catch (error) {
                console.error('メモの削除に失敗しました:', error);
            }
        },
    },
    created() {
        this.editedMemo = {
            ID: '',
            Title: '',
            Content: '',
            ActualModel: '',
            start_time: new Date(),
            end_time: new Date(new Date().getTime() + 30 * 60000),
            ScheduledBox: '',
            ...this.memo
        };
        if (this.editedMemo.ScheduledBox) {
            this.editedMemo.ScheduledBox = this.editedMemo.ScheduledBox.toString();
        }
        this.fetchCastList();
    },
};
</script>
