import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:equatable/equatable.dart';

abstract class ChatEvent extends Equatable {
  const ChatEvent();

  @override
  List<Object?> get props => [];
}

class MessageReceived extends ChatEvent {
  final Message message;

  const MessageReceived(this.message);

  @override
  List<Object?> get props => [message];
}

class SendMessage extends ChatEvent {
  final Message message;

  const SendMessage(this.message);

  @override
  List<Object?> get props => [message];
}
class ListAllMessagesEvent extends ChatEvent{}

class InitiateChatEvent extends ChatEvent{
  final String receiver;
  const InitiateChatEvent(this.receiver);
  @override
  List<Object> get props => [receiver];
  
  }
