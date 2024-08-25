

import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';
import '../../domain/repository/chat_repository.dart';

class ChatRepositoryImpl implements ChatRepository {
  @override
  Future<ChatEntity> chatById(String chatId) {
    // TODO: implement chatById
    throw UnimplementedError();
  }

  @override
  Future<void> deleteChat(String chatId) {
    // TODO: implement deleteChat
    throw UnimplementedError();
  }

  @override
  Future<List<MessageEntity>> getChatMessages(String chatId) {
    // TODO: implement getChatMessages
    throw UnimplementedError();
  }

  @override
  Future<List<ChatEntity>> getChats() {
    // TODO: implement getChats
    throw UnimplementedError();
  }

  @override
  Future<ChatEntity> initiateChat(String userId) {
    // TODO: implement initiateChat
    throw UnimplementedError();
  }
 
}