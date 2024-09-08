<template>
    <div :style="orderStyle" class="order-item" draggable="true" @dragstart="onDragStart" @dragend="onDragEnd"
        @click="onClick">
        <v-card class="order-content" :color="orderColor">
            {{ memo.Title }}
        </v-card>
    </div>
</template>

<script>
export default {
    props: {
        memo: Object,
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
        getMemoStyle(memo) {
            const startTime = new Date(memo.start_time);
            const endTime = new Date(startTime.getTime() + memo.ScheduledBox * 60000);
            const duration = memo.ScheduledBox; // 分単位での期間
            const dayStart = new Date(this.dayStart);
            const dayEnd = new Date(this.dayEnd);
            const startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = (memo.castIndex * 50) + 50;

            const backgroundColor = this.getMemoColor(memo);

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px',
                backgroundColor: backgroundColor,
                border: '2px solid var(--color-border)',
                borderRadius: '4px',
                padding: '2px',
                overflow: 'hidden',
                whiteSpace: 'nowrap',
                textOverflow: 'ellipsis',
            };
        },
        getMemoColor(memo) {
            if (memo.CompletionFlg === "1") {
                return '#9E76B4';
            } else {
                return '#7F1184';
            }
        },
        onDragStart(event) {
            event.dataTransfer.setData('text/plain', this.memo.ID);
            event.dataTransfer.effectAllowed = 'move';
        },
        onDragEnd(event) {
            this.$emit('dragend', event);
        },
        onClick() {
            this.$emit('memo-clicked', this.memo);
        }
    },
    computed: {
        orderStyle() {
            return this.getMemoStyle(this.memo);
        },
        orderColor() {
            return this.getMemoColor(this.memo);
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