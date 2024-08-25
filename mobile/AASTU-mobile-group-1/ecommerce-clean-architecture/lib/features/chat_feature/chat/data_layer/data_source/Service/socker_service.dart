import 'dart:async';
import 'package:ecommerce/features/chat_feature/chat/data_layer/model/message_model.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketService {
  final IO.Socket _socket;
  final StreamController<MessageModel> _messageController =
      StreamController.broadcast();

  SocketService(String token)
      : _socket = IO.io('https://g5-flutter-learning-path-be.onrender.com/', <String, dynamic>{
          'transports': ['websocket'],
          'extraHeaders': {'Authorization': 'Bearer $token'},
          'autoConnect': false,
        }) {
    _initialize();
  }

  void _initialize() {
    print("connect");
    _socket.onConnect((_) {
      print('Connected to WebSocket server');
    });

    _socket.onDisconnect((_) {
      print('Disconnected from WebSocket server');
    });

    _socket.on('message:received', (data) {
      // Assuming `data` is a Map<String, dynamic>
      final message = MessageModel.fromJson(data);
      _messageController.add(message);
    });

    _socket.connect();
  }

  Stream<MessageModel> get messages => _messageController.stream;

  void sendMessage(String chatId, String content, String type) {
  if (_socket.connected) {
    print("lets send");
    _socket.emit('message:send', {
      'chatId': chatId,
      'content': content,
      'type': type,
    });
  } else {
    print("Socket is not connected");
  }
}

  void dispose() {
    _socket.disconnect();
    _messageController.close();
  }
}
