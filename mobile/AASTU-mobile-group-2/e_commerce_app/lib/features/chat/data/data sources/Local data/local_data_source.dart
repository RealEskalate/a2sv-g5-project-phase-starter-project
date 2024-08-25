import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/features/chat/data/data%20sources/Local%20data/local_contrat.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';
import 'package:e_commerce_app/features/chat/data/models/message_model.dart';
import 'dart:convert';

class LocalDataSource implements LocalContrat {
  static const String _chatsKey = 'chats';
  static const String _messagesKey = 'messages';

  @override
  Future<void> cacheGetChatByIdLocal(ChatModel chat) async {
    final prefs = await SharedPreferences.getInstance();
    List<String> encodedChats = prefs.getStringList(_chatsKey) ?? [];

    encodedChats.removeWhere((c) {
      final chatData = json.decode(c) as Map<String, dynamic>;
      return chatData['_id'] == chat.id;
    });

    encodedChats.add(json.encode(chat.toJson()));
    await prefs.setStringList(_chatsKey, encodedChats);
  }

  @override
  Future<Either<Failure, List<MessageModel>>> getChatByIdLocal(String userId) async {
    try {
      final prefs = await SharedPreferences.getInstance();
      List<String>? encodedMessages = prefs.getStringList(_messagesKey);

      if (encodedMessages == null) {
        return Right([]);
      }

      List<MessageModel> messages = encodedMessages
          .map((encodedMessage) => MessageModel.fromJson(json.decode(encodedMessage)))
          .where((message) => message.chatId == userId)
          .toList();

      return Right(messages);
    } catch (e) {
      return Left(Failure('Failed to retrieve messages by chat ID'));
    }
  }

  @override
  Future<void> cacheGetAllChatsLocal(List<ChatModel> chats) async {
    final prefs = await SharedPreferences.getInstance();
    List<String> encodedChats = chats.map((chat) => json.encode(chat.toJson())).toList();
    await prefs.setStringList(_chatsKey, encodedChats);
  }

  @override
  Future<Either<Failure, List<ChatModel>>> getAllChatLocal() async {
    try {
      final prefs = await SharedPreferences.getInstance();
      List<String>? encodedChats = prefs.getStringList(_chatsKey);

      if (encodedChats == null) {
        return Right([]);
      }

      List<ChatModel> chats = encodedChats
          .map((encodedChat) => ChatModel.fromJson(json.decode(encodedChat)))
          .toList();

      return Right(chats);
    } catch (e) {
      return Left(Failure('Failed to retrieve all chats'));
    }
  }
}
