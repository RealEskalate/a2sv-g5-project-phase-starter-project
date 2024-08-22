import 'dart:convert';


import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/product/data/models/product_model.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  const testProductModel = ProductModel(
      id: '1',
      name: 'Product 1',
      description: 'Product 1 description',
      imageUrl: 'product1.jpg',
      price: 100);

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

  test('should return a json map containing proper data', () async {
    //arrange
    final Map<String, dynamic> jsonMap = json
        .decode(readJson('helpers/fixtures/dummy_product_response.json'))[0];

    //act
    final result = testProductModel.toJson();

    //assert
    expect(result, jsonMap);
  });

  test('should throw an exception if fails to cast json to model', () async {
    //arrange
    final Map<String, dynamic> jsonMap = json
        .decode(readJson('helpers/fixtures/dummy_product_response.json'))[1];

    //act and assert
    expect(() => ProductModel.fromJson(jsonMap),
        throwsA(isA<JsonParsingException>()));
  });




test('should convert from json list to model list', () async {
  //arrange
  final List<dynamic> jsonList = json.decode(readJson('helpers/fixtures/dummy_products_cached.json'));
  
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
});}