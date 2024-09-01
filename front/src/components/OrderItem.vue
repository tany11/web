<template>
    <div :style="orderStyle" class="order-item">
        <div class="order-content">
            {{ order.CustomerName }} - {{ formatTime(order.start_time) }}
        </div>
    </div>
</template>

<script>
export default {
    props: {
        order: Object,
        timelineStart: Date,
        casts: Array,
        dayStart: Date,
        dayEnd: Date,
    },
    methods: {
        formatTime(time) {
            return new Date(time).toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
        calculatePosition() {
            const startTime = new Date(this.order.start_time);
            const endTime = new Date(this.order.end_time);
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const dayStart = new Date(this.dayStart);
            const dayEnd = new Date(this.dayEnd);
            let startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            let widthPercentage = (duration / (24 * 60)) * 100;

            if (startTime < dayStart) {
                startPercentage = 0;
                widthPercentage = ((endTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            } else if (endTime > dayEnd) {
                widthPercentage = ((dayEnd - startTime) / (24 * 60 * 60 * 1000)) * 100;
            }
            const top = (this.order.castIndex * 50) + 50;

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px',
                backgroundColor: this.order.ActualModel ? 'var(--vt-c-indigo)' : 'var(--vt-c-divider-light-1)',
                border: '2px solid var(--color-border)',
                borderRadius: '4px',
                padding: '2px',
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
            };
        }
    },
    computed: {
        orderStyle() {
            const startTime = new Date(this.order.start_time);
            const endTime = new Date(this.order.end_time);
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const dayStart = new Date(this.dayStart);
            const dayEnd = new Date(this.dayEnd);
            const startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = (this.order.castIndex * 50) + 50;

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px',
                backgroundColor: this.order.ActualModel ? 'var(--vt-c-indigo)' : 'var(--vt-c-divider-light-1)',
                border: '2px solid var(--color-border)',
                borderRadius: '4px',
                padding: '2px',
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
            };
        }
    }
};
</script>

<style scoped>
.order-item {
    cursor: move;
}

.order-content {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    font-weight: bold;
    color: white;
}
</style>