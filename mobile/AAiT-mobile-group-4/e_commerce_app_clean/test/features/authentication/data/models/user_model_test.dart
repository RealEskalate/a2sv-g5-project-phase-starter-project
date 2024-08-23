import 'dart:convert';

import 'package:application1/features/authentication/data/model/user_model.dart';
import 'package:application1/features/authentication/domain/entities/user_data.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helper/dummy_data/json_reader.dart';

void main() {
  const tUserModel = UserModel(email: 'ley@gmail.com',name: 'ley');
  const String userModelPath = 'helper/dummy_data/auth_model/dummy_user_model.json';
  test(
    'should extend from the User entity',
    () async {
      // Assert
      expect(tUserModel, isA<UserEntity>());
    },
  );

  test(
    'should receive a fromjson function',
    () async {
      //arrange
      final Map<String, dynamic> jsonMap = json
          .decode(readJson(userModelPath));
      //act
      final result = UserModel.fromJson(jsonMap);
      //assert
      expect(tUserModel, result);
    },
  );
}
