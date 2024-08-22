import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../../core/errors/exception.dart';
import '../../domain/entity/auth_entity.dart';
import '../models/auth_model.dart';

abstract class AuthRemoteDataSource {
  Future<AuthResponseEntity> login(AuthEntity authEntity);
  Future<void> register(SentUserEntity userEntity);
  Future<UserEntity> getUserProfile(String accessToken);
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  final http.Client client;
  final String url = 'https://g5-flutter-learning-path-be.onrender.com/api/v2';

  AuthRemoteDataSourceImpl({required this.client});

  @override
  Future<AuthResponseEntity> login(AuthEntity authEntity) async {
    final response = await client.post(
      Uri.parse('$url/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'email': authEntity.email,
        'password': authEntity.password,
      }),
    );

    if (response.statusCode == 201) {
      return AuthResponseModel.fromJson(json.decode(response.body)['data']);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> register(SentUserEntity userEntity) async {
    final response = await client.post(
      Uri.parse('$url/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({
        'name': userEntity.name,
        'email': userEntity.email,
        'password': userEntity.password,
      }),
    );
    print(userEntity.name);
    print(userEntity.email);
    print(userEntity.password);
    print('from auth_remote_datasource.dart');
    print(response.body);
    if (response.statusCode == 201) {
      print('suuuuu');
      // print(SentUserModel.fromJson(json.decode(response.body)));
      return Future.value();
    } else {
      throw ServerException();
    }
  }

  @override
  Future<UserEntity> getUserProfile(String accessToken) async {
    final response = await client.get(
      Uri.parse('$url/users/me'),
      headers: {'Authorization': 'Bearer $accessToken'},
    );

    if (response.statusCode == 200) {
      return UserModel.fromJson(json.decode(response.body)['data']);
    } else {
      throw ServerException();
    }
  }
}
