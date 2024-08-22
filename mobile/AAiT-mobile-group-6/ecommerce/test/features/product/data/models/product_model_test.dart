// ignore_for_file: prefer_single_quotes

import 'dart:convert';

import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  const testProductModel = ProductModel(
    id: '1',
    name: 'Nike Air Max 270',
    price: 300,
    imageUrl: 'images/nike.jpg',
    description:
        'footwear option characterized by its open lacing system, where the shoelace eyelets are sewn on top of the vamp (the upper part of the shoe). This design feature provides a more relaxed and casual look compared to the closed lacing system of oxford shoes. Derby shoes are typically made of high-quality leather, known for its durability and elegance, making them suitable for both formal and casual occasions. With their timeless style and comfortable fit, derby leather shoes are a staple in any well-rounded wardrobe.',
  );

  test('should be a subclass of Product entity', () async {
    // Assert
    expect(testProductModel, isA<ProductModel>());
  });
  test('should return a valid model', () async {
    // Arrange
    final Map<String, dynamic> jsonMap =
        json.decode(readJson('dummy_product_response.json'));
    // Act

    final result = ProductModel.fromJson(jsonMap);
    // Assert
    expect(result, equals(testProductModel));
  });

  test('should retutn json map containing proper data', () async {
    // Act
    final result = testProductModel.toJson();
    // Assert
    final expectedMap = {
      "id": '1',
      "name": 'Nike Air Max 270',
      "price": 300.00,
      "imageUrl": "images/nike.jpg",
      "description":
          "footwear option characterized by its open lacing system, where the shoelace eyelets are sewn on top of the vamp (the upper part of the shoe). This design feature provides a more relaxed and casual look compared to the closed lacing system of oxford shoes. Derby shoes are typically made of high-quality leather, known for its durability and elegance, making them suitable for both formal and casual occasions. With their timeless style and comfortable fit, derby leather shoes are a staple in any well-rounded wardrobe.",
    };
    expect(result, expectedMap);
  });
}
