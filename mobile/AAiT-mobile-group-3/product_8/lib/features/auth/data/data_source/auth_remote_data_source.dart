import 'dart:convert';
import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:http/http.dart' as http;
import '../../../../core/constants/constants.dart';
import '../../../../core/exception/exception.dart';
import '../models/authenticated_model.dart';
import '../models/sign_in_user_model.dart';
import '../models/sign_up_user_model.dart';
import '../models/user_data_model.dart';

abstract class AuthRemoteDataSource {
  Future<AuthenticatedModel> signIn(SignInUserModel signInUserModel);
  Future<Unit> signUp(SignUpUserModel signUpUserModel);
  Future<UserDataModel> getUser(String token);
}

class AuthRemoteDataSourceImpl implements AuthRemoteDataSource {
  final http.Client client;

  AuthRemoteDataSourceImpl({required this.client});

  @override
  Future<AuthenticatedModel> signIn(SignInUserModel signInUserModel) async {
    try {
      final http.Response response = await client.post(
        Uri.parse('${Urls.autUrl}/auth/login'),
        body: json.encode(signInUserModel.toJson()),

        headers: {'Content-Type': 'application/json'},
      );
      print(response.body);
      if (response.statusCode == 201) {
        final data = json.decode(response.body)['data'];

        return AuthenticatedModel.fromJson(data);
      } else if (response.statusCode == 401) {
        throw UnauthorizedException();
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }

  @override
  Future<Unit> signUp(SignUpUserModel signUpUserModel) async {
    try {
      final http.Response response = await client.post(
        Uri.parse('${Urls.autUrl}/auth/register'),
        body: json.encode(signUpUserModel.toJson()
        ),
        headers: {'Content-Type': 'application/json'},
      );
      if (response.statusCode == 201) {
        return unit;
      }else if(response.statusCode == 409){
        throw UserAlreadyExistsException();
      }
       else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }
  
  @override
  Future<UserDataModel> getUser(String token) async {
    try {
      final http.Response response = await client.get(
        Uri.parse('${Urls.autUrl}/users/me'),
        headers: {'Authorization' : 'Bearer $token'}
      );
      if (response.statusCode == 200) {
        final data = json.decode(response.body)['data'];
        return UserDataModel.fromJson(data);
      } else if (response.statusCode == 401) {
        throw UnauthorizedException();
      } else {
        throw ServerException();
      }
    } on SocketException {
      throw const SocketException('No Internet Connection');
    }
  }
}
