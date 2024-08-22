import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exception.dart';

abstract class AuthLocalDataSource {
  Future<String> getToken();
  Future<Unit> cacheToken(String token);
  Future<Unit> logout();
}

const TOKEN = 'token';

class AuthLocalDataSourceImpl extends AuthLocalDataSource {
  final SharedPreferences sharedPreferences;

  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<Unit> cacheToken(String token) {
    try {
      final jsonToken = json.encode(token);
      sharedPreferences.setString(TOKEN, jsonToken);
      return Future.value(unit);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<String> getToken() {
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

  @override
  Future<Unit> logout(){
    try{
      sharedPreferences.remove(TOKEN);
      return Future.value(unit);
    } catch(e){
      throw CacheException();
    }
  }
}
