import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';

class MessageModel extends MessageEntity{
   MessageModel(
    {
      required super.messageId,
      required super.chatEntity,
      required super.content,
      required super.sender,
      required super.type
    }
  );
  factory MessageModel.fromJson(Map<String,dynamic> json){
    return MessageModel(
      chatEntity: chatEntityFromJson(json['chat'] ), 
      content: json['content'],
      sender:jsonToUserEntity(json['sender']),
      type: json['type'] as String,
      messageId: json['_id'] as String,

      );
  } 

   Map<String, dynamic> toJson() {
    return {
      '_id': messageId,
      'sender': userEntityToJson(sender),
      'content': content,
      'type': type,
      'chat': chatEntityToJson(chatEntity),  
    };
  }
}




ChatEntity chatEntityFromJson(Map<String, dynamic> json){
  return ChatEntity(
    chatId: json['chatId'],
    user1: jsonToUserEntity(json['user1']),
    user2: jsonToUserEntity(json['user2']),
  );
}

Map<String, dynamic> chatEntityToJson(ChatEntity entity) {
    return {
      'chatId': entity.chatId,
      'user1': userEntityToJson(entity.user1),
      'user2': userEntityToJson(entity.user2),
    };
  }

// Function to convert UserEntity to JSON
Map<String, dynamic> userEntityToJson(UserEntity user) {
  return {
    'id': user.id,
    'name': user.name,
    'email': user.email,
  };
}

// Function to convert JSON to UserEntity
UserEntity jsonToUserEntity(Map<String, dynamic> json) {
  return UserEntity(
    id: json['id'],
    name: json['name'],
    email: json['email'],
    accessToken: json['accessToken'],
  );
}
