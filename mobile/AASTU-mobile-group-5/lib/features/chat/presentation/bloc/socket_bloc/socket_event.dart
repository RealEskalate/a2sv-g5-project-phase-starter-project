import 'package:equatable/equatable.dart';

abstract class SocketEvent extends Equatable {
  @override
  List<Object> get props => [];
}

class ConnectSocket extends SocketEvent {}

class DisconnectSocket extends SocketEvent {}

class SendMessage extends SocketEvent {
  final String message;

  SendMessage(this.message);

  @override
  List<Object> get props => [message];
}

class ReceiveMessage extends SocketEvent {
  final String message;

  ReceiveMessage(this.message);

  @override
  List<Object> get props => [message];
}
class DeliverMessage extends SocketEvent{
  final String message;
  DeliverMessage(this.message);
  @override
  // TODO: implement props
  List<Object> get props => [message];
}