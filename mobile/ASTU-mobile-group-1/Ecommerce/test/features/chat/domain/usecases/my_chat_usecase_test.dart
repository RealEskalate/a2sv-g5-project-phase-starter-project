import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/chat/domain/entities/chat_entity.dart';
import 'package:product_6/features/chat/domain/usecases/my_chat_usecase.dart';
import 'package:dartz/dartz.dart';

import '../../helpers/test_helpers.mocks.dart';

void main() {
  late MyChatUsecase myChatUsecase;
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
    myChatUsecase = MyChatUsecase(chatRepository: mockChatRepository);
  });

  const UserEntity userEntity1 = UserEntity(
    id: '66c72bd1fc1a63830d084348',
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

  const List<ChatEntity> chatEntityList = [chatEntity];
  test(
      'Should test if the usecase is sending the required datas to the repository',
      () async {
    //arrange
    when(mockChatRepository.myChat())
        .thenAnswer((_) async => const Right(chatEntityList));

    //act
    final result = await myChatUsecase.execute();

    //assert

    expect(result, const Right(chatEntityList));
  });
}
