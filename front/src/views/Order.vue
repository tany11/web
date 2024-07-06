<template>
    <div class="order-container">
        <div class="order-form">
            <!-- ダッシュボードの内容をここに追加 -->
            <h2>受注表</h2>
            <ul>
                <li>
                    <label for="store-name">店名:</label>
                    <select id="store-name" name="store-name">
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
                    <input type="text" id="model-name" name="model-name">
                </li>
                <li>
                    <label for="actual-model">実モデル:</label>
                    <input type="text" id="actual-model" name="actual-model">
                </li>
                <li>
                    <label for="course-name">コース名:</label>
                    <input type="text" id="course-name" name="course-name">
                    <label for="course-duration">分:</label>
                </li>
                <li>
                    <label for="price">料金:</label>
                    <input type="text" id="price" name="price">
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
                    <select id="delivery" name="delivery">
                        <option value="option1">Option 1</option>
                        <option value="option2">Option 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="nomination-fee">指名料:</label>
                    <input type="text" id="nomination-fee" name="nomination-fee">
                </li>
                <li>
                    <label for="transportation-fee">交通費:</label>
                    <input type="text" id="transportation-fee" name="transportation-fee">
                </li>
                <li>
                    <label for="travel-expenses">出張費:</label>
                    <input type="text" id="travel-expenses" name="travel-expenses">
                </li>
                <li>
                    <label for="media">媒体:</label>
                    <select id="media" name="media">
                        <option value="media1">Media 1</option>
                        <option value="media2">Media 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="notes">備考:</label>
                    <select id="notes" name="notes">
                        <option value="note1">Note 1</option>
                        <option value="note2">Note 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="card-handler">カード対応者:</label>
                    <select id="card-handler" name="card-handler">
                        <option value="handler1">Handler 1</option>
                        <option value="handler2">Handler 2</option>
                        <!-- Add more options as needed -->
                    </select>
                </li>
                <li>
                    <label for="order-taker">受注者:</label>
                    <select id="order-taker" name="order-taker">
                        <option value="taker1">Taker 1</option>
                        <option value="taker2">Taker 2</option>
                        <!-- Add more options as needed -->
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
                    <h3>オーダー詳細</h3>
                    <p><strong>店名:</strong> {{ currentOrder.StoreName }}</p>
                    <p><strong>お客様名:</strong> {{ currentOrder.CustomerName }}</p>
                    <p><strong>電話番号:</strong> {{ currentOrder.phoneNumber }}</p>
                    <p><strong>モデル名:</strong> {{ currentOrder.modelName }}</p>
                    <p><strong>実モデル:</strong> {{ currentOrder.actualModel }}</p>
                    <p><strong>コース:</strong> {{ currentOrder.courseName }}</p>
                    <p><strong>料金:</strong> {{ currentOrder.price }}</p>
                    <p><strong>郵便番号:</strong> {{ currentOrder.PostalCode }}</p>
                    <p><strong>住所:</strong> {{ currentOrder.Address }}</p>
                    <p><strong>送り:</strong> {{ currentOrder.delivery }}</p>
                    <p><strong>指名料:</strong> {{ currentOrder.nominationFee }}</p>
                    <p><strong>交通費:</strong> {{ currentOrder.transportationFee }}</p>
                    <p><strong>出張費:</strong> {{ currentOrder.travelExpenses }}</p>
                    <p><strong>媒体:</strong> {{ currentOrder.media }}</p>
                    <p><strong>備考:</strong> {{ currentOrder.notes }}</p>
                    <p><strong>カード対応者:</strong> {{ currentOrder.travelExpenses }}</p>
                    <p><strong>受注者:</strong> {{ currentOrder.travelExpenses }}</p>
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
            orderTaker: '',
            storeName: '',
            modelName: '',
            actualModel: '',
            courseName: '',
            price: '',
            postalCode: '',
            delivery: '',
            nominationFee: '',
            transportationFee: '',
            travelExpenses: '',
            media: '',
            notes: '',
            cardHandler: '',
            orders: [],
            currentPage: 1,
            totalPages: 1,
            loading: true,
        }
    },
    computed: {
        currentOrder() {
            console.log('Current page:', this.currentPage);
            console.log('Orders length:', this.orders.length);
            if (!this.orders || this.orders.length === 0) return null;
            return this.orders[this.currentPage - 1] || null;
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
                    console.log('郵便番号API応答:', response.data);
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
                    courseName: this.courseName,
                    price: this.price,
                    postalCode: this.postalCode,
                    address: this.address,
                    delivery: this.delivery,
                    nominationFee: this.nominationFee,
                    transportationFee: this.transportationFee,
                    travelExpenses: this.travelExpenses,
                    media: this.media,
                    notes: this.notes,
                    cardHandler: this.cardHandler,
                    orderTaker: this.orderTaker,
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
        }
    },
    mounted() {
        this.fetchOrders();
    }
}

</script>
