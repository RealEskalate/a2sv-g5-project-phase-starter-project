import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

class GetAllProductsUsecase implements UseCase<List<Product>, NoParams> {
  final ProductRepository productRepository;

  GetAllProductsUsecase(this.productRepository);

  @override
  Future<Either<Failure, List<Product>>> call(NoParams params) {
    return productRepository.getAllProducts();
  }
}
