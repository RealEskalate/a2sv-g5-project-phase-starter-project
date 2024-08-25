import '../entity/chat.dart';
import '../entity/message.dart';

abstract class ChatRepository {
  Future<List<Chat>>getMyChats();
  Future<Chat> intiateChat(String userId);
  Future<Chat> chatById(String chatId);
  Future<void> deleteMessage(String chatId);
  Future<List<Message>>getChatMessage(String chatId);
  
}
