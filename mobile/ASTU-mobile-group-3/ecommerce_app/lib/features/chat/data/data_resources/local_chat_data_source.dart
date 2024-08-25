
import 'package:shared_preferences/shared_preferences.dart';

abstract class LocalChatDataSource {
  
}

class LocalChatDataSourceImpl implements LocalChatDataSource {
  late final SharedPreferences sharedPreferences;

  LocalChatDataSourceImpl(this.sharedPreferences);
}