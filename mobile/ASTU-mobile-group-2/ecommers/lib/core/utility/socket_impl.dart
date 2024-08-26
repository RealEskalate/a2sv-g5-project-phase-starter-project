import 'dart:async';

import 'package:shared_preferences/shared_preferences.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

import 'global_message_part.dart';

class SocketService {
  late IO.Socket socket;
  late Completer<void> _completer;
  bool _isInitialized = false;
  String messageId = '';

  SocketService._() {
    _completer = Completer<void>();
  }

  static Future<SocketService> create() async {
    var service = SocketService._();
    await service._initializeSocket(); // No BuildContext needed here
    return service;
  }

  Future<void> _initializeSocket() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    String token = sharedPreferences.getString('key') ?? '';

    socket = IO.io(
      'https://g5-flutter-learning-path-be.onrender.com',
      <String, dynamic>{
        'transports': ['websocket'],
        'extraHeaders': {
          'Authorization': 'Bearer $token',
        }
      },
    );

    socket.on('connect', (_) {
      print('Connected to the socket server');
      if (!_completer.isCompleted) {
        _completer.complete();
        _isInitialized = true;
      }
    });

    socket.on('disconnect', (_) {
      print('Disconnected from the socket server');
    });

    socket.on('message:received', (data) {
      
      GlobalMessagePart.gloablMessage[data['_id']]?.add({
        'senderId': data['sender']['_id'],
        'content': data['content']
      });
      print('Message received by recipient: $data');
    });
  }

  Future<void> _ensureInitialized() async {
    if (!_isInitialized) {
      await _completer.future;
    }
  }

  Future<void> connect() async {
    await _ensureInitialized();
    socket.connect();
  }

  void disconnect() {
    socket.disconnect();
  }

  void sendMessage(String chatId, String message) async {
    await _ensureInitialized();
    Map<String, String> messageSend = {
      'chatId': chatId,
      'content': message,
      'type': 'text',
    };
    print(12344444);
    socket.emit('message:send', messageSend);
    print('Message sent: $messageSend');
  }

  void listen(String event, Function(dynamic) callback) async {
    await _ensureInitialized();
    socket.on(event, callback);
  }
}
