import 'package:equatable/equatable.dart';

enum MessageType { text, image, video, audio, file }

class MessageContent {
  final String image;
  final String video;
  final String audio;
  final String file;
  final String text;

  MessageContent(this.image, this.video, this.audio, this.file, this.text);
}

class MessageEntity extends Equatable {
  final MessageType type;
  final MessageContent content;
  final String senderId;
  final String chatId;
  final String? messageId;

  MessageEntity(this.messageId, {
    required this.type,
    required this.content,
    required this.senderId, 
    required this.chatId,
  });

  @override
  List<Object?> get props => [type, content, senderId, chatId, messageId];
}