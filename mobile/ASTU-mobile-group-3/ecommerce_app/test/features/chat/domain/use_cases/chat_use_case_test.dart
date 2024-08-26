import 'package:ecommerce_app/features/auth/domain/entities/user_entity.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/message.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/AcknowledgeMessageDeliveryUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/CreateChatRoomUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveChatRoomsUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveMessagesUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/SendMessageUseCase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late MockChatRepository mockChatRepository;
  late AcknowledgeMessageDeliveryUseCase acknowledgeMessageDeliveryUseCase;
  late CreateChatRoomUseCase createChatRoomUseCase;
  late RetrieveChatRoomsUseCase retrieveChatRoomsUseCase;
  late RetrieveMessagesUseCase retrieveMessagesUseCase;
  late SendMessageUseCase sendMessageUseCase;

  setUp(() {
    mockChatRepository = MockChatRepository();
    acknowledgeMessageDeliveryUseCase =
        AcknowledgeMessageDeliveryUseCase(mockChatRepository);
    createChatRoomUseCase = CreateChatRoomUseCase(mockChatRepository);
    retrieveChatRoomsUseCase = RetrieveChatRoomsUseCase(mockChatRepository);
    retrieveMessagesUseCase = RetrieveMessagesUseCase(mockChatRepository);
    sendMessageUseCase = SendMessageUseCase(mockChatRepository);
  });

  group('AcknowledgeMessageDeliveryUseCase', () {
    final messageId = 'testMessageId';

    test('should call acknowledgeMessageDelivery on repository', () async {
      // Arrange
      when(mockChatRepository.acknowledgeMessageDelivery(messageId))
          .thenAnswer((_) async => Future.value());

      // Act
      await acknowledgeMessageDeliveryUseCase.call(messageId);

      // Assert
      verify(mockChatRepository.acknowledgeMessageDelivery(messageId))
          .called(1);
    });
  });

  group('CreateChatRoomUseCase', () {
    final userId = 'testUserId';

    test('should call createChatRoom on repository', () async {
      // Arrange
      when(mockChatRepository.createChatRoom(userId))
          .thenAnswer((_) async => Future.value());

      // Act
      await createChatRoomUseCase.call(userId);

      // Assert
      verify(mockChatRepository.createChatRoom(userId)).called(1);
    });
  });

  group('RetrieveChatRoomsUseCase', () {
    final chatRooms = [
      const ChatEntity(
        chatId: 'chat1',
        user1: UserEntity(
            name: 'User1',
            email: 'user1@test.com',
            password: 'pass1',
            id: '1',
            v: 1),
        user2: UserEntity(
            name: 'User2',
            email: 'user2@test.com',
            password: 'pass2',
            id: '2',
            v: 2),
      ),
    ];

    test('should return a list of chat rooms from the repository', () async {
      // Arrange
      when(mockChatRepository.retrieveChatRooms())
          .thenAnswer((_) async => chatRooms);

      // Act
      final result = await retrieveChatRoomsUseCase.call();

      // Assert
      expect(result, chatRooms);
      verify(mockChatRepository.retrieveChatRooms()).called(1);
    });
  });

  group('RetrieveMessagesUseCase', () {
    final chatId = 'chat1';
    final messages = [
      const MessageEntity(
        messageId: 'msg1',
        sender: UserEntity(
            name: 'User1',
            email: 'user1@test.com',
            password: 'pass1',
            id: '1',
            v: 1),
        chat: ChatEntity(
          chatId: 'chat1',
          user1: UserEntity(
              name: 'User1',
              email: 'user1@test.com',
              password: 'pass1',
              id: '1',
              v: 1),
          user2: UserEntity(
              name: 'User2',
              email: 'user2@test.com',
              password: 'pass2',
              id: '2',
              v: 2),
        ),
        content: 'Hello!',
      ),
    ];

    test('should return a list of messages from the repository', () async {
      // Arrange
      when(mockChatRepository.retrieveMessages(chatId))
          .thenAnswer((_) async => messages);

      // Act
      final result = await retrieveMessagesUseCase.call(chatId);

      // Assert
      expect(result, messages);
      verify(mockChatRepository.retrieveMessages(chatId)).called(1);
    });
  });

  group('SendMessageUseCase', () {
    final chatId = 'chat1';
    final content = 'Hello!';
    final type = 'text';

    test('should call sendMessage on repository', () async {
      // Arrange
      when(mockChatRepository.sendMessage(chatId, content, type))
          .thenAnswer((_) async => Future.value());

      // Act
      await sendMessageUseCase.call(chatId, content, type);

      // Assert
      verify(mockChatRepository.sendMessage(chatId, content, type)).called(1);
    });
  });
}
