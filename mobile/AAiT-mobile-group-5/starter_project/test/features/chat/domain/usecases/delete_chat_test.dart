import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:starter_project/core/constants/constants.dart';
import 'package:starter_project/core/errors/failure.dart';
import 'package:starter_project/features/chat/domain/usecases/delete_chat.dart';

import '../../../../helpers/helpers.test.mocks.mocks.dart';

void main() {
  late DeleteChatUsecase deleteChatUsecase;
  late MockChatRepository mockChatRepository;

  setUp(() {
    // Initialize the mock repository and the use case before each test
    mockChatRepository = MockChatRepository();
    deleteChatUsecase = DeleteChatUsecase(mockChatRepository);
  });
  const testChatId = '1';

  test('Should Delete The Chat if It Is Successfull', () async {
    // Arrange: Prepare the mock method response
    when(mockChatRepository.deleteChat(testChatId))
        .thenAnswer((_) async => const Right(null));

    // Act: Call the use case method
    final result =
        await deleteChatUsecase(const DeleteChatParams(chatId: testChatId));

    // Assert: Verify the results
    expect(result,
        const Right(null)); // Expect the result to be a successful Right(null)
    verify(mockChatRepository.deleteChat(
        testChatId)); // Verify that the deleteChat method was called with the correct parameter
    verifyNoMoreInteractions(
        mockChatRepository); // Ensure no other interactions with the mock repository
  });

  test('Should Return Failure Id The Server Is Fail', () async {
    // Arrange: Prepare the mock method to return a failure
    when(mockChatRepository.deleteChat(testChatId)).thenAnswer(
        (_) async => const Left(ServerFailure(message: Messages.serverError)));

    // Act: Call the use case method
    final result =
        await deleteChatUsecase(const DeleteChatParams(chatId: testChatId));

    // Assert: Verify the results
    expect(
        result,
        const Left(ServerFailure(
            message: Messages
                .serverError))); // Expect the result to be a failure (ServerFailure)
    verify(mockChatRepository.deleteChat(
        testChatId)); // Verify that the deleteChat method was called with the correct parameter
    verifyNoMoreInteractions(
        mockChatRepository); // Ensure no other interactions with the mock repository
  });
}
