import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exceptions/product_exceptions.dart';
import '../../domain/entities/user_entity.dart';
import '../model/token_model.dart';
import '../model/user_model.dart';

abstract class RemoteAuthDataSource {
  Future<TokenModel> logIn(UserEntity user);
  Future<UserModel> signUp(UserEntity user);
}

class RemoteAuthDataSourceImpl implements RemoteAuthDataSource {
  final http.Client client;

  RemoteAuthDataSourceImpl({required this.client});

  @override
  Future<TokenModel> logIn(UserEntity user) async {
    try {
      final result = await client.post(
        Uri.parse(AppData.logInUser),
        body: {'email': user.email, 'password': user.password},
      ).timeout(const Duration(seconds: 15));

      if (result.statusCode == 201) {
        final Map<String, dynamic> jsonFormat = json.decode(result.body);
        return TokenModel.fromJson(jsonFormat['data']);
      } else if (result.statusCode == 401) {
        throw LoginException();
      } else {
        throw ServerException();
      }
    } on LoginException {
      rethrow;
    } on ServerException {
      rethrow;
    } on Exception {
      rethrow;
    }
  }

  @override
  Future<UserModel> signUp(UserEntity user) async {
    try {
      final result = await client.post(
        Uri.parse(AppData.registerUser),
        body: {
          'name': user.name,
          'email': user.email,
          'password': user.password
        },
      ).timeout(const Duration(seconds: 15));

      if (result.statusCode == 201) {
        final Map<String, dynamic> jsonFormat = json.decode(result.body);
        return UserModel.fromJson(jsonFormat['data']);
      } else if (result.statusCode == 409) {
        throw UserConflictException();
      } else {
        throw ServerException();
      }
    } on UserConflictException {
      rethrow;
    } on ServerException {
      rethrow;
    } on Exception {
      rethrow;
    }
  }
}
