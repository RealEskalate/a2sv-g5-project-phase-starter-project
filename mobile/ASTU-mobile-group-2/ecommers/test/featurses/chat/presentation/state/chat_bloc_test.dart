import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecommers/core/Error/failure.dart';
import 'package:ecommers/features/chat/domain/entity/chat_entity.dart';
import 'package:ecommers/features/chat/domain/entity/message_entity.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_bloc/chat_bloc.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_event/chat_event.dart';
import 'package:ecommers/features/chat/presentation/bloc/chat_state/chat_state.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_hlper.mocks.dart';



void main() {
  late MockChatUsecase mockChatUsecase;
  late ChatBloc chatBloc;

  setUp(() {
    mockChatUsecase = MockChatUsecase();
    chatBloc = ChatBloc(chatUsecase: mockChatUsecase);
  });

  test(
    'test the initial state of the app',
    () {
      expect(chatBloc.state, ChatInitialState());
    },
  );

  MessageEntity messageEntity = MessageEntity(
    messageId: 'msg1',
    messages: [
      {'sender': 'John', 'text': 'Hello!'},
      {'sender': 'Doe', 'text': 'Hi there!'}
    ],
  );

  ChatEntity chatEntity = ChatEntity(
    senderId: '123',
    senderName: 'John',
    recieverId: '456',
    recieverName: 'Doe',
    chatId: 'chat1',
    messages: messageEntity,
  );

  List<ChatEntity> chatList = [
    ChatEntity(
      senderId: '123',
      senderName: 'John',
      recieverId: '456',
      recieverName: 'Doe',
      chatId: 'chat1',
      messages: messageEntity,
    ),
    ChatEntity(
      senderId: '789',
      senderName: 'Alice',
      recieverId: '456',
      recieverName: 'Doe',
      chatId: 'chat2',
      messages: MessageEntity(
        messageId: 'msg2',
        messages: [
          {'sender': 'Alice', 'text': 'Hi!'},
        ],
      ),
    ),
  ];

  group('test the state of the app', () {
    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatByIdSuccess] when fetching chat by ID succeeds.',
      build: () {
        when(mockChatUsecase.getChatById(any))
            .thenAnswer((_) async => Right(chatEntity));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnGetChatById(chatId: 'chat1')),
      expect: () => [
        ChatLoadingState(),
        ChatByIdSuccess(chatEntity: chatEntity),
      ],
    );

    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatErrorState] when fetching chat by ID fails.',
      build: () {
        when(mockChatUsecase.getChatById(any))
            .thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnGetChatById(chatId: 'chat1')),
      expect: () => [
        ChatLoadingState(),
        ChatErrorState(errorMessage: 'try again'),
      ],
    );

    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatMessageGetSuccess] when fetching all chats succeeds.',
      build: () {
        when(mockChatUsecase.getMychats())
            .thenAnswer((_) async => Right(chatList));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnGetAllChat()),
      expect: () => [
        ChatLoadingState(),
        ChatMessageGetSuccess(chatEntity: chatList),
      ],
    );

    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatErrorState] when fetching all chats fails.',
      build: () {
        when(mockChatUsecase.getMychats())
            .thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnGetAllChat()),
      expect: () => [
        ChatLoadingState(),
        ChatErrorState(errorMessage: 'try again'),
      ],
    );

    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatDeleteSuccess] when deleting a chat succeeds.',
      build: () {
        when(mockChatUsecase.deleteChats(any))
            .thenAnswer((_) async => const Right(true));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnDeleteChat(chatId: 'chat1')),
      expect: () => [
        ChatLoadingState(),
        ChatDeleteSuccess(chatDeleted: true),
      ],
    );

    blocTest<ChatBloc, ChatState>(
      'emits [ChatLoadingState, ChatErrorState] when deleting a chat fails.',
      build: () {
        when(mockChatUsecase.deleteChats(any))
            .thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
        return chatBloc;
      },
      act: (bloc) => bloc.add(OnDeleteChat(chatId: 'chat1')),
      expect: () => [
        ChatLoadingState(),
        ChatErrorState(errorMessage: 'try again'),
      ],
    );
  });
}
