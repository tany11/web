<template>
    <div class="order">
        <div class="order-content">
            <!-- ダッシュボードの内容をここに追加 -->
            <h2>最近の活動</h2>
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
            </ul>
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
        }
    }
}
</script>