// import 'package:dartz/dartz.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/exception/exception.dart';
// import 'package:product_8/core/failure/failure.dart';
// import 'package:product_8/features/auth/data/models/sign_up_user_model.dart';
// import 'package:product_8/features/auth/data/models/user_data_model.dart';
// import 'package:product_8/features/auth/data/repository/auth_repository_impl.dart';
// import 'package:product_8/features/auth/domain/entities/sign_in_user_entitiy.dart';
// import 'package:product_8/features/auth/domain/entities/sign_up_user_entitiy.dart';

// import '../../../../helpers/test_helper.mocks.dart';




// void main() {
//   late AuthRepositoryImpl repository;
//   late MockAuthRemoteDataSource mockRemoteDataSource;
//   late MockAuthLocalDataSource mockLocalDataSource;
//   late MockNetworkInfo mockNetworkInfo;

//   setUp(() {
//     mockRemoteDataSource = MockAuthRemoteDataSource();
//     mockLocalDataSource = MockAuthLocalDataSource();
//     mockNetworkInfo = MockNetworkInfo();
//     repository = AuthRepositoryImpl(
//       authRemoteDataSource: mockRemoteDataSource,
//       authLocalDataSource: mockLocalDataSource,
//       networkInfo: mockNetworkInfo,
//     );
//   });

//   group('signUp', () {
//     const tSignUpUserEntity = SignUpUserEntitiy(
//       email: 'test@example.com',
//       password: 'password123',
//       name: 'Test User',
//     );
//     const tSignUpUserModel = SignUpUserModel(
//       email: 'test@example.com',
//       password: 'password123',
//       name: 'Test User',
//     );

//     test('should check if the device is online', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signUp(any))
//           .thenAnswer((_) async => tSignUpUserModel);
//       // act
//       repository.signUp(tSignUpUserEntity);
//       // assert
//       verify(mockNetworkInfo.isConnected);
//     });

//     test('should return SignUpUserModel when call to remote data source is successful', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signUp(any))
//           .thenAnswer((_) async => tSignUpUserModel);
//       // act
//       final result = await repository.signUp(tSignUpUserEntity);
//       // assert
//       verify(mockRemoteDataSource.signUp(any));
//       expect(result, equals(const Right(tSignUpUserModel)));
//     });

//     test('should return ServerFailure when call to remote data source is unsuccessful', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signUp(any)).thenThrow(ServerException());
//       // act
//       final result = await repository.signUp(tSignUpUserEntity);
//       // assert
//       verify(mockRemoteDataSource.signUp(any));
//       expect(result, equals(const Left(ServerFailure(message: 'Server Error'))));
//     });

//     test('should return ConnectionFailure when the device is offline', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
//       // act
//       final result = await repository.signUp(tSignUpUserEntity);
//       // assert
//       expect(result,
//           equals(const Left(ConnectionFailure(message: 'No Internet Connection'))));
//     });
//   });

//   group('signIn', () {
//     const tSignInUserEntity = SignInUserEntitiy(
//       email: 'test@example.com',
//       password: 'password123',
//     );
//     const tUserDataModel = UserDataModel(
//       data: DataModel(name: 'Test User', email: 'test@example.com'),
//       token: '123token',
//     );

//     test('should check if the device is online', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signIn(any))
//           .thenAnswer((_) async => tUserDataModel);
//       // act
//       repository.signIn(tSignInUserEntity);
//       // assert
//       verify(mockNetworkInfo.isConnected);
//     });

//     test('should return UserDataModel when call to remote data source is successful', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signIn(any))
//           .thenAnswer((_) async => tUserDataModel);
//       // act
//       final result = await repository.signIn(tSignInUserEntity);
//       // assert
//       verify(mockRemoteDataSource.signIn(any));
//       verify(mockLocalDataSource.cacheToken(tUserDataModel.token));
//       expect(result, equals(const Right(tUserDataModel)));
//     });

//     test('should return ServerFailure when call to remote data source is unsuccessful', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => true);
//       when(mockRemoteDataSource.signIn(any)).thenThrow(ServerException());
//       // act
//       final result = await repository.signIn(tSignInUserEntity);
//       // assert
//       verify(mockRemoteDataSource.signIn(any));
//       expect(result, equals(const Left(ServerFailure(message: 'Server Error'))));
//     });

//     test('should return ConnectionFailure when the device is offline', () async {
//       // arrange
//       when(mockNetworkInfo.isConnected).thenAnswer((_) async => false);
//       // act
//       final result = await repository.signIn(tSignInUserEntity);
//       // assert
//       expect(result,
//           equals(const Left(ConnectionFailure(message: 'No Internet Connection'))));
//     });
//   });
// }
