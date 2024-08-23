import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../../core/exception/exception.dart';
import '../../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<List<ProductModel>> getProducts();
  Future<void> cacheProducts(List<ProductModel> productToCache);
}

// ignore: constant_identifier_names
const CACHED_PRODUCTS = 'CACHED_PRODUCTS';

class ProductLocalDataSourceImpl extends ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  ProductLocalDataSourceImpl({required this.sharedPreferences});
  @override
  Future<void> cacheProducts(List<ProductModel> productToCache) {
    try {
      final jsonProduct = json.encode(ProductModel.toJsonList(productToCache));
      return sharedPreferences.setString(CACHED_PRODUCTS, jsonProduct);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<List<ProductModel>> getProducts() {
    try {
      final productsString = sharedPreferences.getString(CACHED_PRODUCTS);
      if (productsString != null) {
        final decodedJson = json.decode(productsString);
        return Future.value(ProductModel.fromJsonList(decodedJson));
      } else {
        throw CacheException();
      }
    } catch (e) {
      throw CacheException();
    }
  }
}
