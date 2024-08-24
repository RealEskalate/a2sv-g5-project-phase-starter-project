import 'package:equatable/equatable.dart';
import '../../../data/models/message_model.dart';
import '../../../domain/entities/message_entity.dart';
// import '../../domain/entities/message_entity.dart';

abstract class SocketEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class SendMessage extends SocketEvent {
  final MessageModel message;

  SendMessage(this.message);

  @override
  List<Object> get props => [message];
}

class ReceiveMessage extends SocketEvent {
  final MessageModel message;

  ReceiveMessage(this.message);

  @override
  List<Object> get props => [message];
}

class MessageDelivered extends SocketEvent {
  final MessageModel message;

  MessageDelivered(this.message);

  @override
  List<Object> get props => [message];
}
