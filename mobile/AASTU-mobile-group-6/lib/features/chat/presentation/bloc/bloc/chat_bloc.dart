import 'package:ecommerce_app_ca_tdd/locator.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:ecommerce_app_ca_tdd/core/usecases/usecases.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/entities/message.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/usecases/get_chats_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/domain/usecases/initiate_chat_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_event.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/presentation/bloc/bloc/chat_state.dart';
import 'package:ecommerce_app_ca_tdd/features/chat/socket/socket_manager.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_all_usecase.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  
  final GetChatsUsecase getChatsUsecase;
  final InitiateChatUsecase initiateChatUsecase;

  ChatBloc(this.getChatsUsecase,this.initiateChatUsecase) : super(ChatInitial()) {
    on<ListAllMessagesEvent>((event,emit) async{
      emit(ChatLoading());
      var result = await getChatsUsecase(NoParams());
      result.fold((l)=> emit(ChatError(l.message)), (r)=> emit(ChatLoaded(r)));
    }
    );
    on<InitiateChatEvent>((event,emit) async {
      emit(ChatInitateLoading());
      var result = await initiateChatUsecase(event.receiver);
      result.fold((l)=> emit(ChatInitateFailure(l.message)),(r)=>(ChatInitateLoaded(r)));
    }
    );

  }

}