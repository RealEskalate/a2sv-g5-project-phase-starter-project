import 'package:equatable/equatable.dart';

import '../../../../auth/domain/entities/user.dart';
import 'chat.dart';

class Message extends Equatable {
  final String messageId;
  final UserEntity sender;
  final ChatEntity chat;
  final String type;
  final String content;

  Message(
      {required this.messageId,
      required this.sender,
      required this.chat,
      required this.type,
      required this.content});
  @override
  List<Object?> get props => [messageId, sender, chat, type, content];
}
