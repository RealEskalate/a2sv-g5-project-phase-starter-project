import 'package:ecommerce_app/features/auth/data/model/user_model.dart';
import 'package:ecommerce_app/features/auth/domain/entities/user_entity.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';

void main() {
  test('User data model Should be subclass of user entity ', () {
    expect(AuthData.userEntity, isA<UserEntity>());
  });

  test('Should return model when strings are given', () {
    const result = UserModel(
        name: AuthData.name,
        email: AuthData.email,
        password: AuthData.password);

    expect(result, AuthData.userModel);
  });

  test('Should return an entity from model', () {
    final result = AuthData.userModel.toEntity();

    expect(result, AuthData.userEntity);
  });
}
