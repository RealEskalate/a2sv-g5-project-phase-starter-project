import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/usecase.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class ViewProductUseCase
    implements UseCase<Future<Either<Failure, Product>>, String> {
  final ProductRepository repository;

  ViewProductUseCase(this.repository);

  @override
  Future<Either<Failure, Product>> call(String id) async {
    return await repository.getProductById(id);
  }
}
