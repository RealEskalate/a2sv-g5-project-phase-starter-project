import '../../../authentication/data/model/user_model.dart';
import '../../domain/entities/message.dart';

class MessageModel extends Message {
  const MessageModel({
    super.chatId,
    super.sender,
    required super.content,
    required super.type,
  });

  static MessageType toMessageType(String type) {
    switch (type) {
      case 'text':
        return MessageType.text;
      case 'image':
        return MessageType.image;
      default:
        return MessageType.text;
    }
  }
  
  static String toTypeString(MessageType type) {
    switch (type) {
      case MessageType.text:
        return 'text';
      case MessageType.image:
        return 'image';
      default:
        return 'text';
    }
  }
  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      content: json['content'],
      sender: UserModel.fromJson(json['sender']),
      type: toMessageType(json['type']),
    );
  }

  factory MessageModel.fromEntity(Message message) {
    return MessageModel(
      content: message.content,
      sender: message.sender,
      type: message.type,
      chatId: message.chatId
    );
  }


  Map<String, dynamic> toJson() {
    return {
      'chatId': chatId,
      'content': content,
      'type': toTypeString(type),
    };
  } 

}