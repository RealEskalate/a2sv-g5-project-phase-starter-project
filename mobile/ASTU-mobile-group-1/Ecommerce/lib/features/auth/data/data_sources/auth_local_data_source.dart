import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exception.dart';

abstract class AuthLocalDataSource {
  Future<bool> logOut();
  Future<bool> checkSignedIn();
  Future<bool> cacheToken(String token);
  Future<String> getToken();
}

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final SharedPreferences prefs;

  AuthLocalDataSourceImpl({required this.prefs});

  @override
  Future<String> getToken() async {
    try {
      final token = prefs.getString('accessToken');
      if (token != null) {
        return token;
      }
      throw Exception();
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<bool> logOut() async {
    try {
      await prefs.setBool('isLoggedIn', false);
      await prefs.remove('accessToken');
      return Future.value(true);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<bool> checkSignedIn() {
    try {
      bool isLoggedIn = prefs.getBool('isLoggedIn') ?? false;
      return Future.value(isLoggedIn);
    } catch (e) {
      throw UnknownException();
    }
  }

  @override
  Future<bool> cacheToken(String token) async {
    try {
      await prefs.setString('accessToken', token);
      await prefs.setBool('isLoggedIn', true);

      return Future.value(true);
    } catch (e) {
      throw UnknownException();
    }
  }
}
