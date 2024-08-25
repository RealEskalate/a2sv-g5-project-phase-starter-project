import '../../../auth/data/model/user_model.dart';
import '../../domain/entity/chat.dart';

class ChatModel extends ChatEntity {
  const ChatModel({
    required String chatId,
    required UserModel user1, // Use UserModel directly
    required UserModel user2, // Use UserModel directly
  }) : super(
          chatId: chatId,
          user1: user1,
          user2: user2,
        );

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      chatId: json['_id'],
      user1: UserModel.fromJson(json['user1']),
      user2: UserModel.fromJson(json['user2']),
    );
  }

//   Map<String, dynamic> toJson() {
//     return {
//       '_id': chatId, // Adjusting key to match your JSON
//       'user1': user1.toJson(), // Serialize UserModel directly
//       'user2': user2.toJson(), // Serialize UserModel directly
//     };
//   }
}
