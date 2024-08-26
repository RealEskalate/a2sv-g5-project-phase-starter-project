import 'package:bloc/bloc.dart';
import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/chat/data/data%20repository/data_repository.dart';
import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:meta/meta.dart';

part 'chat_event.dart';
part 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final ChatRepositoryImpl chatRepositoryImpl;

  ChatBloc({required this.chatRepositoryImpl}) : super(ChatInitial()) {
    on<LoadCurrentChats>(_onLoadCurrentChats);
    // on<LoadChatWithUser>(_onLoadChatWithUser);
    on<CreateChat>(_onCreateChat);
    on<DeleteChat>(_onDeleteChat);
  }

  Future<void> _onLoadCurrentChats(
      LoadCurrentChats event, Emitter<ChatState> emit) async {
    emit(LoadingCurrentChats());
    try {
      final Either<Failure, List<ChatEntity>> result =
          await chatRepositoryImpl.getAllChats();
      result.fold(
        (failure) => emit(ChatError(failure.message)),
        (chats) => emit(CurrentChatsLoaded(chats)),
      );
    } catch (e) {
      emit(ChatError(
          'An unexpected error occurred while loading current chats.'));
    }
  }

  // Future<void> _onLoadChatWithUser(
  //     LoadChatWithUser event, Emitter<ChatState> emit) async {
  //   emit(LoadingChatWithUser());
  //   try {
  //     final Either<Failure, List<MessageEntity>> result = await dataRepository.getChatById(event.userId);
  //     result.fold(
  //       (failure) => emit(ChatError(failure.message)),
  //       (messages) => emit(ChatWithUserLoaded(messages)),
  //     );
  //   } catch (e) {
  //     emit(ChatError('An unexpected error occurred while loading chat with user.'));
  //   }
  // }




  Future<void> _onCreateChat(CreateChat event, Emitter<ChatState> emit) async {
    try {
      final Either<Failure, ChatEntity> result =
          await chatRepositoryImpl.createChatById(event.sellerId);
        print(result);
      result.fold(
        (failure) => emit(ChatError(failure.message)),
        (_) => emit(CreateNewChat(event.sellerId)),
      );
    } catch (e) {
      emit(
          ChatError('An unexpected error occurred while creating a new chat.'));
    }
  }

  Future<void> _onDeleteChat(DeleteChat event, Emitter<ChatState> emit) async {
    try {
      final Either<Failure, bool> result =
          await chatRepositoryImpl.deleteChatById(event.chatId);
      result.fold(
        (failure) => emit(ChatError(failure.message)),
        (_) => emit(ChatDeleted(event.chatId)),
      );
    } catch (e) {
      emit(ChatError('An unexpected error occurred while deleting the chat.'));
    }
  }
}
