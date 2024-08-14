<template>
    <div class="customer-list">
        <h2>検索結果</h2>
        <table v-if="customers.length > 0">
            <thead>
                <tr>
                    <th>名前</th>
                    <th>電話番号</th>
                    <th>住所</th>
                    <th>最終利用日</th>
                </tr>
            </thead>
            <tbody>
                <tr v-for="customer in customers" :key="customer.ID">
                    <td><a @click="selectCustomer(customer)" class="customer-link">{{ customer.CustomerName }}</a></td>
                    <td>{{ customer.PhoneNumber }}</td>
                    <td>{{ customer.Address }}</td>
                    <td>{{ formatDate(customer.UpdatedAt) }}</td>
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
        },
        formatDate(dateString) {
            if (!dateString) return '-'
            const date = new Date(dateString)
            return date.toLocaleDateString('ja-JP')
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

.customer-link {
    color: #007bff;
    cursor: pointer;
    text-decoration: underline;
}

.customer-link:hover {
    color: #0056b3;
}
</style>