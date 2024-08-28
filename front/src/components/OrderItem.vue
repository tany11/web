<template>
    <v-tooltip bottom :disabled="isDragging">
        <template v-slot:activator="{ on, attrs }">
            <div class="order-item" :style="orderStyle" v-bind="attrs" v-on="on" @mousedown="startDrag"
                @mouseup="endDrag" @mousemove="drag">
                <div class="order-content">
                    <div class="order-header" :style="{ backgroundColor: getOrderColor(order) }">
                        <span class="order-time">{{ formatTime(order.start_time) }}</span>
                    </div>
                    <div class="order-body">
                        <div class="order-customer">{{ order.CustomerName }}</div>
                        <div class="order-address">{{ order.Address }}</div>
                    </div>
                </div>
            </div>
        </template>
        <OrderTooltip :order="order" />
    </v-tooltip>
</template>

<script>
import OrderTooltip from './OrderTooltip.vue';

export default {
    name: 'OrderItem',
    components: {
        OrderTooltip
    },
    props: {
        order: {
            type: Object,
            required: true,
        },
        timelineStart: {
            type: Date,
            required: true,
        },
        casts: {
            type: Array,
            required: true,
        },
    },
    data() {
        return {
            isDragging: false,
            startX: 0,
            startY: 0,
        };
    },
    computed: {
        orderStyle() {
            const startTime = this.order.start_time;
            const endTime = this.order.end_time;
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const startPercentage = ((startTime - this.timelineStart) / (24 * 60 * 60 * 1000)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = (this.order.ActualModel ? this.casts.findIndex(cast => cast.cast_id === this.order.ActualModel) : this.casts.length) * 50 + 50;

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px',
                backgroundColor: this.getOrderColor(this.order),
                border: '2px solid var(--color-border)',
                borderRadius: '4px',
                padding: '2px',
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
                cursor: 'move',
            };
        },
    },
    methods: {
        startDrag(event) {
            this.isDragging = true;
            this.startX = event.clientX;
            this.startY = event.clientY;
        },
        endDrag() {
            this.isDragging = false;
        },
        drag(event) {
            if (!this.isDragging) return;

            const deltaX = event.clientX - this.startX;
            const deltaY = event.clientY - this.startY;

            const newStartTime = this.addMinutesToTime(this.order.start_time, (deltaX / (window.innerWidth / 24)) * 60);
            const newEndTime = this.addMinutesToTime(this.order.end_time, (deltaX / (window.innerWidth / 24)) * 60);

            const updatedOrder = {
                ...this.order,
                start_time: newStartTime,
                end_time: newEndTime,
                ActualModel: Math.floor(deltaY / 50), // 50pxごとにキャストを変更
            };

            this.$emit('update', updatedOrder);
            this.startX = event.clientX;
            this.startY = event.clientY;
        },
        getOrderColor(order) {
            const colors = ['#3f51b5', '#00bcd4', '#4caf50', '#ffc107', '#e91e63', '#9c27b0', '#2196f3', '#009688'];
            return colors[order.ID % colors.length];
        },
        timeToMinutes(time) {
            const [hours, minutes] = time.split(':').map(Number);
            return hours * 60 + minutes;
        },
        addMinutesToTime(time, minutes) {
            const [hours, mins] = time.split(':').map(Number);
            const totalMinutes = hours * 60 + mins + minutes;
            const newHours = Math.floor(totalMinutes / 60) % 24;
            const newMinutes = totalMinutes % 60;
            return `${newHours.toString().padStart(2, '0')}:${newMinutes.toString().padStart(2, '0')}`;
        },
        formatTime(date) {
            return date.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
    },
};
</script>

<style scoped>
.order-item {
    position: absolute;
    z-index: 3;
    border-radius: 4px;
    overflow: hidden;
    box-shadow: 0 2px 4px rgba(0, 0, 0, 0.1);
    transition: box-shadow 0.3s ease;
}

.order-item:hover {
    box-shadow: 0 4px 8px rgba(0, 0, 0, 0.2);
}

.order-content {
    height: 100%;
    display: flex;
    flex-direction: column;
    background-color: #ffffff;
}

.order-header {
    padding: 2px 4px;
    color: #ffffff;
    font-weight: bold;
    font-size: 0.8rem;
}

.order-body {
    padding: 4px;
    display: flex;
    flex-direction: column;
    justify-content: center;
    flex-grow: 1;
}

.order-customer {
    font-weight: bold;
    font-size: 0.9rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
}

.order-address {
    font-size: 0.8rem;
    white-space: nowrap;
    overflow: hidden;
    text-overflow: ellipsis;
    color: #666;
}

@media (prefers-color-scheme: dark) {
    .order-content {
        background-color: #2c3e50;
    }

    .order-customer {
        color: #ffffff;
    }

    .order-address {
        color: #bdc3c7;
    }
}
</style>