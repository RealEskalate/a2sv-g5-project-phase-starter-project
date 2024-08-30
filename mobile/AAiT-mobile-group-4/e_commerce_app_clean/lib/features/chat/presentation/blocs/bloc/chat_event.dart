part of 'chat_bloc.dart';

sealed class ChatEvent extends Equatable {
  const ChatEvent();

  @override
  List<Object> get props => [];
}

class FetchChatsEvent extends ChatEvent {}

class GetChatMessagesEvent extends ChatEvent {
  final String chatId;

  const GetChatMessagesEvent({required this.chatId});
}

class CreateChatEvent extends ChatEvent {
  final String userId;

  const CreateChatEvent({required this.userId});
}

class SendMessageEvent extends ChatEvent {
  final Message message;
  const SendMessageEvent({required this.message});
}
