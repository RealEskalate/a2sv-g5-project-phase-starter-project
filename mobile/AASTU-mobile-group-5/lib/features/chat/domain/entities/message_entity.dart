



import 'package:equatable/equatable.dart';

import '../../../user/domain/entities/user.dart';

class MessageEntity extends Equatable{
  final String chatId;
  final User sender;
  final String type;
  final String content;
  MessageEntity({
    required this.chatId,
    required this.sender,
    required this.type,
    required this.content,
    
  });
  @override
  
  List<Object?> get props => [chatId,sender,type,content];
  
}