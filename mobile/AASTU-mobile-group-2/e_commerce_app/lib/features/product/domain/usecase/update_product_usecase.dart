import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';

class UpdateProduct {
  
  ProductRepository productRepository;
  UpdateProduct(this.productRepository);

  Future<Either<Failure, void>> execute(ProductEntity updatedProduct) {
    return productRepository.updateProduct(updatedProduct.id, updatedProduct);
  }
}
