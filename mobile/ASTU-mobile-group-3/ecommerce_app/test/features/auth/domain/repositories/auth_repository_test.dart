import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/failures/failure.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/log_in_usecase.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/log_out_usecase.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/sign_up_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late LogInUsecase logInUsecase;
  late SignUpUsecase signUpUseCase;
  late LogOutUsecase logOutUsecase;
  setUp(() {
    mockAuthRepository = MockAuthRepository();
    logInUsecase = LogInUsecase(authRepository: mockAuthRepository);
    signUpUseCase = SignUpUsecase(authRepository: mockAuthRepository);
    logOutUsecase = LogOutUsecase(repository: mockAuthRepository);
  });

  group('logIn usecase', () {
    test('Should return true when user signing in succesfully', () async {
      when(mockAuthRepository.logIn(AuthData.userEntity))
          .thenAnswer((_) async => const Right(true));

      final result = await logInUsecase.execute(AuthData.userEntity);

      verify(mockAuthRepository.logIn(any));

      expect(result, const Right(true));
    });

    test('Should return Failure when failud when user signing in fails',
        () async {
      when(mockAuthRepository.logIn(AuthData.userEntity)).thenAnswer(
          (_) async =>
              Left(ServerFailure(AppData.getMessage(AppData.serverError))));

      final result = await logInUsecase.execute(AuthData.userEntity);

      verify(mockAuthRepository.logIn(any));

      expect(
          result, Left(ServerFailure(AppData.getMessage(AppData.serverError))));
    });
  });

  group('Sign up', () {
    test('Should return Failure when  user registration success', () async {
      when(mockAuthRepository.signUp(any))
          .thenAnswer((_) async => const Right(true));

      final result = await signUpUseCase.execute(AuthData.userEntity);

      verify(mockAuthRepository.signUp(any));

      expect(result, const Right(true));
    });
    test('Should return Failure when failud when user registration fails',
        () async {
      when(mockAuthRepository.signUp(any)).thenAnswer((_) async =>
          Left(ServerFailure(AppData.getMessage(AppData.serverError))));

      final result = await signUpUseCase.execute(AuthData.userEntity);

      verify(mockAuthRepository.signUp(any));

      expect(
          result, Left(ServerFailure(AppData.getMessage(AppData.serverError))));
    });
  });

  group('Log out repo test', () {
    test('Logout repository', () async {
      /// arrange
      when(mockAuthRepository.logOut())
          .thenAnswer((_) async => const Right(true));

      /// action
      final result = await logOutUsecase.execute();

      /// assert
      expect(result, const Right(true));
    });

    test('Logout repository', () async {
      /// arrange
      when(mockAuthRepository.logOut()).thenAnswer((_) async =>
          Left(CacheFailure(AppData.getMessage(AppData.cacheError))));

      /// action
      final result = await logOutUsecase.execute();

      /// assert
      expect(
          result, Left(CacheFailure(AppData.getMessage(AppData.cacheError))));
    });
  });
}
