import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/features/auth/domain/entities/user_data_entity.dart';
import 'package:product_8/features/chat/domain/entities/chat_entity.dart';
import 'package:product_8/features/chat/domain/usecases/initiate_chat_use_case.dart';
import '../../../helpers/test_helper.mocks.dart';

void main() {
  late InitiateChatUseCase initiatechatusecase;
  late MockChatRepository chatrepository;
  late ChatEntity check;

  setUp(() {
    chatrepository = MockChatRepository();
    initiatechatusecase = InitiateChatUseCase(chatrepository);
    check = const ChatEntity(	
			'',
      user1: UserDataEntity(email: '', name: ''),
      user2: UserDataEntity(email: '', name: ''),
    );
  });

  test('Should test the initiate chats use case', () async {
    when(chatrepository.myChatById(any)).thenAnswer((_) async => Right(check));
    await initiatechatusecase.call(const InitiateChatParams(sellerId: 'chatid'));
    verify(chatrepository.myChatById('chatid'));
  });
}
