import 'dart:convert';

import 'package:shared_preferences/shared_preferences.dart';

import '../../../../../core/error/exception.dart';
import '../../models/product_model.dart';
import 'local_data_source.dart';

class ProductLocalDataSourceImpl extends ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  // ignore: non_constant_identifier_names
  final CACHE_PRODUCTS_KEY = 'PRODUCTS';
  ProductLocalDataSourceImpl({required this.sharedPreferences});

  
  @override
  Future<bool> cacheProducts(List<ProductModel> products) {

    final String jsonString = jsonEncode(products); // JsonEncode converts the proudcts to json if they have .toJson() method implicitliy
    return sharedPreferences.setString(CACHE_PRODUCTS_KEY, jsonString);
  }

  @override
  Future<List<ProductModel>> getProducts() async {
    final jsonString = sharedPreferences.getString(CACHE_PRODUCTS_KEY);
    if (jsonString != null) {
      final List<dynamic> jsonDecoded = jsonDecode(jsonString);
      final products = jsonDecoded.map((product) {
        return ProductModel.fromJson(product);
      }).toList();
      return products;
    } else {
      throw CacheException();
    }
  }
}
