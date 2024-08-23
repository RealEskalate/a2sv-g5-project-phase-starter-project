import '../../../auth/data/models/auth_model.dart';
import '../../domain/entity/chat_entity.dart';

class ChatModel extends ChatEntity {
  ChatModel({
    required String chatId,
    required UserModel sender,
    required UserModel receiver,
  }) : super(
          chatId: chatId,
          sender: sender,
          receiver: receiver,
        );

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      chatId: json['chatId'] as String,
      sender: UserModel.fromJson(json['sender']),
      receiver: UserModel.fromJson(json['receiver']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'chatId': chatId,
      'sender': (sender as UserModel).toJson(),
      'receiver': (receiver as UserModel).toJson(),
    };
  }
}
