import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/entities/user_entitiy.dart';

class UserAccess {
  final String access_token;
  const UserAccess({required String access_token,}) : access_token = access_token;


  factory UserAccess.fromJson(Map<String, dynamic> json) {
    return UserAccess(
      // TO be deleteed
      access_token: json['data']['access_token'],
    );
  }
  


// ignore: empty_constructor_bodies
}