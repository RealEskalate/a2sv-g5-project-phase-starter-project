import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/product.dart';

abstract class ProductRepository {
  /// Delete the row with the given [id]
  ///
  /// int to tell how many rows are affected
  Future<Either<Failure, int>> deleteProduct(String id);

  /// Replace the the given product with the given [product] based the the field id
  ///
  /// int to tell how many rows are affected
  Future<Either<Failure, int>> updateProduct(ProductEntity product);

  Future<Either<Failure, List<ProductEntity>>> getAllProducts();

  /// Returns the product of the given [id]
  Future<Either<Failure, ProductEntity>> getProduct(String id);

  /// [product] is inserted into the table
  ///
  /// int to tell to tell the id
  Future<Either<Failure, int>> insertProduct(ProductEntity product);
}