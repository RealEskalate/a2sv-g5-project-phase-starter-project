import 'package:application1/core/error/failure.dart';
import 'package:application1/features/authentication/domain/entities/log_in.dart';
import 'package:application1/features/authentication/domain/usecases/log_in_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late LogInUsecase
   logInUsecase;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    logInUsecase =
        LogInUsecase(authRepository: mockAuthRepository);
  });
  const tLogInEntity =  LogInEntity(
    email: 'ley@gmail.com', password: '1234',
    
  );

  test('should return void when log in is successful',
      () async {
    when(mockAuthRepository.logIn(tLogInEntity))
        // ignore: void_checks
        .thenAnswer((_) async => const Right(unit));

    final result = await logInUsecase(const LogInParams(logInEntity: tLogInEntity));

    expect(result, const Right(unit));
  });
    test('should return a failure when logout  is unsuccessful',
      () async {
    when(mockAuthRepository.logIn(tLogInEntity))
        .thenAnswer((_) async => const Left(ServerFailure('user not found')));

    final result = await logInUsecase(const LogInParams(logInEntity: tLogInEntity));

    expect(result, const Left(ServerFailure('user not found')));
  });
}