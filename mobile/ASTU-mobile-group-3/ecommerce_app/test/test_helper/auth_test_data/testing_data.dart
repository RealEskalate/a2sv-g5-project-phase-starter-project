import 'dart:io';

import 'package:ecommerce_app/features/auth/data/model/signed_up_user_model.dart';
import 'package:ecommerce_app/features/auth/data/model/token_model.dart';
import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:ecommerce_app/features/auth/domain/entities/signed_up_user_entity.dart';
import 'package:ecommerce_app/features/auth/domain/entities/token_entity.dart';
import 'package:ecommerce_app/features/auth/domain/entities/user_entity.dart';
import 'package:flutter_test/flutter_test.dart';

class AuthData {
  static const String token =
      'eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6InVzZXJAZ21haWwuY29tIiwiaWF0IjoxNzIzNzIwNTk3LCJleHAiOjE3MjQxNTI1OTd9.Mz9mIyOPHgyONb3bLvDO9N2wwF562Xb4nQnJEhtV3Fk';
  static const String name = 'Mr. User';
  static const String email = 'user@gmail.com';
  static const String password = 'userpassword';
  static const String id = '66bde36e9bbe07fc39034cdd';
  static const TokenEntity tokenEntity = TokenEntity(token: token);
  static const TokenModel tokenModel = TokenModel(token: token);

  static const signedUpUserEntity =
      SignedUpUserEntity(id: id, email: email, name: name);
  static const signedUpUserModel =
      SignedUpUserModel(id: id, email: email, name: name);

  static const userEntity =
      UserEntity(name: name, email: email, password: password);
  static const userModel =
      UserModel(name: name, email: email, password: password);
  static String readJson() {
    String dir = Directory.current.path;
    if (dir.contains('/test')) {
      dir = dir.replaceAll('/test', '');
    }

    dir = '$dir/test/test_helper/auth_test_data/json_request.json';

    return File(dir).readAsStringSync();
  }
}

void main() {
  test('description', () {
    print(AuthData.readJson());
  });
}
