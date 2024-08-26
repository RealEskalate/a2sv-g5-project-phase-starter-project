import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../../../core/utility/global_message_part.dart';
import '../../../../../core/utility/socket_impl.dart';
import 'socket_event.dart';
import 'socket_state.dart';

class SocketBloc extends Bloc<SocketEvent, SocketState> {
  final SocketService socketService;

  SocketBloc(this.socketService) : super(SocketInitial()) {
    on<ConnectToSocket>(_onConnectToSocket);
    on<DisconnectFromSocket>(_onDisconnectFromSocket);
    on<SendMessage>(_onSendMessage);
    on<ReceiveMessage>(_onReceiveMessage);

    // Listen to messages from the socket
    socketService.listen('message:received', (data) {
      add(ReceiveMessage(data));
    });
  }

  void _onConnectToSocket(ConnectToSocket event, Emitter<SocketState> emit) async {
    await socketService.connect();
    emit(SocketConnected());
  }

  void _onDisconnectFromSocket(DisconnectFromSocket event, Emitter<SocketState> emit) {
    socketService.disconnect();
    emit(SocketDisconnected());
  }

Future<void> _onSendMessage(SendMessage event, Emitter<SocketState> emit) async {
  print(event.message);
  print(event.chatId);
  
  
  socketService.sendMessage(event.chatId, event.message);
  
  // Emit the state after the message is sent
  emit(SocketDisconnected());
  print(12345);
  emit(SocketConnected());
  print(GlobalMessagePart.gloablMessage);
}

  void _onReceiveMessage(ReceiveMessage event, Emitter<SocketState> emit) {
    emit(SocketConnected());
    print(GlobalMessagePart.gloablMessage);
    emit(SocketMessageReceived(event.messageData));

    print(12345);
  }

  
}
