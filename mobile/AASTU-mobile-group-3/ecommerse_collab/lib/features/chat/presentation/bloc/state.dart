import 'package:equatable/equatable.dart';

import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

class ChatState extends Equatable {
  const ChatState();
  
  @override
  List<Object?> get props => [];
}

class ChatInitialState extends ChatState {}

class ChatLoadingState extends ChatState {}

class ChatLoadedState extends ChatState {
  final List<Chat> chats;
  const ChatLoadedState(this.chats);

  @override
  List<Object?> get props => [chats];
}

class MessagesLoadingState extends ChatState {}

class MessagesLoadedState extends ChatState {
  final List<Message> messages;

  const MessagesLoadedState(this.messages);

  @override
  List<Object?> get props => [messages];
}

class MessageSendingState extends ChatState {}

class MessageSentState extends ChatState {
  final String message;
  const MessageSentState(this.message);

  @override
  List<Object?> get props => [message];
}

class MessageReceivingState extends ChatState {}

class ChatTypingIndicatorState extends ChatState {}

class ChatMessageDeletedState extends ChatState {
  final String message;
  const ChatMessageDeletedState(this.message);

  @override
  List<Object?> get props => [message];
}

class ChatReceivedState extends ChatState {
  final String message;
  const ChatReceivedState(this.message);

  @override
  List<Object?> get props => [message];
}

enum NotificationType {
  newMessage,
  systemAlert,
  error,
}

class NotificationState extends ChatState {
  final NotificationType type;
  final String content;
  final DateTime? timestamp;

  const NotificationState({
    required this.type,
    required this.content,
    this.timestamp,
  });

  @override
  List<Object?> get props => [type, content, timestamp];
}

class ChatErrorState extends ChatState {
  final String message;
  const ChatErrorState(this.message);

  @override
  List<Object?> get props => [message];
}
