import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/data/models/message_model.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:equatable/equatable.dart';

abstract class MessageState extends Equatable {
  
  const MessageState();

  @override
  List<Object> get props => [];
}

class MessageInitial extends MessageState {
  const MessageInitial();

  @override
  List<Object> get props => [];
}

class MessageLoadInProgress extends MessageState {
  const MessageLoadInProgress();
}

class MessageLoadSuccess extends MessageState {
  final Stream<MessageModel> messages;
  const MessageLoadSuccess(this.messages);

  @override
  List<Object> get props => [messages];
}

class MessageLoadFailure extends MessageState {
  final String error;
  const MessageLoadFailure(this.error);

  @override
  List<Object> get props => [error];
}

class MessageSentSuccess extends MessageState {
  final String ans;
  const MessageSentSuccess(this.ans);

  @override
  List<Object> get props => [ans];
}

class MessageSentFailure extends MessageState {
  final String error;
  const MessageSentFailure(this.error);

  @override
  List<Object> get props => [error];
}