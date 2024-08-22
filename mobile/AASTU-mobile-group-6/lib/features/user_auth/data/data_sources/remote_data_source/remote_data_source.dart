import 'dart:convert';
import 'dart:math';

import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/core/errors/failure/failures.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_access.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

abstract class UserRemoteDataSource {
  Future<Either<Failure,String>> loginUser(UserEntity user);
  Future<String> registerUser(UserModel user);
}


class UserRemoteDataSourceImpl extends UserRemoteDataSource{
  final http.Client client;
    UserRemoteDataSourceImpl({required this.client});

  @override
  Future<Either<Failure,String>> loginUser(UserEntity user) async{
    
    var body = user.toJsonentitiy();

    final response = await client.post(Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/login'),body: body);

    
      if (response.statusCode==201){
        var data = jsonDecode(response.body);
        var userAccess = data['data']['access_token'];
        var cache = await SharedPreferences.getInstance();
        cache.setString('access_token', userAccess);
        var newt = cache.getString('access_token');
        print(newt);
        return Right('Successfully Logged In');
      
      }
      else{
        return Left(ServerFailure('Authentication Failed'));
      }
    
  }
  @override
  Future<String> registerUser(UserModel user) async{
    
  var body = user.toJson();


    final response = await client.post(Uri.parse('https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/register'),body: body);
    print(response.statusCode);
    if (response.statusCode==201){
      
      return 'Successfully Registered';

    }else{
      return "Authentication Failed";
    }
    
  }

}