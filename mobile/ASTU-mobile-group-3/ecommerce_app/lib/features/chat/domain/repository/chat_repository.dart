import '../entity/chat.dart';
import '../entity/message.dart';

abstract class ChatRepository {
  Future<List<ChatEntity>> retrieveChatRooms();
  Future<List<MessageEntity>> retrieveMessages(String chatId);
  Future<void> createChatRoom(ChatEntity chat);
  Future<void> sendMessage(String chatId, MessageEntity message);
  Future<void> acknowledgeMessageDelivery(String messageId);
  Stream<MessageEntity> onMessageReceived();
}
