<template>
    <v-card v-if="order">
        <v-card-title>{{ isNew ? '新規オーダー' : 'オーダー詳細' }}</v-card-title>
        <v-card-text>
            <v-form @submit.prevent="submitForm" ref="form" v-model="formIsValid">

                <v-row v-if="isNew">
                    <v-col cols="12" sm="6">
                        <v-text-field v-model="order.storeCode" :label="requiredLabel('店舗コード')" @input="filterStores"
                            :rules="[v => !!v || '店舗コードは必須です']"></v-text-field>
                    </v-col>
                    <v-col cols="12" sm="6">
                        <v-select v-model="order.StoreID" :items="filteredStores" item-title="name" item-value="id"
                            :label="requiredLabel('店名')" :rules="[v => !!v || '店名は必須です']"></v-select>
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

                <v-text-field v-model="order.PhoneNumber" :label="requiredLabel('電話番号')"
                    @blur="isNew ? fetchCustomerInfo() : null" :rules="[
                        v => !!v || '電話番号は必須です',
                        v => /^(0\d{1,4}-?\d{1,4}-?\d{4})$/.test(v) || '正しい電話番号の形式で入力してください'
                    ]">
                    <template v-slot:append>
                        <v-btn text small color="primary" @click="showCustomerDetail" v-if="order.PhoneNumber">
                            詳細
                        </v-btn>
                    </template>
                </v-text-field>

                <v-row>
                    <v-col cols="12">
                        <v-radio-group v-model="order.UsageType" inline :rules="[v => !!v || '利用タイプは必須です']" required
                            :label="requiredLabel('利用タイプ')" :class="{ 'required-field': isNew }">
                            <v-radio v-for="usage in usageTypes" :key="usage.ID" :label="usage.DisplayName"
                                :value="usage.ClassificationCode" class="mr-4"></v-radio>
                        </v-radio-group>
                    </v-col>
                </v-row>
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
                    @keydown.down.prevent="adjustValue('Price', 1000)" @keydown.up.prevent="adjustValue('Price', -1000)"
                    @keydown.tab="handleTabPress"></v-text-field>
                <v-text-field v-model="order.City" label="市区町村" lang="ja" v-ime-on ref="cityInput"></v-text-field>
                <v-text-field v-model="order.Address" label="住所" lang="ja"></v-text-field>

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
                    <v-col cols="4">
                        <v-btn color="primary" block type="submit" :disabled="!formIsValid">{{ isNew ? '登録' : '更新'
                            }}</v-btn>
                    </v-col>
                    <v-col cols="4" v-if="!isNew">
                        <v-btn color="success" block @click="showOrderText">テキスト表示</v-btn>
                    </v-col>
                    <v-col cols="4" v-if="!isNew">
                        <v-btn color="error" block @click="confirmDelete">削除</v-btn>
                    </v-col>
                </v-row>
            </v-form>
        </v-card-text>
        <v-card-actions>
            <v-btn @click="closeDialog">閉じる</v-btn>
        </v-card-actions>
    </v-card>
    <v-card v-else>
        <v-card-text>
            オダー情報が見つかりません。
        </v-card-text>
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
            <v-card-text>このオーダーを確定してよろしいですか？</v-card-text>
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
            <v-card-text>このオーダーを削除してもよろしいですか？</v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="deleteConfirmDialog = false; proceedWithDeletion()">はい</v-btn>
                <v-btn color="grey" text @click="deleteConfirmDialog = false">いいえ</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>

    <OrderTextDisplay ref="orderTextDisplay" :orderDetails="formatOrderDetails()" />
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import CustomerDetail from './CustomerDetail.vue';
import { useWebSocket } from '@/utils/websocket';
import OrderTextDisplay from './OrderTextDisplay.vue';

export default {
    props: {
        order: {
            type: Object,
            default: () => ({})
        },
        isNew: {
            type: Boolean,
            default: false
        }
    },
    emits: ['close', 'order-updated'],
    components: {
        CustomerDetail,
        OrderTextDisplay
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
            usageTypes: [],
            formIsValid: false,
            isComposing: false
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
        },
        formattedDateAndCount() {
            const now = new Date();
            const date = this.getAdjustedDate(now);
            const dayOfWeek = ['日', '月', '火', '水', '木', '金', '土'][date.getDay()];
            const month = date.getMonth() + 1;
            const day = date.getDate();
            // 本数のカウントロジックは実際のデータ基づいて実装する必要があます

            return `${month}月${day}日(${dayOfWeek}) 本数：`;
        }
    },
    methods: {
        async fetchDropdownData() {
            try {
                const [stores, casts, staff, media, usage] = await Promise.all([
                    axios.get(`${this.apiBaseUrl}/store/dropdown`),
                    axios.get(`${this.apiBaseUrl}/cast/dropdown`),
                    axios.get(`${this.apiBaseUrl}/staff/dropdown`),
                    axios.get(`${this.apiBaseUrl}/media/dropdown`),
                    axios.get(`${this.apiBaseUrl}/master/usage`)
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

                this.usageTypes = usage.data.data || [];
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
            if (!this.$refs.form.validate()) {
                this.showAlert('必須項目を入力してください。', 'error');
                return;
            }

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

                // WebSocket経由で更新を送信
                if (this.wsIsConnected && this.wsSocket) {
                    this.wsSocket.send(JSON.stringify({
                        type: 'order_update',
                        order: response.data
                    }));
                }
            } catch (error) {
                console.error('オーダーの作成に失敗しました:', error);
                this.showAlert('オーダーの作成に失敗しました。', 'error');
            }
        },
        async updateOrder() {
            if (!this.$refs.form.validate()) {
                this.showAlert('必須項目を入力してください。', 'error');
                return;
            }

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

                const response = await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}`, this.order);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オーダーが正常に更新されました。', 'success');

                // WebSocket経由で更新を送信
                if (this.wsIsConnected && this.wsSocket) {
                    this.wsSocket.send(JSON.stringify({
                        type: 'order_update',
                        order: response.data
                    }));
                }
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
                    this.showAlert('顧客情報が見つかりません。', 'error');
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
            if (this.isNew && this.order.PhoneNumber && this.validatePhoneNumber(this.order.PhoneNumber)) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/phone/${this.order.PhoneNumber}`);
                    if (response.data && response.data.data) {
                        this.$nextTick(() => {
                            this.order.CustomerName = response.data.data.CustomerName || '';
                            if (this.order.UsageType === "1") {
                                this.order.City = response.data.data.City1 || '';
                                this.order.Address = response.data.data.Address1 || '';
                            } else {
                                this.order.City = response.data.data.City3 || '';
                                this.order.Address = response.data.data.Address3 || '';
                            }
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
                const response = await axios.put(`${this.apiBaseUrl}/orders/${this.order.ID}/completion`);
                this.$emit('order-updated');
                this.$emit('close');
                this.showAlert('オーダーが確定されました。', 'success');

                // WebSocket経由で更新を送信
                if (this.wsIsConnected && this.wsSocket) {
                    this.wsSocket.send(JSON.stringify({
                        type: 'order_update',
                        order: response.data
                    }));
                }
            } catch (error) {
                console.error('オーダーの確定に失敗しました:', error);
                this.showAlert('オーダーの確定に敗しました。', 'error');
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

                // WebSocket経由で更新を送信
                if (this.wsIsConnected && this.wsSocket) {
                    this.wsSocket.send(JSON.stringify({
                        type: 'order_delete',
                        orderId: this.order.ID
                    }));
                }
            } catch (error) {
                console.error('オーダーの削除に失敗しました:', error);
                this.showAlert('オーダーの削除に敗しました。', 'error');
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
        showOrderText() {
            this.$refs.orderTextDisplay.show();
        },
        formatOrderDetails() {
            if (!this.order) {
                return ''; // orderがnullの場合は空文字列を返す
            }
            const details = [
                this.formattedDateAndCount,
                `店名: ${this.getStoreName(this.order.StoreID)}`,
                `電話番号: ${this.order.PhoneNumber || ''}`,
                `お客様名: ${this.order.CustomerName || ''}`,
                `モデル名: ${this.order.ModelName || ''}(${this.getCastName(this.order.ActualModel)})`,
                `コース: ${this.order.CourseMin || ''}+${this.order.ExtraCourse || ''}分`,
                `料金: ¥${this.order.Price || ''}`,
                `市区町村: ${this.order.City || ''}`,
                `住所: ${this.order.Address || ''}`,
                `送り: ${this.getDriverName(this.order.DriverID)}`,
                `指名料: ¥${this.order.ReservationFee || ''}`,
                `交通費: ¥${this.order.TransportationFee || ''}`,
                `出張費: ¥${this.order.TravelCost || ''}`,
                `媒体: ${this.getMediaName(this.order.Media)}`,
                `備考: ${this.order.Notes || ''}`,
                `カード対応者: ${this.getCardStaffName(this.order.CardStaffID)}`,
                `受注者: ${this.getOrderStaffName(this.order.OrderStaffID)}`,
                `受付時間: ${this.formatCreatedAtTime(this.order.CreatedAt)}`
            ];
            return details.join('\n');
        },
        formatCreatedAtTime(createdAt) {
            if (!createdAt) return '';
            const date = new Date(createdAt);
            const hours = date.getHours().toString().padStart(2, '0');
            const minutes = date.getMinutes().toString().padStart(2, '0');
            return `${hours}:${minutes}`;
        },
        getStoreName(id) {
            const store = this.storeList.find(s => s.id === id);
            return store ? store.name : '';
        },
        getCastName(id) {
            const cast = this.castList.find(c => c.cast_id === id);
            return cast ? cast.name : '';
        },
        getDriverName(id) {
            const driver = this.driverStaffList.find(d => d.staff_id === id);
            return driver ? driver.name : '';
        },
        getMediaName(id) {
            const media = this.mediaList.find(m => m.id === id);
            return media ? media.name : '';
        },
        getCardStaffName(id) {
            const cardStaff = this.officeStaffList.find(d => d.staff_id === id);
            return cardStaff ? cardStaff.name : '';
        },
        getOrderStaffName(id) {
            const orderStaff = this.officeStaffList.find(d => d.staff_id === id);
            return orderStaff ? orderStaff.name : '';
        },
        getAdjustedDate(date) {
            const hours = date.getHours();
            const minutes = date.getMinutes();

            if (hours < 6 || (hours === 5 && minutes <= 59)) {
                // 6:00未満の場合、前日の日付を返す
                date.setDate(date.getDate() - 1);
            }
            return date;
        },
        closeDialog() {
            this.$emit('close');
        },
        requiredLabel(label) {
            return this.isNew ? `${label} *` : label;
        },
        validatePhoneNumber(phoneNumber) {
            const phoneRegex = /^(0\d{1,4}-?\d{1,4}-?\d{4})$/;
            return phoneRegex.test(phoneNumber);
        },
        handleTabPress(event) {
            if (this.isComposing) {
                event.preventDefault();
                // IME 確定後にフォーカスを移動
                this.$nextTick(() => {
                    this.$refs.cityInput.$el.querySelector('input').focus();
                });
            }
        }
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
        // フォームの有効性を監視
        '$refs.form': {
            handler(form) {
                if (form) {
                    this.formIsValid = form.validate();
                }
            },
            deep: true
        }
    },
    directives: {
        imeOn: {
            inserted: function (el) {
                el.querySelector('input').style.imeMode = 'active';
            }
        }
    },
    setup() {
        const { socket, isConnected } = useWebSocket();
        return { wsSocket: socket, wsIsConnected: isConnected };
    },
    async mounted() {
        await this.fetchDropdownData();

        document.addEventListener('compositionstart', () => {
            this.isComposing = true;
        });

        document.addEventListener('compositionend', () => {
            this.isComposing = false;
        });
    },
    beforeDestroy() {
        document.removeEventListener('compositionstart', this.handleCompositionStart);
        document.removeEventListener('compositionend', this.handleCompositionEnd);
    }
};
</script>

<style scoped>
.required-label::after {
    content: " *";
    color: red;
}

.v-text-field input {
    ime-mode: active;
}
</style>
