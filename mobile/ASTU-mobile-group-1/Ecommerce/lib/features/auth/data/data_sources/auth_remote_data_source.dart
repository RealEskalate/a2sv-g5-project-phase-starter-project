import 'dart:convert';
import 'dart:developer';

import 'package:http/http.dart' as http;

import '../../../../core/constants/constants.dart';
import '../../../../core/error/exception.dart';

import '../models/sign_in_model.dart';
import '../models/sign_up_model.dart';
import '../models/signed_in_model.dart';
import '../models/user_model.dart';
import 'auth_local_data_source.dart';

abstract class AuthRemoteDataSource {
  Future<SignedInModel> signIn(SignInModel signIn);
  Future<bool> signUp(SignUpModel signUp);
  Future<UserModel> getUser();
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  final http.Client _client;
  final AuthLocalDataSource _authLocalDataSource;

  AuthRemoteDataSourceImpl(
      {required http.Client client,
      required AuthLocalDataSource authLocalDataSource})
      : _authLocalDataSource = authLocalDataSource,
        _client = client;

  @override
  Future<UserModel> getUser() async {
    try {
      final token = await _authLocalDataSource.getToken();
      log(token);
      final response = await http.get(
        Uri.parse(Urls.getUser()),
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer $token',
        },
      );
      log(response.statusCode.toString());
      if (response.statusCode == 200) {
        final Map<String, dynamic> userDetails = json.decode(response.body);

        return UserModel.fromJson(userDetails['data'], token);
      } else {
        throw Exception();
      }
    } catch (e) {
      log(e.toString());
      throw UnknownException();
    }
  }

  @override
  Future<SignedInModel> signIn(SignInModel signIn) async {
    try {
      final url = Uri.parse(Urls.signIn());

      final response = await _client.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'email': signIn.email,
          'password': signIn.password,
        }),
      );

      final Map<String, dynamic> jsonResponse = json.decode(response.body);
      if (response.statusCode == 201) {
        _authLocalDataSource.cacheToken(jsonResponse['data']['access_token']);

        return SignedInModel.fromJson(jsonResponse);
      } else {
        throw ServerException();
      }
    } catch (_) {
      throw ServerException(
          message: 'Wrong email or password. Please try again.');
    }
  }

  @override
  Future<bool> signUp(SignUpModel signUp) async {
    try {
      final url = Uri.parse(Urls.signUp());

      final response = await _client.post(
        url,
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'name': signUp.name,
          'email': signUp.email,
          'password': signUp.password,
        }),
      );
      log(response.toString());
      if (response.statusCode == 201) {
        return true;
      } else {
        final Map<String, dynamic> jsonResponse = json.decode(response.body);

        String errorMessage;
        if (jsonResponse['message'] is List) {
          errorMessage = (jsonResponse['message'] as List).join('\n');
        } else {
          errorMessage = jsonResponse['message'] as String;
        }

        throw ServerException(message: errorMessage);
      }
    } on ServerException catch (e) {
      throw ServerException(message: e.message);
    } catch (e) {
      throw ServerException();
    }
  }
}
