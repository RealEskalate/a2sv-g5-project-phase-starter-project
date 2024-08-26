import 'dart:async';
import 'dart:convert';

import 'package:ecommerce_app_ca_tdd/features/chat/data/models/message_model.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

class SocketService {
  late IO.Socket socket;
  final StreamController<MessageModel> _messageController =
      StreamController.broadcast();

  void connectToServer() {
    socket = IO.io('https://g5-flutter-learning-path-be.onrender.com/', <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
    });

    socket.connect();
    socket.on('message:received', (data) {
      // Assuming `data` is a Map<String, dynamic>
      final message = MessageModel.fromJson(data);
      _messageController.add(message);
    });

    // Event listeners
    socket.onConnect((_) {
      var text = 'Connected to server';
      print(text);  
    });

    socket.onDisconnect((_) {
      print('Disconnected from server');
    });

  }
  Stream<MessageModel> get messages => _messageController.stream;

  void sendMessage(String chat, String message, String type) {
    socket.emit('message:send', jsonEncode({
      'chatId': chat,
      'content': message,
      'type': type,
    }));
  }


  void dispose() {
    socket.dispose();
  }
}
