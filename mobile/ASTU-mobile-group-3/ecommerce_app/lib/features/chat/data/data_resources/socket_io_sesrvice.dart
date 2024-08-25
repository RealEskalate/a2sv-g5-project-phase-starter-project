import 'dart:async';
import 'package:socket_io/socket_io.dart';
import '../../domain/entity/message.dart';

abstract class SocketIOService {
  Future<void> emitSendMessage(String chatId, MessageEntity message);
  Future<void> emitMessageDelivered(String messageId);
  Stream<MessageEntity> onMessageReceived();
}

class SocketIOServiceImpl implements SocketIOService {
  final Socket socket;
  final _messageStreamController = StreamController<MessageEntity>.broadcast();

  SocketIOServiceImpl({required this.socket}) {
    _setupListeners();
  }

  void _setupListeners() {
    socket.on('message:received', (data) {
      final message = MessageEntity.fromJson(data);
      _messageStreamController.add(message);
    });
  }

  @override
  Future<void> emitSendMessage(String chatId, MessageEntity message) async {
    socket.emit('send:message', {
      'chatId': chatId,
      'message': {},
    });
  }

  @override
  Future<void> emitMessageDelivered(String messageId) async {
    socket.emit('message:delivered', {'messageId': messageId});
  }

  @override
  Stream<MessageEntity> onMessageReceived() {
    return _messageStreamController.stream;
  }

  void dispose() {
    _messageStreamController.close();
  }
}
