import 'package:bloc/bloc.dart';
import 'package:socket_io_client/socket_io_client.dart' as IO;

import '../../data/data_resources/socket_io_sesrvice.dart';
import '../../domain/usecases/AcknowledgeMessageDeliveryUseCase.dart';
import '../../domain/usecases/CreateChatRoomUseCase.dart';
import '../../domain/usecases/OnMessageReceivedUseCase.dart';
import '../../domain/usecases/RetrieveChatRoomsUseCase.dart';
import '../../domain/usecases/RetrieveMessagesUseCase.dart';
import '../../domain/usecases/SendMessageUseCase.dart';
import 'chat_event.dart';
import 'chat_state.dart';

class ChatBloc extends Bloc<ChatEvent, ChatState> {
  final AcknowledgeMessageDeliveryUseCase acknowledgeMessageDeliveryUseCase;
  final CreateChatRoomUseCase createChatRoomUseCase;
  final OnMessageReceivedUseCase onMessageReceivedUseCase;
  final RetrieveChatRoomsUseCase retrieveChatRoomsUseCase;
  final RetrieveMessagesUseCase retrieveMessagesUseCase;
  final SendMessageUseCase sendMessageUseCase;
  final SocketIOService socketIOService;

  ChatBloc(
    this.acknowledgeMessageDeliveryUseCase,
    this.createChatRoomUseCase,
    this.onMessageReceivedUseCase,
    this.retrieveChatRoomsUseCase,
    this.retrieveMessagesUseCase,
    this.sendMessageUseCase,
    this.socketIOService,
  ) : super(ChatInitial()) {
    on<LoadChatRooms>(_onLoadChatRooms);
    on<LoadMessages>(_onLoadMessages);
    on<CreateChatRoom>(_onCreateChatRoom);
    on<SendMessage>(_onSendMessage);
    on<AcknowledgeMessageDelivery>(_onAcknowledgeMessageDelivery);
    on<MessageReceived>(_onMessageReceived);

    on<ConnectServerEvent>(_connectserver);
  }

  Future<void> _onLoadChatRooms(
    LoadChatRooms event,
    Emitter<ChatState> emit,
  ) async {
    emit(ChatLoading());
    try {
      final chats = await retrieveChatRoomsUseCase.call();
      emit(ChatLoaded(chats));
    } catch (e) {
      emit(ChatError('Failed to load chat rooms'));
    }
  }

  Future<void> _onLoadMessages(
    LoadMessages event,
    Emitter<ChatState> emit,
  ) async {
    emit(ChatLoading());
    try {
      final messages = await retrieveMessagesUseCase.call(event.chatId);
      emit(MessagesLoaded(messages));
    } catch (e) {
      emit(ChatError('Failed to load messages'));
    }
  }

  Future<void> _onCreateChatRoom(
    CreateChatRoom event,
    Emitter<ChatState> emit,
  ) async {
    emit(ChatLoading());
    try {
      await createChatRoomUseCase.call(event.userId);
      add(LoadChatRooms());
    } catch (e) {
      emit(ChatError('Failed to create chat room'));
    }
  }

  Future<void> _onSendMessage(
    SendMessage event,
    Emitter<ChatState> emit,
  ) async {
    try {
      await sendMessageUseCase.call(event.chatId, event.content, event.type);
    } catch (e) {
      emit(ChatError('Failed to send message'));
    }
  }

  Future<void> _onAcknowledgeMessageDelivery(
    AcknowledgeMessageDelivery event,
    Emitter<ChatState> emit,
  ) async {
    try {
      await acknowledgeMessageDeliveryUseCase.call(event.messageId);
    } catch (e) {
      emit(ChatError('Failed to acknowledge message delivery'));
    }
  }

  void _onMessageReceived(
    MessageReceived event,
    Emitter<ChatState> emit,
  ) {
    add(MessageReceived(event.message));
  }

  void _connectserver(
    ConnectServerEvent event,
    Emitter<ChatState> emit,
  ) {
    socketIOService.connect();

    socketIOService.socket.onConnect((_) {
      print('Connected to the socket server');
    });
  }
}
