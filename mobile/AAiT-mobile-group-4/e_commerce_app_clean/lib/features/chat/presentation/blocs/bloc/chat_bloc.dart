import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';

import '../../../../../core/usecase/usecase.dart';
import '../../../domain/entities/chat.dart';
import '../../../domain/entities/message.dart';
import '../../../domain/usecases/create_chat_usecase.dart';
import '../../../domain/usecases/get_chat_usecase.dart';
import '../../../domain/usecases/get_chats_usecase.dart';
import '../../../domain/usecases/send_message_usecase.dart';

part 'chat_event.dart';
part 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final GetChatsUsecase getChatsUsecase;
  final GetChatUsecase getChatUsecase;
  final CreateChatUsecase createChatUsecase;
  final SendMessageUsecase sendMessageUsecase;
  ChatBloc({
    required this.getChatUsecase,
    required this.getChatsUsecase,
    required this.createChatUsecase,
    required this.sendMessageUsecase,
  }) : super(ChatInitial()) {
    on<FetchChatsEvent>((event, emit) async {
      emit(ChatsLoadingState());
      var result = await getChatsUsecase(NoParams());
      result.fold(
        (left) => emit(ChatErrorState(message: left.message)), 
        (chats) => emit(ChatsLoadedState(chats: chats)));
    });
    
    on<CreateChatEvent>((event, emit) async {
      emit(ChatMessageLoadingState());
      var result = await createChatUsecase(CreateChatParams(event.userId));
      result.fold(
        (left) => emit(ChatErrorState(message: left.message)), 
        (chat) => emit(ChatCreatedState(chat: chat)));
    });
    
    on<SendMessageEvent>((event, emit) async {
      emit(ChatMessageLoadingState());
      var result = await sendMessageUsecase(SendMessageParams(event.message));
      result.fold(
        (left) => emit(ChatErrorState(message: left.message)), 
        (chat) => emit(ChatMessageSentState()));
    });
    
    on<GetChatMessagesEvent>((event, emit) async {
      emit(ChatMessageLoadingState());
      var stream = await getChatUsecase(GetChatParams(event.chatId));

      List<Message> messages = [];
      await emit.forEach(stream, onData: (result) {
        return result.fold(
          (fialure) => ChatErrorState(message: fialure.message), 
          (message) {
            if (messages.isNotEmpty && messages[messages.length - 1] == message) {
              return ChatMessageLoadedState(messages: messages);
            }
            return ChatMessageLoadedState(messages: messages..add(message));
          }
        );
      });
    });
  }
}
