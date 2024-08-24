import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:starter_project/core/constants/constants.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/chat/domain/usecases/get_single_chat.dart';
import 'package:starter_project/features/shared/entities/user.dart';

import '../../../../helpers/helpers.test.mocks.mocks.dart';

void main() {
  late GetSingleChat getSingleChatUsecase;
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
    getSingleChatUsecase = GetSingleChat(mockChatRepository);
  });
  const testID = '66c84151b7068ee15142f817';
  const user1 = User(
    name: "string",
    email: "cat@gmail.com",
    password: "password",
  );

  const user2 = User(
    name: "Mr. User",
    email: "user@gmail.com",
    password: "password",
  );

  const chat = Chat(
    chatId: "66c84151b7068ee15142f817",
    user1: user1,
    user2: user2,
  );

  test('should get a single chat from the repository', () async {
    // arrange
    when(mockChatRepository.getSingleChat(any))
        .thenAnswer((_) async => const Right(chat));

    // act
    final result =
        await getSingleChatUsecase(const GetChatParams(chatId: testID));

    // assert
    expect(result, const Right(chat));
    verify(mockChatRepository.getSingleChat(testID));
    verifyNoMoreInteractions(mockChatRepository);
  });

  test('should return failure when repository fails', () async {
    // arrange
    when(mockChatRepository.getSingleChat(any)).thenAnswer(
        (_) async => const Left(ServerFailure(message: Messages.serverError)));

    // act
    final result =
        await getSingleChatUsecase(const GetChatParams(chatId: testID));

    // assert
    expect(result, const Left(ServerFailure(message: Messages.serverError)));
    verify(mockChatRepository.getSingleChat(testID));
    verifyNoMoreInteractions(mockChatRepository);
  });
}
