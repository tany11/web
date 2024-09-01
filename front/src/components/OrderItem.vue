<template>
    <div :style="orderStyle" class="order-item" draggable="true" @dragstart="onDragStart" @dragend="onDragEnd">
        <v-card class="order-content" :color="orderColor">
            {{ order.CustomerName }} - {{ formatTime(order.start_time) }}
        </v-card>
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
    data() {
        return {
            isDragging: false
        };
    },
    methods: {
        formatTime(time) {
            return new Date(time).toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
        getOrderStyle(order) {
            const startTime = new Date(order.start_time);
            const endTime = new Date(order.end_time);
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const dayStart = new Date(this.dayStart);
            const dayEnd = new Date(this.dayEnd);
            const startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = (order.castIndex * 50) + 50;

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px',
                backgroundColor: order.ActualModel ? 'var(--vt-c-indigo)' : 'var(--vt-c-divider-light-1)',
                border: '2px solid var(--color-border)',
                borderRadius: '4px',
                padding: '2px',
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
            };
        },
        getOrderColor(order) {
            const colors = ['indigo', 'cyan', 'green', 'amber', 'pink', 'purple', 'blue', 'teal'];
            return colors[order.ID % colors.length];
        },
        onDragStart(event) {
            event.dataTransfer.setData('text/plain', this.order.ID);
            event.dataTransfer.effectAllowed = 'move';
        },
        onDragEnd(event) {
            this.$emit('dragend', event);
        }
    },
    computed: {
        orderStyle() {
            return this.getOrderStyle(this.order);
        },
        orderColor() {
            return this.getOrderColor(this.order);
        }
    }
};
</script>

<style scoped>
.order-item {
    cursor: move;
    user-select: none;
    transition: opacity 0.2s, box-shadow 0.2s;
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