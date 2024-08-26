import 'package:equatable/equatable.dart';

import '../../../auth/domain/entities/user_entity.dart';
import '../../../auth/presentation/pages/pages.dart';
import '../../domain/entities/chat_entity.dart';
import '../../domain/entities/message_entity.dart';
import '../../domain/usecases/delete_chat.dart';
import '../../domain/usecases/get_chat_message_usecase.dart';
import '../../domain/usecases/initiate_chat_usecase.dart';
import '../../domain/usecases/my_chat.dart';
import '../../domain/usecases/my_chat_by_id.dart';
import '../../domain/usecases/send_message_usecase.dart';

part 'chat_event.dart';
part 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final MyChat myChatUsecase;
  final InitiateChatUsecase initiateChatUsecase;
  final DeleteChat deleteChatUsecase;
  final GetChatMessageUsecase getChatMessageUsecase;
  final MyChatById myChatByIdUsecase;
  SendMessageUsecase sendMessageUsecase;
  ChatBloc(this.myChatUsecase, this.initiateChatUsecase, this.deleteChatUsecase, this.getChatMessageUsecase, this.myChatByIdUsecase,this.sendMessageUsecase) : super(ChatInitial()) {
    on<GetAllChatEvent>((event, emit) async{
      emit(ChatLoadingState());
      final res = await myChatUsecase.myChat();
      res.fold((error){emit(ChatFailureState(message:error.message ));}, 
      (right){
        emit(LoadedAllChatState([], allChats: right));
      });

    });

    on<InitiateChatEvent>((event,emit) async {
      emit(ChatLoadingState());
      final res = await initiateChatUsecase.initiateChat(event.recieverId);
      res.fold((error){emit(ChatFailureState(message:error.message ));}, 
      (right){
        add(GetChatMessageEvent(chatEntity: right));
      });
    });

    on<GetChatMessageEvent>((event, emit) async {
  emit(ChatLoadingState());
  final res = await getChatMessageUsecase.getChatMessages(event.chatEntity.chatId);
  res.fold(
    (error) {
      emit(ChatFailureState(message: error.message));
    },
    (messages) {
      emit(IndividualChatingState(chatEntity: event.chatEntity,chatMessages: messages));
    },
  );
});

  on<SendMessageEvent>((event,emit)async{
      emit(ChatLoadingState());
      final res = await sendMessageUsecase.send(event.chatEntity.chatId, event.message, event.type);
      res.fold((left){
        emit(ChatFailureState(message: left.message));
      }, (right){
        add(GetChatMessageEvent(chatEntity: event.chatEntity));
      });

  });


  }
}
