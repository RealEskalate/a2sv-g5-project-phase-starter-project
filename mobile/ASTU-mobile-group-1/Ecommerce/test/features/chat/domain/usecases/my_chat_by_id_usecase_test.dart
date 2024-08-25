import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/chat/domain/entities/chat_entity.dart';
import 'package:product_6/features/chat/domain/usecases/my_chat_by_id_usecase.dart';

import '../../helpers/test_helpers.mocks.dart';

void main() {
  late MyChatById myChatById;
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
    myChatById = MyChatById(chatRepository: mockChatRepository);
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

  const ChatEntity chatEntity = ChatEntity(
      chatId: '66c767d7944d8f950440bd9e',
      user1: userEntity1,
      user2: userEntity2);

  test('should test if the data is being sent to the repository', () async {
    //arrange
    when(mockChatRepository.myChatbyId(chatId))
        .thenAnswer((_) async => const Right(chatEntity));

    //act
    final result = await myChatById.execute(chatId);

    //assert
    expect(result, const Right(chatEntity));
  });
}
