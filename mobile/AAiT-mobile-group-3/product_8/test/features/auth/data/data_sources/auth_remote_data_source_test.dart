// import 'dart:convert';

// import 'package:http/http.dart' as http;
// import 'package:mockito/annotations.dart';
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/constants/constants.dart';
// import 'package:product_8/core/exception/exception.dart';
// import 'package:product_8/features/auth/data/data_source/auth_remote_data_source.dart';
// import 'package:product_8/features/auth/data/models/sign_in_user_model.dart';
// import 'package:product_8/features/auth/data/models/sign_up_user_model.dart';
// import 'package:product_8/features/auth/data/models/user_data_model.dart';
// import 'package:test/test.dart';

// import '../../../../helpers/test_helper.mocks.dart';

// @GenerateMocks([http.Client])
// void main() {
//   late MockHttpClient mockHttpClient;
//   late AuthRemoteDataSourceImpl remoteDataSource;

//   setUp(() {
//     mockHttpClient = MockHttpClient();
//     remoteDataSource = AuthRemoteDataSourceImpl(client: mockHttpClient);
//   });

//   group('signIn', () {
//     const signInUserModel =
//         SignInUserModel(email: 'test@example.com', password: 'password');
//     const userDataModel = UserDataModel(
//       data: DataModel(name: 'John Doe', email: 'test@example.com'),
//       token: 'someToken123',
//     );

//     test('should return UserDataModel when response code is 200', () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('${Urls.autUrl}/auth/login'),
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer((_) async => http.Response(
//             json.encode({
//               'data': {
//                 'name': 'John Doe',
//                 'email': 'test@example.com',
//               },
//               'token': 'someToken123',
//             }),
//             200,
//           ));

//       // act
//       final result = await remoteDataSource.signIn(signInUserModel);

//       // assert
//       expect(result, equals(userDataModel));
//     });

//     test('should throw ServerException when response code is not 200',
//         () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('${Urls.autUrl}/auth/login'),
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer((_) async => http.Response('Something went wrong', 400));

//       // act
//       final call = remoteDataSource.signIn;

//       // assert
//       expect(() => call(signInUserModel), throwsA(isA<ServerException>()));
//     });
//   });

//   group('signUp', () {
//     const signUpUserModel = SignUpUserModel(
//         email: 'test@example.com', password: 'password', name: 'John Doe');
//     const signUpResponseModel = SignUpUserModel(
//         email: 'test@example.com', password: 'password', name: 'John Doe');

//     test('should return SignUpUserModel when response code is 200', () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('${Urls.autUrl}/users/me'),
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer((_) async => http.Response(
//             json.encode({
//               'data': {
//                 'email': 'test@example.com',
//                 'password': 'password',
//                 'name': 'John Doe',
//               },
//             }),
//             200,
//           ));

//       // act
//       final result = await remoteDataSource.signUp(signUpUserModel);

//       // assert
//       expect(result, equals(signUpResponseModel));
//     });

//     test('should throw ServerException when response code is not 200',
//         () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('${Urls.autUrl}/users/me'),
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer((_) async => http.Response('Something went wrong', 400));

//       // act
//       final call = remoteDataSource.signUp;

//       // assert
//       expect(() => call(signUpUserModel), throwsA(isA<ServerException>()));
//     });
//   });
// }
