import 'package:flutter_bloc/flutter_bloc.dart';
import 'socket_event.dart';
import 'socket_state.dart';
import 'package:web_socket_channel/web_socket_channel.dart';

class SocketBloc extends Bloc<SocketEvent, SocketState> {
  WebSocketChannel? _channel;

  SocketBloc() : super(SocketInitial()) {
    on<SendMessage>(_onSendMessage);
    on<ReceiveMessage>(_onReceiveMessage);
    on<MessageDelivered>(_onMessageDelivered);
  }

  void _onSendMessage(SendMessage event, Emitter<SocketState> emit) {
    try {
      _channel?.sink.add(event.message.toJson());
      emit(SocketMessageDelivered(event.message.content)); 
    } catch (error) {
      emit(SocketError(error.toString()));
    }
  }

  void _onReceiveMessage(ReceiveMessage event, Emitter<SocketState> emit) {
    emit(SocketMessageReceived(event.message.content)); 
  }

  void _onMessageDelivered(MessageDelivered event, Emitter<SocketState> emit) {
   
    emit(SocketMessageDelivered(event.message.content));
  }

  @override
  Future<void> close() {
    _channel?.sink.close();
    return super.close();
  }
}
