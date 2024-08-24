import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;

import '../../../../../core/error/exceptions.dart';
import '../../models/auth_model.dart';

abstract class AuthRemoteDataSource {
  Future<String> signUp(SignUPModel newUser);
  Future<String> logIn(LogInModel oldUser);
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  late final http.Client client;

  AuthRemoteDataSourceImpl({required this.client});
  final baseUrl = 'https://g5-flutter-learning-path-be.onrender.com/api/v3';

  @override
  Future<String> signUp(SignUPModel newUser) async {
    try {
      var url = Uri.parse('$baseUrl/auth/register');
      final user = await client.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(newUser.toJson()),
      );
      if (user.statusCode == 201) {
        return 'user created';
      } else {
        throw ServerException(message: 'Failed to create user');
      }
    } catch (e) {
      throw ServerException(message: 'Unexpected error: ${e.toString()}');
    }
  }

  @override
  Future<String> logIn(LogInModel oldUser) async {
    try {
      var url = Uri.parse('$baseUrl/auth/login');

      final res = await http.post(url,
          headers: {'Content-Type': 'application/json'},
          body: jsonEncode(oldUser.toJson()));

      if (res.statusCode == 201 || res.statusCode == 200) {
        print(jsonDecode(res.body)['data']['access_token']);
        return jsonDecode(res.body)['data']['access_token'];
      } else {
        throw ServerException(message: 'server failure');
      }
    } catch (e) {
      throw ServerException(message: e.toString());
    }
  }
}
