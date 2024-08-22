import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';

class GetOneProduct {
  ProductRepository productRepository;
  String id;
  GetOneProduct(this.productRepository, this.id);

  Future<Either<Failure, ProductEntity>> execute() {
    return productRepository.getOneProduct(id);
  }
}
