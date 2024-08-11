<template>
    <div class="customer-list">
        <h2>検索結果</h2>
        <table v-if="customers.length > 0">
            <thead>
                <tr>
                    <th>名前</th>
                    <th>電話番号（下4桁）</th>
                    <th>キャスト</th>
                    <th>最終来店日</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="customer in customers" :key="customer.id">
                    <td><a @click="selectCustomer(customer)">{{ customer.name }}</a></td>
                    <td>{{ customer.phoneNumber.slice(-4) }}</td>
                    <td>{{ customer.castName }}</td>
                    <td>{{ customer.lastVisitDate }}</td>
                </tr>
            </tbody>
        </table>
        <p v-else>該当する顧客が見つかりません。</p>
    </div>
</template>

<script>
export default {
    name: 'CustomerList',
    props: {
        customers: Array
    },
    methods: {
        selectCustomer(customer) {
            this.$emit('select-customer', customer)
        }
    }
}
</script>

<style scoped>
.customer-list table {
    width: 100%;
    border-collapse: collapse;
}

.customer-list th,
.customer-list td {
    border: 1px solid #ddd;
    padding: 8px;
    text-align: left;
}

.customer-list tr:hover {
    background-color: #f5f5f5;
    cursor: pointer;
}
</style>