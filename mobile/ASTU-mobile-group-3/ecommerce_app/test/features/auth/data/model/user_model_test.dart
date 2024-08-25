import 'dart:convert';

import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:ecommerce_app/features/auth/domain/entities/user_entity.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  test('User data model Should be subclass of user entity ', () {
    expect(AuthData.userEntity, isA<UserEntity>());
  });

  test('Should return model when strings are given', () {
    const result = UserModel(
        name: AuthData.name,
        email: AuthData.email,
        password: AuthData.password,
        id: AuthData.id,
        v: 0);

    expect(result, AuthData.userModel);
  });

  test('Should return an entity from model', () {
    final result = AuthData.userModel.toEntity();

    expect(result, AuthData.userEntity);
  });

  test('Should return valid model from strings', () {
    /// action
    final model = UserModel.fromStrings(AuthData.userModel.name,
        AuthData.userModel.email, AuthData.userModel.password);

    /// datas
    final expectedData = UserModel(
      name: AuthData.userModel.name,
      email: AuthData.userModel.email,
      password: AuthData.userModel.password,
      id: '',
      v: 0,
    );

    ///assert
    expect(model, expectedData);
  });

  test('Should return valid model from json', () {
    Map<String, dynamic> decoded = json.decode(AuthData.readJson());
    final model = UserModel.fromJson(decoded['data']);

    expect(
      model,
      UserModel(
        name: decoded['data']['name'],
        email: decoded['data']['email'],
        id: decoded['data']['id'],
        password: '',
        v: 0,
      ),
    );
  });

  test('Should return approproate seller info', () {
    /// action
    final jsonDecoded = json.decode(TestingDatas.getSingleProduct());
    final result = UserModel.fromSellerJson(jsonDecoded['data']['seller']);

    expect(
      result,
      UserModel(
          name: jsonDecoded['data']['seller']['name'],
          email: jsonDecoded['data']['seller']['email'],
          id: jsonDecoded['data']['seller']['_id'],
          v: jsonDecoded['data']['seller']['__v'],
          password: ''),
    );
  });

  test('Should return approriate model from entity', () {
    final result = UserModel.fromEntity(AuthData.userEntity);
    expect(result, AuthData.userModel);
  });
}
