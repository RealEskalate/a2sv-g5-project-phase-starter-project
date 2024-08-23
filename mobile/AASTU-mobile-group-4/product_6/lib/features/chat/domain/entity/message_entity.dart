import 'package:equatable/equatable.dart';

import '../../../auth/domain/entity/auth_entity.dart';

enum MessageType { text, imageUrl, voiceUrl }

class MessageEntity extends Equatable {
  final String chatId;
  final String? timestamp;
  final MessageType type;
  final String content;
  final UserEntity sender;


  MessageEntity({
    required this.chatId,
    required this.type,
    required this.content,
    this.timestamp,
    required this.sender,
  });

  @override
  List<Object?> get props => [chatId, type, content, timestamp];
}
