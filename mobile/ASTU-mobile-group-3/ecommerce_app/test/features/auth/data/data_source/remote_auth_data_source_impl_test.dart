import 'dart:convert';

import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/exceptions/product_exceptions.dart';
import 'package:ecommerce_app/features/auth/data/data_source/remote_auth_data_source.dart';
import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late RemoteAuthDataSourceImpl remoteAuthDataSourceImpl;
  setUp(() {
    mockHttpClient = MockHttpClient();
    remoteAuthDataSourceImpl = RemoteAuthDataSourceImpl(client: mockHttpClient);
  });

  group('Log in test', () {
    test('Should return true when data is valid', () async {
      /// arrange
      when(mockHttpClient.post(any, body: {
        'email': AuthData.email,
        'password': AuthData.password,
      })).thenAnswer((_) async => http.Response(AuthData.readJson(), 201));

      final result = await remoteAuthDataSourceImpl.logIn(AuthData.userEntity);

      expect(result, AuthData.tokenModel);
    });

    test('Should return throw server exception when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.logInUser),
          body: {
            'email': AuthData.email,
            'password': AuthData.password,
          },
        ),
      ).thenAnswer((_) async => http.Response('Done', 401));

      /// action
      final result = remoteAuthDataSourceImpl.logIn(AuthData.userEntity);

      /// assert
      expect(result, throwsA(isA<LoginException>()));
    });

    test('Should throw server exception when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.logInUser),
          body: {
            'email': AuthData.email,
            'password': AuthData.password,
          },
        ),
      ).thenThrow(ServerException());

      /// action
      final result = remoteAuthDataSourceImpl.logIn;

      /// assert
      expect(() async => result(AuthData.userEntity),
          throwsA(isA<ServerException>()));
    });
  });

  group('Sign up test', () {
    test('Should return true when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.registerUser),
          body: {
            'name': AuthData.name,
            'email': AuthData.email,
            'password': AuthData.password
          },
        ),
      ).thenAnswer((_) async => http.Response(AuthData.readJson(), 201));

      /// action
      final result = await remoteAuthDataSourceImpl.signUp(AuthData.userEntity);

      /// assert
      expect(
          result, UserModel.fromJson(json.decode(AuthData.readJson())['data']));
    });

    test('Should return throw server exception when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.registerUser),
          body: {
            'name': AuthData.name,
            'email': AuthData.email,
            'password': AuthData.password
          },
        ),
      ).thenAnswer((_) async => http.Response('Done', 404));

      /// action
      final result = remoteAuthDataSourceImpl.signUp;

      /// assert
      expect(() async => result(AuthData.userEntity),
          throwsA(isA<ServerException>()));
    });

    test('Should return throw server exception when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.registerUser),
          body: {
            'name': AuthData.name,
            'email': AuthData.email,
            'password': AuthData.password
          },
        ),
      ).thenAnswer((_) async => http.Response('Done', 409));

      /// action
      final result = remoteAuthDataSourceImpl.signUp(AuthData.userEntity);

      expect(result, throwsA(isA<UserConflictException>()));
    });

    test('Should throw server exception when data is valid', () async {
      /// arrange
      when(
        mockHttpClient.post(
          Uri.parse(AppData.registerUser),
          body: {
            'name': AuthData.name,
            'email': AuthData.email,
            'password': AuthData.password
          },
        ),
      ).thenThrow(ServerException());

      /// action
      final result = remoteAuthDataSourceImpl.signUp;

      /// assert
      expect(() async => result(AuthData.userEntity),
          throwsA(isA<ServerException>()));
    });
  });
}
