import 'dart:convert';
import 'dart:io';

import '../../../../core/constants/constants.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatEntity>> getChatRooms();
  Future<List<MessageEntity>> getMessages(String chatId);
  Future<void> createChatRoom(ChatEntity chat);
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
        await httpClient.getUrl(Uri.parse(AppData.getMssagesById(chatId)));
    final response = await request.close();

    if (response.statusCode == HttpStatus.ok) {
      final responseBody = await response.transform(utf8.decoder).join();
      final List<dynamic> data = json.decode(responseBody);
      return data.map((json) => MessageEntity.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load chat rooms');
    }
  }

  @override
  Future<void> createChatRoom(ChatEntity chat) async {
    // Implement API call to POST /api/v3/chats
  }
}
