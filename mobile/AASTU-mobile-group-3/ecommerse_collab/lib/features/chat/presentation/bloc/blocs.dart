// import 'package:bloc/bloc.dart';

// import 'event.dart';
// import 'state.dart';

// class ChatBloc extends Bloc<ChatEvent, ChatState>{

  
//   ChatBloc() : super(ChatInitialState()){
//     on<LoadChatsEvent>(_onLoadChats);
//     on<LoadMessagesEvent>(_onLoadMessagesEvent);
//     on<SendMessageEvent>(_onSendMessageEvent);
//     on<ReceiveMessageEvent>(_onReceiveMessageEvent);
//     on<TypingEvent>(_onTypingEvent);
//     on<DeleteMessageEvent>(_onDeleteMessage);
//     on<NotificationEvent>(_onNotification);
//     on<ErrorEvent>(_onError);
//   } 
  

//   Future<void> _onLoadChats(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onLoadMessagesEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onSendMessageEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onReceiveMessageEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onTypingEvent(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onDeleteMessage(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }

//   Future<void> _onNotification(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }Future<void> _onError(LoadChatsEvent event, Emitter<ChatState> emit) async {
//     emit(ChatLoadingState());

   
//   }


// }