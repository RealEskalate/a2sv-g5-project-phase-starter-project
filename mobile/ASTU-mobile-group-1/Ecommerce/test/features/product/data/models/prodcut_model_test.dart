import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/product/data/models/product_model.dart';
import 'package:product_6/features/product/domain/entities/product_entity.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  const testProductModel = ProductModel(
      id: '1',
      name: 'shoe',
      description: 'best shoes',
      price: 12.0,
      imageUrl: 'http',
      seller: UserEntity.empty);

  test(
    'should a subclass of Product',
    () async {
      expect(testProductModel, isA<ProductEntity>());
    },
  );

  test('should return a valid model from json', () async {
    // arrange
    final Map<String, dynamic> jsonMap =
        json.decode(readJson('helpers/dummy_data/dummy_product_data.json'));

    // act

    final result = ProductModel.fromJson(jsonMap);

    // expect
    expect(result, equals(testProductModel));
  });

  test('should return valid json data', () async {
    final result = testProductModel.toJson();
    final expectedJsonMap = {
      'id': '1',
      'name': 'shoe',
      'description': 'best shoes',
      'price': 12.0,
      'imageUrl': 'http'
    };

    expect(result, equals(expectedJsonMap));
  });
}
