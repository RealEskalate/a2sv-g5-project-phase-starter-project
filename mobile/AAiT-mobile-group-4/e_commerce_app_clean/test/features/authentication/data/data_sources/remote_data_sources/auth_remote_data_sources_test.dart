import 'dart:convert';

import 'package:application1/core/constants/constants.dart';
import 'package:application1/core/error/exception.dart';
import 'package:application1/features/authentication/data/data_sources/remote/auth_remote_datasource_impl.dart';
import 'package:application1/features/authentication/data/model/log_in_model.dart';
import 'package:application1/features/authentication/data/model/sign_up_model.dart';
import 'package:application1/features/authentication/data/model/user_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../../helper/dummy_data/json_reader.dart';
import '../../../../../helper/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late AuthRemoteDatasourceImpl authRemoteDataSourceImpl;
  late MockAuthLocalDataSource mockAuthLocalDataSource;
  setUp(() {
    mockHttpClient = MockHttpClient();
    mockAuthLocalDataSource = MockAuthLocalDataSource();
    authRemoteDataSourceImpl = AuthRemoteDatasourceImpl(
        client: mockHttpClient, authLocalDataSource: mockAuthLocalDataSource);
  });
  const String logInResponsePath =
      '/helper/dummy_data/auth_response/dummy_log_in_response.json';
  const String signUpResponsePath =
      '/helper/dummy_data/auth_response/dummy_sign_up_response.json';
  const String getCurrentUserResponsePath =
      '/helper/dummy_data/auth_response/dummy_get_current_user_response.json';
  const tSignUpModel =
      SignUpModel(email: 'ley@gmail.com', password: '1234', username: 'ley');
  const tLogInModel = LogInModel(email: 'ley@gmail.com', password: '1234');
  const tUserModel = UserModel(email: 'ley@gmail.com', name: 'ley');
  group('sign up', () {
    test('should return a void if the response is 201', () async {
      //arrange
      when(mockHttpClient.post(Uri.parse(Urls2.signUp()),
              body: jsonEncode(tSignUpModel),
              headers: {'Content-Type': 'application/json'}))
          .thenAnswer(
              (_) async => http.Response(readJson(signUpResponsePath), 201));
      //act
      await authRemoteDataSourceImpl.signUp(tSignUpModel);
      //assert
      verify(mockHttpClient.post(Uri.parse(Urls2.signUp()),
          body: jsonEncode(tSignUpModel),
          headers: {'Content-Type': 'application/json'}));
    });

  });
  group('Log in', () {
    test('should return a void if the response is 201', () async {
      //arrange
      when(mockAuthLocalDataSource.cacheToken('mytoken'))
          .thenAnswer((_) async => true);
      when(mockHttpClient.post(Uri.parse(Urls2.login()),
              body: jsonEncode(tLogInModel),
              headers: {'Content-Type': 'application/json'}))
          .thenAnswer(
              (_) async => http.Response(readJson(logInResponsePath), 201));
      //act
      await authRemoteDataSourceImpl.logIn(tLogInModel);
      //assert
      verify(mockHttpClient.post(Uri.parse(Urls2.login()),
          body: jsonEncode(tLogInModel),
          headers: {'Content-Type': 'application/json'}));
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
      when(mockAuthLocalDataSource.cacheToken('mytoken'))
          .thenAnswer((_) => Future.value(true));
      when(mockHttpClient.post(Uri.parse(Urls2.login()),
              body: jsonEncode(tLogInModel),
              headers: {'Content-Type': 'application/json'}))
          .thenAnswer((_) async => http.Response('Not Found', 404));
      //act
      expect(() async {
        await authRemoteDataSourceImpl.logIn(tLogInModel);
      }, throwsA(isA<ServerException>()));
    });
  });

  group('get current user', () {
    test('should return a user model if the response is 200', () async {
      //arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls2.getCurrentUser()), headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer mytoken'
      })).thenAnswer((_) async =>
          http.Response(readJson(getCurrentUserResponsePath), 200));
      //act
      final result = await authRemoteDataSourceImpl.getCurrentUser();
      //assert
      expect(result, tUserModel);
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls2.getCurrentUser()), headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer mytoken'
      })).thenAnswer((_) async => http.Response('Not Found', 404));
      //act
      expect(() async {
        await authRemoteDataSourceImpl.getCurrentUser();
      }, throwsA(isA<ServerException>()));
    });
  });
}
