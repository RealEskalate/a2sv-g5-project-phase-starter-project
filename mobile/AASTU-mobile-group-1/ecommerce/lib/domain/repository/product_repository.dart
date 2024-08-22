import 'package:ecommerce/core/import/import_file.dart';

abstract class ProductRepository {
  Stream<List<ProductEntity>> getAllProducts();
  Future<Either<Failure, ProductEntity>> getProduct(id);
  Future<Either<Failure, ProductEntity>> addProduct(product);
  Future<Either<Failure, ProductEntity>> updateProduct(product);
  Future<Either<Failure, void>> deleteProduct(productId);
}
