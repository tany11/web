<template>
    <div class="customer-detail">
        <h2>顧客詳細</h2>
        <div v-if="customer">
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
                        <td>{{ order.ActualModel }}</td>
                    </tr>
                </tbody>
            </table>
            <p v-else>注文履歴はありません。</p>
        </div>
        <div v-else>
            <p>顧客が見つかりません。</p>
        </div>
    </div>
</template>

<script>
export default {
    name: 'CustomerDetail',
    props: {
        customer: Object
    },
    methods: {
        formatDate(dateString) {
            const date = new Date(dateString);
            return date.toLocaleDateString('ja-JP');
        }
    }
}
</script>

<style scoped>
.customer-detail {
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 4px;
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