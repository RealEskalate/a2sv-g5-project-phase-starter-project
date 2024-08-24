import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity{
   ChatModel(
    {
      required super.chatId,
      required super.user1,
      required super.user2
    }
  );
  factory ChatModel.fromJson(Map<String,dynamic> json){
    return ChatModel(
      chatId: json['id'] as String, 
      user1: json['user1'] as UserEntity, 
      user2: json['user2'] as UserEntity);
  } 

}


