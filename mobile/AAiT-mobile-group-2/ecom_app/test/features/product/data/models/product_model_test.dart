import 'dart:convert';

import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/product/data/models/product_model.dart';
import 'package:ecom_app/features/product/data/models/seller_model.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  const testProductModel = ProductModel(
    id: '66c775b9322c0bf78ca69c59',
    name: 'PC',
    description: 'long description',
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1724347832/images/niqoqvv6rvwuatzuikkr.png',
    price: 123,
    seller: SellerModel(
      id: '66bde36e9bbe07fc39034cdd',
      name: 'Mr. User',
      email: 'user@gmail.com',
    ),
  );

  test('should be a subclass of product entity', () async {
    //assert
    expect(testProductModel, isA<Product>());
  });

  test('should return a valid model from json', () async {
    //arrange
    final Map<String, dynamic> jsonMap = json
        .decode(readJson('helpers/fixtures/dummy_product_response.json'))[0];

    //act
    final result = ProductModel.fromJson(jsonMap);
    //assert
    expect(result, equals(testProductModel));
  });

  test('should throw an exception if fails to cast json to model', () async {
    // Arrange: Create invalid JSON data to trigger an exception
    final Map<String, dynamic> invalidJsonMap = {
      'id': null, // Should be a String
      'name': 'PC',
      'description': 'long description',
      'price': 'invalid', // Should be a num
      'imageUrl': 'https://example.com/image.png',
      'seller': {
        '_id': '66bde36e9bbe07fc39034cdd',
        'name': 'Mr. User',
        'email': 'user@gmail.com',
        '__v': 0
      }
    };

    // Act and Assert: Expecting the JsonParsingException to be thrown
    expect(() => ProductModel.fromJson(invalidJsonMap),
        throwsA(isA<JsonParsingException>()));
  });
  test('should convert from json list to model list', () async {
    //arrange
    final List<dynamic> jsonList =
        json.decode(readJson('helpers/fixtures/dummy_products_cached.json'));

    //act
    final result = ProductModel.fromJsonList(jsonList);

    //assert
    expect(result, isA<List<ProductModel>>());
  });

  test('should convert from model list to json list', () async {
    //arrange
    final List<ProductModel> productList = [testProductModel];

    //act
    final result = ProductModel.toJsonList(productList);

    //assert
    expect(result, isA<List<Map<String, dynamic>>>());
  });

  test('should convert from model to entity', () async {
    //arrange
    final ProductModel productModel = testProductModel;

    //act
    final result = productModel.toEntity();

    //assert
    expect(result, isA<Product>());
  });

  test('should convert from model list to entity list', () async {
    //arrange
    final List<ProductModel> productModelList = [testProductModel];

    //act
    final result = ProductModel.toEntityList(productModelList);

    //assert
    expect(result, isA<List<Product>>());
  });

  test('should convert from entity to model', () async {
    //arrange
    final Product product = testProductModel.toEntity();

    //act
    final result = ProductModel.toModel(product);

    //assert
    expect(result, isA<ProductModel>());
  });
}
