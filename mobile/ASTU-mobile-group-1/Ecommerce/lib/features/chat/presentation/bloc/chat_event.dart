part of 'chat_bloc.dart';

@immutable
sealed class ChatEvent {}

class StartChat extends ChatEvent {}

class SendMessage extends ChatEvent {
  final String chatId;
  final String type;
  final String message;

  SendMessage(
      {required this.chatId, required this.type, required this.message});
}

class NewMessageReceivedEvent extends ChatEvent {
  final MessageEntity message;

  NewMessageReceivedEvent({required this.message});
}
