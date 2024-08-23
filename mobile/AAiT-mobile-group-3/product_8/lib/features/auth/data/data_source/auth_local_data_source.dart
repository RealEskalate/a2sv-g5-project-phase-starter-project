import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/exception/exception.dart';

abstract class AuthLocalDataSource {
  Future<Unit> cacheToken(String token);
  Future<String> getToken();
  Future<Unit> deleteToken();
}

const TOKEN = 'TOKEN';

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final SharedPreferences sharedPreferences;

  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<Unit> cacheToken(String token) async {
    try {
      final jsonToken = json.encode(token);
      sharedPreferences.setString(TOKEN, jsonToken);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<Unit> deleteToken() async {
    try {
      sharedPreferences.remove(TOKEN);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<String> getToken() async {
    try {
      final token = sharedPreferences.getString(TOKEN);
      if (token != null) {
        final decodedJson = json.decode((token));
        return Future.value(decodedJson);
      } else {
        throw CacheException();
      }
    } catch (e) {
      throw CacheException();
    }
  }
}
