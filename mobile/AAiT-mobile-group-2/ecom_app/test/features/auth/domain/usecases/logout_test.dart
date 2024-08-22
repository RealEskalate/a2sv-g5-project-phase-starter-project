import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';
import 'package:ecom_app/features/auth/domain/usecases/logout.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late LogoutUsecase logoutUsecase;
  late MockAuthRepository mockAuthRepository;

  setUp((){
    mockAuthRepository = MockAuthRepository();
    logoutUsecase = LogoutUsecase(mockAuthRepository);
  });

   

  group('LogoutUsecase', () {

    test('should return unit if successful', () async {
      //arrange
      when(mockAuthRepository.logout())
          .thenAnswer((_) async => const Right(unit));

      //act
      final result = await logoutUsecase(NoParams());

      //assert
      expect(result, const Right(unit));
    });
    test('should return failure if unsuccessful', () async {
      //arrange
      when(mockAuthRepository.logout())
          .thenAnswer((_) async => const Left(ServerFailure('test error message')));

      //act
      final result = await logoutUsecase(NoParams());

      //assert
      expect(result, const Left(ServerFailure('test error message')));
    });

  });
}