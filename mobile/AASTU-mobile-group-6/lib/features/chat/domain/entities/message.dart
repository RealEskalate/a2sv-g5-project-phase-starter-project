import 'package:ecommerce_app_ca_tdd/features/chat/data/models/chat_models.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:equatable/equatable.dart';

class Message extends Equatable{
  final String messageid;
  final UserModel sender;
  final ChatModel uniqueChat;
  final String type;
  final String content;
  
  Message({required this.messageid,required this.sender, required this.uniqueChat, required this.type, required this.content});
  
  @override
  // TODO: implement props
  List<Object?> get props => [sender, uniqueChat, type, content];


  Map<String, dynamic> toJson() {
    return {
      '_id': messageid,
      'sender': sender.toJson(),
      'chat': uniqueChat.toJson(),
      'type': 'text',
      'content': content
    };

  }

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      messageid: json['_id'],
      sender: UserModel.fromJson(json['sender']),
      uniqueChat: ChatModel.fromJson(json['chat']),
      type: json['type'],
      content: json['content'],

      );
  }
  
}
class MessageType{
  final String type;
  final String content;
  MessageType({required this.type,required this.content});
  @override
  // TODO: implement props
  Map<String, dynamic> toJson() {
    return {
      'type': 'text',
      'content': content
    };

  }
  List<Object?> get props => [type, content];
                                  
                                
}