import '../../../authentication/domain/entity/user.dart';
import '../entity/chat.dart';

abstract class ChatRepository {
  Future<List<Chat>>getMyChats();
  Future<Chat> intiateChat(String userId);
  Future<User> chatById(String chatId);
  Future<void> deleteMessage(String chatId);
  Future<List<Chat>>getChatMessage(String chatId);
  
}
