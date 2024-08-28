<template>
    <v-container class="customer-list">
        <v-row>
            <v-col>
                <h2 class="text-h5 mb-4">検索結果</h2>
            </v-col>
        </v-row>
        <v-row v-if="customers.length > 0">
            <v-col>
                <v-data-table :headers="headers" :items="customers" :items-per-page="10" class="elevation-1">
                    <template v-slot:item.CustomerName="{ item }">
                        <v-btn text color="primary" @click="selectCustomer(item)">
                            {{ item.CustomerName }}
                        </v-btn>
                    </template>
                    <template v-slot:item.UpdatedAt="{ item }">
                        {{ formatDate(item.UpdatedAt) }}
                    </template>
                </v-data-table>
            </v-col>
        </v-row>
        <v-row v-else>
            <v-col>
                <v-alert type="info">
                    該当する顧客が見つかりません。
                </v-alert>
            </v-col>
        </v-row>
    </v-container>
</template>

<script>
export default {
    name: 'CustomerList',
    props: {
        customers: Array
    },
    data() {
        return {
            headers: [
                { text: '名前', value: 'CustomerName' },
                { text: '電話番号', value: 'PhoneNumber' },
                { text: '住所', value: 'Address' },
                { text: '最終利用日', value: 'UpdatedAt' }
            ]
        }
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