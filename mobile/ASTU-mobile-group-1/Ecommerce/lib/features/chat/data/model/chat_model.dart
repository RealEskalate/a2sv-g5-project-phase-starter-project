import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity {
  const ChatModel({
    required String chatId,
    required UserEntity user1,
    required UserEntity user2,
  }) : super(chatId: chatId, user1: user1, user2: user2);

  // Factory constructor to create a ChatModel from a JSON map
  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      chatId: json['_id'] as String,
      user1: userEntityFromJson(json['user1']),
      user2: userEntityFromJson(json['user2']),
    );
  }

  // Static method to convert a ChatEntity to JSON
  Map<String, dynamic> toJson() {
    return {
      '_id': chatId,
      'user1': userEntityToJson(user1),
      'user2': userEntityToJson(user2),
    };
  }
}

// Function to convert UserEntity to JSON
Map<String, dynamic> userEntityToJson(UserEntity user) {
  return {
    '_id': user.id.toString(),
    'name': user.name.toString(),
    'email': user.email.toString(),
  };
}

// Function to convert JSON to UserEntity
UserEntity userEntityFromJson(Map<String, dynamic> json) {
  return UserEntity(
    id: json['_id'].toString(),
    name: json['name'].toString(),
    email: json['email'].toString(),
  );
}
