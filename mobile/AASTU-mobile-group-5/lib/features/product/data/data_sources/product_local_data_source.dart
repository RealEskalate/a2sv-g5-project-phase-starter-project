import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';

import '../../../../core/error/exceptions.dart';
import '../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<List<ProductModel>> getAllProducts();
  Future<ProductModel> getProductById(String id);
  Future<ProductModel> addProduct(ProductModel product);
  Future<ProductModel> updateProduct(ProductModel product);
  Future<void> deleteProduct(String id);
  Future<void> cacheProducts(List<ProductModel> products);
}

class ProductLocalDataSourceImpl implements ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  ProductLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<List<ProductModel>> getAllProducts() async {
    final jsonString = sharedPreferences.getString('CACHED_PRODUCTS');
    if (jsonString != null) {
      final List<dynamic> jsonResponse = json.decode(jsonString);
      return jsonResponse.map((json) => ProductModel.fromJson(json)).toList();
    } else {
      return throw CacheException();
    }
  }

  @override
  Future<ProductModel> getProductById(String id) async {
    final jsonString = sharedPreferences.getString('CACHED_PRODUCTS');
    if (jsonString != null) {
      final List<dynamic> jsonResponse = json.decode(jsonString);
      final productJson = jsonResponse.firstWhere((json) => json['id'] == id,
          orElse: () => null);
      if (productJson != null) {
        return ProductModel.fromJson(productJson);
      } else {
        throw CacheException();
      }
    } else {
      throw CacheException();
    }
  }

  @override
  Future<ProductModel> addProduct(ProductModel product) async {
    final jsonString = sharedPreferences.getString('CACHED_PRODUCTS');
    if (jsonString != null) {
      final List<dynamic> jsonResponse = json.decode(jsonString);
      jsonResponse.add(product.toJson());
      await sharedPreferences.setString(
          'CACHED_PRODUCTS', json.encode(jsonResponse));
    } else {
      await sharedPreferences.setString(
          'CACHED_PRODUCTS', json.encode([product.toJson()]));
    }
    return product;
  }

  @override
  Future<void> deleteProduct(String id) async {
    final jsonString = sharedPreferences.getString('CACHED_PRODUCTS');
    if (jsonString != null) {
      final List<dynamic> jsonResponse = json.decode(jsonString);
      jsonResponse.removeWhere((json) => json['id'] == id);
      await sharedPreferences.setString(
          'CACHED_PRODUCTS', json.encode(jsonResponse));
    } else {
      throw CacheException();
    }
  }

  @override
  Future<ProductModel> updateProduct(ProductModel product) async {
    final jsonString = sharedPreferences.getString('CACHED_PRODUCTS');
    if (jsonString != null) {
      final List<dynamic> jsonResponse = json.decode(jsonString);
      final index = jsonResponse.indexWhere((json) => json['id'] == product.id);
      if (index != -1) {
        jsonResponse[index] = product.toJson();
        await sharedPreferences.setString(
            'CACHED_PRODUCTS', json.encode(jsonResponse));
        return product;
      } else {
        throw CacheException();
      }
    } else {
      throw CacheException();
    }
  }

  @override
  Future<void> cacheProducts(List<ProductModel> products) async {
    final jsonString =
        json.encode(products.map((product) => product.toJson()).toList());
    await sharedPreferences.setString('CACHED_PRODUCTS', jsonString);
  }
}
