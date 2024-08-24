



import 'package:equatable/equatable.dart';

import '../../../user/domain/entities/user.dart';
import 'chat_entity.dart';

class MessageEntity extends Equatable{
  final String messageId;
  final User sender;
  
  final String content;
  final ChatEntity chat;
  final String type;

  MessageEntity({
    required this.messageId,
    required this.sender,
  
    required this.content,
    required this.chat,
    required this.type

  });
  @override
  
  List<Object?> get props => [messageId,sender,content,chat, type];
  
}