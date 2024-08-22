import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/usecase.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

// 

class UpdateProductUseCase implements UseCase<Future<Either<Failure, Product>>, Product> {
  final ProductRepository repository;

  UpdateProductUseCase(this.repository);

  @override
  Future<Either<Failure, Product>> call(Product product) async {
    return await repository.updateProduct(product, product.id);
  }
}
