import 'package:equatable/equatable.dart';
import '../../../auth/domain/entities/user_entity.dart';
import 'chat.dart';

class MessageEntity extends Equatable {
  final String messageId;
  final UserEntity sender;
  final ChatEntity chat;
  final String content;

  const MessageEntity({
    required this.messageId,
    required this.sender,
    required this.chat,
    required this.content,
  });

  // Method to convert MessageEntity to JSON
  Map<String, dynamic> toJson() {
    return {
      '_id': messageId,
      'sender': sender.toJson(),
      'chat': chat.toJson(),
      'content': content,
    };
  }

  factory MessageEntity.fromJson(Map<String, dynamic> json) {
    return MessageEntity(
      messageId: json['_id'],
      sender: UserEntity.fromJson(json['sender']),
      chat: ChatEntity.fromJson(json['chat']),
      content: json['content'],
    );
  }

  @override
  List<Object?> get props => [messageId, sender, chat, content];
}
