import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/auth/domain/entities/register_entity.dart';
import 'package:ecom_app/features/auth/domain/usecases/register.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late RegisterUsecase registerUsecase;
  late MockAuthRepository mockAuthRepository;

  setUp((){
    mockAuthRepository = MockAuthRepository();
    registerUsecase = RegisterUsecase(mockAuthRepository);
  });

   final tRegisterEntity = RegistrationEntity(
    email: 'sd@gmail.com',
    password: '123456', name: 'Simon',);

  group('RegisterUsecase', () {

    test('should return unit if successful', () async {
      //arrange
      when(mockAuthRepository.register(tRegisterEntity))
          .thenAnswer((_) async => const Right(unit));

      //act
      final result = await registerUsecase(RegisterParams(registrationEntity: tRegisterEntity));

      //assert
      expect(result, const Right(unit));
    });
    test('should return failure if unsuccessful', () async {
      //arrange
      when(mockAuthRepository.register(tRegisterEntity))
          .thenAnswer((_) async => const Left(ServerFailure('test error message')));

      //act
      final result = await registerUsecase(RegisterParams(registrationEntity: tRegisterEntity));

      //assert
      expect(result, const Left(ServerFailure('test error message')));
    });

  });
}