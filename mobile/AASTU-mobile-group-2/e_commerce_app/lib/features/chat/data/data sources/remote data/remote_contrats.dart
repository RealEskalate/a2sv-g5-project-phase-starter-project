import 'package:e_commerce_app/features/chat/data/models/message_model.dart';

import '../../models/chat_model.dart';

abstract class ChatRemoteDataSource {
  Future<List<ChatModel>> getAllChat();
  Future<List<MessageModel>> getMessagesById(String chatId);
  Future<ChatModel> createChatById(String sellerId);
  Future<bool> deleteChatById(String chatId);
  Future<bool> sendMessage(String chatId, String message, String content);
  

}


