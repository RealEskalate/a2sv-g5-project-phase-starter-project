import 'dart:convert';

import 'package:http/http.dart' as http;

import '../../../../../core/constants/constants.dart';
import '../../../../../core/error/exception.dart';
import '../../model/log_in_model.dart';
import '../../model/sign_up_model.dart';
import '../../model/user_model.dart';
import '../local/local_data_source.dart';
import 'auth_remote_data_source.dart';

class AuthRemoteDatasourceImpl implements AuthRemoteDataSource {
  final http.Client client;

  final AuthLocalDataSource authLocalDataSource;
  AuthRemoteDatasourceImpl(
      {required this.client, required this.authLocalDataSource});

  @override
  Future<UserModel> getCurrentUser() async {
    final token = await authLocalDataSource.getToken();
    final response =
        await client.get(Uri.parse(Urls2.getCurrentUser()), headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $token',
    });
    if (response.statusCode == 200) {
      final user = UserModel.fromJson(jsonDecode(response.body)['data']);
      return user;
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> logIn(LogInModel logInModel) async {
    final response = await client.post(Uri.parse(Urls2.login()),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode(logInModel));

    if (response.statusCode == 201) {
      await authLocalDataSource
          .cacheToken(jsonDecode(response.body)['data']['access_token']);
    } else {
      throw ServerException();
    }
  }

  @override
  Future<void> logOut() async {
    try {
      await authLocalDataSource.removeToken();
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<void> signUp(SignUpModel signUpModel) async {
    final response = await client.post(
      Uri.parse(Urls2.signUp()),
      body: jsonEncode(signUpModel),
      headers: {'Content-Type': 'application/json'},
    );
    
    if (response.statusCode != 201) {
      final jsonMap = jsonDecode(response.body)['message'][0];
      throw ServerException(message: jsonMap);
    }
  }
}
