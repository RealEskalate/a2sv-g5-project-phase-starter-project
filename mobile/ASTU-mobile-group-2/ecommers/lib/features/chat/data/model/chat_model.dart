import '../../domain/entity/chat_entity.dart';
import '../../domain/entity/message_entity.dart';

class ChatModel extends ChatEntity {
  ChatModel({
    required super.senderId,
    required super.senderName,
    required super.recieverId,
    required super.recieverName,
    required super.chatId,
    required super.messages

  });

  factory ChatModel.fromJson(Map<String, dynamic> json) => ChatModel(
        senderId: json['senderId'],
        senderName: json['senderName'],
        recieverId: json['recieverId'],
        recieverName: json['recieverName'],
        chatId: json['chatId'],
        messages: json['messages']
      );
    

  

  Map<String, dynamic> toJson() => {};

  ChatEntity toEntity() => ChatEntity(
        senderId: senderId,
        senderName: senderName,
        recieverId: recieverId,
        recieverName: recieverName,
        chatId: chatId,
        messages:messages

      );
}


class MessageModel extends MessageEntity {
  MessageModel({required super.messageId, required super.messages});

  factory MessageModel.fromJson(Map<String, dynamic> json) => MessageModel(
        messageId: json['messageId'],
        messages : json['messages']
      );
  MessageEntity toEntity() => MessageEntity(
    messageId: messageId,
    messages: messages,);
}