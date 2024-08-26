import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:flutter/foundation.dart';



class ChatModel extends ChatEntity {
  final String chatId;
  final UserModel user1;
  final UserModel user2;
  ChatModel({
    required this.chatId,
    required this.user1,
    required this.user2,
  }) : super(chatid: chatId, user1: user1, user2: user2);

  factory ChatModel.fromJson(Map<String, dynamic> json) {
    return ChatModel(
      chatId: json['chatId'],
      user2: UserModel.fromJson(json['user2']),
      user1: UserModel.fromJson(json['user1']),
    );
  }

  Map<String, dynamic> toJson() {
    return {
      'chatId': chatId,
      'user2': (user2 as UserModel).toJson(),
      'user1': (user1 as UserModel).toJson(),
    };
  }
}












// import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
// import 'package:ecommerce_app_ca_tdd/features/product/data/models/seller_model.dart';
// import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';

// class ChatModel{
//   final String id;
//   final ChatEntity chat;
//   final String type;
//   final String content;

//   // get receiver => chat.user1 == sender ? chat.user2 : chat.user1;

//   const ChatModel({
//     required this.id,
//     required this.chat,
//     required this.type,
//     required this.content,
//   });

//   @override
//   List<Object?> get props => [id, chat, type, content];

//   factory ChatModel.fromJson(Map<String, dynamic> json) {
//     return ChatModel(
//       id: json['_id'],
//       chat: ChatEntity.fromJson(json['chat']),
//       type: json['type'],
//       content: json['content'],
      
//     );
//   }
//   // factory ChatModel.fromEntity(ChatEntity chat) {
//   //   return ChatModel(
//   //     id: chat.id,
//   //     user1: UserModel.fromEntity(chat.user1),
//   //     user2: UserModel.fromEntity(chat.user2),
//   //   );
//   // }

//   Map<String, dynamic> toJson(){
//     return {
//       '_id': id,
//       'chat': chat.toJson(),
//       'type': type,
//       'content': content,
//     };

//     }
        
        
// }