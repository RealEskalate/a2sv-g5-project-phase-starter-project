import '../../../auth/domain/entities/user_entity.dart';
import 'chat_entity.dart';

class MessageEntity {
  final String messageId;
  final ChatEntity chatEntity;
  final String content;
  final String type;
  final UserEntity sender;

  MessageEntity(
      {required this.messageId,
      required this.chatEntity,
      required this.content,
      required this.type,
      required this.sender});
}
