import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/chat/domain/entities/chat_entity.dart';
import 'package:product_6/features/chat/domain/entities/message_entity.dart';
import 'package:product_6/features/chat/domain/usecases/get_chat_message_usecase.dart';

import '../../helpers/test_helpers.mocks.dart';

void main() {
  late MockChatRepository mockChatRepository;
  late GetChatMessageUsecase getChatMessageUsecase;

  setUp(() {
    mockChatRepository = MockChatRepository();
    getChatMessageUsecase =
        GetChatMessageUsecase(chatRepository: mockChatRepository);
  });

  const String chatId = '66c767d7944d8f950440bd9e';

  const UserEntity userEntity1 = UserEntity(
      id: '66c72bd1fc1a63830d084348',
      name: 'string',
      email: 'm@gmail.com',
      accessToken: 'accessToken');

  const UserEntity userEntity2 = UserEntity(
      id: '66bde36e9bbe07fc39034cdd',
      name: 'Mr. User',
      email: 'user@gmail.com',
      accessToken: 'accessToken');

  const ChatEntity chatEntity =
      ChatEntity(chatId: chatId, user1: userEntity1, user2: userEntity2);

  const MessageEntity messageEntity = MessageEntity(
      messageId: '66c872f6346255dac2604bab',
      chatEntity: chatEntity,
      content: 'Hello',
      type: 'text',
      sender: userEntity1);

  const List<MessageEntity> messageEntityList = [messageEntity];

  test('should test if the usecase is correctly sending data to the repository',
      () async {
    //arrange
    when(mockChatRepository.getChatMessages(chatId))
        .thenAnswer((_) async => const Right(messageEntityList));

    //act
    final result = await getChatMessageUsecase.execute(chatId);

    //assert
    expect(result, const Right(messageEntityList));
  });
}
