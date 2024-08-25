import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entity/user.dart';
import 'chat.dart';

class Message extends Equatable{
  final String id;
  final User sender;
  final Chat chat;
  final String type;
  final dynamic content;

  Message({
    required this.id,
    required this.sender,
    required this.chat,
    required this.type,
    required this.content,
  });
  
  @override
  // TODO: implement props
  List<Object?> get props => [id, sender, chat, type, content];

  
 
  
}