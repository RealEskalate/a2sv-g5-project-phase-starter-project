import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../model/user_model.dart';

abstract class UserLocalDataSource{
  
  Future<void> cacheUser(UserModel userModel);
  UserModel getUser();
}

class UserLocalDataSourceImpl extends UserLocalDataSource{
  final SharedPreferences sharedPreferences;
  UserLocalDataSourceImpl({required this.sharedPreferences});
  String key = 'user';
  
  @override
  Future<void> cacheUser(UserModel userModel) async {
    final jsonData = json.encode(userModel.toJson());
    bool result = await sharedPreferences.setString(key, jsonData);

    if(result == false){
      throw Exception('Error');
    }
  }
  
  @override
  UserModel getUser() {

    var user = sharedPreferences.getString(key);
    if (user != null) {
      return UserModel.fromJson(json.decode(user));
    }else{
      throw Exception('Error');
    }
    
  }


}