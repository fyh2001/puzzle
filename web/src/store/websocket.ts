import { defineStore } from "pinia";
import { ref } from "vue";
import { baseURL } from "@/api";
import { useMessage } from "naive-ui";
import { useUserStore } from "./user";
import { useNotificationStore } from "./notification";
import type { Message } from "@/types/websocket";

export const useWebsocketStore = defineStore("websocket", () => {
  const Message = useMessage();

  const socket = ref<WebSocket | null>(null);
  const socketConnected = ref(false);

  // 连接
  const connect = () => {
    socket.value = new WebSocket(baseURL.replace("http", "ws") + "/ws/connect");

    socket.value.onopen = () => {
      console.log("websocket connected");
      socketConnected.value = true;

      auth();
    };

    socket.value.onclose = () => {
      console.log("websocket closed");
      socketConnected.value = false;
    };

    socket.value.onerror = () => {
      console.log("websocket error");
      socketConnected.value = false;
    };

    socket.value.onmessage = (event) => {
      const message: Message = JSON.parse(event.data);
      handleMessages(message);
    };

    // 心跳检测
    setInterval(() => {
      if (socket.value) {
        send({
          type: "ping",
          content: "ping",
        });
      }
    }, 10000);
  };

  // 发送消息
  const send = (message: Message) => {
    if (socket.value) {
      const messageStr = JSON.stringify(message);

      socket.value.send(messageStr);
    }
  };

  // 断开连接
  const disconnect = () => {
    if (socket.value) {
      socket.value.close();
      socket.value = null;
    }
  };

  // 重连
  const reconnect = () => {
    disconnect();
    connect();
  };

  // 认证
  const auth = () => {
    const userStore = useUserStore();

    send({
      type: "auth",
      content: userStore.getToken,
    });
  };

  // 处理消息
  const handleMessages = (message: Message) => {
    const notificationStore = useNotificationStore();

    if (message.type === "notification") {
      notificationStore.fetchUnreadCount();
      Message.info("您有 1 条新的通知");
    }

    if (message.type === "auth") {
      Message.info(message.content);
    }

    if (message.type === "pong") {
      console.log("pong");
    }
  };

  return {
    socket,
    socketConnected,
    connect,
    send,
    disconnect,
    reconnect,
    auth,
    handleMessages,
  };
});
