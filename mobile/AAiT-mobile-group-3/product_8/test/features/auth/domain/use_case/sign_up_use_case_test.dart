// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/failure/failure.dart';
// import 'package:product_8/features/auth/domain/entities/sign_up_user_entitiy.dart';
// import 'package:product_8/features/auth/domain/use_case/sign_up_use_case.dart';

// import '../../../../helpers/test_helper.mocks.dart';

// void main() {
//   late SignUpUseCase signUpUseCase;
//   late MockAuthRepository mockAuthRepository;

//   setUp(() {
//     mockAuthRepository = MockAuthRepository();
//     signUpUseCase = SignUpUseCase(authRepository: mockAuthRepository);
//   });

//   const SignUpUserEntitiy signUpUserEntitiy = SignUpUserEntitiy(
//     name: 'John Doe',
//     email: 'test@example.com',
//     password: 'password123',
//   );

//   group('Sign up test', () {
//      test('should sign up user successfully', () async {
//     // Arrange
//     when(mockAuthRepository.signUp(any))
//         .thenAnswer((_) async => const Right(signUpUserEntitiy));

//     // Act
//     final result = await signUpUseCase(signUpUserEntitiy);

//     // Assert
//     expect(result, const Right(signUpUserEntitiy));
//     verify(mockAuthRepository.signUp(signUpUserEntitiy));
//     verifyNoMoreInteractions(mockAuthRepository);
//   });

//   test('should return a failure when sign up fails', () async {
//     // Arrange
//     const failure = AuthorizationFailure(message: 'Email already in use');
//     when(mockAuthRepository.signUp(any))
//         .thenAnswer((_) async => const Left(failure));

//     // Act
//     final result = await signUpUseCase(signUpUserEntitiy);

//     // Assert
//     expect(result, const Left(failure));
//     verify(mockAuthRepository.signUp(signUpUserEntitiy));
//     verifyNoMoreInteractions(mockAuthRepository);
//   });
//   });
// }
