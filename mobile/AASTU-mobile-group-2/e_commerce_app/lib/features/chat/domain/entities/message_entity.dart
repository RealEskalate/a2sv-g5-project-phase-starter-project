import 'package:e_commerce_app/features/auth/domain/entities/user.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:equatable/equatable.dart';

class MessageContent {
  final String image;
  final String video;
  final String audio;
  final String file;
  final String text;

  MessageContent(this.image, this.video, this.audio, this.file, this.text);
}

class MessageEntity extends Equatable {
  final String chatId;
  final User sender;
  final ChatEntity chat;
  final MessageContent content;
  final String messageId;

  MessageEntity({
      required this.messageId,
      required this.content,
      required this.chatId,
      required this.chat,
      required this.sender,
  });

  @override
  List<Object?> get props => [chat, content, sender, chatId, messageId];
}