part of 'message_bloc.dart';

sealed class MessageState extends Equatable {
  const MessageState();
  
  @override
  List<Object> get props => [];
}

final class MessageInitial extends MessageState {}

class MessageLoading extends MessageState {}

class MessageLoaded extends MessageState {
  final List<ChatMessageEntity> messages;

  const MessageLoaded(this.messages);

  @override
  List<Object> get props => [messages];
}

class MessageError extends MessageState {
  final String message;

  const MessageError(this.message);

  @override
  List<Object> get props => [message];
}

class MessageSent extends MessageState {}

class MessageReceived extends MessageState {
  final ChatMessageEntity message;

  const MessageReceived(this.message);

  @override
  List<Object> get props => [message];
}

class MessageDeleted extends MessageState {}
