import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../../../core/usecase/usecase.dart';
import '../../../../domain/entities/chat_entity.dart';
import '../../../../domain/usecases/delete_chat_use_case.dart';
import '../../../../domain/usecases/initiate_chat_use_case.dart';
import '../../../../domain/usecases/my_chat_by_id_use_case.dart';
import '../../../../domain/usecases/my_chats_use_case.dart';

part 'chat_list_event.dart';
part 'chat_list_state.dart';

class ChatListBloc extends Bloc<ChatListEvent, ChatListState> {
  final MyChatsUseCase myChatsUseCase;
  final MyChatByIdUseCase myChatByIdUseCase;
  final InitiateChatUseCase initiateChatUseCase;
  final DeleteChatUseCase deleteChatUseCase;

  ChatListBloc(
      {required this.myChatByIdUseCase,
      required this.initiateChatUseCase,
      required this.deleteChatUseCase,
      required this.myChatsUseCase})
      : super(ChatListInitial()) {

    on<LoadAllChatEvent>((event, emit) async {
      emit(const LoadingAllChatState());
      final result = await myChatsUseCase.call(NoParams());

      result.fold((fail) {
        emit(const LoadingAllChatError(errorMessage: 'error in loading chats'));
      }, (data) {
        emit(LoadAllChatState(chats: data));
      });
    });



     on<GetChatEvent>((event, emit) async {
      emit(const LoadingSingleChat());
      final result = await myChatByIdUseCase.call(MyChatByIdParams(chatId: event.chatId));

      result.fold((fail) {
        emit(const LoadingSingleChatError(errorMessage: 'error in loading this chat'));
      }, (data) {
        emit(LoadedSingleChatState(chat: data));
      });
    });

    on<InitiateChatEvent>((event, emit) async {
      emit(const InitiatingChatState());
      final result = await initiateChatUseCase.call(InitiateChatParams(sellerId: event.sellerId));

      result.fold((fail) {
        emit(const InitiatingChatError(errorMessage: 'error in initiating chat'));
      }, (data) {
        emit(InitiatedChatState(chat: data));
      });
    });

    on<DeleteChatEvent>((event, emit) async {
      emit(const DeletingChatState());
      final result = await deleteChatUseCase.call(DeleteChatParams(chatId: event.chatId));

      result.fold((fail) {
        emit(const DeletingChatError(errorMessage: 'error in deleting chat'));
      }, (data) {
        emit(const DeletedChatState());
      });
    });

  }
}
