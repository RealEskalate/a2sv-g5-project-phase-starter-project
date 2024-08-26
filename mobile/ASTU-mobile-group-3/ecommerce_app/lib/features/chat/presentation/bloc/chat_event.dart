// chat_event.dart
import 'package:equatable/equatable.dart';
import '../../domain/entity/message.dart';

abstract class ChatEvent extends Equatable {
  @override
  List<Object?> get props => [];
}

class LoadChatRooms extends ChatEvent {}

class ConnectServerEvent extends ChatEvent {}

class LoadMessages extends ChatEvent {
  final String chatId;

  LoadMessages(this.chatId);

  @override
  List<Object?> get props => [chatId];
}

class CreateChatRoom extends ChatEvent {
  final String userId;

  CreateChatRoom(this.userId);

  @override
  List<Object?> get props => [userId];
}

class SendMessage extends ChatEvent {
  final String chatId;
  final String content;
  final String type;

  SendMessage(this.chatId, this.content, this.type);

  @override
  List<Object?> get props => [chatId, content, type];
}

class AcknowledgeMessageDelivery extends ChatEvent {
  final String messageId;

  AcknowledgeMessageDelivery(this.messageId);

  @override
  List<Object?> get props => [messageId];
}

class MessageReceived extends ChatEvent {
  final MessageEntity message;

  MessageReceived(this.message);

  @override
  List<Object?> get props => [message];
}
