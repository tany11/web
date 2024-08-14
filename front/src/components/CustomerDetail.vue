<template>
    <div class="modal" v-if="customer" @click.self="closeModal">
        <div class="modal-content">
            <span class="close" @click="closeModal">&times;</span>
            <h2>顧客詳細</h2>
            <div class="customer-detail">
                <p><strong>名前：</strong>{{ customer.CustomerName }}</p>
                <p><strong>電話番号：</strong>{{ customer.PhoneNumber }}</p>
                <p><strong>住所：</strong>{{ customer.Address }}</p>
                <p><strong>メモ：</strong>{{ customer.Memo }}</p>
                <p><strong>合計金額：</strong>{{ customer.TotalPrice }}円</p>
                <p><strong>利用回数：</strong>{{ customer.TotalUseTime }}回</p>

                <h3>注文履歴</h3>
                <table v-if="customer.OrderList && customer.OrderList.length">
                    <thead>
                        <tr>
                            <th>日付</th>
                            <th>コース時間</th>
                            <th>料金</th>
                            <th>キャスト</th>
                        </tr>
                    </thead>
                    <tbody>
                        <tr v-for="order in customer.OrderList" :key="order.ID">
                            <td>{{ formatDate(order.CreatedAt) }}</td>
                            <td>{{ order.CourseMin }}分</td>
                            <td>{{ order.Price }}円</td>
                            <td>{{ getCastName(order.ActualModel) }}</td>
                        </tr>
                    </tbody>
                </table>
                <p v-else>注文履歴はありません。</p>
            </div>
        </div>
    </div>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';


export default {
    name: 'CustomerDetail',
    props: {
        customer: Object
    },
    data() {
        return {
            castList: []
        }
    },
    computed: {
        ...mapState(['apiBaseUrl'])
    },
    methods: {
        formatDate(dateString) {
            const date = new Date(dateString);
            return date.toLocaleDateString('ja-JP');
        },
        closeModal() {
            this.$emit('close');
        },
        getCastName(castId) {
            const cast = this.castList.find(c => c.cast_id === castId);
            return cast ? cast.name : 'Unknown';
        },
        async fetchCastList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.castList = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        }
    },
    mounted() {
        this.fetchCastList();
    }
}
</script>

<style scoped>
.modal {
    position: fixed;
    z-index: 1000;
    left: 0;
    top: 0;
    width: 100%;
    height: 100%;
    overflow: auto;
    background-color: rgba(0, 0, 0, 0.6);
    display: flex;
    align-items: center;
    justify-content: center;
}

.modal-content {
    background-color: #292929;
    padding: 30px;
    border: 1px solid #888;
    width: 80%;
    max-width: 600px;
    position: relative;
    border-radius: 8px;
    box-shadow: 0 4px 6px rgba(0, 0, 0, 0.1);
}

.close {
    color: #4CAF50;
    position: absolute;
    top: 10px;
    right: 20px;
    font-size: 28px;
    font-weight: bold;
    cursor: pointer;
}

.close:hover,
.close:focus {
    color: #45a049;
    text-decoration: none;
}

.customer-detail {
    margin-top: 20px;
}

table {
    width: 100%;
    border-collapse: collapse;
    margin-top: 15px;
}

th,
td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
}

th {
    background-color: #34495E;
}
</style>