import 'package:ecommerce/features/auth/data/model/user_model.dart';

class ChatEntity{
 final String chatId;
  final UserModel user1;
  final UserModel user2;
  const ChatEntity({
    required this.chatId,
    required this.user1,
    required this.user2,
  });
}