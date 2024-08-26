import 'dart:async';
import 'dart:convert';
import 'dart:developer';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../../../../core/constants/constants.dart';
import '../../../../../core/error/exception.dart';
import '../../model/chat_model.dart';
import '../../model/message_model.dart';
import 'remote_data_source.dart';

class RemoteDataSourceImpl extends RemoteDataSource {
  final http.Client client;
  final String accessToken;
  late IO.Socket socket;

  final StreamController<MessageModel> _messageStreamController =
      StreamController<MessageModel>.broadcast();

  RemoteDataSourceImpl({required this.client, required this.accessToken}) {
    _initializeWebSocket();
  }

  @override
  Future<bool> deleteChat(String chatId) async {
    throw UnimplementedError();
  }

  @override
  Future<List<ChatModel>> getAllChats() async {
    throw UnimplementedError();
  }

  @override
  Future<ChatModel> getChatById(String chatId) async {
    throw UnimplementedError();
  }

  @override
  Future<ChatModel> initiateChat(String receiverId) async {
    throw UnimplementedError();
  }

  @override
  Future<List<MessageModel>> getChatMessages(String chatId) async {
    throw UnimplementedError();
  }

  @override
  Future<void> sendMessage(String chatId, String message, String type) async {
    socket.emit('message:send', {
      'chatId': chatId,
      'content': message,
      'type': type,
    });
  }

  @override
  Stream<MessageModel> getMessages() {
    return _messageStreamController.stream;
  }

  void _initializeWebSocket() {
    socket = IO.io(
      'https://g5-flutter-learning-path-be.onrender.com',
      <String, dynamic>{
        'transports': ['websocket'],
        'extraHeaders': {'Authorization': 'Bearer $accessToken'}
      },
    );

    socket.connect();

    socket.onConnect((_) {
      log('Connected to the socket server');
    });

    socket.on('message:received', (data) {
      final message = MessageModel.fromJson(data);
      _messageStreamController.add(message);
    });

    socket.on('message:delivered', (data) {
      log('Message delivered: $data');
    });

    socket.onDisconnect((_) {
      dispose();
    });
  }

  void dispose() {
    _messageStreamController.close();
    socket.disconnect();
  }
}
