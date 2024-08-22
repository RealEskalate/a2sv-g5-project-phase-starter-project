import 'package:shared_preferences/shared_preferences.dart';

class AuthLocalDataSource {
  SharedPreferences sharedPreferences;
  AuthLocalDataSource({required this.sharedPreferences});

  Future<void> cacheToken(String Token) async {
    await sharedPreferences.setString("user", Token);
  }

   static Future<String> getToken() async {
    var sharedPreferences = await SharedPreferences.getInstance();

    final userToken = sharedPreferences.getString("user");
    print(userToken);
    if (userToken != null) {
      return userToken;
    } else {
      throw Exception("no user have logged in");
    }
  }
}
