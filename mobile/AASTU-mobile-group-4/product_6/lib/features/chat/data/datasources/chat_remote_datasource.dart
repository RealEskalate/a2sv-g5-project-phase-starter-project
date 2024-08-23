import 'dart:convert';
import 'package:http/http.dart' as http;

import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exception.dart';
import '../models/chat_models.dart';


abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> getAllChats(String token);
  Future<ChatModel> getChatById(String chatId, String token);
  Future<ChatModel> initiateChat(String userId, String token);
  Future<void> deleteChat(String chatId, String token);
}

class ChatRemoteDataSourceImpl implements ChatRemoteDataSource {
  final http.Client client;

  ChatRemoteDataSourceImpl({required this.client});

  @override
  Future<List<ChatModel>> getAllChats(String token) async {
    final url = Uri.parse('${Urls.chattUrl}');
    final response = await client.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body) as List;
      return data.map((chat) => ChatModel.fromJson(chat)).toList();
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ChatModel> getChatById(String chatId, String token) async {
    final url = Uri.parse('${Urls.chattUrl}/$chatId');
    final response = await client.get(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return ChatModel.fromJson(data);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<ChatModel> initiateChat(String userId, String token) async {
    final url = Uri.parse('${Urls.chattUrl}');
    final response = await client.post(
      url,
      headers: {
        'Authorization': 'Bearer $token',
        'Content-Type': 'application/json',
      },
      body: jsonEncode({'userId': userId}),
    );

    if (response.statusCode == 200) {
      final data = jsonDecode(response.body);
      return ChatModel.fromJson(data);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> deleteChat(String chatId, String token) async {
    final url = Uri.parse('${Urls.chattUrl}/$chatId');
    final response = await client.delete(
      url,
      headers: {'Authorization': 'Bearer $token'},
    );

    if (response.statusCode != 200) {
      throw ServerException();
    }
  }
}
