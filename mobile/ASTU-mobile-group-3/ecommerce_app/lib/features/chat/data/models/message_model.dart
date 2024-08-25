import '../../../auth/data/model/user_model.dart';
import '../../domain/entity/message.dart';
import 'chat_model.dart';

class MessageModel extends MessageEntity {
  const MessageModel({
    required super.messageId,
    required super.sender,
    required super.chat,
    required super.content,
  });

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageId: json['messageId'],
      sender: UserModel.fromSellerJson(
        (json['sender']),
      ),
      chat: ChatModel.fromJson(json['chat']),
      content: json['content'],
    );
  }
}
