import '../../../auth/data/models/auth_model.dart';
import '../../domain/entity/message_entity.dart';


class MessageModel extends MessageEntity {
  MessageModel({
    required String chatId,
    required MessageType type,
    required String content,
    String? timestamp,
    required UserModel sender,
  }) : super(
          chatId: chatId,
          type: type,
          content: content,
          timestamp: timestamp,
          sender: sender,
        );

  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      chatId: json['chatId'] as String,
      type: MessageType.values.firstWhere((e) => e.toString() == 'MessageType.${json['type']}'),
      content: json['content'] as String,
      timestamp: json['timestamp'] as String?,
      sender: UserModel.fromJson(json['sender']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'chatId': chatId,
      'type': type.toString().split('.').last,
      'content': content,
      'timestamp': timestamp,
      'sender': (sender as UserModel).toJson(),
    };
  }
}
