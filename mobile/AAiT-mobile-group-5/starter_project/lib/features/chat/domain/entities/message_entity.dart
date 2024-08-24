import 'package:equatable/equatable.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/shared/entities/user.dart';

class Message extends Equatable {
  final String id;
  final String content;
  final Chat chat;
  final User sender;

  const Message({
    required this.id,
    required this.content,
    required this.chat,
    required this.sender,
  });

  @override
  List<Object?> get props => [id, content, chat, sender];
}
