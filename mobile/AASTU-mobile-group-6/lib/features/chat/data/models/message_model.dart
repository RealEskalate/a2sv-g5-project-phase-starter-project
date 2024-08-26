import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import '../../domain/entities/message.dart';


class MessageModel extends Message {

  MessageModel({
      required super.messageid,
      required super.uniqueChat,
      required super.sender,
      required super.type,
      required super.content,
  });

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageid: json['_id'],
      uniqueChat: ChatModel.fromJson(json['chat']),
      sender: UserModel.fromJson(json['sender']),
      type: json['type'] ?? 'text',
      content: json['content'],
    );
  }

  @override
  Map<String, dynamic> toJson() {
    return {
      'id': messageid,
      'chat': uniqueChat.toJson(),
      'sender': sender.toJson(),
      'type': 'text',
      'content': content,
    };
  }
}