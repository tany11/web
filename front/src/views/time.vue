<template>
    <v-container fluid>
        <v-card class="mx-auto" max-width="1200">
            <v-card-title class="text-h4 font-weight-bold blue-grey lighten-5 py-4">
                詳細スケジュール
            </v-card-title>
            <v-row no-gutters>
                <!-- キャスト名の列 -->
                <v-col cols="2">
                    <v-list dense>
                        <v-list-item v-for="(cast, index) in casts" :key="index" class="px-2">
                            <v-list-item-content>
                                <v-list-item-title>{{ cast.name }}</v-list-item-title>
                            </v-list-item-content>
                        </v-list-item>
                    </v-list>
                </v-col>
                <!-- スケジュール本体 -->
                <v-col cols="10">
                    <v-card-text class="pa-0">
                        <!-- 時間表示 -->
                        <v-row no-gutters class="time-header">
                            <v-col v-for="hour in 24" :key="hour" cols="1" class="text-center">
                                {{ formatHour(hour - 1) }}
                            </v-col>
                        </v-row>
                        <!-- スケジュールグリッド -->
                        <v-row no-gutters class="schedule-grid">
                            <v-col v-for="cast in casts" :key="cast.id" cols="12">
                                <v-row no-gutters class="schedule-row" style="height: 50px;">
                                    <v-col v-for="hour in 24" :key="hour" cols="1" class="hour-cell">
                                        <div v-for="quarter in 4" :key="quarter" class="quarter-cell"></div>
                                    </v-col>
                                </v-row>
                            </v-col>
                        </v-row>
                        <!-- タスク表示 -->
                        <div v-for="task in tasks" :key="task.id" class="task-item" :style="getTaskStyle(task)">
                            <v-tooltip bottom>
                                <template v-slot:activator="{ on, attrs }">
                                    <div v-bind="attrs" v-on="on" :class="`task-content ${task.color}--text`">
                                        {{ task.name }}
                                    </div>
                                </template>
                                <span>{{ task.name }} ({{ formatTime(task.start) }} - {{ formatTime(task.end) }})</span>
                            </v-tooltip>
                        </div>
                    </v-card-text>
                </v-col>
            </v-row>
        </v-card>
    </v-container>
</template>

<script>
export default {
    name: 'DetailedSchedule',
    data() {
        return {
            casts: [
                { id: 1, name: 'キャスト1' },
                { id: 2, name: 'キャスト2' },
                { id: 3, name: 'キャスト3' },
                { id: 4, name: 'キャスト4' },
            ],
            tasks: [
                { id: 1, castId: 1, name: 'タスク1', start: '09:00', end: '11:30', color: 'indigo' },
                { id: 2, castId: 2, name: 'タスク2', start: '10:00', end: '14:00', color: 'cyan' },
                { id: 3, castId: 3, name: 'タスク3', start: '13:00', end: '16:30', color: 'green' },
                { id: 4, castId: 4, name: 'タスク4', start: '15:00', end: '18:00', color: 'amber' },
            ],
        };
    },
    methods: {
        formatHour(hour) {
            return `${hour.toString().padStart(2, '0')}:00`;
        },
        formatTime(time) {
            return time;
        },
        getTaskStyle(task) {
            const startTime = this.timeToMinutes(task.start);
            const endTime = this.timeToMinutes(task.end);
            const duration = endTime - startTime;
            const startPercentage = (startTime / (24 * 60)) * 100;
            const widthPercentage = (duration / (24 * 60)) * 100;
            const top = this.casts.findIndex(cast => cast.id === task.castId) * 50 + 30; // 30px for time header

            return {
                position: 'absolute',
                left: `${startPercentage}%`,
                width: `${widthPercentage}%`,
                top: `${top}px`,
                height: '46px', // 50px - 4px for borders
                backgroundColor: `${task.color} lighten-4`,
                border: `2px solid ${task.color}`,
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
    },
};
</script>

<style scoped>
.time-header {
    height: 30px;
    border-bottom: 1px solid #e0e0e0;
}

.schedule-grid {
    position: relative;
}

.schedule-row {
    border-bottom: 1px solid #e0e0e0;
}

.hour-cell {
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

.task-item {
    position: absolute;
    z-index: 1;
}

.task-content {
    height: 100%;
    display: flex;
    align-items: center;
    justify-content: center;
    font-size: 12px;
    font-weight: bold;
}
</style>