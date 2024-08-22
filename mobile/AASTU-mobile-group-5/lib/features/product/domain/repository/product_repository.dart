import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../data/models/product_model.dart';
import '../entities/product.dart';

abstract class ProductRepository {
  Future<Either<Failure, List<Product>>> getAllProducts();
  Future<Either<Failure, Product>> getProductById(String id);
  Future<Either<Failure, ProductModel>> addProduct(ProductModel product, String imagePath);
  Future<Either<Failure, ProductModel>> updateProduct(ProductModel product);
  Future<Either<Failure, void>> deleteProduct(String id);
}