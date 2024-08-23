



import '../../domain/entities/chat_entity.dart';

class ChatModel extends ChatEntity{
  ChatModel({
    required super.chat_id,
    required super.seller_one,
    required super.seller_two,
  });
  factory ChatModel.fromJson(Map<String,dynamic>Json){
    return ChatModel(
      chat_id: Json['data']['_id'],
      seller_one: Json['data']['user1'],
      seller_two: Json['data']['user2'],
    );
  }
  Map<String,dynamic>toJson(){
    return{
      '_id': chat_id,
      'user1': seller_one,
      'user2': seller_two,
    };
  }

}