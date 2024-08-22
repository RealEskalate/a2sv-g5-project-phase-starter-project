import 'dart:convert';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import '../../../domain/entities/product_entity.dart';

abstract class ProductLocalDataSource {

  Future <ProductModel> getProduct(String id);
  Future<String> addProduct(ProductModel product);
  Future<String> deleteProduct(String id);
  Future<String> updateProduct(ProductModel product);
  Future<List<ProductModel>> getProducts();
 
}

class ProductLocalDataSourceImpl extends ProductLocalDataSource {
  final SharedPreferences sharedPreferences;

  ProductLocalDataSourceImpl({required this.sharedPreferences});

  @override
  Future<ProductModel> getProduct(String id) async {
    final jsonString = sharedPreferences.getString('PRODUCT_$id');
    if (jsonString != null) {
      var ans = ProductModel.fromJson(jsonDecode(jsonString));// Error Place
      print(ans);
      return ans;
    } else {
      throw Exception('Product not found');
    }
  }// Needs to be fixed

  @override
  Future<String> addProduct(ProductModel product) async {
    final jsonString = jsonEncode(product.toJson());
    await sharedPreferences.setString('PRODUCT_${product.name}', jsonString);
    return 'Product added successfully';
  }

  @override
  Future<String> deleteProduct(String id) async {
    await sharedPreferences.remove('PRODUCT_$id');
    if (sharedPreferences.containsKey('PRODUCT_$id')) {
      throw Exception('Failed to delete product');
    } else {
      return 'Product deleted successfully';
    }
  }

  @override
  Future<String> updateProduct(ProductModel product) async {
    final jsonString = jsonEncode(product.toJson());
    await sharedPreferences.setString('PRODUCT_${product.name}', jsonString);
    print(jsonString);
    return 'Product Updated successfully';
  }// Needs to be fixed

  @override
  Future<List<ProductModel>> getProducts() async {
    final List<ProductModel> products = [];
    sharedPreferences.getKeys().forEach((key) {
      if (key.contains('PRODUCT_')) {
        final jsonString = sharedPreferences.getString(key);
        if (jsonString != null) {
          products.add(ProductModel.fromJson(jsonDecode(jsonString)));
        }
      }
    });
    return products;
  }// Needs to be fixed
}
















// import 'dart:async';
// import 'dart:core';

// import 'package:dartz/dartz.dart';
// import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
// import 'package:ecommerce_app_ca_tdd/features/product/domain/entities/product_entity.dart';
// import '../../../../../core/errors/exceptions/exceptions.dart';
// import '../../../../../core/errors/failure/failures.dart';
// import '../../../domain/repositories/product_repository.dart';
// import 'package:shared_preferences/shared_preferences.dart';
// import '../../models/product_models.dart';


// abstract class ProductLocalDataSource {

//   Future <ProductModel> getProduct(String id);
//   Future<String> addProduct(ProductModel product);
//   Future<String> deleteProduct(String id);
//   Future<String> updateProduct(ProductModel product);
//   Future<List<ProductEntity>> getProducts();
 
// }



// class ProductLocalDataSourceImpl extends ProductLocalDataSource {
//   final SharedPreferences sharedPreferences;

//   ProductLocalDataSourceImpl(this.sharedPreferences);
  


//   @override
//   Future<Either<Failure, String>> deleteProduct(String id) async {

//     // Shared Pref instanse
//     var sharedPref = await SharedPreferences.getInstance();

//     await sharedPref.remove(id);
//     if (sharedPref.containsKey(id)) {
//       return Left(LocalDataSourceFailure('Failed to delete product'));
//     } else {
//       return Right('Product Successfully Deleted');
//     }
//   }
//   @override
//   Future<Either<Failure, String>> addProduct(String name, String description, double price, String imagePath) async {
//     // Shared Pref instanse
//     var sharedPref = await SharedPreferences.getInstance();
//     // sharedPref.then((value) => value.setString(name, description));
//     await sharedPref.setString(name, description);
//     if (sharedPref.containsKey(name)) {
//       return Left(LocalDataSourceFailure('Failed to add product'));
//     }
//     return Right('Product Successfully Added');

//   }
//   @override
//   Future<Either<Failure, String>> updateProduct(String id, String name, String description, double price, String imagePath) async{
//     var sharedPref_up = await SharedPreferences.getInstance();

//     await sharedPref_up.setString(id, name);
//     await sharedPref_up.setDouble(id, price);
//     await sharedPref_up.setString(id, description);
//     await sharedPref_up.setString(id, imagePath);

//     if (sharedPref_up.containsKey(id) && sharedPref_up.containsKey(name) && sharedPref_up.containsKey(description) && sharedPref_up.containsKey(imagePath)) {
//       return Right('Product Successfully Updated');
//     } else {
//       return Left(LocalDataSourceFailure('Failed to update product'));
//     }
    
//   }
  
//   @override
//   Future<Either<Failure, ProductEntity>> getProduct(String id) async {
//     var sharedPref_get = await SharedPreferences.getInstance();
//       if (sharedPref_get.containsKey(id)){
//         return Right(sharedPref_get);
//       }else{
//         return Left(LocalDataSourceFailure("Failed to get product"));
//       }
//   }
  
//   @override
//   Future<Either<Failure, List<ProductEntity>>> getProducts() async {
//     var sharedPref_get = await SharedPreferences.getInstance();
//       if (sharedPref_get.containsKey(id)){
//         return Right(sharedPref_get.);
//       }else{
//         return Left(LocalDataSourceFailure("Failed to get product"));
//       }
//   }


// }
