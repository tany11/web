<template>
    <v-card>
        <v-card-title>オーダー詳細</v-card-title>
        <v-card-text>
            <v-form @submit.prevent="updateOrder">
                <v-row>

                    <v-col cols="12" sm="8">
                        <v-select v-model="order.StoreID" :items="storeList" item-title="name" item-value="id"
                            label="店名"></v-select>
                    </v-col>
                </v-row>

                <v-text-field v-model="order.PhoneNumber" label="電話番号"></v-text-field>
                <v-text-field v-model="order.CustomerName" label="お客様名"></v-text-field>
                <v-text-field v-model="order.ModelName" label="モデル名"></v-text-field>

                <v-select v-model="order.ActualModel" :items="castList" item-title="name" item-value="cast_id"
                    label="実モデル"></v-select>

                <v-row>
                    <v-col cols="5">
                        <v-text-field v-model.number="order.CourseMin" label="コース（分）" type="number" :step="10"
                            @keydown.down.prevent="adjustValue('CourseMin', 10)"
                            @keydown.up.prevent="adjustValue('CourseMin', -10)"></v-text-field>
                    </v-col>
                    <v-col cols="2" class="d-flex justify-center align-center">
                        <span class="text-h5">+</span>
                    </v-col>
                    <v-col cols="5">
                        <v-text-field v-model.number="order.ExtraCourse" label="（分）" type="number" :step="5"
                            @keydown.down.prevent="adjustValue('ExtraCourse', 5)"
                            @keydown.up.prevent="adjustValue('ExtraCourse', -5)"></v-text-field>
                    </v-col>
                </v-row>

                <v-text-field v-model.number="order.Price" label="料金" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('Price', 1000)"></v-text-field>
                <v-text-field v-model="order.PostalCode" label="郵便番号"></v-text-field>
                <v-text-field v-model="order.Address" label="住所"></v-text-field>

                <v-select v-model="order.DriverID" :items="driverStaffList" item-title="name" item-value="staff_id"
                    label="送り"></v-select>

                <v-text-field v-model.number="order.ReservationFee" label="指名料" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('ReservationFee', 1000)"></v-text-field>
                <v-text-field v-model.number="order.TransportationFee" label="交通費" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('TransportationFee', 1000)"></v-text-field>
                <v-text-field v-model.number="order.TravelCost" label="出張費" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('TravelCost', 1000)"></v-text-field>

                <v-select v-model="order.Media" :items="mediaList" item-title="name" item-value="id"
                    label="媒体"></v-select>

                <v-textarea v-model="order.Notes" label="備考"></v-textarea>

                <v-select v-model="order.CardstaffID" :items="officeStaffList" item-title="name" item-value="staff_id"
                    label="カード"></v-select>
                <v-select v-model="order.OrderStaffID" :items="officeStaffList" item-title="name" item-value="staff_id"
                    label="受注者"></v-select>
                <!-- 到着時刻の表示を変更 -->
                <v-text-field v-model="formattedScheduledTime" label="到着時刻" readonly></v-text-field>

                <v-btn color="primary" block type="submit">更新</v-btn>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-btn @click="$emit('close')">閉じる</v-btn>
            <v-spacer></v-spacer>
            <v-btn color="primary" @click="openNewMemoModal">メモ作成</v-btn>
        </v-card-actions>
        <v-dialog v-model="showNewMemoModal" max-width="800px">
            <MemoDetail :memo="newMemo" @close="closeNewMemoModal" @memo-updated="handleMemoCreated" />
        </v-dialog>
    </v-card>
</template>

<script>
import MemoDetail from './MemoDetail.vue';
import axios from 'axios';
import { mapState } from 'vuex';

export default {
    props: {
        order: Object
    },
    components: {
        MemoDetail
    },
    data() {
        return {
            storeList: [],
            castList: [],
            driverStaffList: [],
            officeStaffList: [],
            mediaList: [],
            showNewMemoModal: false,
            newMemo: {}
        };
    },
    computed: {
        ...mapState(['apiBaseUrl']),
        formattedScheduledTime() {
            if (!this.order || !this.order.ScheduledTime) return '';
            const date = new Date(this.order.ScheduledTime);
            let hours = date.getHours();
            const minutes = date.getMinutes().toString().padStart(2, '0');

            // 24時以降の処理
            if (hours >= 24) {
                hours = hours % 24;
            }

            // 時間を文字列に変換し、必要に応じて先頭にゼロを追加
            const hoursStr = hours.toString().padStart(2, '0');

            return `${hoursStr}:${minutes}`;
        }
    },
    methods: {
        async fetchDropdownData() {
            try {
                const [stores, casts, drivers, offices, media] = await Promise.all([
                    axios.get(`${this.apiBaseUrl}/store/dropdown`),
                    axios.get(`${this.apiBaseUrl}/cast/dropdown`),
                    axios.get(`${this.apiBaseUrl}/staff/dropdown?role=driver`),
                    axios.get(`${this.apiBaseUrl}/staff/dropdown?role=office`),
                    axios.get(`${this.apiBaseUrl}/media/dropdown`)
                ]);

                this.storeList = stores.data.data || [];
                this.castList = casts.data.data || [];
                this.driverStaffList = drivers.data.data || [];
                this.officeStaffList = offices.data.data || [];
                this.mediaList = media.data.data || [];
            } catch (error) {
                console.error('ドロップダウンデータの取得に失敗しました:', error);
            }
        },
        async updateOrder() {
            try {
                // INT項目の値を整数に変換
                this.order.CourseMin = parseInt(this.order.CourseMin, 10);
                this.order.ExtraCourse = parseInt(this.order.ExtraCourse, 10);
                this.order.Price = parseInt(this.order.Price, 10);
                this.order.ReservationFee = parseInt(this.order.ReservationFee, 10);
                this.order.TransportationFee = parseInt(this.order.TransportationFee, 10);
                this.order.TravelCost = parseInt(this.order.TravelCost, 10);

                // ActualModelが空の場合、空文字列を明示的に設定
                if (!this.order.ActualModel) {
                    this.order.ActualModel = '';
                }

                await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}`, this.order);
                this.$emit('close');
                this.$emit('order-updated'); // 更新イベントを発火
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
            }
        },
        adjustValue(field, amount) {
            this.order[field] = Math.max(0, (this.order[field] || 0) + amount);
        },
        openNewMemoModal() {
            this.newMemo = {
                ActualModel: this.order.ActualModel,
                start_time: this.order.ScheduledTime,
            };
            this.showNewMemoModal = true;
        },
        closeNewMemoModal() {
            this.showNewMemoModal = false;
        },
        handleMemoCreated() {
            this.closeNewMemoModal();
            this.$emit('memo-created');
        }
    },
    async mounted() {
        await this.fetchDropdownData();
    }
};
</script>
