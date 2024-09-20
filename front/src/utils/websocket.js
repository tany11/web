import { ref } from 'vue';
import { useStore } from 'vuex';

const socket = ref(null);
const isConnected = ref(false);

export function useWebSocket() {
    const store = useStore();

    const connect = () => {
        const apiBaseUrl = store.state.apiBaseUrl;
        const token = store.state.token;

        console.log('Connecting with token:', token);

        // apiBaseUrlからホスト部分のみを抽出
        const urlParts = new URL(apiBaseUrl);
        const socketUrl = `${urlParts.protocol === 'https:' ? 'wss:' : 'ws:'}//${urlParts.host}/api/v1/ws`;
        console.log('socketUrl:', socketUrl);

        socket.value = new WebSocket(`${socketUrl}?token=${token}`);

        socket.value.onopen = () => {
            console.log('WebSocket接続が確立されました');
            isConnected.value = true;
        };

        socket.value.onclose = (event) => {
            console.log('WebSocket接続が閉じられました', event.reason);
            isConnected.value = false;
        };

        socket.value.onerror = (error) => {
            console.error('WebSocketエラー:', error);
        };

        socket.value.onmessage = (event) => {
            const data = JSON.parse(event.data);
            console.log('受信したメッセージ:', data);
            
            // メッセージの種類に応じて処理を行う
            if (data.content && typeof data.content === 'object') {
                if (data.content.type === 'order_update') {
                    store.dispatch('updateOrder', data.content.order);
                }
            }
        };
    };

    const disconnect = () => {
        if (socket.value && socket.value.readyState === WebSocket.OPEN) {
            socket.value.close();
        }
    };

    const sendMessage = (message) => {
        if (socket.value && socket.value.readyState === WebSocket.OPEN) {
            socket.value.send(JSON.stringify(message));
        } else {
            console.error('WebSocket is not connected.');
        }
    };

    return {
        socket,
        isConnected,
        connect,
        disconnect,
        sendMessage,
    };
}