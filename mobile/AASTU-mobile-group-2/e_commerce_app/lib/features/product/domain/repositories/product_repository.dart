import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';

abstract class ProductRepository {
  Future<Either<Failure, List<ProductEntity>>> getAllProduct();
  Future<Either<Failure, ProductEntity>> getOneProduct(String id);
  Future<Either<Failure,ProductEntity>> insertProduct(ProductEntity newProduct);
  Future<Either<Failure,String>> deleteProduct(String id);
  Future<Either<Failure,ProductEntity>> updateProduct(String id, ProductEntity updatedProduct);
}
