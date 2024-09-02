<template>
    <v-container fluid class="order-container">
        <v-row>
            <v-col cols="12" md="6">
                <v-card class="order-form">
                    <v-card-title>受注表</v-card-title>
                    <v-card-text>
                        <v-form @submit.prevent="submitOrder">
                            <v-row>
                                <v-col cols="12" sm="6">
                                    <v-text-field v-model="storeCode" label="店舗コード"
                                        @input="filterStores"></v-text-field>
                                </v-col>
                                <v-col cols="12" sm="6">
                                    <v-select v-model="storeID" :items="filteredStores" item-title="name"
                                        item-value="id" label="店名"
                                        :error-messages="showValidation && !isStoreNameValid ? ['店名を選択してください'] : []"></v-select>
                                </v-col>
                            </v-row>

                            <v-text-field v-model="phoneNumber" label="電話番号" @blur="fetchCustomerInfo"
                                :error-messages="showValidation && !isPhoneNumberValid ? ['有効な電話番号を入力してください'] : []">
                                <template v-slot:append>
                                    <v-btn text small color="primary" @click="showCustomerDetail" v-if="phoneNumber">
                                        詳細
                                    </v-btn>
                                </template>
                            </v-text-field>

                            <v-text-field v-model="customerName" label="お客様名"></v-text-field>
                            <v-text-field v-model="modelName" label="モデル名"></v-text-field>

                            <v-select v-model="actualModel" :items="castList" item-title="name" item-value="cast_id"
                                label="実モデル"></v-select>

                            <v-row>
                                <v-col cols="5">
                                    <v-text-field v-model="courseMin" label="コース（分）" type="number" :step="10"
                                        @keydown.up.prevent="adjustCourseMin(-10)"
                                        @keydown.down.prevent="adjustCourseMin(10)">
                                    </v-text-field>
                                </v-col>
                                <v-col cols="2" class="d-flex justify-center align-center">
                                    <span class="text-h5">+</span>
                                </v-col>
                                <v-col cols="5">
                                    <v-text-field v-model="extraCourse" label="（分）" type="number" :step="5"
                                        @keydown.up.prevent="adjustExtraCourse(-5)"
                                        @keydown.down.prevent="adjustExtraCourse(5)">
                                    </v-text-field>
                                </v-col>
                            </v-row>

                            <v-text-field v-model.number="price" label="料金" type="number" prefix="¥" :step="1000"
                                @keydown.up.prevent="adjustPrice(-1000)"
                                @keydown.down.prevent="adjustPrice(1000)"></v-text-field>

                            <v-text-field v-model="postalCode" label="郵便番号" @blur="fetchAddress"></v-text-field>
                            <v-text-field v-model="address" label="住所"></v-text-field>

                            <v-select v-model="driverID" :items="driverStaffList" item-title="name"
                                item-value="staff_id" label="送り"></v-select>

                            <v-text-field v-model.number="reservationFee" label="指名料" type="number" prefix="¥"
                                :step="1000" @keydown.up.prevent="adjustReservationFee(-1000)"
                                @keydown.down.prevent="adjustReservationFee(1000)"></v-text-field>
                            <v-text-field v-model.number="transportationFee" label="交通費" type="number" prefix="¥"
                                :step="1000" @keydown.up.prevent="adjustTransportationFee(-1000)"
                                @keydown.down.prevent="adjustTransportationFee(1000)"></v-text-field>
                            <v-text-field v-model.number="travelCost" label="出張費" type="number" prefix="¥" :step="1000"
                                @keydown.up.prevent="adjustTravelCost(-1000)"
                                @keydown.down.prevent="adjustTravelCost(1000)"></v-text-field>

                            <v-select v-model="media" :items="mediaList" item-title="name" item-value="id"
                                label="媒体"></v-select>

                            <v-textarea v-model="notes" label="備考"></v-textarea>

                            <v-select v-model="cardstaffID" :items="officeStaffList" item-title="name"
                                item-value="staff_id" label="カード"></v-select>

                            <v-select v-model="orderStaffID" :items="officeStaffList" item-title="name"
                                item-value="staff_id" label="受注者"></v-select>

                            <v-text-field v-model="scheduledTime" label="到着時刻" type="datetime-local"
                                :step="900"></v-text-field>

                            <v-btn color="primary" block type="submit" :disabled="!isFormValid">受注</v-btn>
                        </v-form>
                    </v-card-text>
                </v-card>
            </v-col>

            <v-col cols="12" md="6">
                <v-card class="order-display">
                    <v-card-text>
                        <v-progress-circular v-if="loading" indeterminate color="primary"></v-progress-circular>
                        <div v-else-if="orders.length === 0">
                            オーダーデータがありません。
                            <v-btn @click="toggleShowCompleted" color="primary" text>
                                {{ showCompleted ? '未確定のみ表示' : '確定済を表示' }}
                            </v-btn>
                        </div>
                        <div v-else>
                            <v-pagination v-model="currentPage" :length="totalPages" :disabled="isEditing"
                                @input="fetchOrders(currentPage)" v-if="totalPages > 1"></v-pagination>

                            <v-card v-if="currentOrder" class="order-box">
                                <v-card-title class="order-header">
                                    オーダー明細
                                    <v-btn @click="toggleShowCompleted" color="primary" text>
                                        {{ showCompleted ? '確定済を非表示' : '確定済を表示' }}
                                    </v-btn>
                                </v-card-title>

                                <v-card-text>
                                    <v-form @submit.prevent="updateOrder">
                                        <template v-for="(value, key) in displayableFields" :key="`row-${key}`">
                                            <div class="order-detail-row">
                                                <span class="order-detail-label">{{ getFieldLabel(key) }}</span>
                                                <span class="order-detail-value">
                                                    <template v-if="!isEditing">
                                                        {{ getDisplayValue(key, value) }}
                                                    </template>
                                                    <template v-else>
                                                        <v-text-field
                                                            v-if="isEditableField(key) && !['ActualModel', 'StoreID', 'Media', 'CardstaffID', 'OrderStaffID', 'DriverID', 'ScheduledTime'].includes(key)"
                                                            v-model="editedOrder[key]" :type="getInputType(key)"
                                                            :error-messages="getErrorMessages(key)"
                                                            :step="getStepValue(key)"
                                                            @keydown.up.prevent="adjustValue(key, -getStepValue(key))"
                                                            @keydown.down.prevent="adjustValue(key, getStepValue(key))"
                                                            dense></v-text-field>
                                                        <v-select
                                                            v-else-if="['ActualModel', 'StoreID', 'Media', 'CardstaffID', 'OrderStaffID', 'DriverID'].includes(key)"
                                                            v-model="editedOrder[key]" :items="getSelectItems(key)"
                                                            item-title="text" item-value="value"
                                                            :error-messages="getErrorMessages(key)" dense></v-select>
                                                        <v-text-field v-else-if="key === 'ScheduledTime'"
                                                            v-model="editedOrder[key]" type="datetime-local"
                                                            :error-messages="getErrorMessages(key)" dense
                                                            :step="900"></v-text-field>
                                                        <v-textarea v-else-if="key === 'Notes'"
                                                            v-model="editedOrder[key]" rows="3" auto-grow
                                                            dense></v-textarea>
                                                    </template>
                                                </span>
                                            </div>
                                        </template>

                                        <v-btn-group>
                                            <v-btn v-if="!isEditing" @click="toggleEdit" color="primary">編集</v-btn>
                                            <v-btn v-if="isEditing" @click="saveEdit" color="success">保存</v-btn>
                                            <v-btn v-if="isEditing" @click="cancelEdit" color="error">キャンセル</v-btn>
                                            <v-btn @click="confirmOrder" :disabled="isEditing" color="info">確定</v-btn>
                                        </v-btn-group>
                                    </v-form>
                                </v-card-text>

                                <v-card-actions>
                                    <v-spacer></v-spacer>
                                    <v-btn @click="confirmDelete" :disabled="isEditing" color="error">削除</v-btn>
                                </v-card-actions>
                            </v-card>
                            <div v-else>
                                選択されたオーダーがありません。
                            </div>
                        </div>
                    </v-card-text>
                </v-card>
            </v-col>
        </v-row>

        <v-dialog v-model="showCustomerModal" max-width="600px">
            <CustomerDetail :customer="customerDetail" @close="closeCustomerModal" />
        </v-dialog>
    </v-container>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import { defineAsyncComponent } from 'vue';
const CustomerDetail = defineAsyncComponent(() => import('@/components/CustomerDetail.vue'));

export default {
    name: 'Order',
    components: {
        CustomerDetail
    },
    data() {
        return {
            phoneNumber: '',
            customerName: '',
            address: '',
            storeID: null,
            modelName: '',
            actualModel: '',
            courseMin: '',
            extraCourse: '',
            price: 0,
            postalCode: '',
            reservationFee: 0,
            transportationFee: 0,
            travelCost: 0,
            media: null,
            notes: '',
            orders: [],
            currentPage: 1,
            totalPages: 1,
            loading: true,
            driverID: '',
            cardstaffID: '',
            orderStaffID: '',
            staffList: [],
            castList: [],
            isEditing: false,
            editedOrder: {},
            showValidation: false,
            storeList: [],
            mediaList: [],
            showCompleted: false,
            showCustomerModal: false,
            customerDetail: null,
            storeCode: '',
            filteredStores: [],
            editValidation: false,
            scheduledTime: '',
        }
    },
    computed: {
        currentOrder() {
            if (!this.orders || this.orders.length === 0) return null;
            return this.orders[this.currentPage - 1] || null;
        },
        driverStaffList() {
            return this.staffList.filter(staff => staff.driver_flg === "1");
        },
        officeStaffList() {
            return this.staffList.filter(staff => staff.office_flg === "1");
        },
        displayableFields() {
            const excludedFields = ['GroupID', 'CustomerID', 'UpdatedAt', 'CompletionFlg', 'IsDeleted', 'ExtraCourse', 'PostalCode', 'DummyStoreFlg'];
            return this.currentOrder
                ? Object.keys(this.currentOrder).reduce((acc, key) => {
                    if (!excludedFields.includes(key)) {
                        acc[key] = this.currentOrder[key];
                    }
                    return acc;
                }, {})
                : {};
        },
        isPhoneNumberValid() {
            // 日本の電番号の簡単なバリデーション
            const phoneRegex = /^(0\d{1,4}-?\d{1,4}-?\d{4})$/;
            return phoneRegex.test(this.phoneNumber);
        },
        isStoreIDValid() {
            return this.storeID !== null && this.storeID !== '';
        },
        isEditedOrderValid() {
            return this.isEditedStoreIDValid &&
                this.isEditedCustomerNameValid &&
                this.isEditedActualModelValid &&
                this.isEditedCourseMinValid &&
                this.isEditedPriceValid &&
                this.isEditedAddressValid &&
                this.isEditedDriverIDValid &&
                this.isEditedPhoneNumberValid;
        },
        isEditedStoreIDValid() {
            return this.editedOrder.StoreID !== null && this.editedOrder.StoreID !== '';
        },
        isEditedCustomerNameValid() {
            return this.editedOrder.CustomerName && this.editedOrder.CustomerName.trim() !== '';
        },
        isEditedActualModelValid() {
            return this.editedOrder.ActualModel !== '';
        },
        isEditedCourseMinValid() {
            const courseMinInt = parseInt(this.editedOrder.CourseMin, 10);
            return !isNaN(courseMinInt) && courseMinInt > 0;
        },
        isEditedPriceValid() {
            return this.editedOrder.Price !== '' && this.editedOrder.Price > 0;
        },
        isEditedAddressValid() {
            return this.editedOrder.Address && this.editedOrder.Address.trim() !== '';
        },
        isEditedDriverIDValid() {
            return this.editedOrder.DriverID !== '';
        },
        isEditedPhoneNumberValid() {
            const phoneRegex = /^(0\d{1,4}-?\d{1,4}-?\d{4})$/;
            return phoneRegex.test(this.editedOrder.PhoneNumber);
        },
        isFormValid() {
            return this.isStoreIDValid &&
                this.isPhoneNumberValid;
        },
        ...mapState(['apiBaseUrl']),
    },
    methods: {
        async fetchCustomerInfo() {
            if (this.phoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/phone/${this.phoneNumber}`);
                    if (response.data && response.data.data) {
                        this.$nextTick(() => {
                            this.customerName = response.data.data.CustomerName || '';
                            this.address = response.data.data.Address || '';
                            console.log('データが更新されました:', this.customerName, this.address);
                        });
                    }
                } catch (error) {
                    console.error('顧客情報の取得に失敗しました:', error);
                }
            }
        },
        async fetchAddress() {
            if (this.postalCode) {
                try {
                    const response = await axios.get(`https://zipcloud.ibsnet.co.jp/api/search?zipcode=${this.postalCode}`);
                    if (response.data && response.data.results) {
                        const result = response.data.results[0];
                        this.address = `${result.address1}${result.address2}${result.address3}`;
                    } else {
                        console.log('該当する住所が見つかりませんでした');
                    }
                } catch (error) {
                    console.error('住所情報の取得に失敗しました:', error);
                }
            }
        },
        validateForm() {
            this.showValidation = true;
            if (!this.isFormValid) {
                alert('フォームに無効な入力があります。すべての必須フィールドを正しく入力してください。');
                return false;
            }
            return true;
        },
        async submitOrder() {
            if (!this.validateForm()) {
                return;
            }

            try {
                const orderData = {
                    storeID: this.storeID,
                    phoneNumber: this.phoneNumber,
                    customerName: this.customerName,
                    modelName: this.modelName,
                    actualModel: this.actualModel,
                    courseMin: parseInt(this.courseMin, 10),
                    extraCourse: parseInt(this.extraCourse, 10) || 0,
                    price: parseInt(this.price, 10),
                    postalCode: this.postalCode,
                    address: this.address,
                    driverID: this.driverID,
                    reservationFee: parseInt(this.reservationFee, 10),
                    transportationFee: parseInt(this.transportationFee, 10),
                    travelCost: parseInt(this.travelCost, 10),
                    media: this.media,
                    notes: this.notes,
                    cardstaffID: this.cardstaffID,
                    orderStaffID: this.orderStaffID,
                    scheduledTime: this.formatDateTimeForServer(this.scheduledTime),
                };

                const response = await axios.post(`${this.apiBaseUrl}/orders`, orderData);
                console.log('注文が正常に送信されました:', response.data);

                // 成功メッセージを表示
                alert('注文が正常に送信されました。');

                // フォームをリセットし、オーダー一覧を更新
                this.resetForm();
                await this.fetchOrders();

                // 最新のオーダーを表示するめに currentPage を 1 にセット
                this.currentPage = 1;
            } catch (error) {
                console.error('注文の送信に失敗しました:', error);
                alert('注文の送信に失敗しました。もう一度お試しください。');
            }
        },
        resetForm() {
            // フォームの各フィールドをリセット
            this.phoneNumber = '';
            this.customerName = '';
            this.address = '';
            this.storeID = null;
            this.modelName = '';
            this.actualModel = '';
            this.courseMin = '';
            this.extraCourse = '';
            this.price = '';
            this.postalCode = '';
            this.transportationFee = '';
            this.media = null;
            this.notes = '';
            this.driverID = '';
            this.cardstaffID = '';
            this.orderStaffID = '';
            this.reservationFee = '';
            this.travelCost = '';
            this.showValidation = false;
            this.scheduledTime = this.getInitialScheduledTime();
        },
        async fetchOrders(page = 1, limit = 10) {
            this.loading = true;
            try {
                const endpoint = this.showCompleted ? '/orders' : '/orders/reserved';
                const response = await axios.get(`${this.apiBaseUrl}${endpoint}?page=${page}&limit=${limit}`);
                this.orders = response.data.data || [];
                this.totalPages = Math.ceil(this.orders.length); // 修正: totalPagesの計算をordersの長さで行う
                this.currentPage = page;
                console.log('Fetched orders:', this.orders);
                console.log('Current page after fetch:', this.currentPage);
                console.log('Total pages after fetch:', this.totalPages);
            } catch (error) {
                console.error('オーダーの取得に失敗しました:', error);
                this.orders = [];
                this.totalPages = 0;
                this.currentPage = 0;
            } finally {
                this.loading = false;
            }
        },
        prevPage() {
            if (this.currentPage > 1 && !this.isEditing) {
                this.currentPage--;
            }
        },
        nextPage() {
            if (this.currentPage < this.totalPages && !this.isEditing) {
                this.currentPage++;
            }
        },

        //プルダウンリストを取得
        async fetchStaffList() {
            if (this.staffList.length > 0) return; // キャッシュがある場合はスキップ
            try {
                const response = await axios.get(`${this.apiBaseUrl}/staff/dropdown`);
                this.staffList = response.data.data || [];
                console.log('スタッフリスト:', this.staffList);
            } catch (error) {
                console.error('タッフリの取得に失敗しました:', error);
            }
        },
        async fetchCastList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.castList = response.data.data || [];
                console.log('キャストリスト:', this.castList);
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        },
        async fetchStoreList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/store/dropdown`);
                this.storeList = response.data.data || [];
                this.filteredStores = [...this.storeList]; // 初期状態では全ての店舗を表示
                console.log('店舗リスト:', this.storeList);
            } catch (error) {
                console.error('店舗リストの取得に失敗しました:', error);
            }
        },
        async fetchMediaList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/media/dropdown`);
                this.mediaList = response.data.data || [];
                console.log('媒体リスト:', this.mediaList);
            } catch (error) {
                console.error('媒体リスト取得に失敗また:', error);
            }
        },

        // 明細で使用するリストを取得
        getCastName(castId) {
            const cast = this.castList.find(c => c.cast_id === castId);
            return cast ? cast.name : 'Unknown';
        },
        getStaffName(staffId) {
            const staff = this.staffList.find(s => s.staff_id === staffId);
            return staff ? staff.name : 'Unknown';
        },
        getStoreName(storeId) {
            const store = this.storeList.find(s => s.id === storeId);
            return store ? store.name : 'Unknown';
        },
        getMediaName(mediaId) {
            const media = this.mediaList.find(m => m.id === mediaId);
            return media ? media.name : 'Unknown';
        },
        toggleEdit() {
            if (!this.isEditing) {
                this.editedOrder = { ...this.currentOrder };
            }
            this.isEditing = !this.isEditing;
        },
        saveEdit() {
            this.editValidation = true;
            if (this.isEditedOrderValid) {
                this.updateOrder();
            } else {
                alert('フォームに無効な入力があります。すべての必須フィールドを正しく入力してください。');
            }
        },
        cancelEdit() {
            this.editedOrder = { ...this.currentOrder };
            this.isEditing = false;
            this.editValidation = false;
        },
        async updateOrder() {
            if (!this.currentOrder) {
                console.error('現在のオーダーが見つかりません');
                console.error(this.currentOrder);
                return;
            }

            const orderId = this.currentOrder.ID;
            if (!orderId) {
                console.error('現在のオーダーIDが見つかりません');
                console.error(orderId);
                return;
            }

            // INT項目の値を整数に変換
            const intFields = ['CourseMin', 'ExtraTime', 'ExtraCourse', 'Price', 'ReservationFee', 'TransportationFee', 'TravelCost'];
            intFields.forEach(field => {
                if (this.editedOrder[field] !== undefined) {
                    this.editedOrder[field] = Math.floor(Number(this.editedOrder[field]));
                }
            });

            // 日時フィールドをフォーマット
            if (this.editedOrder.ScheduledTime) {
                this.editedOrder.ScheduledTime = this.formatDateTimeForServer(this.editedOrder.ScheduledTime);
            }

            try {
                await axios.put(`${this.apiBaseUrl}/orders/${orderId}`, this.editedOrder);
                this.fetchOrders();
                this.isEditing = false;
            } catch (error) {
                console.error('オーダーの新に失敗しました:', error);
            }
        },
        async confirmDelete() {
            if (this.isEditing) return; // 編集中は処理を実行しない

            if (confirm('このオーダーを削除してもよろしいですか？')) {
                try {
                    await axios.put(`${this.apiBaseUrl}/orders/${this.currentOrder.ID}/delete`);
                    this.fetchOrders();
                } catch (error) {
                    console.error('オーダーの削除に失敗しました:', error);
                }
            }
        },
        async confirmOrder() {
            if (this.isEditing) return; // 編集中は処理を実行しない

            if (confirm('このオーダーを確定してもよろしいですか？')) {
                try {
                    await axios.put(`${this.apiBaseUrl}/orders/${this.currentOrder.ID}/completion`);
                    this.fetchOrders();
                } catch (error) {
                    console.error('オーダーの確定に失敗しまた:', error);
                }
            }
        },
        isEditableField(key) {
            const editableFields = [
                'CustomerName',
                'PhoneNumber',
                'ModelName',
                'ActualModel',
                'CourseMin',
                'ExtraTime',
                'Price',
                'Address',
                'ReservationFee',
                'TransportationFee',
                'TravelCost',
                'Notes',
                'CardstaffID',
                'DriverID',
                'StoreID',
                'Media',
                'OrderStaffID'
            ];
            return editableFields.includes(key);
        },
        isPriceField(key) {
            return ['Price', 'ReservationFee', 'TransportationFee', 'TravelCost'].includes(key);
        },
        isIntField(key) {
            return ['CourseMin', 'ExtraTime', 'ExtraCourse'].includes(key);
        },
        getStepValue(key) {
            if (this.isPriceField(key)) {
                return 1000;
            } else if (key === 'CourseMin') {
                return 10;
            } else if (key === 'ExtraTime' || key === 'ExtraCourse') {
                return 30;
            }
            return 1;
        },
        validateIntInput(key) {
            if (this.isIntField(key)) {
                this.editedOrder[key] = Math.floor(this.editedOrder[key]);
            }
        },
        getInputType(key) {
            if (this.isPriceField(key)) {
                return 'number';
            } else if (key === 'ScheduledTime') {
                return 'datetime-local';
            }
            return 'text';
        },
        getSelectItems(key) {
            if (key === 'ActualModel') {
                return this.castList.map(cast => ({ text: cast.name, value: cast.cast_id }));
            } else if (key === 'StoreID') {
                return this.storeList.map(store => ({ text: store.name, value: store.id }));
            } else if (key === 'Media') {
                return this.mediaList.map(media => ({ text: media.name, value: media.id }));
            } else if (key === 'CardstaffID' || key === 'OrderStaffID' || key === 'DriverID') {
                return this.staffList.filter(staff => staff.office_flg === "1").map(staff => ({ text: staff.name, value: staff.staff_id }));
            }
            return [];
        },
        getFieldLabel(key) {
            const labels = {
                StoreID: '店名　　',
                CustomerName: 'お客様名',
                PhoneNumber: '電話番号',
                ModelName: 'モデル名',
                ActualModel: '実モデル',
                CourseMin: 'コース　',
                ExtraTime: '延長時間',
                Price: '料金　　',
                Address: '住所　　',
                DriverID: '送り　　',
                ReservationFee: '指名料　',
                TransportationFee: '交通費　',
                TravelCost: '出張費　',
                Media: '媒体　　',
                Notes: '備考　　',
                CardstaffID: 'カード　',
                OrderStaffID: '受注者　',
                CreatedAt: '受注日　',
                ScheduledTime: '予定時間',
            };
            return labels[key] || key;
        },
        formatDate(dateString) {
            if (!dateString) return '';
            const date = new Date(dateString);
            const japanTime = new Date(date.toLocaleString("en-US", { timeZone: "Asia/Tokyo" }));

            const year = String(japanTime.getFullYear()).slice(-2);
            const month = String(japanTime.getMonth() + 1).padStart(2, '0');
            const day = String(japanTime.getDate()).padStart(2, '0');
            const hours = String(japanTime.getHours()).padStart(2, '0');
            const minutes = String(japanTime.getMinutes()).padStart(2, '0');

            return `${year}/${month}/${day} ${hours}:${minutes}`;
        },
        getDisplayValue(key, value) {
            if (value === null || value === undefined) return '';

            if (key === 'ActualModel') {
                return this.getCastName(value);
            } else if (key === 'DriverID' || key === 'CardstaffID' || key === 'OrderStaffID') {
                return this.getStaffName(value);
            } else if (key === 'StoreID') {
                return this.getStoreName(value);
            } else if (key === 'Media') {
                return this.getMediaName(value);
            } else if (this.isPriceField(key)) {
                return `¥${value.toLocaleString()}`;
            } else if (key === 'Notes') {
                return value || '';
            } else if (key === 'CreatedAt' || key === 'ScheduledTime') {
                return this.formatDate(value);
            } else if (key === 'CourseMin') {
                const extraCourse = this.currentOrder.ExtraCourse || 0;
                return `${value}分 + ${extraCourse}分`;
            } else if (key === 'ExtraTime') {
                return `${value}分`;
            } else if (key === 'PhoneNumber') {
                return `184${value}`;
            }
            return value;
        },
        getErrorMessages(key) {
            if (this.editValidation) {
                if (key === 'StoreID' && !this.isEditedStoreIDValid) {
                    return ['店舗を選択してください'];
                } else if (key === 'CustomerName' && !this.isEditedCustomerNameValid) {
                    return ['お客様名を入力してください'];
                } else if (key === 'ActualModel' && !this.isEditedActualModelValid) {
                    return ['実モデルを選択してください'];
                } else if (key === 'CourseMin' && !this.isEditedCourseMinValid) {
                    return ['有効なコース時間を入力してください'];
                } else if (key === 'Price' && !this.isEditedPriceValid) {
                    return ['有効な料金を入力してください'];
                } else if (key === 'Address' && !this.isEditedAddressValid) {
                    return ['住所を入力してください'];
                } else if (key === 'DriverID' && !this.isEditedDriverIDValid) {
                    return ['ドライバを選択してください'];
                } else if (key === 'PhoneNumber' && !this.isEditedPhoneNumberValid) {
                    return ['有効な電話番号を入力してください'];
                }
            }
            return [];
        },
        async toggleShowCompleted() {
            this.showCompleted = !this.showCompleted;
            await this.fetchOrders();
            if (this.orders.length === 0) {
                this.currentPage = 0;
                this.totalPages = 0;
            } else {
                this.currentPage = 1;
                this.totalPages = this.orders.length;
            }
        },
        async showCustomerDetail() {
            if (this.phoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/detail/${this.phoneNumber}`);
                    if (response.data && response.data.data) {
                        this.customerDetail = response.data.data;
                        this.showCustomerModal = true;
                    }
                } catch (error) {
                    console.error('顧客詳細の取得に失敗しました:', error);
                }
            }
        },
        closeCustomerModal() {
            this.showCustomerModal = false;
        },
        filterStores() {
            if (this.storeCode === '') {
                this.filteredStores = [...this.storeList];
            } else {
                this.filteredStores = this.storeList.filter(store =>
                    store.store_code.toString().startsWith(this.storeCode)
                );
            }

            // フィルタリング後に店舗が1つだけ残った場合、自動的にその店舗を選択
            if (this.filteredStores.length === 1) {
                this.storeID = this.filteredStores[0].id;
            } else {
                this.storeID = null; // 複数の選択肢がある場合はクリア
            }
        },
        adjustPrice(amount) {
            this.price = Math.max(0, (this.price || 0) + amount);
        },
        adjustReservationFee(amount) {
            this.reservationFee = Math.max(0, (this.reservationFee || 0) + amount);
        },
        adjustTransportationFee(amount) {
            this.transportationFee = Math.max(0, (this.transportationFee || 0) + amount);
        },
        adjustTravelCost(amount) {
            this.travelCost = Math.max(0, (this.travelCost || 0) + amount);
        },
        adjustValue(key, amount) {
            if (this.isEditableField(key)) {
                const currentValue = parseInt(this.editedOrder[key], 10) || 0;
                this.editedOrder[key] = Math.max(0, currentValue + amount);
            }
        },
        getInitialScheduledTime() {
            const now = new Date(new Date().toLocaleString("en-US", { timeZone: "Asia/Tokyo" }));
            now.setMinutes(Math.ceil(now.getMinutes() / 15) * 15);
            now.setSeconds(0);
            now.setMilliseconds(0);
            return now.toISOString().slice(0, 16);
        },
        formatDateTimeForServer(dateTimeString) {
            if (!dateTimeString) return null;
            const date = new Date(dateTimeString);
            return date.toISOString();
        },
        adjustCourseMin(amount) {
            this.courseMin = Math.max(0, (this.courseMin || 0) + amount);
        },
        adjustExtraCourse(amount) {
            this.extraCourse = Math.max(0, (this.extraCourse || 0) + amount);
        },
    },
    async mounted() {
        this.loading = true;
        await this.fetchOrders();
        this.loading = false;

        // 他のデータ取得を非同期で行う
        await Promise.all([
            this.fetchStaffList(),
            this.fetchCastList(),
            this.fetchStoreList(),
            this.fetchMediaList()
        ]);
    }
}
</script>

<style scoped>
.order-detail-row {
    display: flex;
    align-items: center;
    margin-bottom: 8px;
}

.order-detail-label {
    flex: 0 0 100px;
    font-weight: bold;
}

.order-detail-value {
    flex: 1;
    padding-left: 8px;
}
</style>