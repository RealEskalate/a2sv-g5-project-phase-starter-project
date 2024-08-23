import 'package:shared_preferences/shared_preferences.dart';

Future<bool> userValidation() async {
  try {
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    final value = sharedPreferences.getString('key');
   
    if (value == null) {
      return false;
    }
    return true;
  } catch (e) {
    return false;
  }
}
