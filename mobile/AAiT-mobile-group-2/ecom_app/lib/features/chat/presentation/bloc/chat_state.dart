part of 'chat_bloc.dart';

abstract class ChatState extends Equatable {
  const ChatState();  

  @override
  List<Object> get props => [];
}
class ChatInitial extends ChatState {}

class ChatLoading extends ChatState {}

class ChatLoaded extends ChatState {
  final List<UserChatEntity> chats;

  ChatLoaded({required this.chats});

  @override
  List<Object> get props => [chats];
}

class ChatError extends ChatState {
  final String message;

  ChatError({required this.message});

  @override
  List<Object> get props => [message];
}

class ChatMessageLoaded extends ChatState {
  final List<MessageEntity> messages;

  ChatMessageLoaded({required this.messages});

  @override
  List<Object> get props => [messages];
}

class ChatMessageError extends ChatState {
  final String message;

  ChatMessageError({required this.message});

  @override
  List<Object> get props => [message];
}

class ChatMessageLoading extends ChatState {}

class ChatInitiated extends ChatState {
  final UserChatEntity chat;

  ChatInitiated({required this.chat});

  @override
  List<Object> get props => [chat];
}

class ChatInitiateError extends ChatState {
  final String message;

  ChatInitiateError({required this.message});

  @override
  List<Object> get props => [message];
}

class ChatInitiateLoading extends ChatState {}