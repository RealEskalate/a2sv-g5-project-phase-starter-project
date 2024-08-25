import 'dart:math';

import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/features/chat/domain/usecases/delete_chat_usecase.dart';

import '../../helpers/test_helpers.mocks.dart';

void main() {
  late DeleteChatUsecase deleteChatUsecase;
  late MockChatRepository mockChatRepository;

  const String chatId = '66c767d7944d8f950440bd9e';

  setUp(() {
    mockChatRepository = MockChatRepository();
    deleteChatUsecase = DeleteChatUsecase(chatRepository: mockChatRepository);
  });

  test('check if the data is succesfully passing to the repository', () async {
    //arrange

    when(mockChatRepository.deleteChat(chatId))
        .thenAnswer((_) async => const Right(true));

    //act
    final result = await deleteChatUsecase.execute(chatId);

    //assert
    expect(result, const Right(true));
  });
}
