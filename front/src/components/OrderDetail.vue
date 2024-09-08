<template>
    <v-card>
        <v-card-title>{{ isNew ? '新規オーダー' : 'オーダー詳細' }}</v-card-title>
        <v-card-text>
            <v-form @submit.prevent="submitForm">
                <v-row v-if="isNew">
                    <v-col cols="12" sm="6">
                        <v-text-field v-model="order.storeCode" label="店舗コード" @input="filterStores"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                        <v-select v-model="order.StoreID" :items="filteredStores" item-title="name" item-value="id"
                            label="店名"></v-select>
                    </v-col>
                </v-row>
                <v-row v-else>
                    <v-col cols="12" sm="8">
                        <v-select v-model="order.StoreID" :items="storeList" item-title="name" item-value="id"
                            label="店名"></v-select>
                    </v-col>
                    <v-col cols="12" sm="4" class="d-flex align-center">
                        <v-btn v-if="order.CompletionFlg !== '1'" color="success" @click="confirmOrder">確定</v-btn>
                    </v-col>
                </v-row>

                <v-text-field v-model="order.PhoneNumber" label="電話番号" @blur="isNew ? fetchCustomerInfo() : null">
                    <template v-slot:append>
                        <v-btn text small color="primary" @click="showCustomerDetail" v-if="order.PhoneNumber">
                            詳細
                        </v-btn>
                    </template>
                </v-text-field>
                <v-text-field v-model="order.CustomerName" label="お客様名"></v-text-field>
                <v-text-field v-model="order.ModelName" label="モデル名"></v-text-field>

                <v-select v-model="order.ActualModel" :items="castList" item-title="name" item-value="cast_id"
                    label="実モデル"></v-select>

                <v-row>
                    <v-col cols="4" v-if="isNew">
                        <v-text-field v-model.number="order.CourseMin" label="コース（分）" type="number" :step="10"
                            @keydown.down.prevent="adjustValue('CourseMin', 10)"
                            @keydown.up.prevent="adjustValue('CourseMin', -10)"></v-text-field>
                    </v-col>
                    <v-col cols="1" v-if="isNew" class="d-flex justify-center align-center">
                        <span class="text-h5">+</span>
                    </v-col>
                    <v-col cols="3" v-if="isNew">
                        <v-text-field v-model.number="order.ExtraCourse" label="（分）" type="number" :step="5"
                            @keydown.down.prevent="adjustValue('ExtraCourse', 5)"
                            @keydown.up.prevent="adjustValue('ExtraCourse', -5)"></v-text-field>
                    </v-col>
                    <template v-else>
                        <v-col cols="4">
                            <v-text-field v-model.number="order.CourseMin" label="コース（分）" type="number" :step="10"
                                @keydown.down.prevent="adjustValue('CourseMin', 10)"
                                @keydown.up.prevent="adjustValue('CourseMin', -10)"></v-text-field>
                        </v-col>

                        <v-col cols="1" class="d-flex justify-center align-center">
                            <span class="text-h5">+</span>
                        </v-col>
                        <v-col cols="3">
                            <v-text-field v-model.number="order.ExtraCourse" label="（分）" type="number" :step="5"
                                @keydown.down.prevent="adjustValue('ExtraCourse', 5)"
                                @keydown.up.prevent="adjustValue('ExtraCourse', -5)"></v-text-field>
                        </v-col>
                        <v-col cols="1" class="d-flex justify-center align-center">
                            <span class="text-h5">+</span>
                        </v-col>
                        <v-col cols="3">
                            <v-text-field v-model.number="order.ExtraTime" label="延長（分）" type="number" :step="5"
                                @keydown.down.prevent="adjustValue('ExtraTime', 30)"
                                @keydown.up.prevent="adjustValue('ExtraTime', -30)"></v-text-field>
                        </v-col>
                    </template>
                </v-row>

                <v-text-field v-model.number="order.Price" label="料金" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('Price', 1000)"
                    @keydown.up.prevent="adjustValue('Price', -1000)">
                </v-text-field>
                <v-text-field v-model="order.City" label="市区町村"></v-text-field>
                <v-text-field v-model="order.Address" label="住所"></v-text-field>

                <v-select v-model="order.DriverID" :items="filteredDriverStaffList" item-title="name"
                    item-value="staff_id" label="送り">
                    <template v-slot:prepend-inner>
                        <v-text-field v-model="driverFilter" label="フィルター" dense hide-details
                            @input="filterDrivers"></v-text-field>
                    </template>
                </v-select>

                <v-text-field v-model.number="order.ReservationFee" label="指名料" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('ReservationFee', 1000)"
                    @keydown.up.prevent="adjustValue('ReservationFee', -1000)"></v-text-field>
                <v-text-field v-model.number="order.TransportationFee" label="交通費" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('TransportationFee', 1000)"
                    @keydown.up.prevent="adjustValue('TransportationFee', -1000)"></v-text-field>
                <v-text-field v-model.number="order.TravelCost" label="出張費" type="number" prefix="¥" :step="1000"
                    @keydown.down.prevent="adjustValue('TravelCost', 1000)"
                    @keydown.up.prevent="adjustValue('TravelCost', -1000)"></v-text-field>

                <v-select v-model="order.Media" :items="mediaList" item-title="name" item-value="id"
                    label="媒体"></v-select>

                <v-textarea v-model="order.Notes" label="備考"></v-textarea>

                <v-select v-model="order.CardstaffID" :items="filteredOfficeStaffList" item-title="name"
                    item-value="staff_id" label="カード">
                    <template v-slot:prepend-inner>
                        <v-text-field v-model="cardStaffFilter" label="フィルター" dense hide-details
                            @input="filterOfficeStaff"></v-text-field>
                    </template>
                </v-select>
                <v-select v-model="order.OrderStaffID" :items="filteredOfficeStaffList" item-title="name"
                    item-value="staff_id" label="受注者">
                    <template v-slot:prepend-inner>
                        <v-text-field v-model="orderStaffFilter" label="フィルター" dense hide-details
                            @input="filterOfficeStaff"></v-text-field>
                    </template>
                </v-select>
                <v-text-field v-if="isNew" v-model="order.ScheduledTime" label="到着時刻"
                    type="datetime-local"></v-text-field>
                <v-text-field v-else v-model="formattedScheduledTime" label="到着時刻" readonly></v-text-field>

                <v-row>
                    <v-col>
                        <v-btn color="primary" block type="submit">{{ isNew ? '登録' : '更新' }}</v-btn>
                    </v-col>
                    <v-col v-if="!isNew">
                        <v-btn color="error" block @click="confirmDelete">削除</v-btn>
                    </v-col>
                </v-row>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-btn @click="$emit('close')">閉じる</v-btn>
        </v-card-actions>
    </v-card>

    <v-dialog v-model="showCustomerModal" max-width="600px">
        <CustomerDetail :customer="customerDetail" @close="closeCustomerModal" />
    </v-dialog>

    <v-snackbar v-model="showSnackbar" :color="snackbarColor" :timeout="3000">
        {{ snackbarText }}
        <template v-slot:action="{ attrs }">
            <v-btn text v-bind="attrs" @click="showSnackbar = false">
                閉じる
            </v-btn>
        </template>
    </v-snackbar>

    <v-dialog v-model="confirmDialog" max-width="300">
        <v-card>
            <v-card-title>確認</v-card-title>
            <v-card-text>このオーダーを確定してもよろしいですか？</v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="confirmDialog = false; proceedWithConfirmation()">はい</v-btn>
                <v-btn color="grey" text @click="confirmDialog = false">いいえ</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>

    <v-dialog v-model="deleteConfirmDialog" max-width="300">
        <v-card>
            <v-card-title>確認</v-card-title>
            <v-card-text>このオーダーを削除しても���ろしいですか？</v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="deleteConfirmDialog = false; proceedWithDeletion()">はい</v-btn>
                <v-btn color="grey" text @click="deleteConfirmDialog = false">いいえ</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import CustomerDetail from './CustomerDetail.vue';

export default {
    props: {
        order: Object,
        isNew: {
            type: Boolean,
            default: false
        }
    },
    emits: ['close', 'order-updated'],
    components: {
        CustomerDetail
    },
    data() {
        return {
            storeList: [],
            castList: [],
            driverStaffList: [],
            officeStaffList: [],
            mediaList: [],
            showCustomerModal: false,
            customerDetail: null,
            filteredStores: [],
            showSnackbar: false,
            snackbarText: '',
            snackbarColor: 'info',
            driverFilter: '',
            cardStaffFilter: '',
            orderStaffFilter: '',
            filteredDriverStaffList: [],
            filteredOfficeStaffList: [],
            confirmDialog: false,
            deleteConfirmDialog: false,
        };
    },
    computed: {
        ...mapState(['apiBaseUrl']),
        formattedScheduledTime() {
            if (!this.order || !this.order.ScheduledTime) return '';
            const date = new Date(this.order.ScheduledTime);
            let hours = date.getHours();
            const minutes = date.getMinutes().toString().padStart(2, '0');

            if (hours >= 24) {
                hours = hours % 24;
            }

            const hoursStr = hours.toString().padStart(2, '0');

            return `${hoursStr}:${minutes}`;
        }
    },
    methods: {
        async fetchDropdownData() {
            try {
                const [stores, casts, staff, media] = await Promise.all([
                    axios.get(`${this.apiBaseUrl}/store/dropdown`),
                    axios.get(`${this.apiBaseUrl}/cast/dropdown`),
                    axios.get(`${this.apiBaseUrl}/staff/dropdown`),
                    axios.get(`${this.apiBaseUrl}/media/dropdown`)
                ]);

                this.storeList = stores.data.data || [];
                this.castList = casts.data.data || [];
                this.mediaList = media.data.data || [];

                // スタッフリストをドライバーと事務所スタッフに分類
                const allStaff = staff.data.data || [];
                this.driverStaffList = allStaff.filter(s => s.driver_flg === "1");
                this.officeStaffList = allStaff.filter(s => s.office_flg === "1");

                this.filteredDriverStaffList = [...this.driverStaffList];
                this.filteredOfficeStaffList = [...this.officeStaffList];
            } catch (error) {
                console.error('ドロップダウンデータの取得に失敗しました:', error);
            }
        },
        submitForm() {
            if (this.isNew) {
                this.createOrder();
            } else {
                this.updateOrder();
            }
        },
        async createOrder() {
            try {
                // 日時フィールドをフォーマット
                if (this.order.ScheduledTime) {
                    this.order.ScheduledTime = this.formatDateTimeForServer(this.order.ScheduledTime);
                }

                // INT項目の値を整数に変換
                const intFields = ['CourseMin', 'ExtraTime', 'ExtraCourse', 'Price', 'ReservationFee', 'TransportationFee', 'TravelCost'];
                intFields.forEach(field => {
                    if (this.order[field] !== undefined) {
                        this.order[field] = Math.floor(Number(this.order[field]));
                    }
                });

                const response = await axios.post(`${this.apiBaseUrl}/orders`, this.order);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オーダーが正常に作成されました。', 'success');
            } catch (error) {
                console.error('オーダーの作成に失敗しました:', error);
                this.showAlert('オーダーの作成に失敗しました。', 'error');
            }
        },
        async updateOrder() {
            try {
                // INT項目の値を整数に変換
                const intFields = ['CourseMin', 'ExtraTime', 'ExtraCourse', 'Price', 'ReservationFee', 'TransportationFee', 'TravelCost'];
                intFields.forEach(field => {
                    if (this.order[field] !== undefined) {
                        this.order[field] = Math.floor(Number(this.order[field]));
                    }
                });

                if (!this.order.ActualModel) {
                    this.order.ActualModel = '';
                }

                await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}`, this.order);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オー���ーが正常に更新されました。', 'success');
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
                this.showAlert('オーダーの更新に失敗しました。', 'error');
            }
        },
        adjustValue(field, amount) {
            this.order[field] = Math.max(0, (this.order[field] || 0) + amount);
        },
        showAlert(message, color = 'info') {
            this.snackbarText = message;
            this.snackbarColor = color;
            this.showSnackbar = true;
        },
        async showCustomerDetail() {
            if (this.order.PhoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/detail/${this.order.PhoneNumber}`);
                    if (response.data && response.data.data) {
                        this.customerDetail = response.data.data;
                        this.showCustomerModal = true;
                    } else {
                        this.showAlert('この電話番号に関連する顧客情報が見つかりませんでした。', 'warning');
                    }
                } catch (error) {
                    console.error('顧客詳細の取得に失敗しました', error);
                    this.showAlert('顧��情報が見つかりません。', 'error');
                }
            } else {
                this.showAlert('電話番号が入力されていません。', 'warning');
            }
        },
        closeCustomerModal() {
            this.showCustomerModal = false;
        },
        formatDateTimeForServer(dateTimeString) {
            if (!dateTimeString) return null;
            const date = new Date(dateTimeString);
            return date.toISOString();
        },
        filterStores() {
            if (this.order.storeCode === '') {
                this.filteredStores = [...this.storeList];
            } else {
                this.filteredStores = this.storeList.filter(store =>
                    store.store_code.toString() === this.order.storeCode
                );
            }

            if (this.filteredStores.length === 1) {
                this.order.StoreID = this.filteredStores[0].id;
            } else {
                this.order.StoreID = null;
            }
        },
        async fetchCustomerInfo() {
            if (this.isNew && this.order.PhoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/phone/${this.order.PhoneNumber}`);
                    if (response.data && response.data.data) {
                        this.$nextTick(() => {
                            this.order.CustomerName = response.data.data.CustomerName || '';
                            this.order.Address = response.data.data.Address || '';
                            console.log('データが更新されました:', this.order.CustomerName, this.order.Address);
                        });
                    }
                } catch (error) {
                    console.error('顧客情報の取得に失敗しました:', error);
                }
            }
        },
        async confirmOrder() {
            this.confirmDialog = true;
        },
        async proceedWithConfirmation() {
            try {
                await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}/completion`);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オーダーが確定されました。', 'success');
            } catch (error) {
                console.error('オーダーの確定に失敗しました:', error);
                this.showAlert('オーダーの確定に失敗しました。', 'error');
            }
        },
        async confirmDelete() {
            this.deleteConfirmDialog = true;
        },
        async proceedWithDeletion() {
            try {
                await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}/delete`);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オーダーが削除されました。', 'success');
            } catch (error) {
                console.error('オーダーの削除に失敗しました:', error);
                this.showAlert('オーダーの削除に失敗しました。', 'error');
            }
        },
        filterDrivers() {
            this.filteredDriverStaffList = this.driverStaffList.filter(staff =>
                staff.name.toLowerCase().includes(this.driverFilter.toLowerCase())
            );
        },
        filterOfficeStaff() {
            this.filteredOfficeStaffList = this.officeStaffList.filter(staff =>
                staff.name.toLowerCase().includes(this.cardStaffFilter.toLowerCase()) ||
                staff.name.toLowerCase().includes(this.orderStaffFilter.toLowerCase())
            );
        },
    },
    watch: {
        driverFilter() {
            this.filterDrivers();
        },
        cardStaffFilter() {
            this.filterOfficeStaff();
        },
        orderStaffFilter() {
            this.filterOfficeStaff();
        },
    },
    async mounted() {
        await this.fetchDropdownData();
    }
};
</script>
