import '../../../authentication/data/model/user_model.dart';
import '../../domain/entities/chat.dart';

class ChatModel extends Chat {
  ChatModel({
    required super.id, 
    required super.user
  });

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id: json['_id'],
      user: UserModel.fromJson(json['user2']),
    );
  }
}