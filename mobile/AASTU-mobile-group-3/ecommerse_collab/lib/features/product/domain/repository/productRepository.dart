

import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/product.dart';

abstract class ProductRepository{
  Future<Either<Failure, List<Product>>> getAllProduct();
  Future<Either<Failure, void>> deleteProduct(String id);
  Future<Either<Failure, Product>> updateProduct({required String id, required String name, required int price, required String description});
  Future<Either<Failure, void>> addProduct(Product product);
  Future <Either<Failure, Product>> getProduct(String id);
  
}


