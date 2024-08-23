import '../../../../auth/data/model/user_model.dart';
import '../../domain/entity/message.dart';

class MessageModel extends Message {
  MessageModel(
      {required super.sender,
      required super.chatId,
      required super.type,
      required super.content});

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      sender: UserModel.fromJson(json['sender']),
      chatId: json['chatId'],
      type: json['type'],
      content: json['content'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'sender': UserModel.fromEntity(sender).toJson(),
      'chatId': chatId,
      'type': type,
      'content': content,
    };
  }

  factory MessageModel.fromEntity(Message entity) {
    return MessageModel(
      sender: UserModel.fromEntity(entity.sender),
      chatId: entity.chatId,
      type: entity.type,
      content: entity.content,
    );
  }
}
