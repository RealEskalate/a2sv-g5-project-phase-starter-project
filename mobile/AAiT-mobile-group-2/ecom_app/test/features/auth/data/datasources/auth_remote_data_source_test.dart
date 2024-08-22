import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/auth/data/datasources/auth_remote_data_source.dart';
import 'package:ecom_app/features/auth/data/models/authenticated_model.dart';
import 'package:ecom_app/features/auth/data/models/login_model.dart';
import 'package:ecom_app/features/auth/data/models/register_model.dart';
import 'package:ecom_app/features/auth/data/models/user_data_model.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late AuthRemoteDataSourceImpl authRemoteDataSourceImpl;

  setUp(() {
    mockHttpClient = MockHttpClient();
    authRemoteDataSourceImpl = AuthRemoteDataSourceImpl(client: mockHttpClient);
  });

  final tLoginModel = LoginModel(email: 'email', password: 'password');
  final tRegisterModel =
      RegisterModel(email: 'email', password: 'password', name: 'name');

  group('Login Impl', () {
    final jsonLoginResponse = jsonEncode({
      'statusCode': 201,
      'message': '',
      'data': {
        'access_token':
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ0MjQzLCJleHAiOjE3MjQ1NzYyNDN9.oyC9gsD5ozRSCRMsC8M5WE8Wwxyzsbcn6-l7dLS8fsQ'
      }
    });

    final tAuthModel = AuthenticatedModel(
        token:
            'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ0MjQzLCJleHAiOjE3MjQ1NzYyNDN9.oyC9gsD5ozRSCRMsC8M5WE8Wwxyzsbcn6-l7dLS8fsQ');

    test('should return an AuthenticatedModel when successful', () async {
      //arrange
      when(mockHttpClient.post(any, body: tLoginModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonLoginResponse, 201));

      //act
      final result = await authRemoteDataSourceImpl.login(tLoginModel);

      //assert
      expect(result, tAuthModel);
    });
    test('should throw a unauthorized exception when statuscode is 401',
        () async {
      //arrange
      when(mockHttpClient.post(any, body: tLoginModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonLoginResponse, 401));

      //assert
      expect(() => authRemoteDataSourceImpl.login(tLoginModel),
          throwsA(isA<UnauthorizedException>()));
    });
    test('should throw a server exception when unsuccessful', () async {
      //arrange
      when(mockHttpClient.post(any, body: tLoginModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonLoginResponse, 400));

      //assert
      expect(() => authRemoteDataSourceImpl.login(tLoginModel),
          throwsA(isA<ServerException>()));
    });
  });
  group('Register Impl', () {
    final jsonRegisterResponse = jsonEncode({
      'statusCode': 201,
      'message': '',
      'data': {
        'id': '66c45cdf6198f150e643c9e2',
        'name': 'Mr. User15',
        'email': 'user15@gmail.com'
      }
    });

    test('should return an AuthenticatedModel when successful', () async {
      //arrange
      when(mockHttpClient.post(any, body: tRegisterModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonRegisterResponse, 201));

      //act
      final result = await authRemoteDataSourceImpl.register(tRegisterModel);

      //assert
      expect(result, unit);
    });
    test('should throw a unauthorized exception when statuscode is 401',
        () async {
      //arrange
      when(mockHttpClient.post(any, body: tRegisterModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonRegisterResponse, 409));

      //assert
      expect(() => authRemoteDataSourceImpl.register(tRegisterModel),
          throwsA(isA<UserAlreadyExistsException>()));
    });
    test('should throw a server exception when unsuccessful', () async {
      //arrange
      when(mockHttpClient.post(any, body: tRegisterModel.toJson()))
          .thenAnswer((_) async => http.Response(jsonRegisterResponse, 400));

      //assert
      expect(() => authRemoteDataSourceImpl.register(tRegisterModel),
          throwsA(isA<ServerException>()));
    });
  });


  group('Get User Impl', () {
    final jsonUserResponse = jsonEncode({
    'statusCode': 200,
    'message': '',
    'data': {
        'id': '66bde36e9bbe07fc39034cdd',
        'email': 'user@gmail.com',
        'name': 'Mr. User'
    }
});
  final tUserDataModel = UserDataModel(email: 'user@gmail.com', name: 'Mr. User');
  final tToken = 'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwic3ViIjoiNjZiZGUzNmU5YmJlMDdmYzM5MDM0Y2RkIiwiaWF0IjoxNzI0MTQ3MzA2LCJleHAiOjE3MjQ1NzkzMDZ9.G2g2i_llh6EdyO4lD08pBCUYSwBvoqhbSMLPHW-jHlw';
    test('should return an AuthenticatedModel when successful', () async {
      //arrange
      when(mockHttpClient.get(any, headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response(jsonUserResponse, 200));

      //act
      final result = await authRemoteDataSourceImpl.getUser(tToken);

      //assert
      expect(result, tUserDataModel);
    });
    test('should throw a unauthorized exception when statuscode is 401',
        () async {
      //arrange
      when(mockHttpClient.get(any, headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response(jsonUserResponse, 401));

      //assert
      expect(() => authRemoteDataSourceImpl.getUser(tToken),
          throwsA(isA<UnauthorizedException>()));
    });
    test('should throw a server exception when unsuccessful', () async {
      //arrange
      when(mockHttpClient.get(any, headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response(jsonUserResponse, 400));

      //assert
      expect(() => authRemoteDataSourceImpl.getUser(tToken),
          throwsA(isA<ServerException>()));
    });
  });
}
