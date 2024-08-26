// import 'package:shared_preferences/shared_preferences.dart';

// class AuthServices {
//   static Future<String> getToken() async {
//     var sharedPreferences = await SharedPreferences.getInstance();

//     final userToken = sharedPreferences.getString("user");
//     print(userToken);
//     if (userToken != null) {
//       return userToken;
//     } else {
//       throw Exception("no user have logged in");
//     }
//   }
// }
import 'package:shared_preferences/shared_preferences.dart';
import 'package:logger/logger.dart';

class AuthServices {
  static final logger = Logger();

  static Future<String> getToken() async {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();

    final userToken = sharedPreferences.getString("user") ?? "";
    logger.d(userToken);
    if (userToken.isNotEmpty) {
      return userToken;
    } else {
      throw Exception("no user has logged in");
    }
  }
}
