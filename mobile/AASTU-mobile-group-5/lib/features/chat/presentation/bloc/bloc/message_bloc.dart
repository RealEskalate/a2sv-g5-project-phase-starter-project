import 'package:bloc/bloc.dart';
import 'package:equatable/equatable.dart';
import 'package:flutter_bloc/flutter_bloc.dart';
import 'package:dartz/dartz.dart';

import '../../../../../core/failure/failure.dart';
import '../../../domain/entities/chat_message_entity.dart';
import '../../../domain/use_case/delete_chat_messages.dart';
import '../../../domain/use_case/load_chat_messages.dart';
import '../../../domain/use_case/send_chat_messages.dart';

part 'message_event.dart';
part 'message_state.dart';




class MessageBloc extends Bloc<MessageEvent, MessageState> {
  final LoadChatMessagesUseCase loadChatMessagesUseCase;
  final SendChatMessageUseCase sendChatMessageUseCase;
  final DeleteChatMessageUseCase deleteChatMessageUseCase;

  MessageBloc({
    required this.loadChatMessagesUseCase,
    required this.sendChatMessageUseCase,
    required this.deleteChatMessageUseCase,
  }) : super(MessageInitial()) {
    on<LoadMessages>(_onLoadMessages);
    on<SendMessage>(_onSendMessage);
    on<ReceiveMessage>(_onReceiveMessage);
    on<DeleteMessage>(_onDeleteMessage);
  }

  void _onLoadMessages(LoadMessages event, Emitter<MessageState> emit) async {
    emit(MessageLoading());

    final Either<Failure, List<ChatMessageEntity>> result =
        await loadChatMessagesUseCase(LoadParams(chatId: event.chatId));

    result.fold(
      (failure) => emit(MessageError("Failed to load messages")),
      (messages) => emit(MessageLoaded(messages)),
    );
  }

  void _onSendMessage(SendMessage event, Emitter<MessageState> emit) async {
    final Either<Failure, void> result =
        await sendChatMessageUseCase(SendParams(chatId: event.chatId, message: event.message));

    result.fold(
      (failure) => emit(MessageError("Failed to send message")),
      (_) => emit(MessageSent()),
    );
  }

  void _onReceiveMessage(ReceiveMessage event, Emitter<MessageState> emit) {
    emit(MessageReceived(event.message));
  }

  void _onDeleteMessage(DeleteMessage event, Emitter<MessageState> emit) async {
    final Either<Failure, void> result =
        await deleteChatMessageUseCase(DeleteParams(messageId: event.messageId));

    result.fold(
      (failure) => emit(MessageError("Failed to delete message")),
      (_) => emit(MessageDeleted()),
    );
  }
}
