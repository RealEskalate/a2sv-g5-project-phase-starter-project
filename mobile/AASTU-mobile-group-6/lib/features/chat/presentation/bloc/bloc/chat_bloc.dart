import 'package:bloc/bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/socket/socket_manager.dart';
import 'package:meta/meta.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final SocketManager socketManager;

  ChatBloc(this.socketManager) : super(ChatInitial()) {
    on<MessageReceived>(_onMessageReceived);
    on<SendMessage>(_onSendMessage);

    _initializeSocketListeners();
  }

  void _initializeSocketListeners() {
    socketManager.receiveMessage((data) {
      final message = Message.fromJson(data);
      add(MessageReceived(message));
    });
  }

  void _onMessageReceived(MessageReceived event, Emitter<ChatState> emit) {
    if (state is ChatLoaded) {
      final currentState = state as ChatLoaded;
      emit(ChatLoaded(List.from(currentState.messages)..add(event.message)));
    } else {
      emit(ChatLoaded([event.message]));
    }
  }

  void _onSendMessage(SendMessage event, Emitter<ChatState> emit) {
    socketManager.s.emit('send_message', event.message.toJson());
    if (state is ChatLoaded) {
      final currentState = state as ChatLoaded;
      emit(ChatLoaded(List.from(currentState.messages)..add(event.message)));
    } else {
      emit(ChatLoaded([event.message]));
    }
  }

  @override
  Future<void> close() {
    socketManager.s.dispose();
    return super.close();
  }
}
