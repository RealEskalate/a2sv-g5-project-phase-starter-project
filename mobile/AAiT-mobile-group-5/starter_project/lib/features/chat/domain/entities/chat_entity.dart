import 'package:equatable/equatable.dart';
import 'package:starter_project/features/shared/entities/user.dart';

class Chat extends Equatable {
  final String chatId;
  final User user1;
  final User user2;

  const Chat({
    required this.chatId,
    required this.user1,
    required this.user2,
  });

  @override
  List<Object?> get props => [chatId, user1, user2];
}
