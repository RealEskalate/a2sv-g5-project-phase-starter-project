import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/chat_entity.dart';
import 'package:equatable/equatable.dart';


abstract class MessageEvent extends Equatable {
  const MessageEvent();

  @override
  List<Object> get props => [];
}

class MessageSocketConnectionRequested extends MessageEvent {
  final ChatEntity chat;

  const MessageSocketConnectionRequested(this.chat);
}

class MessageSent extends MessageEvent {
  final ChatEntity chat;
  final String content;
  final String type;

  const MessageSent(
    this.chat,
    this.content,
    this.type,
  );
}