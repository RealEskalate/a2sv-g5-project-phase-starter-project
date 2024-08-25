import 'dart:convert';
import 'dart:io';

import '../../../../core/constants/constants.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatEntity>> getChatRooms();
  Future<List<MessageEntity>> getMessages(String chatId);
  Future<void> createChatRoom(ChatEntity chat);
  Future<void> sendMessage(String chatId, MessageEntity message);
  Future<void> acknowledgeMessageDelivery(String messageId);
  Stream<MessageEntity> onMessageReceived();
}

class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final HttpClient httpClient;

  ChatRemoteDataSourceImpl({required this.httpClient});

  @override
  Future<List<ChatEntity>> getChatRooms() async {
    final request = await httpClient.getUrl(Uri.parse(AppData.chat));
    final response = await request.close();

    if (response.statusCode == HttpStatus.ok) {
      final responseBody = await response.transform(utf8.decoder).join();
      final List<dynamic> data = json.decode(responseBody);
      return data.map((json) => ChatEntity.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load chat rooms');
    }
  }

  @override
  Future<List<MessageEntity>> getMessages(String chatId) async {
    final request =
        await httpClient.getUrl(Uri.parse('/api/v3/chats/$chatId/messages'));
    final response = await request.close();

    if (response.statusCode == HttpStatus.ok) {
      final responseBody = await response.transform(utf8.decoder).join();
      final List<dynamic> data = json.decode(responseBody);
      return data.map((json) => MessageEntity.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load messages');
    }
  }

  @override
  Future<void> createChatRoom(ChatEntity chat) async {
    final request = await httpClient.postUrl(Uri.parse('/api/v3/chats'));
    request.headers.contentType = ContentType.json;
    request.write(json.encode(chat.toJson()));
    final response = await request.close();

    if (response.statusCode != HttpStatus.created) {
      throw Exception('Failed to create chat room');
    }
  }

  @override
  Future<void> sendMessage(String chatId, MessageEntity message) async {
    final request =
        await httpClient.postUrl(Uri.parse('/api/v3/chats/$chatId/messages'));
    request.headers.contentType = ContentType.json;
    request.write(json.encode(message.toJson()));
    final response = await request.close();

    if (response.statusCode != HttpStatus.ok) {
      throw Exception('Failed to send message');
    }
  }

  @override
  Future<void> acknowledgeMessageDelivery(String messageId) async {
    final request = await httpClient
        .postUrl(Uri.parse('/api/v3/messages/$messageId/delivered'));
    final response = await request.close();

    if (response.statusCode != HttpStatus.ok) {
      throw Exception('Failed to acknowledge message delivery');
    }
  }

  @override
  Stream<MessageEntity> onMessageReceived() {
    // TODO: implement onMessageReceived
    throw UnimplementedError();
  }

  // @override
  // Stream<MessageEntity> onMessageReceived() {
  //   // Assuming WebSocket for real-time updates, modify as needed
  //   final websocket = WebSocket.connect('ws://your-websocket-url');
  //   return websocket.map((data) {
  //     final json = jsonDecode(data);
  //     return MessageEntity.fromJson(json);
  //   });
  // }
}
