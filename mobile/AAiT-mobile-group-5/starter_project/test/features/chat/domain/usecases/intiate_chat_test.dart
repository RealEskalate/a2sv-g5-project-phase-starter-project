// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:starter_project/core/constants/constants.dart';
// import 'package:starter_project/core/errors/failure.dart';
// import 'package:starter_project/features/chat/domain/entities/chat_entity.dart';
// import 'package:starter_project/features/chat/domain/usecases/intiate_chat.dart';
// import 'package:starter_project/features/shared/entities/user.dart';

// import '../../../../helpers/helpers.test.mocks.mocks.dart';

// void main() {
//   late InitiateChatUsecase initiateChatUsecase;
//   late MockChatRepository mockChatRepository;

//   setUp(() {
//     mockChatRepository = MockChatRepository();
//     initiateChatUsecase = InitiateChatUsecase(mockChatRepository);
//   });

//   const user1 = User(
//     name: "string",
//     email: "cat@gmail.com",
//     password: "password",
//   );

//   const user2 = User(
//     name: "Mr. User",
//     email: "user@gmail.com",
//     password: "password",
//   );

//   const chat = Chat(
//     chatId: "66c84151b7068ee15142f817",
//     user1: user1,
//     user2: user2,
//   );
//   const testID = '66c730840740f8c2bae904e0';

//   test(
//       'should initiate a chat and return the created chat entity from the repository',
//       () async {
//     // arrange
//     when(mockChatRepository.intiateChat(testID))
//         .thenAnswer((_) async => const Right(chat));

//     // act
//     final result =
//         await initiateChatUsecase(const IntiateChatParams(userId: testID));

//     // assert
//     expect(result, const Right(chat));
//     verify(mockChatRepository.intiateChat(testID));
//     verifyNoMoreInteractions(mockChatRepository);
//   });

//   test('should return failure when repository fails to initiate chat',
//       () async {
//     // arrange
//     when(mockChatRepository.intiateChat(testID)).thenAnswer(
//         (_) async => const Left(ServerFailure(message: Messages.serverError)));

//     // act
//     final result =
//         await initiateChatUsecase(const IntiateChatParams(userId: testID));

//     // assert
//     expect(result, const Left(ServerFailure(message: Messages.serverError)));
//     verify(mockChatRepository.intiateChat(testID));
//     verifyNoMoreInteractions(mockChatRepository);
//   });
// }
