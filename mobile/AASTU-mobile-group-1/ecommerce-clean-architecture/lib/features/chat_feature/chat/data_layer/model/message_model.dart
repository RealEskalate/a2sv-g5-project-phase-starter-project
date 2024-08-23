import '../../../../auth/data/model/user_model.dart';
import '../../domain/entity/message.dart';
import 'chat_model.dart';

class MessageModel extends Message {
  MessageModel(
      {required super.messageId,
      required super.sender,
      required super.chat,
      required super.type,
      required super.content});

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageId: json['_id'],
      sender: UserModel.forSeller(json['sender']),
      chat: ChatModel.fromJson(json['chat']),
      type: json['type'],
      content: json['content'],
    );
  }
  
  factory MessageModel.fromEntity(Message entity) {
    return MessageModel(
      messageId: entity.messageId,
      sender: UserModel.fromEntity(entity.sender),
      chat:ChatModel.fromEntity(entity.chat),
      type: entity.type,
      content: entity.content,
    );
  }

  Map<String, dynamic> toJson() => {
        '_id': messageId,
        'sender': UserModel.fromEntity(sender).toJsonForSeller(),
        'chat': ChatModel.fromEntity(chat).toJson(),
        'type': type,
        'content': content,
      };
}
