import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exceptions.dart';
import '../../domain/entitity/product.dart';
import '../model/product_model.dart';

abstract class ProductLocalDataSource {
  Future<Product> getLastProduct();

  Future<void> cacheProduct(ProductModel productToCache);
}

class ProductLocalDataSourceImpl implements ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  ProductLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<void> cacheProduct(Product product) async {
    final jsonString = json.encode(ProductModel.fromProduct(product).toJson());
    await sharedPreferences.setString('cachedProduct', jsonString);
  }

  @override
  Future<Product> getLastProduct() async {
    final jsonString = sharedPreferences.getString('cachedProduct');
    if (jsonString != null) {
      return ProductModel.fromJson(json.decode(jsonString));
    } else {
      throw CacheException();
    }
  }
}
