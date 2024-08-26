abstract class SocketEvent {}

class ConnectToSocket extends SocketEvent {}

class DisconnectFromSocket extends SocketEvent {}

class SendMessage extends SocketEvent {
  final String chatId;
  final String message;

  SendMessage(this.chatId, this.message);
}

class ReceiveMessage extends SocketEvent {
  final Map<String, dynamic> messageData;

  ReceiveMessage(this.messageData);
}
