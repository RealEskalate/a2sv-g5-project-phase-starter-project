
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../../core/error/exception.dart';
import 'local_data_source.dart';

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final SharedPreferences sharedPreferences;

  // ignore: non_constant_identifier_names
  final String CACHED_TOKEN = 'CACHED_TOKEN';
  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<bool> cacheToken(String token) async {
    final result = await sharedPreferences.setString(CACHED_TOKEN, token);
    if (result != true) {
      throw CacheException();
    } else {
      return result;
    }
  }

  @override
  Future<String> getToken() async {
    final token = sharedPreferences.getString(CACHED_TOKEN);
    if (token != null) {
      return token;
    } else {
      throw CacheException();
    }
  }

  @override
  Future<bool> removeToken() async {
    final val = await sharedPreferences.remove(CACHED_TOKEN);
    if (val != true) {
      throw CacheException();
    } else{
      return true;
    }
  }
}
