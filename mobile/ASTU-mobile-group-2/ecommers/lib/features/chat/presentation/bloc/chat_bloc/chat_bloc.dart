import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/usecase/chat_usecase.dart';
import '../chat_event/chat_event.dart';
import '../chat_state/chat_state.dart';

class ChatBloc extends Bloc<ChatEvent,ChatState> {

  final ChatUsecase chatUsecase;

  ChatBloc ({
    required this.chatUsecase
  }):super(ChatInitialState()){

    on<OnGetAllChat>(
      (event,emit) async{
        emit(ChatLoadingState());
        final result = await chatUsecase.getMychats();

        result.fold(
          (failur){
            emit(ChatErrorState(errorMessage: failur.message));
          }, 
          (data){
            emit(ChatMessageGetSuccess(chatEntity: data));
          },
          );
      }
    );

    on<OnDeleteChat>(
      (event,emit) async{
        emit(ChatLoadingState());
        final result = await chatUsecase.deleteChats(event.chatId);
        result.fold(
          (failur){
            emit(ChatErrorState(errorMessage: failur.message));
          }, 
          (data){
            if(data == false){
              emit(ChatErrorState(errorMessage: 'try again'));
            }
            else{
            emit(ChatDeleteSuccess(chatDeleted: data));}
          });
      }
    );

    on<OnGetChatById>(
      (event,emit) async {
        emit(ChatLoadingState());
        final result = await chatUsecase.getChatById(event.chatId);

        result.fold(
          (failur){
            emit(ChatErrorState(errorMessage: failur.message));
          },
          (data){
            emit(ChatByIdSuccess(chatEntity: data));
          });
      }
    );
  }
}