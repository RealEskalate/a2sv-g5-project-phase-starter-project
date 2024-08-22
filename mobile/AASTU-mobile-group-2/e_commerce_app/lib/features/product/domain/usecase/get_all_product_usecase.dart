import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';

class GetAllProductUsecase {
  ProductRepository productRepository;

  GetAllProductUsecase(this.productRepository);

  Future<Either<Failure, List<ProductEntity>>> execute() {
    return productRepository.getAllProduct();
  }
}
