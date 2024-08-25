
import 'message_entity.dart';

class ChatEntity {
  final String senderId;
  final String senderName;
  final String recieverId;
  final String recieverName;
  final String chatId;
  final MessageEntity messages;

  

  ChatEntity({
    required this.senderId,
    required this.senderName,
    required this.recieverId,
    required this.recieverName,
    required this.chatId,
    required this.messages

  });

}



