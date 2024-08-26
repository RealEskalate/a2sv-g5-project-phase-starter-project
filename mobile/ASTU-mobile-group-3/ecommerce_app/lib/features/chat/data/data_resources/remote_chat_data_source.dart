import 'dart:convert';
import 'package:http/http.dart' as http;
import '../../../../core/constants/constants.dart';
import '../../../auth/data/data_source/auth_local_data_source.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatEntity>> getChatRooms();
  Future<List<MessageEntity>> getMessages(String chatId);
  Future<void> createChatRoom(String userId);
}

class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final http.Client httpClient;
  final AuthLocalDataSource authLocalDataSource;

  ChatRemoteDataSourceImpl(
      {required this.httpClient, required this.authLocalDataSource});

  @override
  Future<List<ChatEntity>> getChatRooms() async {
    final token = await authLocalDataSource.getToken();
    List<ChatEntity> chatRooms = [];

    final response = await httpClient.get(
      Uri.parse(AppData.chat),
      headers: {
        'Authorization': 'Bearer ${token.token}',
      },
    );
    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      data['data']
          .forEach((el) async => {chatRooms.add(ChatEntity.fromJson(el))});
      return chatRooms;
    } else {
      throw Exception('Failed to load chat rooms');
    }
  }

  @override
  Future<List<MessageEntity>> getMessages(String chatId) async {
    List<MessageEntity> messageLists = [];
    final token = await authLocalDataSource.getToken();

    final response = await httpClient
        .get(Uri.parse(AppData.getMssagesById(chatId)), headers: {
      'Authorization': 'Bearer ${token.token}',
    });

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      data['data'].forEach(
          (el) async => {messageLists.add(MessageEntity.fromJson(el))});
      return messageLists;
    } else {
      throw Exception('Failed to load messages');
    }
  }

  @override
  Future<void> createChatRoom(String userId) async {
    final token = await authLocalDataSource.getToken();
    final response = await httpClient.post(
      Uri.parse(AppData.chat),
      headers: {
        'Authorization': 'Bearer ${token.token}',
      },
      body: {'userId': userId},
    );

    if (response.statusCode != 201) {
      throw Exception('Failed to create chat room');
    }
  }
}
