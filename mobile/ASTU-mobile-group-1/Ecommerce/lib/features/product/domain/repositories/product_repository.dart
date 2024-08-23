import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/product_entity.dart';
import '../../../auth/domain/entities/user_entity.dart';

abstract class ProductRepository {
  Future<Either<Failure, List<ProductEntity>>> getAllProducts();
  Future<Either<Failure, List<ProductEntity>>> getProductsByCategory(
      String category);

  Future<Either<Failure, ProductEntity>> getProductById(String id);

  Future<Either<Failure, bool>> insertProduct(ProductEntity product);

  Future<Either<Failure, bool>> updateProduct(ProductEntity product);

  Future<Either<Failure, bool>> deleteProduct(String id);
}
