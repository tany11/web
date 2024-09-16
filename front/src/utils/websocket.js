import { ref } from 'vue';
import { useStore } from 'vuex';
import { io } from 'socket.io-client';

const socket = ref(null);
const isConnected = ref(false);

export function useWebSocket() {
    const store = useStore();

    const connect = () => {
        const apiBaseUrl = store.state.apiBaseUrl;
        const token = store.state.token;

        console.log('Connecting with token:', token); // トークンをログ出力

        const socketUrl = `${apiBaseUrl}/socket.io`;  // /api/v1 プレフィックスを削除

        console.log('Attempting to connect to Socket.IO:', socketUrl);

        socket.value = io(socketUrl, {
            transports: ['websocket'],
            extraHeaders: {
                Authorization: `Bearer ${token}`
            }
        });

        socket.value.on('connect', () => {
            console.log('Socket.IO接続が確立されました');
            isConnected.value = true;
        });

        socket.value.on('disconnect', (reason) => {
            console.log('Socket.IO接続が閉じられました', reason, socketUrl);
            isConnected.value = false;
        });

        socket.value.on('connect_error', (error) => {
            console.error('Socket.IOエラー:', error);
        });

        socket.value.on('message', (data) => {
            console.log('受信したメッセージ:', data);
            // ここでメッセージを処理します
        });
    };

    const disconnect = () => {
        if (socket.value) {
            socket.value.disconnect();
        }
    };

    const sendMessage = (message) => {
        if (socket.value && socket.value.connected) {
            socket.value.emit('message', message);
        } else {
            console.error('Socket.IO is not connected.');
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