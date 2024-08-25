// chat_state.dart
import 'package:equatable/equatable.dart';
import '../../domain/entity/chat.dart';
import '../../domain/entity/message.dart';

abstract class ChatState extends Equatable {
  @override
  List<Object?> get props => [];
}

class ChatInitial extends ChatState {}

class ChatLoading extends ChatState {}

class ChatConnected extends ChatState {}

class ChatLoaded extends ChatState {
  final List<ChatEntity> chats;

  ChatLoaded(this.chats);

  @override
  List<Object?> get props => [chats];
}

class MessagesLoaded extends ChatState {
  final List<MessageEntity> messages;

  MessagesLoaded(this.messages);

  @override
  List<Object?> get props => [messages];
}

class ChatError extends ChatState {
  final String message;

  ChatError(this.message);

  @override
  List<Object?> get props => [message];
}
