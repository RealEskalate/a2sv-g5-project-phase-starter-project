import 'dart:convert';
import 'package:http/http.dart' as http;
import '../../../../core/constants/constants.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatEntity>> getChatRooms();
  Future<List<MessageEntity>> getMessages(String chatId);
  Future<void> createChatRoom(ChatEntity chat);
}

class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final http.Client httpClient;

  ChatRemoteDataSourceImpl({required this.httpClient});

  @override
  Future<List<ChatEntity>> getChatRooms() async {
    final response = await httpClient.get(Uri.parse(AppData.chat));

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body) as List;
      return data.map((json) => ChatEntity.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load chat rooms');
    }
  }

  @override
  Future<List<MessageEntity>> getMessages(String chatId) async {
    final response =
        await httpClient.get(Uri.parse(AppData.getMssagesById(chatId)));

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body) as List;
      return data.map((json) => MessageEntity.fromJson(json)).toList();
    } else {
      throw Exception('Failed to load messages');
    }
  }

  @override
  Future<void> createChatRoom(ChatEntity chat) async {
    // final response = await httpClient.post(
    //   Uri.parse(AppData.createChatRoom),
    //   headers: {
    //     'Content-Type': 'application/json',
    //   },
    //   body: jsonEncode(chat.toJson()),
    // );

    // if (response.statusCode != 201) {
    //   throw Exception('Failed to create chat room');
    // }
  }
}
