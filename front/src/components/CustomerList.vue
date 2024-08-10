<template>
    <div class="customer-list">
        <h2>検索結果</h2>
        <table>
            <thead>
                <tr>
                    <th>名前</th>
                    <th>電話番号</th>
                    <th>キャスト</th>
                    <th>最終来店日</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="customer in customers" :key="customer.id" @click="selectCustomer(customer)">
                    <td>{{ customer.name }}</td>
                    <td>{{ customer.phoneNumber }}</td>
                    <td>{{ customer.castName }}</td>
                    <td>{{ customer.lastVisitDate }}</td>
                </tr>
            </tbody>
        </table>
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