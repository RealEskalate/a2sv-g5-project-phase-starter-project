

import '../../../authentication/data/model/user_model.dart';
import '../../domain/entity/message.dart';
import 'chat_model.dart';

class MessageModel extends Message {
  final String id;  
  final UserModel sender;
  final ChatModel chat;
  final String type;
  final dynamic content;
  

  MessageModel({
    required this.id,
    required this.sender,
    required this.chat,
    required this.type,
    required this.content,
  }) : super(id: id, sender: sender, chat: chat, type: type, content: content);

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      id: json['_id'],
      chat: ChatModel.fromJson(json['chat']),
      sender: UserModel.fromJson(json['sender']),
      type: json['type'],
      content: json['content'],
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'chat': chat.toJson(),
      'sender': sender.toJson(),
      'type': type,
      'content': content,
    };
  }
}