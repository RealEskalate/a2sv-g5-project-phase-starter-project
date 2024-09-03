import 'dart:convert';

import 'package:flutter/material.dart';
import 'package:socket_io_client/socket_io_client.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/network/http.dart';
import '../models/chat_model.dart';
import '../models/message_model.dart';
import 'stream_socket.dart';

abstract class ChatDataSource {
  Future<List<ChatModel>> getChats();
  Future<Stream<MessageModel>> getChat(String chatId);
  Future<ChatModel> createChat(String userId);
  Future<void> deleteChat(String chatId);
  Future<void> sendMessage(MessageModel message);
}

class ChatDataSourceImpl implements ChatDataSource {
  final CustomHttp client;

  ChatDataSourceImpl({required this.client});

  @override
  Future<ChatModel> createChat(String userId) async {
    final response = await client.post(Uri.parse('${Urls3.baseUrl}/chats'), body: json.encode({
      'userId':userId,
    }));
    if (response.statusCode == 201) {
      return ChatModel.fromJson(jsonDecode(response.body)['data']);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> deleteChat(String chatId) async {
    final response = await client.delete(Uri.parse('${Urls3.baseUrl}/chats/$chatId'));
    if (response.statusCode != 200) {
      throw ServerException();
    }
  }

  @override
  Future<Stream<MessageModel>> getChat(String chatId) async {
    final stream = StreamSocket();
    var response = await client.get(Uri.parse('${Urls3.baseUrl}/chats/$chatId/messages'));
    if (response.statusCode == 200) {
      List<dynamic> lis = json.decode(response.body)['data'];
      for (var action in lis) {
        stream.addResponse(MessageModel.fromJson(action));
      }
    } else {
      throw ServerException(message: response.body);
    }
    client.socket.connect();
    client.socket.onConnect((_) {
      debugPrint('connected');
    });
    client.socket.on('message:recieved', (data) {
      stream.addResponse(MessageModel.fromJson(data));
    });
    client.socket.on('message:delivered', (data) {
      stream.addResponse(MessageModel.fromJson(data));
    });
    return stream.getResponse;
  }

  @override
  Future<List<ChatModel>> getChats() async {
    final response = await client.get(Uri.parse('${Urls3.baseUrl}/chats'));
    if (response.statusCode == 200) {
      List<dynamic> lis = json.decode(response.body)['data'];
      return lis.map((element) => ChatModel.fromJson(element)).toList();
    } else {
      throw ServerException(message: 'Server Error');
    }
  }

  @override
  Future<void> sendMessage(MessageModel message) async {
    client.socket.emit('message:send', message.toJson());
  }

}
