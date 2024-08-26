import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entities/user_data.dart';

enum MessageType {
  text,
  image,
}

class Message extends Equatable {
  final String? chatId;
  final UserEntity? sender;
  final String content;
  final MessageType type;

  const Message({
    this.chatId,
    this.sender,
    required this.content,
    required this.type,
  });
  
  @override
  List<Object?> get props => [chatId, sender, content, type];
}