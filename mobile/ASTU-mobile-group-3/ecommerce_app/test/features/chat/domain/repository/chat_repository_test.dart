import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/message.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
  });

  group('ChatRepository Tests', () {
    test('retrieveChatRooms should return a list of ChatEntity', () async {
      // Arrange
      final chatRooms = [
        ChatEntity(chatId: 'chat1', user1: mockUser1, user2: mockUser2),
        ChatEntity(chatId: 'chat2', user1: mockUser1, user2: mockUser2),
      ];

      when(mockChatRepository.retrieveChatRooms())
          .thenAnswer((_) async => chatRooms);

      // Act
      final result = await mockChatRepository.retrieveChatRooms();

      // Assert
      expect(result, chatRooms);
      verify(mockChatRepository.retrieveChatRooms()).called(1);
    });

    test('retrieveMessages should return a list of MessageEntity', () async {
      // Arrange
      final chatId = 'chat1';
      final messages = [
        MessageEntity(messageId: 'msg1', content: 'Hello', senderId: 'user1'),
        MessageEntity(messageId: 'msg2', content: 'Hi', senderId: 'user2'),
      ];

      when(mockChatRepository.retrieveMessages(chatId))
          .thenAnswer((_) async => messages);

      // Act
      final result = await mockChatRepository.retrieveMessages(chatId);

      // Assert
      expect(result, messages);
      verify(mockChatRepository.retrieveMessages(chatId)).called(1);
    });

    test('createChatRoom should call repository with correct parameters',
        () async {
      // Arrange
      final chat =
          ChatEntity(chatId: 'chat1', user1: mockUser1, user2: mockUser2);

      when(mockChatRepository.createChatRoom(chat))
          .thenAnswer((_) async => Future.value());

      // Act
      await mockChatRepository.createChatRoom(chat);

      // Assert
      verify(mockChatRepository.createChatRoom(chat)).called(1);
    });

    test('sendMessage should call repository with correct parameters',
        () async {
      // Arrange
      final chatId = 'chat1';
      final message =
          MessageEntity(messageId: 'msg1', content: 'Hello', senderId: 'user1');

      when(mockChatRepository.sendMessage(chatId, message))
          .thenAnswer((_) async => Future.value());

      // Act
      await mockChatRepository.sendMessage(chatId, message);

      // Assert
      verify(mockChatRepository.sendMessage(chatId, message)).called(1);
    });

    test(
        'acknowledgeMessageDelivery should call repository with correct parameters',
        () async {
      // Arrange
      final messageId = 'msg1';

      when(mockChatRepository.acknowledgeMessageDelivery(messageId))
          .thenAnswer((_) async => Future.value());

      // Act
      await mockChatRepository.acknowledgeMessageDelivery(messageId);

      // Assert
      verify(mockChatRepository.acknowledgeMessageDelivery(messageId))
          .called(1);
    });

    test('onMessageReceived should return a stream of MessageEntity', () {
      // Arrange
      final message =
          MessageEntity(messageId: 'msg1', content: 'Hello', senderId: 'user1');
      final messageStream = Stream<MessageEntity>.fromIterable([message]);

      when(mockChatRepository.onMessageReceived())
          .thenAnswer((_) => messageStream);

      // Act
      final result = mockChatRepository.onMessageReceived();

      // Assert
      expect(result, emitsInOrder([message]));
    });
  });
}
