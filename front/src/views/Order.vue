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
                    <input type="text" id="price" name="price" v-model="price">
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
                    <input type="text" id="nomination-fee" name="nomination-fee" v-model="reservationFee">
                </li>
                <li>
                    <label for="transportation-fee">交通費:</label>
                    <input type="text" id="transportation-fee" name="transportation-fee" v-model="transportationFee">
                </li>
                <li>
                    <label for="travel-expenses">出張費:</label>
                    <input type="text" id="travel-expenses" name="travel-expenses" v-model="travelCost">
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
                    <p><strong>店名:</strong> {{ currentOrder.StoreName }}</p>
                    <p><strong>お客様名:</strong> {{ currentOrder.CustomerName }}</p>
                    <p><strong>電話番号:</strong> {{ currentOrder.PhoneNumber }}</p>
                    <p><strong>モデル名:</strong> {{ currentOrder.ModelName }}</p>
                    <p><strong>実モデル:</strong> {{ getCastName(currentOrder.ActualModel) }}</p>
                    <p><strong>コース:</strong> {{ currentOrder.CourseMin }}</p>
                    <p><strong>料金:</strong> {{ currentOrder.Price }}</p>
                    <p><strong>郵便番号:</strong> {{ currentOrder.PostalCode }}</p>
                    <p><strong>住所:</strong> {{ currentOrder.Address }}</p>
                    <p><strong>送り:</strong> {{ getStaffName(currentOrder.DriverID) }}</p>
                    <p><strong>指名料:</strong> {{ currentOrder.ReservationFee }}</p>
                    <p><strong>交通費:</strong> {{ currentOrder.TransportationFee }}</p>
                    <p><strong>出張費:</strong> {{ currentOrder.TravelCost }}</p>
                    <p><strong>媒体:</strong> {{ currentOrder.Media }}</p>
                    <p><strong>備考:</strong> {{ currentOrder.Notes }}</p>
                    <p><strong>カード対応者:</strong> {{ getStaffName(currentOrder.CardstaffID) }}</p>
                    <p><strong>受注者:</strong> {{ getStaffName(currentOrder.OrderStaffID) }}</p>
                    <!-- 他のオーダー詳細 -->
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
                // ここで成功メッセージを表示したり、フォームをリセットしたりできます
            } catch (error) {
                console.error('注文の送信に失敗しました:', error);
                // ここでエラーメッセージを表示できます
            }
        },
        async fetchOrders() {
            this.loading = true;
            try {
                const response = await axios.get('http://localhost:3000/api/v1/orders');
                this.orders = response.data.data || []; // データが data プロパティ内にある場合
                this.totalPages = this.orders.length;
                this.currentPage = this.orders.length > 0 ? 1 : 0;
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
        }
    },
    mounted() {
        this.fetchOrders();
        this.fetchStaffList();
        this.fetchCastList();
    }
}

</script>