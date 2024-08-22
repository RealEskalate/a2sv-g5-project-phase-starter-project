import 'package:shared_preferences/shared_preferences.dart';

abstract class UserLocalDataSource {
  Future<void> saveAccessToken(String token);
  Future<String?> getAccessToken();
  Future<void> deleteAccessToken();
}

class UserLocalDataSourceImpl implements UserLocalDataSource {
  static const String _accessTokenKey = 'ACCESS_TOKEN';
  SharedPreferences sharedPreferences;
  UserLocalDataSourceImpl({required this.sharedPreferences});
  

  @override
  Future<void> saveAccessToken(String token) async {
    
    await sharedPreferences.setString(_accessTokenKey, token);
  }

  @override
  Future<String?> getAccessToken() async {
    
    return sharedPreferences.getString(_accessTokenKey);
  }

  @override
  Future<void> deleteAccessToken() async {
    await sharedPreferences.remove(_accessTokenKey);
  }
}
