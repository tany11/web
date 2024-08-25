<template>
    <div class="time-board">
        <h2>タイムボード</h2>
        <div class="board-content">
            <div class="cast-names">
                <div class="cast-name-header">キャスト</div>
                <div v-for="cast in casts" :key="cast.cast_id" class="cast-name">
                    {{ cast.name }}
                </div>
                <div class="cast-name">フリー枠</div>
            </div>
            <div class="timeline-container">
                <div class="time-scale">
                    <div v-for="hour in 24" :key="hour" class="hour-marker">
                        {{ hour - 1 }}:00
                    </div>
                </div>
                <div class="cast-rows">
                    <div v-for="cast in casts" :key="cast.cast_id" class="cast-row">
                        <div class="order-timeline">
                            <div v-for="order in getOrdersForCast(cast.cast_id)" :key="order.ID" class="order-item"
                                :style="getOrderStyle(order)" @mouseover="showTooltip(order, $event)"
                                @mouseleave="hideTooltip">
                                {{ order.Address }}
                            </div>
                        </div>
                    </div>
                    <div class="cast-row">
                        <div class="order-timeline">
                            <div v-for="order in getFreeOrders()" :key="order.ID" class="order-item free-order"
                                :style="getOrderStyle(order)" @mouseover="showTooltip(order, $event)"
                                @mouseleave="hideTooltip">
                                {{ order.Address }}
                            </div>
                        </div>
                    </div>
                </div>
            </div>
        </div>
        <OrderTooltip v-if="selectedOrder" :order="selectedOrder" :style="tooltipStyle" />
    </div>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import OrderTooltip from '@/components/OrderTooltip.vue';

export default {
    name: 'TimeBoard',
    components: {
        OrderTooltip
    },
    data() {
        return {
            casts: [],
            orders: [],
            selectedOrder: null,
            tooltipStyle: {
                position: 'fixed',
                display: 'none',
                zIndex: 1000
            }
        };
    },
    computed: {
        ...mapState(['apiBaseUrl']),
    },
    methods: {
        async fetchCasts() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.casts = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        },
        async fetchOrders() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/orders`);
                this.orders = response.data.data || [];
            } catch (error) {
                console.error('オーダーの取得に失敗しました:', error);
            }
        },
        getOrdersForCast(castId) {
            return this.orders.filter(order => order.ActualModel === castId);
        },
        getFreeOrders() {
            return this.orders.filter(order => !this.casts.some(cast => cast.cast_id === order.ActualModel));
        },
        getOrderStyle(order) {
            const startTime = new Date(order.CreatedAt);
            const startHour = startTime.getHours() + startTime.getMinutes() / 60;
            const duration = order.CourseMin / 60;

            return {
                left: `${startHour * (100 / 24)}%`,
                width: `${duration * (100 / 24)}%`,
            };
        },
        showTooltip(order, event) {
            this.selectedOrder = order;
            this.tooltipStyle = {
                position: 'fixed',
                display: 'block',
                top: `${event.clientY + 10}px`,
                left: `${event.clientX + 10}px`,
                zIndex: 1000
            };
        },
        hideTooltip() {
            this.selectedOrder = null;
            this.tooltipStyle.display = 'none';
        }
    },
    mounted() {
        this.fetchCasts();
        this.fetchOrders();
    },
};
</script>

<style scoped>
.time-board {
    width: 100%;
    overflow-x: auto;
    background-color: var(--color-background);
    color: var(--color-text);
    padding: 20px;
    border-radius: 8px;
    box-shadow: 0 2px 8px rgba(0, 0, 0, 0.1);
}

.board-content {
    display: flex;
}

.cast-names {
    width: 120px;
    flex-shrink: 0;
}

.cast-name-header {
    height: 40px;
    font-weight: bold;
    display: flex;
    align-items: center;
    justify-content: flex-end;
    padding-right: 15px;
    color: var(--color-heading);
}

.cast-name {
    height: 100px;
    /* タイムラインの縦幅を倍にした */
    padding-right: 15px;
    text-align: right;
    font-weight: bold;
    color: var(--color-heading);
    display: flex;
    align-items: center;
    justify-content: flex-end;
}

.timeline-container {
    flex-grow: 1;
    overflow-x: auto;
    width: 200%;
    /* 追加：タイムラインの幅を2倍に */
}

.time-scale {
    display: flex;
    border-bottom: 1px solid var(--color-border);
    padding-bottom: 10px;
    margin-bottom: 15px;
    height: 40px;
    width: 100%;
    /* 200%から100%に変更 */
}

.hour-marker {
    flex: 1;
    text-align: center;
    font-size: 12px;
    color: var(--color-text-light-2);
    width: calc(100% / 24);
    /* 追加：各時間マーカーの幅を調整 */
}

.cast-rows {
    position: relative;
    width: 100%;
    /* 追加：幅を100%に設定 */
}

.cast-row {
    height: 100px;
    /* タイムラインの縦幅を倍にした */
    margin-bottom: 10px;
}

.order-timeline {
    height: 100%;
    position: relative;
    border-bottom: 1px solid var(--color-border);
    width: 100%;
    /* 200%から100%に変更 */
}

.order-item {
    position: absolute;
    height: 80px;
    /* オーダーアイテムの高さも調整 */
    top: 10px;
    background-color: var(--vt-c-indigo);
    color: var(--vt-c-white);
    font-size: 12px;
    padding: 4px 8px;
    border-radius: 4px;
    overflow: hidden;
    text-overflow: ellipsis;
    white-space: nowrap;
    cursor: pointer;
    transition: background-color 0.3s;
}

.order-item:hover {
    background-color: var(--vt-c-indigo-dark);
}

.free-order {
    background-color: var(--vt-c-divider-light-1);
}

@media (prefers-color-scheme: dark) {
    .time-board {
        background-color: var(--color-background-soft);
    }

    .order-item {
        background-color: var(--vt-c-indigo);
    }

    .free-order {
        background-color: var(--vt-c-divider-dark-1);
    }
}

@media screen and (max-width: 640px) {
    .cast-names {
        width: 80px;
    }

    .cast-name {
        font-size: 12px;
    }

    .order-item {
        font-size: 10px;
        padding: 2px 4px;
    }
}
</style>