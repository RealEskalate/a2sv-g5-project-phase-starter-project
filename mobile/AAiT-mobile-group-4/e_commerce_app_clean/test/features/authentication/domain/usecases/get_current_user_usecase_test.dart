import 'package:application1/core/error/failure.dart';
import 'package:application1/core/usecase/usecase.dart';
import 'package:application1/features/authentication/domain/entities/user_data.dart';
// ignore: unused_import
import 'package:application1/features/authentication/domain/repositories/auth_repo.dart';
import 'package:application1/features/authentication/domain/usecases/get_current_user_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late GetCurrentUserUsecase getCurrentUserUsecase;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    getCurrentUserUsecase =
        GetCurrentUserUsecase(authRepository: mockAuthRepository);
  });
  const tUserEntity =  UserEntity(
    email: 'ley@gmail.com',
    name: 'ley',
  );

  test('should return a userEntity when the user data retrieval is successful',
      () async {
    when(mockAuthRepository.getCurrentUser())
        .thenAnswer((_) async => const Right(tUserEntity));

    final result = await getCurrentUserUsecase(NoParams());

    expect(result, const Right(tUserEntity));
  });
    test('should return a failure when the user data retrieval is unsuccessful',
      () async {
    when(mockAuthRepository.getCurrentUser())
        .thenAnswer((_) async => const Left(ServerFailure('user not found')));

    final result = await getCurrentUserUsecase(NoParams());

    expect(result, const Left(ServerFailure('user not found')));
  });
}
