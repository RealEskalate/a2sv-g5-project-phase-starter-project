import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_entity.dart';

class MessageEntity extends Equatable {
  final String messageId;
  final UserEntity sender;
  final ChatEntity chat;
  final String content;


  const MessageEntity({
    required this.messageId,
    required this.sender,
    required this.chat,
    required this.content
    
  });

  @override
  List<Object?> get props => [messageId,sender,chat,content];
}


