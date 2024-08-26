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
  String accessToken;
  late IO.Socket socket;

  final StreamController<MessageModel> _messageStreamController =
      StreamController<MessageModel>.broadcast();

  RemoteDataSourceImpl({required this.client, required this.accessToken}) {
    _initializeWebSocket();
  }

  @override
  Future<void> updateAccessToken(String token) async {
    accessToken = token;
  }

  @override
  Future<List<ChatModel>> getAllChats() async {
    try {
      final response = await client.get(
        Uri.parse(Urls.baseChat),
        headers: {
          'Authorization': 'Bearer $accessToken',
          'Content-Type': 'application/json',
        },
      );
        print(response.statusCode);
      if (response.statusCode == 200) {
        final result = json.decode(response.body)['data'];
        final List<ChatModel> answer = [];
        result.forEach((json) {
          answer.add(ChatModel.fromJson(json));
        });
        // print(answer);
        return answer;
      } else {
        // print(response.body.toString);
        throw Exception();
      }
    } catch (e) {
      // print(e.toString());
      throw ServerException();
    }
  }

  @override
  Future<ChatModel> getChatById(String chatId) async {
    try {
      final response = await http.get(
        Uri.parse(Urls.getChatById(chatId)),
        headers: {
          'Authorization': 'Bearer $accessToken',
          'Content-Type': 'application/json',
        },
      );
      if (response.statusCode == 200) {
        final result = json.decode(response.body)['data'];
        return ChatModel.fromJson(result);
      } else {
        throw Exception();
      }
    } catch (e) {
      throw Exception(e.toString());
    }
  }

  @override
  Future<ChatModel> initiateChat(String recieverId) async {
    try {
      final response = await client.post(Uri.parse(Urls.baseChat),
          headers: {
            'Authorization': 'Bearer $accessToken',
            'Content-Type': 'application/json',
          },
          body: json.encode({'userId': recieverId}));
      if (response.statusCode == 201) {
        final result = json.decode(response.body)['data'];
        return ChatModel.fromJson(result);
      } else {
        throw Exception();
      }
    } catch (e) {
      throw Exception(e.toString());
    }
  }

  Future<List<MessageModel>> getChatMessages(String chatId) async {
    try {
      final response = await client.get(
        Uri.parse(Urls.getChatMessages(chatId)),
        headers: {
          'Authorization': 'Bearer $accessToken',
          'Content-Type': 'application/json',
        },
      );
      if (response.statusCode == 200) {
        final result = json.decode(response.body)['data'];

        final List<MessageModel> answer = [];
        result.forEach((json) {
          answer.add(MessageModel.fromJson(json));
        });
        return answer;
      } else {
        throw Exception();
      }
    } catch (e) {
      throw Exception(e.toString());
    }
  }

  @override
  Future<void> sendMessage(String chatId, String message, String type) async {
    socket.emit('message:send', {
      'chatId': chatId,
      'content': message,
      'type': type,
    });
    log('message sent');

    socket.on('message:delivered', (data) {
      log('Message delivered: $data');
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

  @override
  Future<bool> deleteChat(String chatId) {
    // TODO: implement deleteChat
    throw UnimplementedError();
  }
}
