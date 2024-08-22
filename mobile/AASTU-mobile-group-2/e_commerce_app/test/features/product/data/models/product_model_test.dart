import 'dart:convert';

import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:flutter/foundation.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  ProductModel testProductModel = ProductModel(
      description: "Explore anime characters.",
      id: "6672752cbd218790438efdb0",
      imageUrl:
          "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg",
      name: "Anime website",
      price: 123);

  test('success model test', () async {
    expect(testProductModel, isA<ProductEntity>());
  });

  test('success for valid from json', () async {
    final Map<String, dynamic> jsonMap = json.decode(
      readJson('helpers/dummy/dummy_product.json'),
    )["data"][0];

    final result = ProductModel.fromJson(jsonMap);

    expect(result, testProductModel);
  });
  test('success for valid to json', () async {
    final expectedJson = {
      "id": "6672752cbd218790438efdb0",
      "name": "Anime website",
      "description": "Explore anime characters.",
      "price": 123,
      "imageUrl":
          "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg"
    };

    final result = testProductModel.toJson();

    expect(result, expectedJson);
  });
}
