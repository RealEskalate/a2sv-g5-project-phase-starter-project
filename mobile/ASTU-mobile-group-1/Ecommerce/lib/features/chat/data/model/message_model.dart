import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';

class MessageModel extends MessageEntity {
  const MessageModel({
    required String messageId,
    required ChatEntity chat,
    required String content,
    required UserEntity sender,
    required String type,
  }) : super(
          messageId: messageId,
          chatEntity: chat,
          content: content,
          sender: sender,
          type: type,
        );

  // Factory constructor to create a MessageModel from a JSON map
  factory MessageModel.fromJson(Map<String, dynamic> json) {
    return MessageModel(
      messageId: json['_id'] as String,
      chat: chatEntityFromJson(json['chat'] as Map<String, dynamic>),
      content: json['content'] as String,
      sender: userEntityFromJson(json['sender'] as Map<String, dynamic>),
      type: json['type'] as String,
    );
  }

  // Method to convert a MessageModel to a JSON map
  Map<String, dynamic> toJson() {
    return {
      '_id': messageId,
      'chat': chatEntityToJson(chatEntity),
      'content': content,
      'sender': userEntityToJson(sender),
      'type': type,
    };
  }
}

// Function to convert JSON to ChatEntity
ChatEntity chatEntityFromJson(Map<String, dynamic> json) {
  return ChatEntity(
    chatId: json['_id'] as String,
    user1: userEntityFromJson(json['user1'] as Map<String, dynamic>),
    user2: userEntityFromJson(json['user2'] as Map<String, dynamic>),
  );
}

// Function to convert ChatEntity to JSON
Map<String, dynamic> chatEntityToJson(ChatEntity entity) {
  return {
    '_id': entity.chatId,
    'user1': userEntityToJson(entity.user1),
    'user2': userEntityToJson(entity.user2),
  };
}

// Function to convert UserEntity to JSON
Map<String, dynamic> userEntityToJson(UserEntity user) {
  return {
    '_id': user.id,
    'name': user.name,
    'email': user.email,
  };
}

// Function to convert JSON to UserEntity
UserEntity userEntityFromJson(Map<String, dynamic> json) {
  return UserEntity(
    id: json['_id'] as String,
    name: json['name'] as String,
    email: json['email'] as String,
  );
}
