import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';
import 'package:ecom_app/features/auth/domain/entities/user_data_entity.dart';
import 'package:ecom_app/features/auth/domain/usecases/get_user.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late GetUserUsecase getUserUsecase;
  late MockAuthRepository mockAuthRepository;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    getUserUsecase = GetUserUsecase(mockAuthRepository);
  });

  final tUserDataEntity = UserDataEntity(email: 'email', name: 'name');

  group('LogoutUsecase', () {
    test('should return userdataentity if successful', () async {
      //arrange
      when(mockAuthRepository.getUser())
          .thenAnswer((_) async => Right(tUserDataEntity));

      //act
      final result = await getUserUsecase(NoParams());

      //assert
      expect(result, Right(tUserDataEntity));
    });
    test('should return failure if unsuccessful', () async {
      //arrange
      when(mockAuthRepository.getUser()).thenAnswer(
          (_) async => const Left(ServerFailure('test error message')));

      //act
      final result = await getUserUsecase(NoParams());

      //assert
      expect(result, const Left(ServerFailure('test error message')));
    });
  });
}
