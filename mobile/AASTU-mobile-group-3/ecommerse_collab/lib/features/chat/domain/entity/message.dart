import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entity/user.dart';

class Message extends Equatable{
  final String chatId;
  final User sender;
  final String type;
  final dynamic content;

  Message({
    required this.chatId,
    required this.sender,
    required this.type,
    required this.content,
  });
  
  @override
  List<Object?> get props => [chatId, sender, type, content];
  
}