import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exception.dart';
import '../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<List<ProductModel>> getCachedProducts();
  Future<bool> cacheProducts(List<ProductModel> productsToCache);
  Future<String> getToken();
}

class ProductLocalDataSourceImpl extends ProductLocalDataSource {
  final SharedPreferences prefs;

  ProductLocalDataSourceImpl({required this.prefs});
  static const cachedProductsKey = 'CACHED_PRODUCTS';

  @override
  Future<String> getToken() async {
    try {
      final token = prefs.getString('accessToken');
      if (token != null) {
        return token;
      }
      throw Exception();
    } catch (e) {
      throw CacheException();
    }
  }

  @override
  Future<bool> cacheProducts(List<ProductModel> productsToCache) async {
    try {
      final List<String> jsonProducts = productsToCache
          .map((product) => json.encode(product.toJson()))
          .toList();

      return await prefs.setStringList(cachedProductsKey, jsonProducts);
    } catch (_) {
      throw CacheException();
    }
  }

  @override
  Future<List<ProductModel>> getCachedProducts() {
    final jsonList = prefs.getStringList(cachedProductsKey);

    if (jsonList != null) {
      return Future.value(jsonList
          .map((jsonString) => ProductModel.fromJson(jsonDecode(jsonString)))
          .toList());
    }

    throw CacheException();
  }
}
