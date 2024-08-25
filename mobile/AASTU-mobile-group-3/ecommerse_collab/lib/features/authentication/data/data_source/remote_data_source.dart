import 'dart:convert';

import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../domain/entity/user.dart';
import '../model/user_model.dart';

abstract class UserRemoteDataSource{
  Future<UserModel> logIn(String email, String password);  
  Future<void> logOut();
  Future<UserModel> signUp(String username, String password, String email);
  Future<UserModel> getMe();
}


class UserRemoteDataSourceImpl extends UserRemoteDataSource{

  final http.Client client;
  UserRemoteDataSourceImpl(this.client);
  
  @override
  Future<void> logOut() async {
    try{
    SharedPreferences sharedPreferences = await SharedPreferences.getInstance();
    sharedPreferences.remove('token');} catch (e){
      throw (e);
    }
    }

  @override
  Future<UserModel> signUp(String username, String password, String email) async {
    print("Remote data source");
    
    final response = await client.post(
      Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v3/auth/register'),
      headers: {
        'Content-Type': 'application/json',
      },
      body: json.encode({
        'name': username,
        'email': email,
        'password': password,
      }),
    );
    print("from remote ${response.body}");
    print('Response status: ${response.statusCode}');
    print('Response body: ${response.body}');

    if (response.statusCode == 201) {
      // final jsonData = json.encode(response.body) as Map<String, dynamic>;
      final jsonData = json.decode(response.body) as Map<String, dynamic>;
      final jsonFinal = jsonData['data'];
      return UserModel.fromJson(jsonFinal);      

    } else {
      throw Exception('Failed to load user');
    }
    
  }
  


  @override
  Future<UserModel> logIn(String email, String password) async {
    final response = await client.post(
      Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v3/auth/login'),
      body: {
        'email': email,
        'password': password,
      },
    );
    if (response.statusCode == 201){
    SharedPreferences prefs = await SharedPreferences.getInstance();
    var result = prefs.setString('token', json.decode(response.body)['data']['access_token']);
    final data = getMe();
    print("data $data");
    return data;
    }else{
      throw Exception('Failed to load user');}
    
  }
  
  @override
  Future<UserModel> getMe() async {
    
      final newURL = Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v3/users/me');
      SharedPreferences prefs = await SharedPreferences.getInstance();
      var headers = {
        'Authorization': "Bearer ${prefs.getString('token')}",
        'Content-Type': 'application/json',
         };
      final response_2 = await client.get(newURL, headers: headers);
      
      print(response_2.body);
      
      if (response_2.statusCode != 200) {
        throw Exception('Failed to load user');
      }
      final jsonFinal =  json.decode(response_2.body) as Map<String, dynamic>;
      print("json final ${jsonFinal['data']}");

      return UserModel.fromJson(jsonFinal['data']);

    }
  }
