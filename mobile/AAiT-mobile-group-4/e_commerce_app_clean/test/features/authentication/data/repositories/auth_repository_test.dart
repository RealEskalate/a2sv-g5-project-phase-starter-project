import 'package:application1/core/error/exception.dart';
import 'package:application1/core/error/failure.dart';
import 'package:application1/features/authentication/data/model/log_in_model.dart';
import 'package:application1/features/authentication/data/model/sign_up_model.dart';
import 'package:application1/features/authentication/data/model/user_model.dart';
import 'package:application1/features/authentication/data/repositories/auth_repo_impl.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockAuthRemoteDataSource mockAuthRemoteDataSource;
  late AuthRepositoryImpl authRepositoryImpl;

  setUp(() {
    mockAuthRemoteDataSource = MockAuthRemoteDataSource();
    authRepositoryImpl = AuthRepositoryImpl(
      authRemoteDataSource: mockAuthRemoteDataSource,
    );
  });
  const tUserModel = UserModel(
    name: 'ley',
    email: 'ley@gmail.com',
  );
  group('get current user', () {
    test(
        'should return a user entity when the call to remote data source is successful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.getCurrentUser())
          .thenAnswer((_) async => tUserModel);
      //act
      final result = await authRepositoryImpl.getCurrentUser();
      //assert
      expect(result, Right(tUserModel.toUserEntity()));
    });

    test(
      'should return a server failure when a call to remote data source is unsuccessful',
      () async {
        //arrange
        when(mockAuthRemoteDataSource.getCurrentUser())
            .thenThrow(ServerException());
        //act
        final result = await authRepositoryImpl.getCurrentUser();
        //assert
        expect(result, const Left(ServerFailure('An error has occurred')));
      },
    );
  });

  const tLogInModel = LogInModel(email: 'ley@gmail.com', password: '1234');
  group('log in', () {
    test(
        'should return a product when a call to remote data source is successful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.logIn(tLogInModel))
          .thenAnswer((_) async => unit);
      //act
      final result = await authRepositoryImpl.logIn(tLogInModel);
      //assert
      expect(result, const Right(unit));
    });

    test(
        'should return a server failure when a call to remote data source is unsuccessful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.logIn(tLogInModel))
          .thenThrow(ServerException());
      //act
      final result = await authRepositoryImpl.logIn(tLogInModel);
      //assert
      expect(result, const Left(ServerFailure('cannot login')));
    });
  });
  const tSignUpModel =
      SignUpModel(email: 'ley@gmail.com', password: '1234', username: 'ley');
  group('sign up', () {
    test(
        'should return a unit  when a call to remote data source is successful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.signUp(tSignUpModel))
          .thenAnswer((_) async => unit);
      //act
      final result = await authRepositoryImpl.signUp(tSignUpModel);
      //assert
      expect(result, const Right(unit));
    });

    test(
        'should return a server failure when a call to remote data source is unsuccessful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.signUp(tSignUpModel))
          .thenThrow(ServerException());
      //act
      final result = await authRepositoryImpl.signUp(tSignUpModel);
      //assert
      expect(result, const Left(ServerFailure('Unknown error occurred')));
    });
  });

 group('log out', () {
    test(
        'should return a unit when a call to remote data source is successful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.logOut())
          .thenAnswer((_) async => unit);
      //act
      final result = await authRepositoryImpl.logOut();
      //assert
      expect(result, const Right(unit));
    });

    test(
        'should return a server failure when a call to remote data source is unsuccessful',
        () async {
      //arrange
      when(mockAuthRemoteDataSource.logOut())
          .thenThrow(ServerException());
      //act
      final result = await authRepositoryImpl.logOut();
      //assert
      expect(result, const Left(ServerFailure('can not logout')));
    });
  });
}
