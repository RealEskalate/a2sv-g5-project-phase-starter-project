import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/error/exceptions.dart';
import 'package:task_9/features/user/data/data_sources/user_remote_data_source.dart';
import 'package:task_9/features/user/data/models/user_model.dart';

import '../../../product/data/datasources/product_remote_data_source_test.mocks.dart';


@GenerateMocks([http.Client])
void main() {
  late UserRemoteDataSourceImpl dataSource;
  late MockClient mockHttpClient;

  setUp(() {
    mockHttpClient = MockClient();
    dataSource = UserRemoteDataSourceImpl(client: mockHttpClient);
  });

  group('loginUser', () {
    const tEmail = 'test@example.com';
    const tPassword = 'password';
    const tAccessToken = 'mock_access_token';

    test('should perform a POST request on the login URL', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response(
          json.encode({'data': {'access_token': tAccessToken}}), 201));

      // Act
      await dataSource.loginUser(tEmail, tPassword);

      // Assert
      verify(mockHttpClient.post(
        Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/login'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({'email': tEmail, 'password': tPassword}),
      ));
    });

    test('should return access token when the response code is 201', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response(
          json.encode({'data': {'access_token': tAccessToken}}), 201));

      // Act
      final result = await dataSource.loginUser(tEmail, tPassword);

      // Assert
      expect(result, equals(tAccessToken));
    });

    test('should throw a ServerException when the response code is not 201', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response('Something went wrong', 400));

      // Act
      final call = dataSource.loginUser(tEmail, tPassword);

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });
  });

  group('registerUser', () {
    const tEmail = 'test@example.com';
    const tPassword = 'password';
    const tName = 'Test Name';
    const tUserModel = UserModel(email: tEmail, name: tName);

    test('should perform a POST request on the register URL', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response(
          json.encode({'data': tUserModel.toJson()}), 201));

      // Act
      await dataSource.registerUser(tEmail, tPassword, tName);

      // Assert
      verify(mockHttpClient.post(
        Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/register'),
        headers: {'Content-Type': 'application/json'},
        body: json.encode({'email': tEmail, 'password': tPassword, 'name': tName}),
      ));
    });

    test('should return UserModel when the response code is 201', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response(
          json.encode({'data': tUserModel.toJson()}), 201));

      // Act
      final result = await dataSource.registerUser(tEmail, tPassword, tName);

      // Assert
      expect(result, equals(tUserModel));
    });

    test('should throw a ServerException when the response code is not 201', () async {
      // Arrange
      when(mockHttpClient.post(
        any,
        headers: anyNamed('headers'),
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response('Something went wrong', 400));

      // Act
      final call = dataSource.registerUser(tEmail, tPassword, tName);

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });
  });
}
