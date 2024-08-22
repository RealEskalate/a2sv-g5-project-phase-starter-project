import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../models/product_model.dart';

abstract class LocalDataSource {
  List<ProductModel> getAllProduct();
  Future<void> addProduct(ProductModel productModel);
  Future<void> cacheProducts(List<ProductModel> cachedProducts);
  Future<ProductModel> updateProduct(ProductModel productModel);
  Future<bool> deleteProduct(String id);
  ProductModel getProduct(String id);
}

class LocalDataSourceImpl implements LocalDataSource {
  final SharedPreferences sharedPreferences;
  var keyName = 'cache';
  LocalDataSourceImpl({required this.sharedPreferences});

  List<ProductModel> _toProductModelList(String? response) {
    List<ProductModel> productModels = [];
    if (response != null) {
      var jsonList = json.decode(response);
      for (var data in jsonList) {
        productModels.add(ProductModel.fromJson(data));
      }

      return productModels;
    }
    return [];
  }

  @override
  Future<void> addProduct(ProductModel productModel) {
    // TODO: implement addProduct
    throw UnimplementedError();
  }

  @override
  Future<void> cacheProducts(List<ProductModel> cachedProducts) async {
    var mapped = cachedProducts.map((model) => (model.toJson())).toList();
    var jsonMap = json.encode(mapped);
    bool response = await sharedPreferences.setString(keyName, jsonMap);

    if (response == false) {
      throw Exception('Error');
    }
  }

  @override
  Future<bool> deleteProduct(String id) async {
    var response = sharedPreferences.getString(keyName);

    if (response != null) {
      var productModels = _toProductModelList(response);

      int index = 0;
      for (var productModel in productModels) {
        if (productModel.id == id) {
          productModels.remove(productModels[index]);
          await cacheProducts(productModels);
          return true;
        }
        index += 1;
      }
      throw Exception("Can't find product with this id");
    } else {
      throw Exception('Problem with shared preference');
    }
  }

  @override
  List<ProductModel> getAllProduct() {
    var response = sharedPreferences.getString(keyName);
    if (response != null) {
      final productModels = _toProductModelList(response);
      return productModels;
    }
    throw Exception('Cache Missed');
  }

  @override
  ProductModel getProduct(String id) {
    var response = sharedPreferences.getString(keyName);
    if (response != null){
      final productModels = _toProductModelList(response);
      for (var productModel in productModels){
        if (productModel.id == id){
          return productModel;
        }
      }
      throw Exception('No Product Found with this ID');
    }
    throw Exception('Problem with shared preference');
  }

  @override
  Future<ProductModel> updateProduct(ProductModel newProductModel) async {
    var response = sharedPreferences.getString(keyName);
    if (response != null){
      int index = 0;
      final productModels = _toProductModelList(response);
      for (var productModel in productModels){
        if (productModel.id == newProductModel.id){
          productModels[index] = newProductModel;
          await cacheProducts(productModels);
          return productModels[index];
        }
        index += 1;
      }
      throw Exception('No Product Found with this ID');
    }
    throw Exception('Problem with shared preference');
  }
}
