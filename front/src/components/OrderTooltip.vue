<template>
    <div class="order-tooltip">
        <h3>{{ order.CustomerName }}</h3>
        <p>コース時間: {{ order.CourseMin }}分</p>
        <p>開始時間: {{ formatTime(order.CreatedAt) }}</p>
        <p>終了時間: {{ formatTime(getEndTime(order)) }}</p>
        <p>キャスト: {{ order.ActualModelName || '未割り当て' }}</p>
        <p>オプション: {{ order.OptionName || 'なし' }}</p>
        <p>料金: {{ order.Price }}円</p>
    </div>
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
    background-color: var(--color-background);
    color: var(--color-text);
    border: 1px solid var(--color-border);
    border-radius: 8px;
    padding: 15px;
    box-shadow: 0 4px 12px rgba(0, 0, 0, 0.1);
    min-width: 250px;
    font-size: 14px;
}

.order-tooltip h3 {
    color: var(--color-heading);
    margin-bottom: 10px;
    font-size: 18px;
}

.order-tooltip p {
    margin: 5px 0;
}

@media (prefers-color-scheme: dark) {
    .order-tooltip {
        background-color: var(--color-background-soft);
        border-color: var(--color-border-hover);
    }
}

@media screen and (max-width: 640px) {
    .order-tooltip {
        font-size: 12px;
        padding: 10px;
        min-width: 200px;
    }

    .order-tooltip h3 {
        font-size: 16px;
    }
}
</style>