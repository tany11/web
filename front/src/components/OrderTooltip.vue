<template>
    <v-card class="order-tooltip">
        <v-card-title>{{ order.CustomerName }}</v-card-title>
        <v-card-text>
            <v-list dense>
                <v-list-item>
                    <v-list-item-title>コース時間:</v-list-item-title>
                    <v-list-item-subtitle>{{ order.CourseMin }}分</v-list-item-subtitle>
                </v-list-item>
                <v-list-item>
                    <v-list-item-title>開始時間:</v-list-item-title>
                    <v-list-item-subtitle>{{ formatTime(order.CreatedAt) }}</v-list-item-subtitle>
                </v-list-item>
                <v-list-item>
                    <v-list-item-title>終了時間:</v-list-item-title>
                    <v-list-item-subtitle>{{ formatTime(getEndTime(order)) }}</v-list-item-subtitle>
                </v-list-item>
                <v-list-item>
                    <v-list-item-title>キャスト:</v-list-item-title>
                    <v-list-item-subtitle>{{ order.ActualModelName || '未割り当て' }}</v-list-item-subtitle>
                </v-list-item>
                <v-list-item>
                    <v-list-item-title>オプション:</v-list-item-title>
                    <v-list-item-subtitle>{{ order.OptionName || 'なし' }}</v-list-item-subtitle>
                </v-list-item>
                <v-list-item>
                    <v-list-item-title>料金:</v-list-item-title>
                    <v-list-item-subtitle>{{ order.Price }}円</v-list-item-subtitle>
                </v-list-item>
            </v-list>
        </v-card-text>
    </v-card>
</template>

<script>
export default {
    name: 'OrderTooltip',
    props: {
        order: {
            type: Object,
            required: true
        }
    },
    methods: {
        formatTime(dateString) {
            const date = new Date(dateString);
            return date.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
        getEndTime(order) {
            const startTime = new Date(order.CreatedAt);
            return new Date(startTime.getTime() + order.CourseMin * 60000);
        }
    }
}
</script>

<style scoped>
.order-tooltip {
    min-width: 250px;
}

@media screen and (max-width: 640px) {
    .order-tooltip {
        min-width: 200px;
    }
}
</style>