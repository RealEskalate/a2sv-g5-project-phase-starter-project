import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import '../../../../core/errors/failure.dart';
import '../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<List<ProductModel>> getProducts();
  Future<ProductModel> getProductById(String id);
  Future<bool> createProduct(ProductModel product);
  Future<bool> updateProduct(ProductModel product);
  Future<bool> deleteProduct(String id);
}

const String productDataKey = 'local_product_data';

class ProductLocalDataSourceImpl implements ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  const ProductLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<List<ProductModel>> getProducts() async {
    final jsonString = sharedPreferences.getString(productDataKey);
    
    if (jsonString != null) {
      final List<dynamic> result = json.decode(jsonString);
      return Future.value(result.map((json) => ProductModel.fromJson(json)).toList());
    } else {
      throw const CacheException('No products found');
    }
  }

  @override
  Future<ProductModel> getProductById(String id) async {
    final jsonString = sharedPreferences.getString(productDataKey);
    
    if (jsonString != null) {
      final List<dynamic> result = json.decode(jsonString);
      final products = result.map((json) => ProductModel.fromJson(json)).toList();
      final product = products.firstWhere((product) => product.id == id, orElse: () => throw const CacheException('Product not found'));
      return Future.value(product);
    } else {
      throw const CacheException('No products found');
    }
  }

  @override
  Future<bool> createProduct(ProductModel product) async {
    final jsonString = sharedPreferences.getString(productDataKey);
    List<ProductModel> products = [];
    
    if (jsonString != null) {
      final List<dynamic> result = json.decode(jsonString);
      products = result.map((json) => ProductModel.fromJson(json)).toList();
    }

    products.add(product);
    final updatedJsonString = json.encode(products.map((product) => product.toJson()).toList());
    return sharedPreferences.setString(productDataKey, updatedJsonString);
  }

  @override
  Future<bool> updateProduct(ProductModel product) async {
    final jsonString = sharedPreferences.getString(productDataKey);
    
    if (jsonString != null) {
      final List<dynamic> result = json.decode(jsonString);
      List<ProductModel> products = result.map((json) => ProductModel.fromJson(json)).toList();
      
      final index = products.indexWhere((p) => p.id == product.id);
      if (index != -1) {
        products[index] = product;
        final updatedJsonString = json.encode(products.map((product) => product.toJson()).toList());
        return sharedPreferences.setString(productDataKey, updatedJsonString);
      } else {
        throw const CacheException('Product not found');
      }
    } else {
      throw const CacheException('No products found');
    }
  }

  @override
  Future<bool> deleteProduct(String id) async {
    final jsonString = sharedPreferences.getString(productDataKey);
    
    if (jsonString != null) {
      final List<dynamic> result = json.decode(jsonString);
      List<ProductModel> products = result.map((json) => ProductModel.fromJson(json)).toList();
      
      products.removeWhere((product) => product.id == id);
      final updatedJsonString = json.encode(products.map((product) => product.toJson()).toList());
      return sharedPreferences.setString(productDataKey, updatedJsonString);
    } else {
      throw const CacheException('No products found');
    }
  }
}

