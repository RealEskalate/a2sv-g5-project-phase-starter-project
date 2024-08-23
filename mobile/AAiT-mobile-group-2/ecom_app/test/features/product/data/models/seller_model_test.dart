import 'dart:convert';

import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/product/data/models/product_model.dart';
import 'package:ecom_app/features/product/data/models/seller_model.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/domain/entities/seller.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/json_reader.dart';

void main() {
  const testSellerModel =
      SellerModel(id: '1', name: 'John', email: 'john@email.com');

  test('should be a subclass of seller entity', () async {
    //assert
    expect(testSellerModel, isA<Seller>());
  });

  test('should return a valid model from json', () async {
    //arrange
    final Map<String, dynamic> jsonMap =
        json.decode(readJson('helpers/fixtures/dummy_seller_response.json'))[0];

    //act
    final result = SellerModel.fromJson(jsonMap);

    //assert
    expect(result, equals(testSellerModel));
  });

  test('should throw an exception if fails to cast json to model', () async {
    //arrange
    final Map<String, dynamic> jsonMap = json
        .decode(readJson('helpers/fixtures/dummy_product_response.json'))[0];

    //act and assert
    expect(() => SellerModel.fromJson(jsonMap),
        throwsA(isA<JsonParsingException>()));
  });

  test('should convert from json list to model list', () async {
    //arrange
    final List<dynamic> jsonList =
        json.decode(readJson('helpers/fixtures/dummy_seller_response.json'));

    //act
    final result = SellerModel.fromJsonList(jsonList);

    //assert
    expect(result, isA<List<SellerModel>>());
  });

  test('should convert from model list to json list', () async {
    //arrange
    final List<SellerModel> sellerList = [testSellerModel];

    //act
    final result = SellerModel.toJsonList(sellerList);

    //assert
    expect(result, isA<List<Map<String, dynamic>>>());
  });

  test('should convert from model to entity', () async {
    //arrange
    const SellerModel sellerModel = testSellerModel;

    //act
    final result = sellerModel.toEntity();

    //assert
    expect(result, isA<Seller>());
  });

  test('should convert from model list to entity list', () async {
    //arrange
    final List<SellerModel> sellerModelList = [testSellerModel];

    //act
    final result = SellerModel.toEntityList(sellerModelList);

    //assert
    expect(result, isA<List<Seller>>());
  });

  test('should convert from entity to model', () async {
    //arrange
    final Seller seller = testSellerModel.toEntity();

    //act
    final result = SellerModel.toModel(seller);

    //assert
    expect(result, isA<SellerModel>());
  });
}
