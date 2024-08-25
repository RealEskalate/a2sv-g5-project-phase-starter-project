import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../models/auth_model.dart';

abstract class UserLogInLocalDataSource {
  Future<void> cacheToken(String token);
  Future<String> getToken();
  Future<void> deleteUser();
  Future<void> setUser(SignUPModel signUpModel);
  Future<SignUPModel?> getUser();
}

const tokenKey = '';

class UserLogInLocalDataSourceImpl implements UserLogInLocalDataSource {
  final SharedPreferences sharedPreferences;
  UserLogInLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<void> cacheToken(String token) async {
    await sharedPreferences.setString('tokenKey', token);
  }

  @override
  Future<void> deleteUser() {
    return sharedPreferences.remove('tokenKey');
  }

  @override
  Future<String> getToken() async {
    final token = sharedPreferences.getString(tokenKey);
    if (token == null) {
      throw Exception('token not found');
    }
    return token;
  }

  @override
  Future<SignUPModel?> getUser() async {
    final String? userDataString = sharedPreferences.getString('userData');

    if (userDataString != null) {
      final SignUPModel signUPModel =
          SignUPModel.fromJson(jsonDecode(userDataString));
      return signUPModel;
    } else {
      throw Exception('No user data found in sharedpreferences');
    }
  }

  @override
  Future<void> setUser(SignUPModel signUpModel) {
    final String userDataString = jsonEncode(signUpModel.toJson());
    return sharedPreferences.setString('userData', userDataString);
  }
}
