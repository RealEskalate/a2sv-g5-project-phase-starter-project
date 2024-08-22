import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';

class DeleteProduct {
  ProductRepository productRepository;
  
  DeleteProduct(this.productRepository);

  Future<Either<Failure, void>> execute(ProductEntity product) {
    
    return productRepository.deleteProduct(product.id);
  }
}
