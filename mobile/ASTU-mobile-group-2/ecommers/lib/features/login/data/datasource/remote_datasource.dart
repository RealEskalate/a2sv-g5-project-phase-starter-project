


import 'dart:convert';
import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../core/Error/failure.dart';
import '../../../../core/const/const.dart';
import '../../../../core/network/check_connectivity.dart';
import '../../domain/entity/login_entity.dart';
import '../model/model.dart';

abstract class RemoteDatasource {
  

   Future<Either<Failure,LoginEntity>>  login(String email, String password);
   Future<Either<Failure,bool>>  register(String email, String password, String fullName);
    

}


class RemoteDatasourceImpl implements RemoteDatasource {
  final http.Client client;
  final NetworkInfo networkInfo;
  final SharedPreferences sharedPreferences;
  RemoteDatasourceImpl({
    required this.client,
    required this.networkInfo,
    required this.sharedPreferences,
    });
  @override
   Future<Either<Failure,LoginEntity>>  login(String email, String password) async {
    try{
      final response = await http.post(Uri.parse(LoginApi.loginApi), body: {
        'email': email,
        'password': password,
      });
      
      if (response.statusCode == 201) {
  
        final data = jsonDecode(response.body);
        final findMe = await http.get(
          Uri.parse(LoginApi.findMe),
           headers: {
              'Authorization': 'Bearer ${data['data']['access_token']}',
            },
          );
          
   
        if (findMe.statusCode == 200) {
          final findMeData = jsonDecode(findMe.body);

          findMeData['data']['accessToken'] = data['data']['access_token'];
          
          sharedPreferences.setString('key', data['data']['access_token']);
          sharedPreferences.setString('name',findMeData['data']['name']);
          sharedPreferences.setString('email', findMeData['data']['email']);
          
          return Right(LoginModel.fromJson(findMeData).toEntity());
        } else {
          return left(const ServerFailure(message: 'server error'));
        }
      } else {
        
        return left(const ServerFailure(message: 'password or email is incorrect'));

      }
    } catch (e){
      return left(const ConnectionFailur(message:'try again'));
    }
  }
  
  @override
  Future<Either<Failure, bool>> register(String email, String password, String fullName) async{
    try {
      final networkCherck = await networkInfo.isConnected;
      if(networkCherck == false){
        return left(const ConnectionFailur(message: 'No internet connection'));
      }

      final response = await http.post(Uri.parse(LoginApi.registerApi), body: {
        'email': email,
        'password': password,
        'name': fullName,
      });
   
      final data = jsonDecode(response.body);
 
      if (response.statusCode == 201) {
        return right(true);
      } else {
        return left(ServerFailure(message: data['message']));
      }
    } catch (e) {
      return left(const ConnectionFailur(message: 'try again'));
    }
  }
}