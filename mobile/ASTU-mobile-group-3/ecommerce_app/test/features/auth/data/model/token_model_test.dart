import 'package:ecommerce_app/features/auth/data/model/token_model.dart';
import 'package:ecommerce_app/features/auth/domain/entities/token_entity.dart';
import 'package:flutter_test/flutter_test.dart';

void main() {
  const TokenModel tokenModel = TokenModel(token: 'token');
  test('Should be sublcass of entity', () {
    expect(tokenModel, isA<TokenEntity>());
  });

  test('Should return valid model', () {
    final result = TokenModel.fromJson(const {'access_token': 'token'});
    expect(result, tokenModel);
  });
}
