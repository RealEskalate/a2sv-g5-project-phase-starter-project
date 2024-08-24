part of 'chat_bloc.dart';

@immutable
sealed class ChatEvent extends Equatable {
  const ChatEvent();
  @override
  List<Object> get props => [];
}

class InitiateChatEvent extends ChatEvent {
  final String userId;
  const InitiateChatEvent({required this.userId});

  @override
  List<Object> get props => [userId];
}

class GetSpecificChatEvent extends ChatEvent {
  final String chatId;
  const GetSpecificChatEvent({required this.chatId});

  @override
  List<Object> get props => [chatId];
}

class GetAllChatEvent extends ChatEvent {}

class DeleteSpecificChatEvent extends ChatEvent {
  final String chatId;
  const DeleteSpecificChatEvent({required this.chatId});

  @override
  List<Object> get props => [chatId];
}

class  GetMessageByChatId extends ChatEvent{
  final String messageId;
  const GetMessageByChatId({required this.messageId});

  @override
  List<Object> get props => [messageId];
}