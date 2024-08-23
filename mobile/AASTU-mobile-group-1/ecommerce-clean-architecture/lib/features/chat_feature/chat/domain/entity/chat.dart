import 'package:ecommerce/features/auth/data/model/user_model.dart';

import '../../../../auth/domain/entities/user.dart';

class ChatEntity{
 final String chatId;
  final UserEntity user1;
  final UserEntity user2;
  const ChatEntity({
    required this.chatId,
    required this.user1,
    required this.user2,
  });
}