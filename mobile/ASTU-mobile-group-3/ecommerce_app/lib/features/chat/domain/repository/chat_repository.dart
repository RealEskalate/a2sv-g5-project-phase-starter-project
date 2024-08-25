
import '../../../auth/domain/entities/user_entity.dart';
import '../entity/chat.dart';
import '../entity/message.dart';

abstract class ChatRepository {
  Future<List<ChatEntity>> getChatRooms();
  Future<List<MessageEntity>> getMessagesForChat(String chatId);
  Future<ChatEntity> createChatRoom(UserEntity user1, UserEntity user2);
  Future<void> sendMessage(String chatId, String message);
  Future<void> acknowledgeMessageDelivery(String messageId);
}

