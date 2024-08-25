
import '../entity/chat.dart';
import '../entity/message.dart';

abstract class ChatRepository {
  Future<List<ChatEntity>> getChats();
  Future<ChatEntity>chatById(String chatId);
  Future<List<MessageEntity>>getChatMessages(String chatId);
  Future<ChatEntity>initiateChat(String userId);
  Future<void>deleteChat(String chatId);
 
}