import 'package:equatable/equatable.dart';

class MessageEntity extends Equatable {
  final String messageId;
  final UserEntity sender;
  final ChatEntity chat;
  final String content;


  const ChatEntity({
    required this.messageId,
    required this.sender,
    required this.chat,
    required this.content
    
  });

  @override
  List<Object?> get props => [
        messageId,
        sender,
        chat,
        content
      ];
}


