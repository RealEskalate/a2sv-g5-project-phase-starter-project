import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/failure/failure.dart';
import 'package:task_9/features/user/domain/entities/user.dart';
import 'package:task_9/features/user/domain/repositories/user_repository.dart';
import 'package:task_9/features/user/domain/use_case/register_user.dart';

import 'register_user_test.mocks.dart';


@GenerateMocks([UserRepository])
void main() {
  late RegisterUser registerUser;
  late MockUserRepository mockUserRepository;

  setUp(() {
    mockUserRepository = MockUserRepository();
    registerUser = RegisterUser(mockUserRepository);
  });

  const tName = 'Test User';
  const tEmail = 'test@example.com';
  const tPassword = 'password';
  final tRegisterParams = RegisterParams(name: tName, email: tEmail, password: tPassword);
  const tUser = User(id: '1', name: tName, email: tEmail);

  test('should return User when the call to repository is successful', () async {
    // Arrange
    when(mockUserRepository.registerUser(any, any, any))
        .thenAnswer((_) async => const Right(tUser));

    // Act
    final result = await registerUser(tRegisterParams);

    // Assert
    expect(result, const Right(tUser));
    verify(mockUserRepository.registerUser(tEmail, tPassword, tName));
    verifyNoMoreInteractions(mockUserRepository);
  });

  test('should return failure when the call to repository is unsuccessful', () async {
    // Arrange
    final tFailure = ServerFailure(message: 'Server Failed');
    when(mockUserRepository.registerUser(any, any, any))
        .thenAnswer((_) async => Left(tFailure));

    // Act
    final result = await registerUser(tRegisterParams);

    // Assert
    expect(result, Left(tFailure));
    verify(mockUserRepository.registerUser(tEmail, tPassword, tName));
    verifyNoMoreInteractions(mockUserRepository);
  });
}
