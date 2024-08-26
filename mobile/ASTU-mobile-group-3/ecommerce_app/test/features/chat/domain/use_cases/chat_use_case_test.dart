import 'package:ecommerce_app/features/auth/domain/entities/user_entity.dart';
import 'package:ecommerce_app/features/chat/domain/entity/chat.dart';
import 'package:ecommerce_app/features/chat/domain/entity/message.dart';
import 'package:ecommerce_app/features/chat/domain/repository/chat_repository.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/AcknowledgeMessageDeliveryUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/CreateChatRoomUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveChatRoomsUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/RetrieveMessagesUseCase.dart';
import 'package:ecommerce_app/features/chat/domain/usecases/SendMessageUseCase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';

// Mock class for ChatRepository
// class MockChatRepository extends Mock implements ChatRepository {}

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
    test('should acknowledge message delivery', () async {
      // Arrange
      const messageId = 'test-message-id';
      when(mockChatRepository.acknowledgeMessageDelivery(any))
          .thenAnswer((_) async => {});

      // Act
      await acknowledgeMessageDeliveryUseCase.call(messageId);

      // Assert
      verify(mockChatRepository.acknowledgeMessageDelivery(messageId))
          .called(1);
    });
  });

  group('CreateChatRoomUseCase', () {
    test('should create a chat room', () async {
      // Arrange
      final user1 = UserEntity(
          id: 'user1',
          name: 'User One',
          email: 'user1@example.com',
          password: 'password',
          v: 1);
      final user2 = UserEntity(
          id: 'user2',
          name: 'User Two',
          email: 'user2@example.com',
          password: 'password',
          v: 1);
      final chat = ChatEntity(chatId: 'chat1', user1: user1, user2: user2);
      when(mockChatRepository.createChatRoom(any)).thenAnswer((_) async => {});

      // Act
      await createChatRoomUseCase.call(chat);

      // Assert
      verify(mockChatRepository.createChatRoom()).called(1);
    });
  });

  group('RetrieveChatRoomsUseCase', () {
    test('should retrieve chat rooms', () async {
      // Arrange
      final user1 = UserEntity(
          id: 'user1',
          name: 'User One',
          email: 'user1@example.com',
          password: 'password',
          v: 1);
      final user2 = UserEntity(
          id: 'user2',
          name: 'User Two',
          email: 'user2@example.com',
          password: 'password',
          v: 1);
      final chatRooms = [
        ChatEntity(chatId: 'chat1', user1: user1, user2: user2)
      ];
      when(mockChatRepository.retrieveChatRooms())
          .thenAnswer((_) async => chatRooms);

      // Act
      final result = await retrieveChatRoomsUseCase.call();

      // Assert
      expect(result, chatRooms);
      verify(mockChatRepository.retrieveChatRooms()).called(1);
    });
  });
}
