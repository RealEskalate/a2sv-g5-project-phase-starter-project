import 'dart:convert';

import 'package:ecommerce/features/auth/data/model/user_model.dart';
import 'package:ecommerce/features/auth/domain/entities/user.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class UserLocalDataSource {
  Future<void> cacheUser(UserModel userTocache);
  Future<UserModel> getCachedUser();
}

class UserLocalDataSourceImpl implements UserLocalDataSource {
  var keyname = 'cachedusers';
  SharedPreferences sharedPreferences;
  UserLocalDataSourceImpl({required this.sharedPreferences});
  @override
  Future<void> cacheUser(UserModel userTocache) async {
    final jsonUser = jsonEncode(userTocache.toJson());
    await sharedPreferences.setString(keyname, jsonUser);
  }
  Future<UserModel> getCachedUser() async {
    final jsonUser = sharedPreferences.getString(keyname);
    if (jsonUser != null) {
      return UserModel.fromJson(jsonDecode(jsonUser));
    } else {
      throw Exception('No user cached');
    }
  }

}