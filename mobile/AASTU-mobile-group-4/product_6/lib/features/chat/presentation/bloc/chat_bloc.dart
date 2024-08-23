import 'package:flutter_bloc/flutter_bloc.dart';
import '../../domain/usecase/delete_chart.dart';
import '../../domain/usecase/get_all_chats.dart';
import '../../domain/usecase/get_chat_byid.dart';
import '../../domain/usecase/intiate_chat.dart';
import 'chat_event.dart';
import 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final GetAllChatsUsecase getAllChatsUsecase;
  final GetChatByIdUsecase getchatbyidusecase;
  final InitiateChatUsecase initiateChatUsecase;
  final DeleteChatUsecase deleteChatUsecase;

  ChatBloc({
    required this.getAllChatsUsecase,
    required this.getchatbyidusecase,
    required this.initiateChatUsecase,
    required this.deleteChatUsecase,
  }) : super(InitialChatState()) {
    on<GetAllChatsEvent>(_onGetAllChatsEvent);
    on<GetChatByIdEvent>(_onGetChatByIdEvent);
    on<InitiateChatEvent>(_onInitiateChatEvent);
    on<DeleteChatEvent>(_onDeleteChatEvent);
  }

  Future<void> _onGetAllChatsEvent(
      GetAllChatsEvent event, Emitter<ChatState> emit) async {
    emit(LoadingChatState());
    final result = await getAllChatsUsecase.getAllChats();
    result.fold(
      (failure) => emit(ErrorChatState(errorMessage: "Unable to fetch chats")),
      (chats) => emit(LoadedChatState(chats: chats)),
    );
  }

  Future<void> _onGetChatByIdEvent(
      GetChatByIdEvent event, Emitter<ChatState> emit) async {
    emit(LoadingChatState());
    final result = await getchatbyidusecase.getChatById(event.chatId);
    result.fold(
      (failure) => emit(ErrorChatState(errorMessage: 'Failed to fetch chat')),
      (chat) => emit(SingleChatLoadedState(chat: chat)),
    );
  }

  Future<void> _onInitiateChatEvent(
      InitiateChatEvent event, Emitter<ChatState> emit) async {
    emit(LoadingChatState());
    final result = await initiateChatUsecase.initiateChat(event.userId);
    result.fold(
      (failure) =>
          emit(ErrorChatState(errorMessage: 'Failed to initiate chat')),
      (chat) => emit(InitiatedChatState(chat: chat)),
    );
  }

  Future<void> _onDeleteChatEvent(
      DeleteChatEvent event, Emitter<ChatState> emit) async {
    emit(LoadingChatState());
    final result = await deleteChatUsecase.deleteChat(event.chatId);
    result.fold(
      (failure) => emit(ErrorChatState(errorMessage: 'Failed to delete chat')),
      (_) => emit(SuccessChatState(message: 'Chat deleted successfully')),
    );
  }
}
