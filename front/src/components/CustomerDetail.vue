<template>
    <div class="customer-detail">
        <h2>顧客詳細</h2>
        <div>
            <p><strong>名前：</strong>{{ customer.name }}</p>
            <p><strong>電話番号：</strong>{{ customer.phoneNumber }}</p>
            <p><strong>キャスト：</strong>{{ customer.castName }}</p>
            <p><strong>最終来店日：</strong>{{ customer.lastVisitDate }}</p>
            <!-- 他の詳細情報を追加 -->
        </div>
    </div>
</template>

<script>
export default {
    name: 'CustomerDetail',
    props: {
        customer: Object
    }
}
</script>

<style scoped>
.customer-detail {
    padding: 20px;
    border: 1px solid #ddd;
    border-radius: 4px;
}
</style>