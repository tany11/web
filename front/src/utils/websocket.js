import { ref } from 'vue';
import { useStore } from 'vuex';

const socket = ref(null);
const isConnected = ref(false);

export function useWebSocket() {
    const store = useStore();

    const connect = () => {
        const apiBaseUrl = store.state.apiBaseUrl;
        const wsUrl = apiBaseUrl.replace(/^http/, 'ws') + '/ws';

        socket.value = new WebSocket(wsUrl);

        socket.value.onopen = () => {
            console.log('WebSocket接続が確立されました');
            isConnected.value = true;
        };

        socket.value.onclose = () => {
            console.log('WebSocket接続が閉じられました');
            isConnected.value = false;
        };

        socket.value.onerror = (error) => {
            console.error('WebSocketエラー:', error);
        };
    };

    const disconnect = () => {
        if (socket.value) {
            socket.value.close();
        }
    };

    return {
        socket,
        isConnected,
        connect,
        disconnect,
    };
}