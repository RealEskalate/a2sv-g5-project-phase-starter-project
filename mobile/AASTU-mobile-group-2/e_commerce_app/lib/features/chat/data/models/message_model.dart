import 'package:e_commerce_app/features/auth/data/models/user_model.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';

class MessageModel extends MessageEntity {
  MessageModel({
    required String messageId,
    required String contentType,
    required String chatId,
    required ChatModel chat,
    required UserModel sender,
    required String message,
  }) : super(
          messageId: messageId,
          contentType: contentType,
          chatId: chatId,
          chat: chat,
          sender: sender,
          message: message,
        );

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageId: json['_id'],
      contentType: json['type'], 
      message: json['content'],
      chatId: json['chat']['_id'], 
      chat: ChatModel.fromJson(json['chat']),
      sender: UserModel.fromJson(json['sender']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': messageId,
      'type': contentType, 
      'content': message,
      'chat': (chat as ChatModel).toJson(),
      'sender': (sender as UserModel).toJson(),
      'chatId': chatId,
    };
  }
}
