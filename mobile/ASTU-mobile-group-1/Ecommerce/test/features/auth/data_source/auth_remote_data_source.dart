// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/http.dart' as http;
// import 'package:mockito/mockito.dart';
// import 'package:product_6/core/constants/constants.dart';
// import 'package:product_6/core/error/exception.dart';
// import 'package:product_6/features/auth/data/data_sources/auth_remote_data_source.dart';
// import 'package:product_6/features/auth/data/models/sign_in_model.dart';
// import 'package:product_6/features/auth/data/models/sign_up_model.dart';
// import 'package:product_6/features/auth/data/models/signed_in_model.dart';
// import 'package:product_6/features/auth/data/models/user_model.dart';

// import '../../../helpers/test_helper.mocks.dart';

// void main() {
//   late MockHttpClient mockHttpClient;
//   late MockAuthLocalDataSource mockAuthLocalDataSource;
//   late AuthRemoteDataSourceImpl authRemoteDataSourceImpl;

//   setUp(() {
//     mockHttpClient = MockHttpClient();
//     mockAuthLocalDataSource = MockAuthLocalDataSource();
//     authRemoteDataSourceImpl = AuthRemoteDataSourceImpl(
//       client: mockHttpClient,
//       authLocalDataSource: mockAuthLocalDataSource,
//     );
//   });

//   const testEmail = 'test@example.com';
//   const testPassword = 'password123';
//   const testToken = 'test_token';

//   final signInModel = SignInModel(email: testEmail, password: testPassword);
//   final signUpModel = SignUpModel(
//     name: 'Test User',
//     email: testEmail,
//     password: testPassword,
//   );

//   group('signIn', () {
//     test('should return SignedInModel when the response code is 201', () async {
//       // arrange
//       when(mockHttpClient.post(
//         any,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer(
//         (_) async => http.Response(
//           readJson('helpers/dummy_data/dummy_sign_in_response.json'),
//           201,
//         ),
//       );
//       when(mockAuthLocalDataSource.cacheToken(any)).thenAnswer((_) async {});

//       // act
//       final result = await authRemoteDataSourceImpl.signIn(signInModel);

//       // assert
//       expect(result, isA<SignedInModel>());
//       verify(mockAuthLocalDataSource.cacheToken(any)).called(1);
//     });

//     test(
//         'should throw ServerException when the response code is not 201 or other error occurs',
//         () async {
//       // arrange
//       when(mockHttpClient.post(
//         any,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer(
//         (_) async => http.Response('Unauthorized', 401),
//       );

//       // act
//       final call = authRemoteDataSourceImpl.signIn;

//       // assert
//       expect(() => call(signInModel), throwsA(isA<ServerException>()));
//     });
//   });

//   group('signUp', () {
//     test('should return true when the response code is 201', () async {
//       // arrange
//       when(mockHttpClient.post(
//         any,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer(
//         (_) async => http.Response(
//           '{"message": "Account created successfully."}',
//           201,
//         ),
//       );

//       // act
//       final result = await authRemoteDataSourceImpl.signUp(signUpModel);

//       // assert
//       expect(result, true);
//     });

//     test(
//         'should throw ServerException when the response code is not 201 or error occurs',
//         () async {
//       // arrange
//       when(mockHttpClient.post(
//         any,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer(
//         (_) async => http.Response(
//           '{"message": "Email already exists."}',
//           400,
//         ),
//       );

//       // act
//       final call = authRemoteDataSourceImpl.signUp;

//       // assert
//       expect(() => call(signUpModel), throwsA(isA<ServerException>()));
//     });
//   });

//   group('getUser', () {
//     test('should return UserModel when the response code is 200', () async {
//       // arrange
//       when(mockAuthLocalDataSource.getToken())
//           .thenAnswer((_) async => testToken);
//       when(mockHttpClient.get(
//         any,
//         headers: anyNamed('headers'),
//       )).thenAnswer(
//         (_) async => http.Response(
//           readJson('helpers/dummy_data/dummy_user_response.json'),
//           200,
//         ),
//       );

//       // act
//       final result = await authRemoteDataSourceImpl.getUser();

//       // assert
//       expect(result, isA<UserModel>());
//     });

//     test('should throw UnknownException when the response code is not 200',
//         () async {
//       // arrange
//       when(mockAuthLocalDataSource.getToken())
//           .thenAnswer((_) async => testToken);
//       when(mockHttpClient.get(
//         any,
//         headers: anyNamed('headers'),
//       )).thenAnswer(
//         (_) async => http.Response(
//           'Not found',
//           404,
//         ),
//       );

//       // act
//       final call = authRemoteDataSourceImpl.getUser;

//       // assert
//       expect(() => call(), throwsA(isA<UnknownException>()));
//     });
//   });
// }
