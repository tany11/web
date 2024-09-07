<template>
    <v-container fluid>
        <v-card class="mx-auto" max-width="1200">
            <!-- メモ機能が欲しいよ！！！！！！！！ -->
            <v-card-title class="text-h4 font-weight-bold blue-grey lighten-5 py-4">
                タイムボード
            </v-card-title>
            <v-col cols="1" class="memo-column">
                <v-btn color="primary" @click="openNewMemoModal">
                    メモ追加
                </v-btn>
            </v-col>
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
                        </v-list-item>
                        <v-list-item v-for="cast in casts" :key="cast.cast_id" class="px-2 cast-item">
                            <v-list-item-title>{{ cast.name }}</v-list-item-title>
                        </v-list-item>
                        <v-list-item class="px-2 cast-item">
                            <v-list-item-title>フリー枠</v-list-item-title>
                        </v-list-item>
                    </v-list>
                </v-col>
                <!-- スケジュール本体 -->
                <v-col cols="9" class="schedule-column" :key="'schedule-column'">
                    <div class="schedule-scroll-container">
                        <div class="schedule-content">
                            <v-card-text class="pa-0">
                                <!-- 時間表示 -->
                                <div class="time-header">
                                    <div v-for="hour in 24" :key="hour" class="hour-label">
                                        <div v-if="(hour + 5) % 24 === 0" class="date-label">
                                            {{ formatDate(getDateForHour(hour)) }}
                                        </div>
                                        {{ formatHour(hour) }}
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
                                <draggable v-model="ordersWithCastAssignment" group="orders" :animation="200"
                                    item-key="ID" @end="onDragEnd">
                                    <template #item="{ element }">
                                        <OrderItem :order="element" :timelineStart="timelineStart"
                                            :dayStart="getDayStart()" :dayEnd="getDayEnd()" :casts="casts"
                                            @update="updateOrder" @order-clicked="handleOrderClicked" />
                                    </template>
                                </draggable>
                                <!-- メモ表示 -->
                                <draggable v-model="memosWithCastAssignment" group="memos" :animation="200"
                                    item-key="ID" @end="onMemoDragEnd">
                                    <template #item="{ element }">
                                        <MemoItem :memo="element" :timelineStart="timelineStart"
                                            :dayStart="getDayStart()" :dayEnd="getDayEnd()" :casts="casts"
                                            @update="updateMemoSchedule" @memo-clicked="handleMemoClicked" />
                                    </template>
                                </draggable>
                            </v-card-text>
                        </div>
                    </div>
                </v-col>
                <!-- メモボタン -->

            </v-row>
        </v-card>
        <v-overlay :value="isLoading">
            <v-progress-circular indeterminate size="64"></v-progress-circular>
        </v-overlay>
        <v-dialog v-model="showOrderModal" max-width="800px">
            <OrderDetail :order="selectedOrder" @close="closeOrderModal" @order-updated="fetchOrders" />
        </v-dialog>
        <v-dialog v-model="showMemoModal" max-width="800px">
            <MemoDetail :memo="selectedMemo" @close="closeMemoModal" @memo-updated="fetchMemos" />
        </v-dialog>
        <v-dialog v-model="showNewMemoModal" max-width="800px">
            <MemoDetail :memo="newMemo" @close="closeNewMemoModal" @memo-updated="handleNewMemoCreated" />
        </v-dialog>
    </v-container>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import draggable from 'vuedraggable/src/vuedraggable'
import OrderItem from '@/components/OrderItem.vue';
import OrderDetail from '@/components/OrderDetail.vue';
import MemoItem from '@/components/MemoItem.vue';
import MemoDetail from '@/components/MemoDetail.vue';

export default {
    name: 'TimeBoard',
    components: {
        draggable,
        OrderItem,
        OrderDetail,
        MemoItem,
        MemoDetail
    },
    data() {
        return {
            casts: [],
            orders: [],
            currentDate: new Date(),
            isLoading: false,
            weekStart: null,
            selectedDateOffset: 0, // 初期値を0に設定
            showOrderModal: false,
            selectedOrder: null,
            memos: [],
            showMemoModal: false,
            selectedMemo: null,
            showNewMemoModal: false,
            newMemo: {},
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
            currentDateStart.setHours(6, 0, 0, 0);
            const currentDateEnd = new Date(currentDateStart);
            currentDateEnd.setDate(currentDateEnd.getDate() + 1);
            currentDateEnd.setHours(5, 59, 59, 999);

            return this.orders.filter(order => {
                const orderStart = new Date(order.start_time);
                const orderEnd = new Date(order.end_time);

                // 日付をまたぐオーダーの処理
                if (orderEnd < orderStart) {
                    orderEnd.setDate(orderEnd.getDate() + 1);
                }

                return (orderStart < currentDateEnd && orderEnd > currentDateStart);
            });
        },
        filteredMemos() {
            if (this.selectedDateOffset === null) {
                return [];
            }
            const currentDateStart = new Date(this.currentDate);
            currentDateStart.setHours(6, 0, 0, 0);
            const currentDateEnd = new Date(currentDateStart);
            currentDateEnd.setDate(currentDateEnd.getDate() + 1);
            currentDateEnd.setHours(5, 59, 59, 999)
            console.log(this.memos);

            return this.memos.filter(memo => {
                const memoStart = new Date(memo.start_time);
                const memoEnd = new Date(memo.end_time);

                // 日付をまたぐメモの処理
                if (memoEnd < memoStart) {
                    memoEnd.setDate(memoEnd.getDate() + 1);
                }

                return (memoStart < currentDateEnd && memoEnd > currentDateStart);
            });
        },
        currentWeekOffset() {
            return this.selectedDateOffset;
        },
        ordersWithCastAssignment() {
            return this.filteredOrders.map(order => {
                const castIndex = this.casts.findIndex(cast => cast.cast_id === order.ActualModel);
                const startTime = new Date(order.start_time);
                const endTime = new Date(order.end_time);

                // 日付をまたぐオーダーの処理
                if (endTime < startTime) {
                    endTime.setDate(endTime.getDate() + 1);
                }

                return {
                    ...order,
                    castIndex: castIndex !== -1 ? castIndex : this.casts.length,
                    start_time: startTime,
                    end_time: endTime
                };
            });
        },
        memosWithCastAssignment() {
            return this.filteredMemos.map(memo => {
                const castIndex = this.casts.findIndex(cast => cast.cast_id === memo.ActualModel);
                const startTime = new Date(memo.start_time);
                const endTime = new Date(memo.end_time);

                // 日付をまたぐメモの処理
                if (endTime < startTime) {
                    endTime.setDate(endTime.getDate() + 1);
                }
                console.log(startTime, endTime);

                return {
                    ...memo,
                    castIndex: castIndex !== -1 ? castIndex : this.casts.length,
                    start_time: startTime,
                    end_time: endTime
                };
            });
        }
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
                startDate.setHours(6, 0, 0, 0);
                const endDate = new Date(startDate);
                endDate.setDate(endDate.getDate() + 7);
                endDate.setHours(5, 59, 59, 999);

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
                // オーダーを取得した後、現在の日付に対応するオーダーをフィルタリング
                this.setDate(this.selectedDateOffset);
            } catch (error) {
                console.error('オーダーの取得に失敗ました:', error);
                this.orders = [];
            } finally {
                this.isLoading = false;
            }
        },
        async updateOrder(updatedOrder) {
            const index = this.orders.findIndex(order => order.ID === updatedOrder.ID);
            if (index !== -1) {
                this.orders[index] = updatedOrder;
            }

            try {
                await axios.put(`${this.apiBaseUrl}/orders/${updatedOrder.ID}/schedule`, null, {
                    params: {
                        actual_model: updatedOrder.ActualModel,
                        scheduled_time: updatedOrder.start_time.toISOString()
                    }
                });
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
            }
        },
        formatHour(hour) {
            return `${(hour + 5) % 24}:00`;
        },
        formatTime(time) {
            return time.toLocaleTimeString('ja-JP', { hour: '2-digit', minute: '2-digit' });
        },
        getOrderStyle(order) {
            const startTime = order.start_time;
            const endTime = order.end_time;
            const duration = (endTime - startTime) / (1000 * 60); // 分単位での期間
            const dayStart = new Date(this.currentDate);
            dayStart.setHours(6, 0, 0, 0);
            const dayEnd = new Date(dayStart);
            dayEnd.setDate(dayEnd.getDate() + 1);
            dayEnd.setHours(5, 59, 59, 999);

            let startPercentage = ((startTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            let widthPercentage = (duration / (24 * 60)) * 100;

            // 日付をまたぐオーダーの処理
            if (startTime < dayStart) {
                startPercentage = 0;
                widthPercentage = ((endTime - dayStart) / (24 * 60 * 60 * 1000)) * 100;
            } else if (endTime > dayEnd) {
                widthPercentage = ((dayEnd - startTime) / (24 * 60 * 60 * 1000)) * 100;
            }

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
            if (hour >= 19) {
                date.setDate(date.getDate() + 1);
            }
            date.setHours((hour + 5) % 24, 0, 0, 0);
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
        getDayStart() {
            const start = new Date(this.currentDate);
            start.setHours(6, 0, 0, 0);
            return start;
        },
        getDayEnd() {
            const end = new Date(this.currentDate);
            end.setDate(end.getDate() + 1);
            end.setHours(5, 59, 59, 999);
            return end;
        },
        onDragEnd(event) {
            const newOrder = this.ordersWithCastAssignment[event.newIndex];
            const newStartTime = this.calculateNewStartTime(event);
            const newCastId = this.calculateNewCastId(event);

            newOrder.start_time = newStartTime;
            newOrder.end_time = new Date(newStartTime.getTime() + newOrder.CourseMin * 60000);
            newOrder.ActualModel = newCastId;

            console.log('ドラッグ終了:', {
                newOrder,
                newStartTime: newStartTime.toISOString(),
                newCastId
            });

            this.updateOrderSchedule(newOrder);
        },
        calculateNewStartTime(event) {
            const scheduleContent = this.$el.querySelector('.schedule-content');
            const rect = scheduleContent.getBoundingClientRect();
            const offsetX = event.originalEvent.clientX - rect.left;

            const cellWidth = rect.width / (24 * 4); // 1日24時間、15分ごとのセル
            const quarterHours = Math.floor(offsetX / cellWidth);

            const newStartTime = new Date(this.currentDate);
            newStartTime.setHours(6, 0, 0, 0); // 日の開始時刻を6:00に設定
            newStartTime.setMinutes(newStartTime.getMinutes() + quarterHours * 15);

            return newStartTime;
        },
        calculateNewCastId(event) {
            const scheduleContent = this.$el.querySelector('.schedule-content');
            const rect = scheduleContent.getBoundingClientRect();
            const offsetY = event.originalEvent.clientY - rect.top;

            const rowHeight = 50; // キャスト行の高さ（ピクセル）
            const castIndex = Math.floor(offsetY / rowHeight) - 1; // -1 は時間ヘッダーの分

            return (castIndex >= 0 && castIndex < this.casts.length) ? this.casts[castIndex].cast_id : '';
        },
        async updateOrderSchedule(updatedOrder) {
            try {
                await axios.put(`${this.apiBaseUrl}/orders/${updatedOrder.ID}/schedule`, null, {
                    params: {
                        actual_model: updatedOrder.ActualModel,
                        scheduled_time: updatedOrder.start_time.toISOString()
                    }
                });
                // 成功した場合、ローカルのオーダーデータを更新
                const index = this.orders.findIndex(order => order.ID === updatedOrder.ID);
                if (index !== -1) {
                    this.orders.splice(index, 1, updatedOrder);
                }
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
                // エラー処理（例：ユーザーに通知）
            }
        },
        handleOrderClicked(order) {
            this.selectedOrder = order;
            this.showOrderModal = true;
        },
        closeOrderModal() {
            this.showOrderModal = false;
            this.selectedOrder = null;
        },
        async fetchMemos() {
            try {
                const startDate = new Date(this.weekStart);
                startDate.setHours(6, 0, 0, 0);
                const endDate = new Date(startDate);
                endDate.setDate(endDate.getDate() + 7);
                endDate.setHours(5, 59, 59, 999);

                const response = await axios.get(`${this.apiBaseUrl}/tips/schedule`, {
                    params: {
                        start_time: startDate.toISOString(),
                        end_time: endDate.toISOString(),
                    }
                });

                console.log('API Response:', response.data);

                this.memos = (response.data.data || []).map(memo => {
                    const start_time = new Date(memo.ScheduledTime);
                    const end_time = new Date(new Date(memo.ScheduledTime).getTime() + memo.ScheduledBox * 60000);
                    console.log("てすと", start_time, end_time);
                    return {
                        ...memo,
                        start_time,
                        end_time
                    };
                });

                console.log('Processed Memos:', this.memos);
            } catch (error) {
                console.error('メモの取得に失敗しました:', error);
                this.memos = [];
            }
        },
        handleMemoClicked(memo) {
            this.selectedMemo = memo;
            this.showMemoModal = true;
        },
        closeMemoModal() {
            this.showMemoModal = false;
            this.selectedMemo = null;
        },
        onMemoDragEnd(event) {
            const newMemo = this.memosWithCastAssignment[event.newIndex];
            const newStartTime = this.calculateNewStartTime(event);
            const newCastId = this.calculateNewCastId(event);

            // 開始時間を更新
            newMemo.start_time = newStartTime;
            // 終了時間を更新（ScheduledBoxを使用）
            newMemo.end_time = new Date(newStartTime.getTime() + newMemo.ScheduledBox * 60000);
            newMemo.ActualModel = newCastId;

            this.updateMemoSchedule(newMemo);
        },
        async updateMemoSchedule(updatedMemo) {
            try {
                await axios.put(`${this.apiBaseUrl}/tips/${updatedMemo.ID}/schedule`, null, {
                    params: {
                        actual_model: updatedMemo.ActualModel,
                        scheduled_time: updatedMemo.start_time.toISOString(),
                        scheduled_box: updatedMemo.ScheduledBox,
                        title: updatedMemo.Title,
                        notes: updatedMemo.Notes
                    }
                });
                // 成功した場合、ローカルのメモデータを更新
                const index = this.memos.findIndex(memo => memo.ID === updatedMemo.ID);
                if (index !== -1) {
                    this.memos.splice(index, 1, {
                        ...updatedMemo,
                        end_time: new Date(updatedMemo.start_time.getTime() + updatedMemo.ScheduledBox * 60000)
                    });
                }
            } catch (error) {
                console.error('メモの更新に失敗しました:', error);
            }
        },
        openNewMemoModal() {
            this.newMemo = {
                ID: null,  // IDをnullに設定
                start_time: new Date(this.currentDate),
                end_time: new Date(new Date(this.currentDate).getTime() + 30 * 60000), // 30分後を終了時間とする
                ActualModel: '',
                Title: '',
                Notes: '',
                ScheduledBox: 30, // デフォルトで30分
            };
            this.showNewMemoModal = true;
        },
        closeNewMemoModal() {
            this.showNewMemoModal = false;
        },
        async handleNewMemoCreated() {
            await this.fetchMemos();
            this.closeNewMemoModal();
        },
    },
    async mounted() {
        this.weekStart = this.getMonday(new Date());
        this.currentDate = new Date(this.weekStart);
        this.selectedDateOffset = this.getCurrentDayOffset();
        await this.fetchCasts();
        await this.fetchOrders();
        await this.fetchMemos();
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
    position: relative;
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
    background-color: var(--color-background);
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
    z-index: 1;
    /* スケジュールグリッドのz-indexを1に設定 */
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

.order-items-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 2;
    /* オーダーアイテムのz-indexを2に設定 */
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
    position: relative;
    top: 0;
    left: 0;
}

.schedule-content {
    width: 1200px;
    position: relative;
}

.time-header,
.hour-label {
    background-color: var(--color-background);
}

.time-header {
    display: flex;
    border-bottom: 1px solid #e0e0e0;
    position: sticky;
    top: 0;
    background-color: var(--color-background);
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
    z-index: 1;
    /* スケジュールグリッドのz-indexを1に設定 */
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

.order-items-container {
    position: absolute;
    top: 0;
    left: 0;
    width: 100%;
    height: 100%;
    z-index: 2;
    /* オーダーアイテムのz-indexを2に設定 */
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

.memo-column {
    display: flex;
    justify-content: center;
    align-items: flex-start;
    padding-top: 10px;
}
</style>