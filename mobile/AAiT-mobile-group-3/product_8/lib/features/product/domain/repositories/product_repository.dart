import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';

import '../entities/product_entity.dart';

abstract class ProductRepositories {
  Future<Either<Failure, List<Product>>> getProducts();
  Future<Either <Failure , Product>> getProduct(String id);
  Future<Either <Failure , Product>> createProduct(Product product);
  Future<Either <Failure , Product>> updateProduct(Product product);
  Future<Either<Failure , void>> deleteProduct(String id);
} 