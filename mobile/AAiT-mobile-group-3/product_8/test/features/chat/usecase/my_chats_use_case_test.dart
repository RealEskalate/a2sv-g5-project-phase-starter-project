import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/usecase/usecase.dart';
import 'package:product_8/features/chat/domain/entities/chat_entity.dart';
import 'package:product_8/features/chat/domain/usecases/my_chats_use_case.dart';
import '../../../helpers/test_helper.mocks.dart';

void main() {
  late MyChatsUseCase mychatusecase;
  late MockChatRepository chatrepository;
	late List<ChatEntity> check = [];

  setUp(() {
    chatrepository = MockChatRepository();
    mychatusecase = MyChatsUseCase(chatrepository);
  });

  test('Should test the my chats use case', () async {

		when(chatrepository.myChats()).thenAnswer((_) async => Right(check));
		await mychatusecase.call(NoParams());
		verify(chatrepository.myChats());
	});
}
