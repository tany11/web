<template>
    <v-card class="free-orders-grid">
        <v-card-title>フリー枠</v-card-title>
        <v-card-text>
            <div class="grid-container">
                <draggable v-model="localFreeOrders" :group="group" :animation="200" @end="onDragEnd" item-key="ID"
                    class="grid-layout">
                    <template #item="{ element }">
                        <div class="grid-item">
                            <OrderItem :order="element" :casts="casts" @update="updateOrder"
                                @order-clicked="$emit('order-clicked', $event)" />
                        </div>
                    </template>
                </draggable>
            </div>
        </v-card-text>
    </v-card>
</template>

<script>
import draggable from 'vuedraggable';
import OrderItem from '@/components/OrderItem.vue';

export default {
    name: 'FreeOrdersGrid',
    components: {
        draggable,
        OrderItem
    },
    props: {
        freeOrders: {
            type: Array,
            required: true
        },
        casts: {
            type: Array,
            required: true
        },
        group: {
            type: Object,
            default: () => ({ name: 'orders', pull: true, put: true })
        }
    },
    emits: ['update', 'order-clicked'],
    data() {
        return {
            localFreeOrders: this.freeOrders
        };
    },
    methods: {
        onDragEnd(evt) {
            if (evt.added) {
                // 新しいアイテムが追加された場合
                const newOrder = evt.added.element;
                newOrder.ActualModel = ''; // フリー枠に移動したのでActualModelをクリア
                newOrder.start_time = null;
                newOrder.end_time = null;
                this.localFreeOrders.splice(evt.added.newIndex, 1, newOrder);
            } else if (evt.removed) {
                // アイテムが削除された場合（他の場所にドラッグされた）
                // 何もしない（親コンポーネントで処理される）
            } else {
                // 順序が変更された場合
                // 何もしない（ローカルの順序は自動的に更新される）
            }
            this.$emit('update', this.localFreeOrders);
        },
        updateOrder(updatedOrder) {
            const index = this.localFreeOrders.findIndex(order => order.ID === updatedOrder.ID);
            if (index !== -1) {
                this.localFreeOrders.splice(index, 1, updatedOrder);
                this.$emit('update', this.localFreeOrders);
            }
        }
    },
    watch: {
        freeOrders: {
            handler(newValue) {
                this.localFreeOrders = newValue;
            },
            deep: true
        }
    }
};
</script>

<style scoped>
.free-orders-grid {
    margin-top: 20px;
}

.grid-container {
    width: 100%;
    min-height: 200px;
}

.grid-layout {
    display: flex;
    flex-wrap: wrap;
    gap: 10px;
}

.grid-item {
    width: calc(25% - 10px);
    min-width: 200px;
    margin-bottom: 10px;
}

@media (max-width: 1200px) {
    .grid-item {
        width: calc(33.33% - 10px);
    }
}

@media (max-width: 900px) {
    .grid-item {
        width: calc(50% - 10px);
    }
}

@media (max-width: 600px) {
    .grid-item {
        width: 100%;
    }
}
</style>