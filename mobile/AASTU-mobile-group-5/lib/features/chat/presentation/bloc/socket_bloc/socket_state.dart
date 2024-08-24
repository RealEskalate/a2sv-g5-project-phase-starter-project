import 'package:equatable/equatable.dart';

abstract class SocketState extends Equatable {
  @override
  List<Object> get props => [];
}

class SocketInitial extends SocketState {}

class SocketConnecting extends SocketState {}

class SocketConnected extends SocketState {}

class SocketDisconnected extends SocketState {}

class SocketMessageReceived extends SocketState {
  final String message;

  SocketMessageReceived(this.message);

  @override
  List<Object> get props => [message];
}
class SocketMessageDelivered extends SocketState{
  final String message;
  SocketMessageDelivered(this.message);
  @override
  // TODO: implement props
  List<Object> get props => [message];
}

class SocketError extends SocketState {
  final String error;

  SocketError(this.error);

  @override
  List<Object> get props => [error];
}
