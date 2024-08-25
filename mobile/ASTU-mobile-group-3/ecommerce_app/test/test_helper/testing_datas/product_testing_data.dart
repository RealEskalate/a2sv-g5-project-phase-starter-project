import 'dart:convert';
import 'dart:io';

import 'package:ecommerce_app/features/product/data/models/product_model.dart';
import 'package:ecommerce_app/features/product/domain/entities/product.dart';
import 'package:flutter/cupertino.dart';

import '../auth_test_data/testing_data.dart';

class TestingDatas {
  /// Data for testing with id
  static const String id = '66c778ec10de84211de27fe1';

  /// Testing data for  prodcut Entity
  static ProductEntity testDataEntity = ProductEntity(
      id: '66c778ec10de84211de27fe1',
      name: 'PC',
      description: 'long description',
      price: 123,
      imageUrl:
          'https://res.cloudinary.com/g5-mobile-track/image/upload/v1724348651/images/pug5nl6gv14np2bqdofu.png',
      seller: AuthData.sellerData.toEntity());

  /// Testing data for  product models
  static ProductModel testDataModel = ProductModel(
    id: '66c778ec10de84211de27fe1',
    name: 'PC',
    description: 'long description',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1724348651/images/pug5nl6gv14np2bqdofu.png',
    seller: AuthData.sellerData,
  );

  /// Testing data for list of ProductEntity
  static List<ProductEntity> productEntityList = [
    TestingDatas.testDataEntity,
  ];

  /// Testing data for product model
  static List<ProductModel> productModelList = [
    TestingDatas.testDataModel,
  ];

  /// shared preference testing list data

  static String sharedPrefTest = '{"6672752cbd218790438efdb0" : 1}';

  static String readJson() {
    String dir = Directory.current.path;
    if (dir.contains('/test')) {
      dir = dir.replaceAll('/test', '');
    }

    dir = '$dir/test/test_helper/testing_datas/single_api_response_data.json';

    return File(dir).readAsStringSync();
  }

  /// Below data's are data that are used as exact data from the api
  static String getAllProductResponce() {
    String dir = Directory.current.path;
    if (dir.contains('/test')) {
      dir = dir.replaceAll('/test', '');
    }

    dir = '$dir/test/test_helper/testing_datas/all_products_responce.json';

    return File(dir).readAsStringSync();
  }

  static String readProductV3() {
    String dir = Directory.current.path;
    if (dir.contains('/test')) {
      dir = dir.replaceAll('/test', '');
    }

    dir = '$dir/test/test_helper/testing_datas/all_products_responce_v3.json';

    return File(dir).readAsStringSync();
  }

  static String getSingleProduct() {
    String dir = Directory.current.path;
    if (dir.contains('/test')) {
      dir = dir.replaceAll('/test', '');
    }

    dir = '$dir/test/test_helper/testing_datas/single_api_responce.json';

    return File(dir).readAsStringSync();
  }

  static const String apiId = '6672776eb905525c145fe0bb';
}

void main() {
  debugPrint(json.decode(TestingDatas.getSingleProduct()).toString());
}
