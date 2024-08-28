<template>
    <v-container fluid>
        <v-card class="mx-auto" max-width="1200">
            <v-card-title class="text-h4 font-weight-bold blue-grey lighten-5 py-4">
                タイムボード
            </v-card-title>
            <!-- 日付切り替えボタンを修正 -->
            <v-card-text>
                <v-row align="center" justify="space-between">
                    <v-col cols="auto">
                        <v-btn @click="changeDate(-7)">前の週</v-btn>
                    </v-col>
                    <v-col cols="auto" v-for="offset in [0, 1, 2, 3, 4, 5, 6]" :key="offset">
                        <v-btn :color="offset === currentWeekOffset ? 'primary' : ''" @click="setDate(offset)">
                            {{ formatDateShort(getOffsetDate(offset)) }}
                        </v-btn>
                    </v-col>
                    <v-col cols="auto">
                        <v-btn @click="changeDate(7)">次の週</v-btn>
                    </v-col>
                </v-row>
            </v-card-text>
            <v-row no-gutters>
                <!-- キャスト名 -->
                <v-col cols="2" class="cast-column">
                    <v-list dense class="cast-list">
                        <v-list-item class="px-2 time-header-placeholder">
                            <v-list-item-content></v-list-item-content>
                        </v-list-item>
                        <v-list-item v-for="cast in casts" :key="cast.cast_id" class="px-2 cast-item">
                            <v-list-item-content>
                                <v-list-item-title>{{ cast.name }}</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                        <v-list-item class="px-2 cast-item">
                            <v-list-item-content>
                                <v-list-item-title>フリー枠</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
                </v-col>
                <!-- スケジュール本体 -->
                <v-col cols="10" class="schedule-column" :key="'schedule-column'">
                    <div class="schedule-scroll-container">
                        <div class="schedule-content">
                            <v-card-text class="pa-0">
                                <!-- 時間表示 -->
                                <div class="time-header">
                                    <div v-for="hour in 24" :key="hour" class="hour-label">
                                        <div v-if="(hour + 5) % 24 === 0" class="date-label">
                                            {{ formatDate(getDateForHour(hour)) }}
                                        </div>
                                        {{ formatHour((hour + 5) % 24) }}
                                    </div>
                                </div>
                                <!-- スケジュールグリッド -->
                                <div class="schedule-grid">
                                    <div v-for="cast in [...casts, { cast_id: 'free' }]" :key="cast.cast_id"
                                        class="schedule-row">
                                        <div v-for="hour in 24" :key="hour" class="hour-cell">
                                            <div v-for="quarter in 4" :key="quarter" class="quarter-cell"></div>
                                        </div>
                                    </div>
                                </div>
                                <!-- オーダー表示 -->
                                <OrderItem v-for="order in filteredOrders" :key="order.ID" :order="order"
                                    :timelineStart="timelineStart" :casts="casts" @update="updateOrder" />
                            </v-card-text>
                        </div>
                    </div>
                </v-col>
            </v-row>
        </v-card>
        <v-overlay :value="isLoading">
            <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>
    </v-container>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import draggable from 'vuedraggable/src/vuedraggable'
import OrderItem from '@/components/OrderItem.vue';

export default {
    name: 'TimeBoard',
    components: {
        draggable,
        OrderItem,
    },
    data() {
        return {
            casts: [],
            orders: [],
            currentDate: new Date(),
            isLoading: false,
            weekStart: null,
            selectedDateOffset: null,
        };
    },
    computed: {
        ...mapState(['apiBaseUrl']),
        freeOrders() {
            return this.orders.filter(order => !order.ActualModel);
        },
        castOrders() {
            const ordersByCast = {};
            this.casts.forEach(cast => {
                ordersByCast[cast.cast_id] = this.orders.filter(order => order.ActualModel === cast.cast_id);
            });
            return ordersByCast;
        },
        timelineStart() {
            const start = new Date(this.currentDate);
            start.setHours(6, 0, 0, 0);
            return start;
        },
        weekDates() {
            const dates = [];
            for (let i = 0; i < 7; i++) {
                const date = new Date(this.weekStart);
                date.setDate(date.getDate() + i);
                dates.push(date);
            }
            return dates;
        },
        filteredOrders() {
            if (this.selectedDateOffset === null) {
                return [];
            }
            const currentDateStart = new Date(this.currentDate);
            currentDateStart.setHours(0, 0, 0, 0);
            const currentDateEnd = new Date(this.currentDate);
            currentDateEnd.setHours(23, 59, 59, 999);

            return this.orders.filter(order => {
                return order.start_time >= currentDateStart && order.start_time <= currentDateEnd;
            });
        },
        currentWeekOffset() {
            return this.selectedDateOffset;
        },
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
            this.isLoading = true;
            try {
                const startDate = new Date(this.weekStart);
                const endDate = new Date(startDate);
                endDate.setDate(endDate.getDate() + 7);

                const response = await axios.get(`${this.apiBaseUrl}/orders/scheduled`, {
                    params: {
                        start_time: startDate.toISOString(),
                        end_time: endDate.toISOString(),
                    }
                });
                this.orders = (response.data.data || []).map(order => ({
                    ...order,
                    start_time: new Date(order.ScheduledTime),
                    end_time: new Date(new Date(order.ScheduledTime).getTime() + order.CourseMin * 60000)
                }));
            } catch (error) {
                console.error('オーダーの取得に失敗しました:', error);
                this.orders = [];
            } finally {
                this.isLoading = false;
            }
        },
        updateOrder(updatedOrder) {
            const index = this.orders.findIndex(order => order.ID === updatedOrder.ID);
            if (index !== -1) {
                this.$set(this.orders, index, updatedOrder);
            }
            // ここでAPIを呼び出してサーバー側のデータを更新する
            // オーダーの到着予想時間と実モデルを更新する
        },
        formatHour(hour) {
            return `${(hour + 24) % 24}:00`;
        },
        formatTime(time) {
            return time.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
        getOrderStyle(order) {
            const startTime = order.start_time;
            const endTime = order.end_time;
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const dayStart = new Date(startTime);
            dayStart.setHours(6, 0, 0, 0);
            const startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = (order.ActualModel ? this.casts.findIndex(cast => cast.cast_id === order.ActualModel) : this.casts.length) * 50 + 50;

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
            // ここでオーダーに応じて色を返す処理を実装
            // 例: オーダーのタイプや状態に基づいて色を決定
            const colors = ['indigo', 'cyan', 'green', 'amber', 'pink', 'purple', 'blue', 'teal'];
            return colors[order.ID % colors.length];
        },
        timeToMinutes(time) {
            const [hours, minutes] = time.split(':').map(Number);
            return hours * 60 + minutes;
        },
        formatDate(date) {
            const options = { month: 'long', day: 'numeric' };
            return date.toLocaleDateString('ja-JP', options);
        },
        getDateForHour(hour) {
            const date = new Date(this.currentDate);
            date.setHours(hour + 5);
            return date;
        },
        async changeDate(days) {
            const newDate = new Date(this.weekStart);
            newDate.setDate(newDate.getDate() + days);
            this.weekStart = this.getMonday(newDate);
            this.currentDate = new Date(this.weekStart);
            this.selectedDateOffset = this.getCurrentDayOffset();
            await this.fetchOrders();
        },
        async setDate(offset) {
            this.selectedDateOffset = offset;
            this.currentDate = new Date(this.weekStart);
            this.currentDate.setDate(this.currentDate.getDate() + offset);
        },
        getOffsetDate(offset) {
            const date = new Date(this.weekStart);
            date.setDate(date.getDate() + offset);
            return date;
        },
        formatDateShort(date) {
            const options = { weekday: 'short', month: 'numeric', day: 'numeric' };
            return date.toLocaleDateString('ja-JP', options);
        },
        getMonday(date) {
            const day = date.getDay();
            const diff = date.getDate() - day + (day === 0 ? -6 : 1);
            return new Date(date.setDate(diff));
        },
        getCurrentDayOffset() {
            const today = new Date();
            const currentWeekDay = this.weekDates.findIndex(date =>
                date.getDate() === today.getDate() &&
                date.getMonth() === today.getMonth() &&
                date.getFullYear() === today.getFullYear()
            );
            return currentWeekDay === -1 ? this.selectedDateOffset : currentWeekDay;
        },
    },
    async mounted() {
        this.weekStart = this.getMonday(new Date());
        this.currentDate = new Date(this.weekStart);
        await this.fetchCasts();
        await this.fetchOrders();
    },
};
</script>

<style scoped>
.cast-column {
    position: sticky;
    left: 0;
    z-index: 2;
    background-color: white;
    border-right: 1px solid #e0e0e0;
}

.cast-list {
    padding-top: 0;
}

.time-header-placeholder {
    height: 50px;
    border-bottom: 1px solid #e0e0e0;
}

.cast-item {
    height: 50px;
    border-bottom: 1px solid #e0e0e0;
}

.schedule-column {
    overflow: hidden;
    position: relative;
}

.schedule-scroll-container {
    width: 100%;
    height: 100%;
    overflow-x: auto;
    overflow-y: hidden;
    position: absolute;
    top: 0;
    left: 0;
}

.schedule-content {
    width: 1200px;
    position: relative;
}

.time-header {
    display: flex;
    border-bottom: 1px solid #e0e0e0;
    position: sticky;
    top: 0;
    background-color: white;
    z-index: 1;
    height: 50px;
}

.hour-label {
    width: 50px;
    text-align: center;
    font-size: 12px;
    display: flex;
    flex-direction: column;
    justify-content: flex-end;
    height: 100%;
    padding-bottom: 2px;
    border-right: 1px solid #e0e0e0;
}

.date-label {
    font-size: 10px;
    font-weight: bold;
    margin-bottom: 2px;
}

.schedule-grid {
    position: relative;
    margin-top: 0;
}

.schedule-row {
    display: flex;
    height: 50px;
    border-bottom: 1px solid #e0e0e0;
}

.hour-cell {
    width: 100px;
    border-right: 1px solid #e0e0e0;
    display: flex;
}

.quarter-cell {
    flex: 1;
    border-right: 1px dashed #e0e0e0;
}

.quarter-cell:last-child {
    border-right: none;
}

.order-item {
    position: absolute;
    z-index: 3;
    top: calc(var(--top) + 50px);
}

.order-content {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    font-weight: bold;
    color: var(--vt-c-indigo-darker);
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

    .cast-column,
    .time-header {
        background-color: var(--color-background-soft);
    }
}

@media screen and (max-width: 640px) {
    .cast-column {
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

.date-header {
    display: none;
}
</style>