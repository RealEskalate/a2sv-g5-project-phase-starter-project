import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:equatable/equatable.dart';

abstract class MessageEvent extends Equatable {
  const MessageEvent();

  @override
  List<Object> get props => [];
}

class MessageConnection extends MessageEvent {
  final ChatEntity chat;

  const MessageConnection(this.chat);
}

class MessageSent extends MessageEvent {
  final String chatId;
  final String content;
  final String type;

  MessageSent(
    this.chatId,
    this.content,
    this.type,
  );
}