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
      chatEntity: json['messageId'] as ChatEntity, 
      content: json['content'],
      sender:json['sender'] as UserEntity,
      type: json['type'] as String,
      messageId: json['messageId'] as String,

    
      );
  } 

}


