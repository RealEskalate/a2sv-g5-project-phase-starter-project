import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/features/auth/domain/entities/user_data_entity.dart';
import 'package:product_8/features/chat/domain/entities/chat_entity.dart';
import 'package:product_8/features/chat/domain/usecases/my_chat_by_id_use_case.dart';
import '../../../helpers/test_helper.mocks.dart';

void main() {
  late MyChatByIdUseCase mychatusecase;
  late MockChatRepository chatrepository;
  late ChatEntity check;

  setUp(() {
    chatrepository = MockChatRepository();
    mychatusecase = MyChatByIdUseCase(chatrepository);
    check = const ChatEntity(	
			'',
      user1: UserDataEntity(email: '', name: ''),
      user2: UserDataEntity(email: '', name: ''),
    );
  });

  test('Should test the my chats use case', () async {
    when(chatrepository.myChatById(any)).thenAnswer((_) async => Right(check));
    await mychatusecase.call(const MyChatByIdParams(chatId: 'chatid'));
    verify(chatrepository.myChatById('chatid'));
  });
}
