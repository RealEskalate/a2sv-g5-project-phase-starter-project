import 'dart:convert';

import 'package:application1/features/authentication/data/model/log_in_model.dart';
import 'package:application1/features/authentication/domain/entities/log_in.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helper/dummy_data/json_reader.dart';

void main() {
  const tLogInModel =
      LogInModel(email: 'ley@gmail.com', password: '1234');
  const String signUpModelPath =
      'helper/dummy_data/auth_model/dummy_log_in_model.json';
  test(
    'should extend from the User entity',
    () async {
      // Assert
      expect(tLogInModel, isA<LogInEntity>());
    },
  );

  test(
    'should return a json file',
    () async {
      //arrange
      final expectedJsonString = json.decode(readJson(signUpModelPath));
      //act
      final result = tLogInModel.toJson();
      //assert
      expect(result, expectedJsonString);
    },
  );
}
