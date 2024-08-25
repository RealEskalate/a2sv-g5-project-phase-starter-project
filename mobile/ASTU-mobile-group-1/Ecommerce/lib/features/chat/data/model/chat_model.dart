import '../../../auth/data/models/user_model.dart';
import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity{
   const ChatModel(
    {
      required super.chatId,
      required super.user1,
      required super.user2
    }
  );
  factory ChatModel.fromJson(Map<String,dynamic> json){
    return ChatModel(
      chatId: json['id'] , 
      user1: UserModel.fromJson(json, token), 
      user2: json['user2'] as UserEntity);
  } 

  factory ChatModel.fromEntity(ChatEntity chatEntity)=>
        ChatModel(chatId: chatEntity.chatId, user1: chatEntity.user1, user2: chatEntity.user2);
  
  factory 

}


