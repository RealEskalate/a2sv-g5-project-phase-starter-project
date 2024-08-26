import 'package:e_commerce_app/features/auth/data/models/user_model.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity {
  const ChatModel({
    required String id,
    required UserModel sender,
    required UserModel receiver, 
  }) : super(
          id: id,
          sender: sender,
          reciever: receiver
        );

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      id: json['_id'],
      sender: UserModel.fromJson2(json['user2']), 
      receiver: UserModel.fromJson2(json['user1']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      '_id': id,
      'user1': (sender as UserModel).toJson(), 
      'user2': (reciever as UserModel).toJson(),
    };
  }
}
