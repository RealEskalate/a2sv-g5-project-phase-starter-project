import 'dart:convert';

import 'package:application1/features/authentication/data/model/sign_up_model.dart';
import 'package:application1/features/authentication/domain/entities/sign_up.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helper/dummy_data/json_reader.dart';

void main() {
  const tSignUpModel =
      SignUpModel(email: 'ley@gmail.com', username: 'ley', password: '1234');
  const String signUpModelPath =
      'helper/dummy_data/auth_model/dummy_sign_up_model.json';
  test(
    'should extend from the User entity',
    () async {
      // Assert
      expect(tSignUpModel, isA<SignUpEntity>());
    },
  );

  test(
    'should return a json file',
    () async {
      //arrange
      final expectedJsonString = json.decode(readJson(signUpModelPath));
      //act
      final result = tSignUpModel.toJson();
      //assert
      expect(result, expectedJsonString);
    },
  );
}
