import '../../../authentication/data/model/user_model.dart';
import '../../domain/entities/chat.dart';

class ChatModel extends Chat {
  ChatModel({
    required super.id, 
    required super.user1,
    required super.user2,
  });

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id: json['_id'],
      user1: UserModel.fromJson(json['user1']),
      user2: UserModel.fromJson(json['user2']),
    );
  }
}