import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../../core/error/exceptions.dart';
import '../models/user_model.dart';

abstract class UserRemoteDataSource {
  Future<String> loginUser(String email, String password);
  Future<UserModel> registerUser(String email, String password, String name);
}

class UserRemoteDataSourceImpl implements UserRemoteDataSource {
  final http.Client client;

  UserRemoteDataSourceImpl({required this.client});

  @override
  Future<String> loginUser(String email, String password) async {
    final response = await client.post(
      Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/login'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({'email': email, 'password': password}),
    );

    if (response.statusCode == 201) {
      return json.decode(response.body)['data']['access_token'];
    } else {
      throw ServerException();
    }
  }

  @override
  Future<UserModel> registerUser(String email, String password, String name) async {
    final response = await client.post(
      Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/register'),
      headers: {'Content-Type': 'application/json'},
      body: json.encode({'email': email, 'password': password, 'name': name}),
    );

    if (response.statusCode == 201) {
      return UserModel.fromJson(json.decode(response.body)['data']);
    } else {
      throw ServerException();
    }
  }
}
