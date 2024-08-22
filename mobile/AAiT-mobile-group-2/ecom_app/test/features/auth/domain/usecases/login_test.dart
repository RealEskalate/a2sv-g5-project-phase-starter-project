import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/auth/domain/entities/login_entity.dart';
import 'package:ecom_app/features/auth/domain/usecases/login.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late LoginUsecase loginUsecase;
  late MockAuthRepository mockAuthRepository;

  setUp((){
    mockAuthRepository = MockAuthRepository();
    loginUsecase = LoginUsecase(mockAuthRepository);
  });

   final tLoginEntity = LoginEntity(
    email: 'sd@gmail.com',
    password: '123456',);

  group('LoginUsecase', () {

    test('should return unit if successful', () async {
      //arrange
      when(mockAuthRepository.login(tLoginEntity))
          .thenAnswer((_) async => const Right(unit));

      //act
      final result = await loginUsecase(LoginParams(loginEntity: tLoginEntity));

      //assert
      expect(result, const Right(unit));
    });
    test('should return failure if unsuccessful', () async {
      //arrange
      when(mockAuthRepository.login(tLoginEntity))
          .thenAnswer((_) async => const Left(ServerFailure('test error message')));

      //act
      final result = await loginUsecase(LoginParams(loginEntity: tLoginEntity));

      //assert
      expect(result, const Left(ServerFailure('test error message')));
    });

  });
}