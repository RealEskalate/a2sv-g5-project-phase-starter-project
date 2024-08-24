import '../../domain/entity/chat_entity.dart';

class ChatModel extends ChatEntity {
  const ChatModel({
    required super.senderId,
    required super.senderName,
    required super.recieverId,
    required super.recieverName,
    required super.chatId,
    required super.messages

  });

  factory ChatModel.fromJson(Map<String, dynamic> json) => ChatModel(
        senderId: json['_id'],
        senderName: json['name'],
        recieverId: json['email'],
        recieverName: json['email'],
        chatId: json['email'],
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
