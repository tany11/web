<template>
    <div class="order-container">
        <div class="order-form">
            <!-- ダッシュボードの内容をここに追加 -->
            <h2>受注表</h2>
            <ul>
                <li>
                    <label for="store-name">店名:</label>
                    <select id="store-name" name="store-name" v-model="storeName">
                        <option value="store1">Store 1</option>
                        <option value="store2">Store 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="phone-number">電話番号:</label>
                    <input type="text" id="phone-number" name="phone-number" v-model="phoneNumber"
                        @blur="fetchCustomerInfo">
                </li>
                <li>
                    <label for="customer-name">お客様名:</label>
                    <input type="text" id="customer-name" name="customer-name" v-model="customerName">
                </li>
                <li>
                    <label for="model-name">モデル名:</label>
                    <input type="text" id="model-name" name="model-name" v-model="modelName">
                </li>
                <li>
                    <label for="actual-model">実モデル:</label>
                    <select id="actual-model" name="actual-model" v-model="actualModel">
                        <option value="">選択してください</option> <!-- デフォルトオプション -->
                        <option v-for="cast in castList" :key="cast.cast_id" :value="cast.cast_id">
                            {{ cast.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="course-name">コース名:</label>
                    <input type="text" id="course-name" name="course-name" v-model="courseMin">
                    <label for="course-duration">分:</label>
                </li>
                <li>
                    <label for="price">料金:</label>
                    <input type="number" id="price" name="price" v-model.number="price" step="1000" min="0">
                </li>
                <!-- 自宅orホテルのプルダウンもほしいか -->
                <li>
                    <label for="postal-code">郵便番号:</label>
                    <input type="text" id="postal-code" name="postal-code" v-model="postalCode" @blur="fetchAddress">
                </li>
                <!-- 住所を使用してHomesとGoogleMapのURLもほしい -->
                <li>
                    <label for="address">住所:</label>
                    <input type="text" id="address" name="address" v-model="address">
                </li>
                <li>
                    <label for="delivery">送り:</label>
                    <select id="delivery" name="delivery" v-model="driverID">
                        <option v-for="staff in driverStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="nomination-fee">指名料:</label>
                    <input type="number" id="nomination-fee" name="nomination-fee" v-model.number="reservationFee"
                        step="1000" min="0">
                </li>
                <li>
                    <label for="transportation-fee">交通費:</label>
                    <input type="number" id="transportation-fee" name="transportation-fee"
                        v-model.number="transportationFee" step="1000" min="0">
                </li>
                <li>
                    <label for="travel-expenses">出張費:</label>
                    <input type="number" id="travel-expenses" name="travel-expenses" v-model.number="travelCost"
                        step="1000" min="0">
                </li>
                <li>
                    <label for="media">媒体:</label>
                    <select id="media" name="media" v-model="media">
                        <option value="media1">Media 1</option>
                        <option value="media2">Media 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="notes">備考:</label>
                    <select id="notes" name="notes" v-model="notes">
                        <option value="note1">Note 1</option>
                        <option value="note2">Note 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="card-handler">カード対応者:</label>
                    <select id="card-handler" name="card-handler" v-model="cardstaffID">
                        <option v-for="staff in officeStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="order-taker">受注者:</label>
                    <select id="order-taker" name="order-taker" v-model="orderStaffID">
                        <option v-for="staff in officeStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <button @click="submitOrder">受注</button>
                </li>
            </ul>
        </div>
        <div class="order-display">
            <div v-if="loading">データを読み込んでいます...</div>
            <div v-else-if="orders.length === 0">
                オーダーデータがありません。
            </div>
            <div v-else>
                <div class="pagination">
                    <span>{{ currentPage }}/{{ totalPages }}</span>
                    <button @click="prevPage" :disabled="currentPage === 1">前へ</button>
                    <button @click="nextPage" :disabled="currentPage === totalPages">次へ</button>
                </div>
                <div v-if="currentOrder" class="order-box">
                    <h3>オーダー明細</h3>
                    <form @submit.prevent="updateOrder">
                        <template v-for="(value, key) in displayableFields" :key="key">
                            <p v-if="!isEditing || !isEditableField(key)">
                                <strong>{{ getFieldLabel(key) }}:</strong> {{ getDisplayValue(key, value) }}
                            </p>
                            <p v-else>
                                <strong>{{ getFieldLabel(key) }}:</strong>
                                <template v-if="key === 'ActualModel'">
                                    <select v-model="editedOrder[key]">
                                        <option v-for="cast in castList" :key="cast.cast_id" :value="cast.cast_id">
                                            {{ cast.name }}
                                        </option>
                                    </select>
                                </template>
                                <template v-else-if="isPriceField(key)">
                                    <input v-model.number="editedOrder[key]" type="number" step="1000" min="0">
                                </template>
                                <input v-else-if="isEditableField(key)" v-model="editedOrder[key]"
                                    :type="getInputType(key)">
                                <span v-else>{{ getDisplayValue(key, value) }}</span>
                            </p>
                        </template>
                        <div class="button-group">
                            <button type="button" @click="toggleEdit">{{ isEditing ? '編集完了' : '編集' }}</button>
                            <button type="button" @click="confirmDelete">削除</button>
                            <button type="button" @click="confirmOrder">確定</button>
                        </div>
                    </form>
                </div>
                <div v-else>
                    選択されたオーダーがありません。
                </div>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
export default {
    name: 'Order',
    data() {
        return {
            phoneNumber: '',
            customerName: '',
            address: '',
            storeName: '',
            modelName: '',
            actualModel: '',
            courseMin: '',
            price: '',
            postalCode: '',
            transportationFee: '',
            media: '',
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
            const excludedFields = ['group_id', 'StoreID', 'CreatedAt', 'UpdatedAt', 'IsDeleted'];
            return Object.keys(this.currentOrder || {}).reduce((acc, key) => {
                if (!excludedFields.includes(key)) {
                    acc[key] = this.currentOrder[key];
                }
                return acc;
            }, {});
        }
    },
    methods: {
        async fetchCustomerInfo() {
            if (this.phoneNumber) {
                try {
                    const response = await axios.get(`http://localhost:3000/api/v1/customers/phone/${this.phoneNumber}`);
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
        async submitOrder() {
            try {
                const orderData = {
                    storeName: this.storeName,
                    phoneNumber: this.phoneNumber,
                    customerName: this.customerName,
                    modelName: this.modelName,
                    actualModel: this.actualModel,
                    courseMin: this.courseMin,
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
                };

                const response = await axios.post('http://localhost:3000/api/v1/orders', orderData);
                console.log('注文が正常に送信されました:', response.data);

                // 成功メッセージを表示
                alert('注文が正常に送信されました。');

                // フォームをリセットし、オーダー一覧を更新
                this.resetForm();
                await this.fetchOrders();

                // 最新のオーダーを表示するために currentPage を 1 にセット
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
            this.storeName = '';
            this.modelName = '';
            this.actualModel = '';
            this.courseMin = '';
            this.price = '';
            this.postalCode = '';
            this.transportationFee = '';
            this.media = '';
            this.notes = '';
            this.driverID = '';
            this.cardstaffID = '';
            this.orderStaffID = '';
            this.reservationFee = '';
            this.travelCost = '';
        },
        async fetchOrders() {
            this.loading = true;
            try {
                const response = await axios.get('http://localhost:3000/api/v1/orders');
                this.orders = response.data.data || []; // データが data プロパティ内にある場合
                this.totalPages = this.orders.length;
                this.currentPage = this.orders.length > 0 ? 1 : 0; // 最新のオーダーを表示するために1ページ目にセット
                console.log('Fetched orders:', this.orders);
                console.log('Current page after fetch:', this.currentPage);
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
            if (this.currentPage > 1) {
                this.currentPage--;
            }
        },
        nextPage() {
            if (this.currentPage < this.totalPages) {
                this.currentPage++;
            }
        },
        async fetchStaffList() {
            try {
                const response = await axios.get('http://localhost:3000/api/v1/staff/dropdown');
                this.staffList = response.data.data || [];
            } catch (error) {
                console.error('スタッフリストの取得に失敗しました:', error);
            }
        },
        async fetchCastList() {
            try {
                const response = await axios.get('http://localhost:3000/api/v1/cast/dropdown');
                this.castList = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        },
        getCastName(castId) {
            const cast = this.castList.find(c => c.cast_id === castId);
            return cast ? cast.name : 'Unknown';
        },
        getStaffName(staffId) {
            const staff = this.staffList.find(s => s.staff_id === staffId);
            return staff ? staff.name : 'Unknown';
        },
        toggleEdit() {
            if (this.isEditing) {
                this.updateOrder();
            } else {
                this.editedOrder = { ...this.currentOrder };
            }
            this.isEditing = !this.isEditing;
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

            try {
                await axios.put(`http://localhost:3000/api/v1/orders/${orderId}`, this.editedOrder);
                this.fetchOrders();
                this.isEditing = false;
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
            }
        },
        async confirmDelete() {
            if (confirm('このオーダーを削除してもよろしいですか？')) {
                try {
                    await axios.delete(`http://localhost:3000/api/v1/orders/${this.currentOrder.id}`);
                    this.fetchOrders();
                } catch (error) {
                    console.error('オーダーの削除に失敗しました:', error);
                }
            }
        },
        async confirmOrder() {
            try {
                await axios.post(`http://localhost:3000/api/v1/orders/${this.currentOrder.id}/confirm`);
                this.fetchOrders();
            } catch (error) {
                console.error('オーダーの確定に失敗しました:', error);
            }
        },
        isEditableField(key) {
            const editableFields = ['CustomerName', 'PhoneNumber', 'ModelName', 'ActualModel', 'CourseMin', 'Price', 'Address', 'ReservationFee', 'TransportationFee', 'TravelCost', 'Notes'];
            return editableFields.includes(key);
        },
        isPriceField(key) {
            return ['Price', 'ReservationFee', 'TransportationFee', 'TravelCost'].includes(key);
        },
        getInputType(key) {
            return this.isPriceField(key) ? 'number' : 'text';
        },
        getFieldLabel(key) {
            const labels = {
                StoreName: '店名',
                CustomerName: 'お客様名',
                PhoneNumber: '電話番号',
                ModelName: 'モデル名',
                ActualModel: '実モデル',
                CourseMin: 'コース',
                Price: '料金',
                PostalCode: '郵便番号',
                Address: '住所',
                DriverID: '送り',
                ReservationFee: '指名料',
                TransportationFee: '交通費',
                TravelCost: '出張費',
                Media: '媒体',
                Notes: '備考',
                CardstaffID: 'カード対応者',
                OrderStaffID: '受注者',
            };
            return labels[key] || key;
        },
        getDisplayValue(key, value) {
            if (key === 'ActualModel') {
                return this.getCastName(value);
            } else if (key === 'DriverID' || key === 'CardstaffID' || key === 'OrderStaffID') {
                return this.getStaffName(value);
            } else if (this.isPriceField(key)) {
                return `¥${value.toLocaleString()}`;
            }
            return value;
        },
    },
    mounted() {
        this.fetchOrders();
        this.fetchStaffList();
        this.fetchCastList();
    }
}

</script>

<style scoped>
.button-group {
    margin-top: 20px;
}

.button-group button {
    margin-right: 10px;
}
</style>