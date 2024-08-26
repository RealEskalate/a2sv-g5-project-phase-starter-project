part of 'chat_bloc.dart';

@immutable
sealed class ChatState {}

final class ChatInitial extends ChatState {}

final class ChatLoaded extends ChatState {
  final List<MessageEntity> messages;

  ChatLoaded({required this.messages});
}

final class ChatErrorState extends ChatState {
  final String message;

  ChatErrorState({required this.message});
}
