import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/product.dart';

abstract class ProductRepository {

  Future<Either<Failure,Product>> getCurrentProduct(String id);
  Future<Either<Failure,List<Product>>> getAllProducts();
  Future<Either<Failure,Product>> createProduct(Product product);
  Future<Either<Failure,Product>> updateProduct(Product product);
  Future<Either<Failure,void>> deleteProduct(String id);
}