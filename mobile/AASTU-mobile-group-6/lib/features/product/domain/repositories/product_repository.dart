import 'package:dartz/dartz.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/models/user_model.dart';
import '../../data/models/product_models.dart';
import '../entities/product_entity.dart';
import '../../../../core/errors/failure/failures.dart';


abstract class ProductRepository {
  Future<Either<Failure, List<ProductModel>>> getProducts();
  Future<Either<Failure, ProductModel>> getProduct(String id);
  Future<Either<Failure, String>> addProduct(ProductEntity product);
  Future<Either<Failure, String>> updateProduct(ProductModel product);
  Future<Either<Failure, String>> deleteProduct(String id);
  Future<Either<Failure,UserModel>> getUserInfo();
} 