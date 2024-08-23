import 'package:application1/core/error/failure.dart';
import 'package:application1/core/usecase/usecase.dart';
import 'package:application1/features/authentication/domain/usecases/log_out_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late LogOutUsecase
   logOutUsecase;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    logOutUsecase =
        LogOutUsecase(authRepository: mockAuthRepository);
  });

  test('should return void when user logs out',
      () async {
    when(mockAuthRepository.logOut())
        // ignore: void_checks
        .thenAnswer((_) async => const Right(unit));

    final result = await logOutUsecase(NoParams());

    expect(result, const Right(unit));
  });
    test('should return a failure when we cannot logout',
      () async {
    when(mockAuthRepository.logOut())
        .thenAnswer((_) async => const Left(ServerFailure('Log out failure')));

    final result = await logOutUsecase(NoParams());

    expect(result, const Left(ServerFailure('Log out failure')));
  });
}