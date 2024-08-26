// import 'dart:async';
// import 'dart:developer';

// import 'package:socket_io_client/socket_io_client.dart' as IO;

// import '../../features/chat/data/model/message_model.dart';

// class WebSocketService {
//   late IO.Socket socket;
//   final StreamController<MessageModel> _messageStreamController =
//       StreamController<MessageModel>.broadcast();
//   final String accessToken;

//   WebSocketService({required this.socketk, required this.accessToken}) {
//     _initializeWebSocket();
//   }

//   Stream<MessageModel> get messages => _messageStreamController.stream;

//   void _initializeWebSocket() {
//     socket.on('message:received', (data) {
//       final message = MessageModel.fromJson(data);
//       _messageStreamController.add(message);
//     });

//     socket.on('message:delivered', (data) {
//       log('Message delivered: $data');
//     });

//     socket.onDisconnect((_) {
//       dispose();
//     });
//   }

//   void dispose() {
//     _messageStreamController.close();
//     socket.disconnect();
//   }

//   void _initializeWebSocket() {
//     socket = IO.io(
//       'https://g5-flutter-learning-path-be.onrender.com',
//       <String, dynamic>{
//         'transports': ['websocket'],
//         'extraHeaders': {'Authorization': 'Bearer $accessToken'}
//       },
//     );

//     socket.connect();

//     socket.onConnect((_) {
//       log('Connected to the socket server');
//     });

//     socket.onConnectError((error) {
//       print("Connection Error: $error");
//       // Add reconnection logic if needed
//       _reconnect();
//     });

//     socket.onDisconnect((_) {
//       print("Disconnected from WebSocket server");
//       _reconnect();
//     });

//     socket.on('message:received', (data) {
//       final message = MessageModel.fromJson(data);
//       _messageStreamController.add(message);
//     });
//   }

//   void _reconnect() {
//     // Reconnect logic: could include exponential backoff, retries, etc.
//     Future.delayed(Duration(seconds: 5), () {
//       if (!socket.connected) {
//         socket.connect();
//       }
//     });
//   }

//   void sendMessage(String chatId, String message, String type) {
//     socket.emit('message:send', {
//       'chatId': chatId,
//       'content': message,
//       'type': type,
//     });
//   }

//   void dispose() {
//     _messageStreamController.close();
//     socket.disconnect();
//   }
// }
