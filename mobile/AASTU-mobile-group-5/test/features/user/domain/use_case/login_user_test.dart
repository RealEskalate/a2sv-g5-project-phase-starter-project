import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/failure/failure.dart';
import 'package:task_9/features/user/domain/repositories/user_repository.dart';
import 'package:task_9/features/user/domain/use_case/login_user.dart';

import 'login_user_test.mocks.dart';

@GenerateMocks([UserRepository])
void main() {
  late LoginUser loginUser;
  late MockUserRepository mockUserRepository;

  setUp(() {
    mockUserRepository = MockUserRepository();
    loginUser = LoginUser(mockUserRepository);
  });

  const tEmail = 'test@example.com';
  const tPassword = 'password';
  final tLoginParams = LoginParams(email: tEmail, password: tPassword);
  const tToken = 'test_token';

  test('should return token when the call to repository is successful', () async {
    // Arrange
    when(mockUserRepository.loginUser(any, any))
        .thenAnswer((_) async => const Right(tToken));

    // Act
    final result = await loginUser(tLoginParams);

    // Assert
    expect(result, const Right(tToken));
    verify(mockUserRepository.loginUser(tEmail, tPassword));
    verifyNoMoreInteractions(mockUserRepository);
  });

  test('should return failure when the call to repository is unsuccessful', () async {
    // Arrange
    final tFailure = ServerFailure(message: 'Server Failure');
    when(mockUserRepository.loginUser(any, any))
        .thenAnswer((_) async => Left(tFailure));

    // Act
    final result = await loginUser(tLoginParams);

    // Assert
    expect(result, Left(tFailure));
    verify(mockUserRepository.loginUser(tEmail, tPassword));
    verifyNoMoreInteractions(mockUserRepository);
  });
}
