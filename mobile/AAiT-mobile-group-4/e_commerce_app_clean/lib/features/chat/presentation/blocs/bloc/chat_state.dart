part of 'chat_bloc.dart';

sealed class ChatState extends Equatable {
  const ChatState();
  
  @override
  List<Object> get props => [];
}

final class ChatInitial extends ChatState {}

class ChatsLoadingState extends ChatState {}

class ChatsLoadedState extends ChatState {
  final List<Chat> chats;

  const ChatsLoadedState({required this.chats});
}

class ChatMessageLoadingState extends ChatState {}

class ChatMessageLoadedState extends ChatState {
  final List<Message> messages;
  const ChatMessageLoadedState({required this.messages});
}

class ChatCreatedState extends ChatState {
  final Chat chat;
  const ChatCreatedState({required this.chat});
}

class ChatErrorState extends ChatState {
  final String message;
  const ChatErrorState({required this.message});
}

class ChatMessageSentState extends ChatState {}
