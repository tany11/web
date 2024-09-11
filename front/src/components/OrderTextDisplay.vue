<template>
    <v-dialog v-model="dialog" max-width="500px">
        <v-card>
            <v-card-title>オーダー詳細</v-card-title>
            <v-card-text>
                <pre>{{ formattedOrderDetails }}</pre>
            </v-card-text>
            <v-card-actions>
                <v-spacer></v-spacer>
                <v-btn color="primary" text @click="dialog = false">閉じる</v-btn>
            </v-card-actions>
        </v-card>
    </v-dialog>
</template>

<script>
export default {
    props: {
        orderDetails: {
            type: String,
            required: true
        }
    },
    data() {
        return {
            dialog: false
        }
    },
    computed: {
        formattedOrderDetails() {
            return this.orderDetails.trim();
        }
    },
    methods: {
        show() {
            this.dialog = true;
        }
    }
}
</script>

<style scoped>
pre {
    white-space: pre-wrap;
    word-wrap: break-word;
}
</style>