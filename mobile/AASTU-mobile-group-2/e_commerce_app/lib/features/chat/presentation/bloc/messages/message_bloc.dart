import 'package:e_commerce_app/features/chat/domain/entities/chat_entity.dart';
import 'package:e_commerce_app/features/chat/domain/entities/message_entity.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';

import '../../../domain/usecases/get_chat_by_user_id_usecase.dart';
import '../../../domain/usecases/get_messages_by_id.dart';
import '../../../domain/usecases/send_message_usecase.dart';
import 'package:bloc_concurrency/bloc_concurrency.dart';
part 'message_event.dart';
part 'message_state.dart';

class MessageBloc extends Bloc<MessageEvent, MessageState> {
  final GetMessagesByIdUsecase getMessagesById;
  final SendMessageUseCase sendMessage;

  MessageBloc({
    required this.getMessagesById,
    required this.sendMessage,
  }) : super(const MessageInitial([])) {
    on<MessageSocketConnectionRequested>(_onLoadMessageRequested,
        transformer: restartable());
    on<MessageSent>(_onSendMessageRequested);
  }

  Future<void> _onLoadMessageRequested(MessageSocketConnectionRequested event,
      Emitter<MessageState> emit) async {
    emit(const MessagesMessageLoadInProgress([]));

    // final messages =
    //     await getChatMessages(GetChatMessagesParams(event.chat.id));
    final messages = await getMessagesById.execute(event.chat.id);
    emit(messages.fold((l) => MessageLoadFailure([]), (r) {
      return MessageLoadSuccess(r);
    }));
    // await emit.forEach(messages, onData: (data) {
    //   final result = data.fold((l) {
    //     return MessageLoadFailure([...state.messages]);
    //   }, (r) {
    //     if (state.messages.isNotEmpty &&
    //         state.messages[state.messages.length - 1] == r) {
    //       return MessageLoadSuccess(r);
    //     }
    //     return MessageLoadSuccess([...state.messages, r]);
    //   });

    //   return result;
    // });
  }

  Future<void> _onSendMessageRequested(
      MessageSent event, Emitter<MessageState> emit) async {
    emit(MessagesMessageLoadInProgress(state.messages));

    final result = await sendMessage.execute(
        event.chat, event.content, event.type);

    emit(result.fold((l) => MessageSentFailure(state.messages),
        (r) => MessageSentSuccess(state.messages)));
  }
}
