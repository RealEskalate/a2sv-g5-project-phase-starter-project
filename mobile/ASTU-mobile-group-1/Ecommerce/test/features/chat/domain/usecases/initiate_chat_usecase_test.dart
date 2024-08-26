import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/chat/domain/entities/chat_entity.dart';
import 'package:product_6/features/chat/domain/usecases/initiate_chat_usecase.dart';

import '../../helpers/test_helpers.mocks.dart';

void main() {
  late InitiateChatUsecase initiateChatUsecase;
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
    initiateChatUsecase =
        InitiateChatUsecase(chatRepository: mockChatRepository);
  });

  const String userId = '66c72bd1fc1a63830d084348';

  const UserEntity userEntity1 = UserEntity(
    id: userId,
    name: 'string',
    email: 'm@gmail.com',
  );

  const UserEntity userEntity2 = UserEntity(
    id: '66bde36e9bbe07fc39034cdd',
    name: 'Mr. User',
    email: 'user@gmail.com',
  );

  const ChatEntity chatEntity = ChatEntity(
      chatId: '66c767d7944d8f950440bd9e',
      user1: userEntity1,
      user2: userEntity2);

  test('should test if the usecase is sending data to the repository properly',
      () async {
    //arrange
    when(mockChatRepository.initiateChat(userId))
        .thenAnswer((_) async => const Right(chatEntity));

    //act
    final result = await initiateChatUsecase.execute(userId);

    //assert
    expect(result, const Right(chatEntity));
  });
}
