import '../../../user/data/models/user_model.dart';
import '../../domain/entities/message_entity.dart';
import 'chat_model.dart';

class MessageModel extends MessageEntity{
  MessageModel({
    required super.messageId,
    required super.sender,
    required super.content,
    required super.chat,
    required super.type,

  });
  factory MessageModel.fromJson(Map<String,dynamic>Json){
    return MessageModel(
      messageId: Json['data']['_id'],
      sender: UserModel.fromJson(Json['data']['sender']),
      chat:ChatModel.fromJson(Json['data']['chat']),
      content: Json['data']['content'], 
      type: Json['data']['type'],
    );
  }
  Map<String,dynamic>toJson(){
    return{
      '_id': messageId,
      'sender':(sender as UserModel).toJson(),
      'content': content,
      'chat': (chat as ChatModel).toJson(),
      'type': type,
    };
  }

}