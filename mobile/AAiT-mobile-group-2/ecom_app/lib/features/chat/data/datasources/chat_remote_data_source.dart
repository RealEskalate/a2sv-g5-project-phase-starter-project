import 'dart:convert';
import 'dart:io';

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';
import '../../../../core/network/custom_client.dart';
import '../models/message_model.dart';
import '../models/user_chat_model.dart';

abstract class ChatRemoteDataSource {
  Future<List<UserChatModel>> getChats(String email);
  Future<List<MessageModel>> getMessages(String chatID);
  Future<UserChatModel> initiateChat(String UserID);
}

class ChatRemoteDataSourceImpl extends ChatRemoteDataSource {
  final CustomHttpClient client;

  ChatRemoteDataSourceImpl({required this.client});

  @override
  Future<List<UserChatModel>> getChats(String email) async {
    var uri = ''; // Replace with your actual endpoint
    try {
      final response = await client.get(uri);
      if (response.statusCode == 200) {
        // Parse the response body
        final responseBody = json.decode(response.body);
        final data = responseBody['data'] as List;

        // Filter the data based on email and map to UserChatModel
        final filteredChats = data.map<UserChatModel>((chat) {
          // Determine the other user
          final user1 = chat['user1'];
          final user2 = chat['user2'];
          final otherUser = user1['email'] != email ? user1 : user2;

          return UserChatModel(
            ChatID: chat['_id'],
            name: otherUser['name'],
          );
        }).toList();

        return filteredChats
            .cast<UserChatModel>(); // Cast the list to the correct type
      } else {
        throw ServerException();
      }
    } catch (e) {
      throw Exception('Failed to load chats: $e');
    }
  }

  @override
  Future<List<MessageModel>> getMessages(String chatID) async {
    var uri =
        'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats/$chatID/messages';

    try {
      final response = await client.get(uri);

      if (response.statusCode == 200) {
        final responseBody = json.decode(response.body);
        final data = responseBody['data'] as List;

        // Map the JSON data to a list of MessageModel
        final messages = data.map<MessageModel>((messageJson) {
          return MessageModel.fromJson(messageJson);
        }).toList();

        return messages;
      } else {
        throw ServerException();
      }
    } catch (e) {
      throw Exception('Failed to load messages: $e'); // Catch any other errors
    }
  }

  @override
  Future<UserChatModel> initiateChat(String UserID) async {
    var uri = 'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
    final jsonBody = jsonEncode({'userId': UserID});
    try {
      final response = await client.post(uri, body: jsonBody);

      if (response.statusCode == 201) {
        final responseBody = json.decode(response.body);
        final data = responseBody['data'];

        // Create UserChatModel with ChatID and empty name since we dont need the name
        final userChat = UserChatModel(
          ChatID: data['_id'],
          name: '',
        );

        return userChat;
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException(ErrorMessages.socketError);
    }
  }
}
