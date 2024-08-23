import 'dart:convert';
import 'dart:io';

import 'package:ecommerce_app/features/product/data/models/product_model.dart';
import 'package:ecommerce_app/features/product/domain/entities/product.dart';
import 'package:flutter/cupertino.dart';

class TestingDatas {
  /// Data for testing with id
  static const String id = '6672752cbd218790438efdb0';

  /// Testing data for  prodcut Entity
  static const testDataEntity = ProductEntity(
    id: '6672752cbd218790438efdb0',
    name: 'Anime website',
    description: 'Explore anime characters.',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',
  );

  /// Testing data for  product models
  static const testDataModel = ProductModel(
    id: '6672752cbd218790438efdb0',
    name: 'Anime website',
    description: 'Explore anime characters.',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',
  );

  /// Testing data for list of ProductEntity
  static const List<ProductEntity> productEntityList = [
    TestingDatas.testDataEntity,
  ];

  /// Testing data for product model
  static const List<ProductModel> productModelList = [
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
