import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/core/errors/failures/failure.dart';
import 'package:ecommerce_app/features/auth/domain/usecases/get_me_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late MockAuthRepository mockAuthRepository;
  late GetMeUsecase getMeUsecase;

  setUp(() {
    mockAuthRepository = MockAuthRepository();
    getMeUsecase = GetMeUsecase(authRepository: mockAuthRepository);
  });
  test('get me usecase ...', () async {
    when(mockAuthRepository.getMe())
        .thenAnswer((_) async => const Right(AuthData.userEntity));

    final result = await getMeUsecase.execute();

    expect(result, const Right(AuthData.userEntity));
  });

  test('get me usecase ...', () async {
    when(mockAuthRepository.getMe())
        .thenAnswer((_) async => const Left(CacheFailure('fail')));

    final result = await getMeUsecase.execute();

    expect(result, const Left(CacheFailure('fail')));
  });
}
