import 'dart:async';

import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:meta/meta.dart';

import '../../../auth/domain/entities/user_entity.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/usecases/get_message_usecase.dart';
import '../../domain/usecases/send_message_usecase.dart';

import '../../domain/usecases/my_chat_usecase.dart';

part 'chat_event.dart';
part 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final GetMessageUsecase getMessagesUseCase;
  final SendMessageUsecase sendMessageUseCase;
  final MyChatUsecase myChatUsecase;
  

  StreamSubscription<MessageEntity>? _messageSubscription;

  ChatBloc({
    required this.getMessagesUseCase,
    required this.sendMessageUseCase,
    required this.myChatUsecase,
  }) : super(ChatInitial()) {
    // Handle StartChat event to listen to incoming messages

    on<StartChat>((event, emit) async {
      final eitherResult = await getMessagesUseCase();
      eitherResult.fold(
        (failure) {
          // Handle failure, e.g., emit an error state
          emit(ChatErrorState(message: failure.message));
        },
        (messageStream) {
          // If successful, listen to the stream
          _messageSubscription = messageStream.listen((message) {
            add(NewMessageReceivedEvent(message: message));
          });
        },
      );
    });

    // Handle NewMessageReceivedEvent to update the state with new messages
    on<NewMessageReceivedEvent>((event, emit) {
      if (state is ChatLoaded) {
        final currentState = state as ChatLoaded;
        final updatedMessages = List<MessageEntity>.from(currentState.messages)
          ..add(event.message);
        emit(ChatLoaded(messages: updatedMessages));
      }
    });
  
    // Handle SendMessage event to send a message to the server
    on<SendMessage>((event, emit) async {
      await sendMessageUseCase(
          chatId: event.chatId, message: event.message, type: event.type);
    });
  
  
   on<GetAllChatEvent>((event, emit) async{
      emit(ChatLoadingState());
      final res = await myChatUsecase.execute();
      res.fold((error){emit(ChatFailureState(message:error.message ));}, 
      (right){
        print(right);
        emit(LoadedAllChatState( allChats: right));
      });

    });
  }

    


  @override
  Future<void> close() {
    _messageSubscription?.cancel();
    return super.close();
  }
}
