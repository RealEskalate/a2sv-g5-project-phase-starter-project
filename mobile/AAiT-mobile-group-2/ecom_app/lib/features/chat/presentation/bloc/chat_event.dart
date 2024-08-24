part of 'chat_bloc.dart';

abstract class ChatEvent extends Equatable {
  const ChatEvent();

  @override
  List<Object> get props => [];
}

class GetChatsEvent extends ChatEvent {}

class GetMessagesEvent extends ChatEvent {
  final String ChatID;

  GetMessagesEvent({required this.ChatID});

  @override
  List<Object> get props => [ChatID];
}

class InitiateChatEvent extends ChatEvent {
  final String UserID;

  InitiateChatEvent({required this.UserID});

  @override
  List<Object> get props => [UserID];
}