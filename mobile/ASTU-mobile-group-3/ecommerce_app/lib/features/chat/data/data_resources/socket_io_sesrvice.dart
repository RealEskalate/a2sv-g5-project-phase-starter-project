import 'dart:async';
import 'dart:developer';

import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../../../core/constants/constants.dart';
import '../../../auth/data/data_source/auth_local_data_source.dart';
import '../../domain/entity/message.dart';

abstract class SocketIOService {
  IO.Socket get socket;
  void connect();
  void disconnect();
  Future<void> emitSendMessage(String chatId, String content, String type);
}

class SocketIOServiceImpl implements SocketIOService {
  late IO.Socket socket;
  final AuthLocalDataSource authLocalDataSource;

  final StreamController<MessageEntity> _messageStreamController =
      StreamController<MessageEntity>.broadcast();

  SocketIOServiceImpl({required this.authLocalDataSource}) {
    socket = IO.io(AppData.chatserver, <String, dynamic>{
      'transports': ['websocket'],
      'autoConnect': false,
      'extraHeaders': {
        'Authorization': 'Bearer ${authLocalDataSource.getToken()}'
      }
    });
  }

  @override
  void connect() {
    socket.connect();
  }

  // eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImRhZ2ltQGdtYWlsLmNvbSIsInN1YiI6IjY2Y2EyODBhZjk4ZDMyYzY4ZWFkY2UxNSIsImlhdCI6MTcyNDUyNDU4NCwiZXhwIjoxNzI0OTU2NTg0fQ.P435ttt-_a53CUjJ7ZFeoaDvm-MNmcZapRBUkkqY7eM

  @override
  Future<void> emitSendMessage(
      String chatId, String content, String type) async {
    final messagePayload = {'chatId': chatId, 'content': content, 'type': type};

    socket.emit('message:send', messagePayload);

    socket.on('message:delivered', (data) {
      log('${data} message deliverd');
    });
  }

  void _setupMessageListener() {
    socket.on('message:received', (data) {
      final message = _parseMessage(data);
      _messageStreamController.add(message);
    });
  }

  MessageEntity _parseMessage(dynamic data) {
    return MessageEntity(
      messageId: data['messageId'],
      sender: data['sender'],
      chat: data['chat'],
      content: data['content'],
    );
  }

  Stream<MessageEntity> get messageStream => _messageStreamController.stream;

  @override
  void disconnect() {
    socket.disconnect();
  }
}
