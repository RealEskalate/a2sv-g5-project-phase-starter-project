import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/constants/constants.dart';
import '../../../../core/errors/exceptions/product_exceptions.dart';
import '../model/token_model.dart';

abstract class AuthLocalDataSource {
  Future<bool> saveToken(TokenModel token);
  Future<TokenModel> getToken();
  Future<bool> clearToken();
}

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final SharedPreferences sharedPreferences;

  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<bool> saveToken(TokenModel token) async {
    try {
      await sharedPreferences.setString(AppData.tokenPlacement, token.token);
      return true;
    } on Exception {
      throw CacheException();
    }
  }

  @override
  Future<TokenModel> getToken() async {
    try {
      final result = sharedPreferences.getString(AppData.tokenPlacement);
      if (result == null) {
        throw CacheException();
      } else {
        return TokenModel(token: result);
      }
    } on Exception {
      throw CacheException();
    }
  }

  @override
  Future<bool> clearToken() async {
    try {
      await sharedPreferences.remove(AppData.tokenPlacement);
      return true;
    } on Exception {
      throw CacheException();
    }
  }
}
