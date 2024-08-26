import 'package:bloc/bloc.dart';

import 'event.dart';
import 'state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState>{

  
  ChatBloc() : super(ChatInitialState()){
    on<LoadChatsEvent>(_onLoadChats);
    on<LoadMessagesEvent>(_onLoadMessagesEvent as EventHandler<LoadMessagesEvent, ChatState>);
    on<SendMessageEvent>(_onSendMessageEvent as EventHandler<SendMessageEvent, ChatState>);
    on<ReceiveMessageEvent>(_onReceiveMessageEvent as EventHandler<ReceiveMessageEvent, ChatState>);
    on<TypingEvent>(_onTypingEvent as EventHandler<TypingEvent, ChatState>);
    on<DeleteMessageEvent>(_onDeleteMessage as EventHandler<DeleteMessageEvent, ChatState>);
    on<NotificationEvent>(_onNotification as EventHandler<NotificationEvent, ChatState>);
    on<ErrorEvent>(_onError as EventHandler<ErrorEvent, ChatState>);
  } 
  

  Future<void> _onLoadChats(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());
  }

  Future<void> _onLoadMessagesEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }

  Future<void> _onSendMessageEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }

  Future<void> _onReceiveMessageEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }

  Future<void> _onTypingEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }

  Future<void> _onDeleteMessage(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }

  Future<void> _onNotification(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }Future<void> _onError(LoadChatsEvent event, Emitter<ChatState> emit) async {
    emit(ChatLoadingState());

   
  }


}