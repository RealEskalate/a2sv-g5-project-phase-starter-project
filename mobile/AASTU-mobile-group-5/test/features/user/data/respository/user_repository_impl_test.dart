import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/error/exceptions.dart';
import 'package:task_9/core/failure/failure.dart';
import 'package:task_9/features/user/data/data_sources/user_local_data_source.dart';
import 'package:task_9/features/user/data/data_sources/user_remote_data_source.dart';
import 'package:task_9/features/user/data/models/user_model.dart';
import 'package:task_9/features/user/data/repositories/user_respository_impl.dart';

import 'user_repository_impl_test.mocks.dart';

@GenerateMocks([UserLocalDataSource, UserRemoteDataSource])
void main() {
  late UserRepositoryImpl repository;
  late MockUserRemoteDataSource mockRemoteDataSource;
  late MockUserLocalDataSource mockLocalDataSource;

  setUp(() {
    mockRemoteDataSource = MockUserRemoteDataSource();
    mockLocalDataSource = MockUserLocalDataSource();
    repository = UserRepositoryImpl(
      remoteDataSource: mockRemoteDataSource,
      localDataSource: mockLocalDataSource,
    );
  });

  group('loginUser', () {
    const tEmail = 'test@example.com';
    const tPassword = 'password123';
    const tAccessToken = 'access_token';

    test('should return access token when the call to remote data source is successful', () async {
      when(mockRemoteDataSource.loginUser(tEmail, tPassword))
          .thenAnswer((_) async => tAccessToken);

      final result = await repository.loginUser(tEmail, tPassword);

      verify(mockRemoteDataSource.loginUser(tEmail, tPassword));
      verify(mockLocalDataSource.saveAccessToken(tAccessToken));
      expect(result, const Right(tAccessToken));
    });

    test('should return ServerFailure when the call to remote data source is unsuccessful', () async {
      when(mockRemoteDataSource.loginUser(tEmail, tPassword))
          .thenThrow(ServerException());

      final result = await repository.loginUser(tEmail, tPassword);

      verify(mockRemoteDataSource.loginUser(tEmail, tPassword));
      verifyZeroInteractions(mockLocalDataSource);
      expect(result, Left(ServerFailure(message: 'Failed to login')));
    });
  });

  group('registerUser', () {
    const tEmail = 'test@example.com';
    const tPassword = 'password123';
    const tName = 'Test User';
    const tUserModel = UserModel(id: '1', name: tName, email: tEmail);

    test('should return User when the call to remote data source is successful', () async {
      when(mockRemoteDataSource.registerUser(tEmail, tPassword, tName))
          .thenAnswer((_) async => tUserModel);

      final result = await repository.registerUser(tEmail, tPassword, tName);

      verify(mockRemoteDataSource.registerUser(tEmail, tPassword, tName));
      expect(result, const Right(tUserModel));
    });

    test('should return ServerFailure when the call to remote data source is unsuccessful', () async {
      when(mockRemoteDataSource.registerUser(tEmail, tPassword, tName))
          .thenThrow(ServerException());

      final result = await repository.registerUser(tEmail, tPassword, tName);

      verify(mockRemoteDataSource.registerUser(tEmail, tPassword, tName));
      expect(result, Left(ServerFailure(message: 'Failed to register')));
    });
  });
}