import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exception.dart';
import '../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<List<ProductModel>> getAllProducts();
  Future<void> cacheAllProducts(List<ProductModel> products);
}

const CACHED_PRODUCTS = 'CACHED_PRODUCTS';

class ProductLocalDataSourceImpl extends ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  ProductLocalDataSourceImpl({required this.sharedPreferences});
  @override
  Future<void> cacheAllProducts(List<ProductModel> productsToCache) {
    try {
      final jsonProduct = json.encode(ProductModel.toJsonList(productsToCache));

      return sharedPreferences.setString(CACHED_PRODUCTS, jsonProduct);
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<List<ProductModel>> getAllProducts() {
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
