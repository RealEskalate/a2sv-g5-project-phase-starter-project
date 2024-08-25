import 'package:equatable/equatable.dart';

class Message extends Equatable{
  late final String senderId;
  late final String chatId;
  late final String type;
  late final String content;
  
  Message({required this.senderId, required this.chatId, required this.type, required this.content});
  
  @override
  // TODO: implement props
  List<Object?> get props => [senderId, chatId, type, content];


  Map<String, String> toJson() {
    return {
      'chatId': chatId,
      'type': type='text',
      'content': content
    };

  }

  factory Message.fromJson(Map<String, dynamic> json) {
    return Message(
      senderId: json['id'],
      chatId: json['chatId'],
      type: json['type'],
      content: json['message'],

      );
  }
  
}