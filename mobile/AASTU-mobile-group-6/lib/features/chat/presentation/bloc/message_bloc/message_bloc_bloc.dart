import 'package:ecommerce_app_ca_tdd/features/chat/domain/usecases/get_messages_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/usecases/send.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/message_bloc/message_bloc_state.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

class MessageBloc extends Bloc<MessageEvent, MessageState> {
  final GetMessagesUsecase getChatMessages;
  final SendUseCase sendMessage;
  MessageBloc(this.getChatMessages,this.sendMessage) : super(MessageInitial()){
  on<MessageConnection> ((event,emit) async{
    emit(MessageLoadInProgress());
    var result = await getChatMessages(event.chat.chatid);
    result.fold((l)=> emit(MessageLoadFailure(l.message)), (r)=>emit(MessageLoadSuccess(r)));

  });
  on<MessageSent>((event,emit) async{
    emit(MessageLoadInProgress());
    var result = await sendMessage(SendUseCaseParams(event.chatId,event.type,event.content));
    result.fold((l)=> emit(MessageSentFailure(l.message)), (r)=>emit(MessageSentSuccess(r.toString())));
   

  });
  }
}