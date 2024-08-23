import 'package:application1/core/error/failure.dart';
import 'package:application1/features/authentication/domain/entities/sign_up.dart';
import 'package:application1/features/authentication/domain/usecases/sign_up_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late SignUpUsecase signUpUsecase;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    signUpUsecase = SignUpUsecase(authRepository: mockAuthRepository);
  });
  const signUpEntity =
      SignUpEntity(email: 'ley@gmail.com', password: '1234', username: 'ley');
  test('should return void when user signs up successfully', () async {
    when(mockAuthRepository.signUp(signUpEntity))
        // ignore: void_checks
        .thenAnswer((_) async => const Right(unit));

    final result = await signUpUsecase(const GetParams(signUpEntity: signUpEntity));

    expect(result, const Right(unit));
  });
  test('should return a failure when user signing up is unsuccessful', () async {
    when(mockAuthRepository.signUp(signUpEntity))
        .thenAnswer((_) async => const Left(ServerFailure('Log out failure')));

    final result = await signUpUsecase(const GetParams(signUpEntity: signUpEntity));

    expect(result, const Left(ServerFailure('Log out failure')));
  });
}
