part of 'message_bloc.dart';

sealed class MessageEvent extends Equatable {
  const MessageEvent();

  @override
  List<Object> get props => [];
}

class LoadMessages extends MessageEvent {
  final String chatId;

  const LoadMessages(this.chatId);

  @override
  List<Object> get props => [chatId];
}

class SendMessage extends MessageEvent {
  final String chatId;
  final ChatMessageEntity message;

  const SendMessage(this.chatId, this.message);

  @override
  List<Object> get props => [chatId, message];
}

class ReceiveMessage extends MessageEvent {
  final ChatMessageEntity message;

  const ReceiveMessage(this.message);

  @override
  List<Object> get props => [message];
}

class DeleteMessage extends MessageEvent {
  final String messageId;

  const DeleteMessage(this.messageId);

  @override
  List<Object> get props => [messageId];
}
