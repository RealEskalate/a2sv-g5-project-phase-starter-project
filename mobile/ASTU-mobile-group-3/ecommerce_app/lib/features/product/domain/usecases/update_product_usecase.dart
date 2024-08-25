import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

class UpdateProductUsecase {
  final ProductRepository updateProductRepository;
  UpdateProductUsecase(this.updateProductRepository);

  Future<Either<Failure, int>> execute(ProductEntity product) async {
    return await updateProductRepository.updateProduct(product);
  }
}