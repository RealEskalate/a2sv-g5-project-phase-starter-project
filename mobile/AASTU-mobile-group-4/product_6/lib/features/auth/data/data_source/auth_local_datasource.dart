import 'package:shared_preferences/shared_preferences.dart';

abstract class AuthLocalDataSource {
  Future<void> cacheAccessToken(String token);
  Future<String?> getAccessToken();
  Future<void> clearAccessToken();
}


const CACHED_ACCESS_TOKEN = 'CACHED_ACCESS_TOKEN';

class AuthLocalDataSourceImpl implements AuthLocalDataSource {
  final SharedPreferences sharedPreferences;

  AuthLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<void> cacheAccessToken(String token) async {
    await sharedPreferences.setString(CACHED_ACCESS_TOKEN, token);
  }

  @override
  Future<String?> getAccessToken() async {
    return sharedPreferences.getString(CACHED_ACCESS_TOKEN);
  }

  @override
  Future<void> clearAccessToken() async {
    await sharedPreferences.remove(CACHED_ACCESS_TOKEN);
  }

  
}
