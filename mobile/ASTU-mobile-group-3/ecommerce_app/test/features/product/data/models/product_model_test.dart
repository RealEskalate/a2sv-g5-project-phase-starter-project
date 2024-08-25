import 'dart:convert';

import 'package:ecommerce_app/features/product/data/models/product_model.dart';
import 'package:ecommerce_app/features/product/domain/entities/product.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  final data = TestingDatas.testDataModel;
  test('Product model should subclass of product Entity', () {
    /// assert
    expect(data, isA<ProductEntity>());
  });

  test('This test is to make sure the Product is implemented correctly', () {
    /// action
    ///
    final result = ProductModel.fromJson(
        json.decode(TestingDatas.getSingleProduct())['data']);

    /// assertion

    expect(result, TestingDatas.testDataModel);
  });

  test('This is test to test the toJson of product class', () {
    /// arrange
    final expectedJson = json.decode(TestingDatas.getSingleProduct());

    /// action
    final result = TestingDatas.testDataModel.toJson();

    /// assertion
    expect(result, expectedJson['data']);
  });

  group('Checking the conversion of model to entity and vice versa', () {
    test('Should return valid ProductEntity from Model', () {
      /// assertion

      expect(ProductModel.fromEntity(TestingDatas.testDataEntity),
          TestingDatas.testDataModel);
    });

    test('Should return valid return from model', () {
      /// asssert

      expect(data.toEntity(), TestingDatas.testDataEntity);
    });

    test('Should return list of entities from list of models', () {
      /// assert
      expect(ProductModel.allToEntity(TestingDatas.productModelList),
          TestingDatas.productEntityList);
    });
  });
}
