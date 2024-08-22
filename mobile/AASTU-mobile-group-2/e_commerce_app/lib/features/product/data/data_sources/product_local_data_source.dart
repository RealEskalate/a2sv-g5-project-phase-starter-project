import 'dart:convert';

import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:equatable/equatable.dart';
import 'package:shared_preferences/shared_preferences.dart';

abstract class ProductLocalDatasource {
  Future<List<ProductModel>> getAllProduct();
  Future<ProductModel> getOneProduct(String id);
  Future<void> cacheProduct(ProductModel product);
  Future<void> cachecProducts(List<ProductModel> products);
}

class ProductLocalDataSourceImpl implements ProductLocalDatasource {
  final productCacheKey = "PRODUCTS";

  final SharedPreferences sharedPreferences;
  ProductLocalDataSourceImpl({required this.sharedPreferences});

  String getProductCacheKey(String id) => '${productCacheKey}_$id';

  @override
  Future<void> cacheProduct(ProductModel product) async {
    await sharedPreferences.setString(
        getProductCacheKey(product.id), jsonEncode(product.toJson()));
  }

  @override
  Future<void> cachecProducts(List<ProductModel> products) async {
    await sharedPreferences.setString(productCacheKey,
        jsonEncode(products.map((product) => product.toJson()).toList()));
  }

  @override
  Future<List<ProductModel>> getAllProduct() async {
    final cachedProducts = sharedPreferences.getString(productCacheKey);
    if (cachedProducts != null) {
      return (ProductModel.fromJsonList(jsonDecode(cachedProducts)));
    } else {
      throw Exception("Products not found");
    }
  }

  @override
  Future<ProductModel> getOneProduct(String id) async {
    final cachedProduct = sharedPreferences.getString(getProductCacheKey(id));
    if (cachedProduct != null) {
      return ProductModel.fromJson(jsonDecode(cachedProduct));
    } else {
      throw Exception("product with that id doesnt exist");
    }
  }
}
