import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_entity.dart';
import 'chat_entity.dart';

class MessageEntity extends Equatable{
  final String messageId;
  final ChatEntity chatEntity;
  final String content;
  final String type;
  final UserEntity sender;

  const MessageEntity(
      {required this.messageId,
      required this.chatEntity,
      required this.content,
      required this.type,
      required this.sender});
      
        @override
        List<Object?> get props => [messageId,chatEntity,content,type,sender];
}
  