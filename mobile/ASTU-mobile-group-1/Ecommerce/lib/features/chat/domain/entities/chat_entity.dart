import '../../../auth/domain/entities/user_entity.dart';

class ChatEntity {
  final String chatId;
  final UserEntity user1;
  final UserEntity user2;

  const ChatEntity(
      {required this.chatId, required this.user1, required this.user2});
}
