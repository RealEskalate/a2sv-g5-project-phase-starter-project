import 'package:equatable/equatable.dart';

class ChatEntity extends Equatable {
  final String chatId;
  final UserEntity user1;
  final UserEntity user2;

  const ChatEntity({
    required this.chatId,
    required this.user1,
    required this.user2,
    
  });

  @override
  List<Object?> get props => [
        chatId
        user1,
        user2
      ];
}


