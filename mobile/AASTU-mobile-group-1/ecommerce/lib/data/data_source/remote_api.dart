import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/failure/failure.dart';
import 'package:ecommerce/data/model/product_model.dart';
import 'package:equatable/equatable.dart';

abstract class Api extends Equatable{
  Stream<List<ProductModel>> getAllProducts();
  Future<Either<Failure, ProductModel>> getProduct(String id);
  Future<Either<Failure, ProductModel>> addProduct(ProductModel product);
  Future<Either<Failure, ProductModel>> updateProduct(ProductModel product);
  Future<Either<Failure, void>> deleteProduct(String id);

}