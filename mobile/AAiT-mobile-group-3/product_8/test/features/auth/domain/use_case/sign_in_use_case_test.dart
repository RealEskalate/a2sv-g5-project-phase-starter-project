// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/failure/failure.dart';
// import 'package:product_8/features/auth/domain/entities/sign_in_user_entitiy.dart';
// import 'package:product_8/features/auth/domain/entities/sign_up_user_entitiy.dart';
// import 'package:product_8/features/auth/domain/entities/user_data_entity.dart';
// import 'package:product_8/features/auth/domain/use_case/sign_in_use_case.dart';

// import '../../../../helpers/test_helper.mocks.dart';

// void main() {
//   late SignInUseCase signInUseCase;
//   late MockAuthRepository mockAuthRepository;

//   setUp(() {
//     mockAuthRepository = MockAuthRepository();
//     signInUseCase = SignInUseCase(authRepository: mockAuthRepository);
//   });

//   const SignInUserEntitiy signInUserEntitiy = SignInUserEntitiy(
//     email: 'test@example.com',
//     password: 'password123',
//   );

//   const Data data = Data(
//     name: 'John Doe',
//     email: 'test@example.com',
   
//   );
//   const String token = 'token';
//   const UserDataEntity userDataEntity = UserDataEntity(data: data, token: token);

//   group('sign in test', () {
//     test('should sign in user successfully', () async {
//     // Arrange
//     when(mockAuthRepository.signIn(any))
//         .thenAnswer((_) async => const Right(userDataEntity));

//     // Act
//     final result = await signInUseCase(signInUserEntitiy);

//     // Assert
//     expect(result, const Right(userDataEntity));
//     verify(mockAuthRepository.signIn(signInUserEntitiy));
//     verifyNoMoreInteractions(mockAuthRepository);
//   });

//    test('should return a failure when sign in fails', () async {
//     // Arrange
//     const failure = AuthenticationFailure(message: 'Invalid credentials');
//     when(mockAuthRepository.signIn(any))
//         .thenAnswer((_) async => const Left(failure));

//     // Act
//     final result = await signInUseCase(signInUserEntitiy);

//     // Assert
//     expect(result, const Left(failure));
//     verify(mockAuthRepository.signIn(signInUserEntitiy));
//     verifyNoMoreInteractions(mockAuthRepository);
//   });


//   });
// }
