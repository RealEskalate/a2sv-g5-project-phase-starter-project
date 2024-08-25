import 'package:ecommerce_app/features/auth/data/model/signed_up_user_model.dart';
import 'package:ecommerce_app/features/auth/domain/entities/signed_up_user_entity.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  const SignedUpUserModel signedUpUserModel = SignedUpUserModel(
      id: '66bde36e9bbe07fc39034cdd',
      name: 'Mr. User',
      email: 'user@gmail.com');

  test('should be subclass of entity', () {
    expect(signedUpUserModel, isA<SignedUpUserEntity>());
  });

  test('Should extract data from json', () {
    final result = SignedUpUserModel.fromJson(const {
      'id': '66bde36e9bbe07fc39034cdd',
      'name': 'Mr. User',
      'email': 'user@gmail.com'
    });

    expect(result, signedUpUserModel);
  });
}
