import 'package:equatable/equatable.dart';

import '../../../authentication/domain/entity/user.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

sealed class ChatEvent extends Equatable{
  const ChatEvent();

  @override
  List<Object?> get props => [];
}

class LoadChatsEvent extends ChatEvent{
  final Chat chats;
  
  const LoadChatsEvent(this.chats);
  @override
  List<Object> get props => [chats];
}

class LoadMessagesEvent extends ChatEvent{

  final String chatId;

  const LoadMessagesEvent(this.chatId);

  @override
  List<Object> get props => [chatId];

}

class SendMessageEvent extends ChatEvent{

  final Message message;
  const SendMessageEvent(this.message);
}

class ReceiveMessageEvent extends ChatEvent{

  final Message message;
  const ReceiveMessageEvent(this.message);
}

class TypingEvent extends ChatEvent {
  final String chatId;
  final User typingUser;

  const TypingEvent(this.chatId, this.typingUser);

  @override
  List<Object> get props => [chatId, typingUser];
}

class DeleteMessageEvent extends ChatEvent {
  final String messageId;
  final String chatId;

  const DeleteMessageEvent(this.messageId, this.chatId);

  @override
  List<Object> get props => [messageId, chatId];
}

enum NotificationType {
  newMessage,
  systemAlert,
  error,
}
class NotificationEvent extends ChatEvent {
  final NotificationType type;
  final String content;
  final DateTime? timestamp;

  const NotificationEvent({
    required this.type,
    required this.content,
    this.timestamp,
  });

  @override
  List<Object?> get props => [type, content, timestamp];
}
class ErrorEvent extends ChatEvent {
  final String errorMessage;

  const ErrorEvent(this.errorMessage);

  @override
  List<Object> get props => [errorMessage];
}