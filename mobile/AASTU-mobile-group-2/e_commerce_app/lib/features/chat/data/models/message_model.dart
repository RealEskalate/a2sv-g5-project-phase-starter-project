import 'package:e_commerce_app/features/auth/data/models/user_model.dart';
import 'package:e_commerce_app/features/chat/data/models/chat_model.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';

class MessageModel extends MessageEntity {
  MessageModel({
    required String messageId,
    required MessageContent content,
    required String chatId,
    required ChatModel chat,
    required UserModel sender,
  }) : super(
          messageId: messageId,
          content: content,
          chatId: chatId,
          chat: chat,
          sender: sender,
        );

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageId: json['_id'],
      content: MessageContent(
        json['content']['image'] ?? '',
        json['content']['video'] ?? '',
        json['content']['audio'] ?? '',
        json['content']['file'] ?? '',
        json['content']['text'] ?? '',
      ),
      chatId: json['chat']['_id'],
      chat: ChatModel.fromJson(json['chat']),
      sender: UserModel.fromJson(json['sender']),
    );
  }



  Map<String, dynamic> toJson() {
    final Map<String, dynamic> contentMap = {};
    if (content.text.isNotEmpty) {
      contentMap['text'] = content.text;
    }
    if (content.image.isNotEmpty) {
      contentMap['image'] = content.image;
    }
    if (content.video.isNotEmpty) {
      contentMap['video'] = content.video;
    }
    if (content.audio.isNotEmpty) {
      contentMap['audio'] = content.audio;
    }
    if (content.file.isNotEmpty) {
      contentMap['file'] = content.file;
    }

    return {
      '_id': messageId,
      'content': contentMap,
      'chat': (chat as ChatModel).toJson(),
      'sender': (sender as UserModel).toJson(),
      'chatId': chatId,
    };
  }
}
