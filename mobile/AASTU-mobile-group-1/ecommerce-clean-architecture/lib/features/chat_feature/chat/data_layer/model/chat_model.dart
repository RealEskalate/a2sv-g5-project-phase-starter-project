import 'package:ecommerce/features/auth/data/model/user_model.dart';

import '../../domain/entity/chat.dart';

class ChatModel extends ChatEntity {
  ChatModel(
      {required super.chatId, required super.user1, required super.user2});

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      chatId: json['_id'],
      user1: UserModel.forSeller(json['user1']),
      user2: UserModel.forSeller(json['user2']),
    );
  }

  ChatModel.fromEntity(ChatEntity entity)
      : super(
            chatId: entity.chatId,
            user1: UserModel.fromEntity(entity.user1),
            user2: UserModel.fromEntity(entity.user2));

  Map<String, dynamic> toJson() => {
        '_id': chatId,
        'user1': UserModel.fromEntity(user1).toJsonForSeller(),
        'user2': UserModel.fromEntity(user2).toJsonForSeller(),
  };
}
