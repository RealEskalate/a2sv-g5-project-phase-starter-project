import '../entity/chat.dart';
import '../entity/message.dart';

abstract class ChatRepository {
  Future<List<ChatEntity>> retrieveChatRooms();
  Future<List<MessageEntity>> retrieveMessages(String chatId);
  Future<void> createChatRoom(String userId);
  Future<void> sendMessage(String chatId, String content, String type);
  Future<void> acknowledgeMessageDelivery(String messageId);
  Stream<MessageEntity> onMessageReceived();
}