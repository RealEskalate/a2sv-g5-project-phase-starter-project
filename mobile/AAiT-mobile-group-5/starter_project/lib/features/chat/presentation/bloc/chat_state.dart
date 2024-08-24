part of 'chat_bloc.dart';

@immutable
sealed class ChatState extends Equatable {
  const ChatState();
  @override
  List<Object> get props => [];
}

final class Empty extends ChatState {}

final class ChatLoadingState extends ChatState {}

final class ErrorState extends ChatState {
  final String message;
  const ErrorState({required this.message});

  @override
  List<Object> get props => [message];
}
