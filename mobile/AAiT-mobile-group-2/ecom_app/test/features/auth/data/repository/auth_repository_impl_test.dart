import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/constants/constants.dart';
import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/auth/data/models/authenticated_model.dart';
import 'package:ecom_app/features/auth/data/models/login_model.dart';
import 'package:ecom_app/features/auth/data/models/register_model.dart';
import 'package:ecom_app/features/auth/data/models/user_data_model.dart';
import 'package:ecom_app/features/auth/data/repositories/auth_repository_impl.dart';
import 'package:ecom_app/features/auth/domain/entities/login_entity.dart';
import 'package:ecom_app/features/auth/domain/entities/register_entity.dart';
import 'package:ecom_app/features/auth/domain/entities/user_data_entity.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockAuthRemoteDataSource mockAuthRemoteDataSource;
  late MockAuthLocalDataSource mockAuthLocalDataSource;
  late MockNetworkInfo mockNetworkInfo;
  late AuthRepositoryImpl authRepositoryImpl;

  setUp(() {
    mockAuthRemoteDataSource = MockAuthRemoteDataSource();
    mockAuthLocalDataSource = MockAuthLocalDataSource();
    mockNetworkInfo = MockNetworkInfo();
    authRepositoryImpl = AuthRepositoryImpl(
        remoteDataSource: mockAuthRemoteDataSource,
        localDataSource: mockAuthLocalDataSource,
        networkInfo: mockNetworkInfo);
  });

  final tLoginModel = LoginModel(email: 'email', password: 'password');
  final tLoginEntity = LoginEntity(email: 'email', password: 'password');
  final tRegisterModel =
      RegisterModel(email: 'email', password: 'password', name: 'name');
  final tRegisterEntity =
      RegistrationEntity(email: 'email', password: 'password', name: 'name');
  final tUserDataModel = UserDataModel(name: 'name', email: 'email');
  final tUserDataEntity = UserDataEntity(name: 'name', email: 'email');


  group('LoginImpl', () {
    test('should return unit after successful login', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.login(tLoginModel))
          .thenAnswer((_) async => AuthenticatedModel(token: 'token'));
      when(mockAuthLocalDataSource.cacheToken('token'))
          .thenAnswer((_) async => unit);

      //act
      final result = await authRepositoryImpl.login(tLoginEntity);

      //assert
      expect(result, const Right(unit));
    });
    test('should return an unauthorized failure after unsuccessful login',
        () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.login(tLoginModel))
          .thenThrow(UnauthorizedException());

      //act
      final result = await authRepositoryImpl.login(tLoginEntity);

      //assert
      expect(result, Left(UnauthorizedFailure(ErrorMessages.forbiddenError)));
    });
    test('should return an server failure after unsuccessful login', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.login(tLoginModel))
          .thenThrow(ServerException());

      //act
      final result = await authRepositoryImpl.login(tLoginEntity);

      //assert
      expect(result, const Left(ServerFailure(ErrorMessages.serverError)));
    });
    test('should return an connection failure when no internet', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      //act
      final result = await authRepositoryImpl.login(tLoginEntity);

      //assert
      expect(result, const Left(ConnectionFailure(ErrorMessages.noInternet)));
    });
  });
  group('RegisterImpl', () {
    test('should return unit after successful register', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.register(tRegisterModel))
          .thenAnswer((_) async => unit);

      //act
      final result = await authRepositoryImpl.register(tRegisterEntity);

      //assert
      expect(result, const Right(unit));
    });
    test(
        'should return an user already exists failure after unsuccessful login',
        () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.register(tRegisterModel))
          .thenThrow(UserAlreadyExistsException());

      //act
      final result = await authRepositoryImpl.register(tRegisterEntity);

      //assert
      expect(result,
          Left(UserAlreadyExistsFailure(ErrorMessages.userAlreadyExists)));
    });
    test('should return an server failure after unsuccessful login', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthRemoteDataSource.register(tRegisterModel))
          .thenThrow(ServerException());

      //act
      final result = await authRepositoryImpl.register(tRegisterEntity);

      //assert
      expect(result, const Left(ServerFailure(ErrorMessages.serverError)));
    });
    test('should return an connection failure when no internet', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      //act
      final result = await authRepositoryImpl.register(tRegisterEntity);

      //assert
      expect(result, const Left(ConnectionFailure(ErrorMessages.noInternet)));
    });
  });

  group('GetUserImpl', () {
    test('should return userdataentity after successful get request', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthLocalDataSource.getToken()).thenAnswer((_) async => 'token');
      when(mockAuthRemoteDataSource.getUser('token'))
          .thenAnswer((_) async => tUserDataModel);

      //act
      final result = await authRepositoryImpl.getUser();

      //assert
      expect(result, Right(tUserDataEntity));
    });
    
    test('should return an server failure after unsuccessful login', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
      when(mockAuthLocalDataSource.getToken()).thenAnswer((_) async => 'token');
      when(mockAuthRemoteDataSource.getUser('token'))
          .thenThrow(ServerException());

      //act
      final result = await authRepositoryImpl.getUser();

      //assert
      expect(result, const Left(ServerFailure(ErrorMessages.serverError)));
    });
    test('should return an connection failure when no internet', () async {
      //arrange
      when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);

      //act
      final result = await authRepositoryImpl.getUser();

      //assert
      expect(result, const Left(ConnectionFailure(ErrorMessages.noInternet)));
    });
  });
  group('LogoutImpl', () {
    test('should return unit after successful logout', () async {
      //arrange
      when(mockAuthLocalDataSource.logout()).thenAnswer((_) async => unit);
      

      //act
      final result = await authRepositoryImpl.logout();

      //assert
      expect(result, const Right(unit));
    });
    
    test('should return an cache failure after unsuccessful login', () async {
      //arrange
     when(mockAuthLocalDataSource.logout()).thenThrow(CacheException());

      //act
      final result = await authRepositoryImpl.logout();

      //assert
      expect(result, const Left(CacheFailure(ErrorMessages.cacheError)));
    });
  });


}
