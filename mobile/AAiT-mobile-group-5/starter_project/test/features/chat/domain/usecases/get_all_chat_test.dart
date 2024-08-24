import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:starter_project/core/constants/constants.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/core/usecases/usecases.dart';
import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
import 'package:starter_project/features/chat/domain/usecases/get_all_chat.dart';
import 'package:starter_project/features/shared/entities/user.dart';

import '../../../../helpers/helpers.test.mocks.mocks.dart';

void main() {
  late GetAllChatUsecase getAllChatUsecase;
  late MockChatRepository mockChatRepository;

  setUp(() {
    mockChatRepository = MockChatRepository();
    getAllChatUsecase = GetAllChatUsecase(mockChatRepository);
  });

  const user1 =
      User(name: "Abebe", email: "Abbe@gmail.com", password: "12345678");
  const user2 =
      User(name: "Abebe2", email: "Abbe2@gmail.com", password: "12345678");

  const chatList = [
    Chat(chatId: '1', user1: user1, user2: user2),
    Chat(chatId: '2', user1: user1, user2: user2)
  ];

  test('Should Get All Chat List From Repository', () async {
    // arrange and set up the return type as the test data
    when(mockChatRepository.getAllChat())
        .thenAnswer((_) async => const Right(chatList));

    // act
    //  accept the actual  returndata from the
    final result = await getAllChatUsecase(NoParams());

    // assert
    expect(result, const Right(chatList));
    verify(mockChatRepository.getAllChat());
    verifyNoMoreInteractions(mockChatRepository);
  });

  test('should return failure when repository fails', () async {
    // arrange
    when(mockChatRepository.getAllChat()).thenAnswer(
        (_) async => const Left(ServerFailure(message: Messages.serverError)));

    // act
    final result = await getAllChatUsecase(NoParams());

    // assert
    expect(result, const Left(ServerFailure(message: Messages.serverError)));
    verify(mockChatRepository.getAllChat());
    verifyNoMoreInteractions(mockChatRepository);
  });
}
