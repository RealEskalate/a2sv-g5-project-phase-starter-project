import 'package:dartz/dartz.dart';

import '../../../../core/errors/failure.dart';
import '../../../../core/usecases/no_param_use_cases.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class ViewAllProductsUseCase
    implements NoParamsUseCase<Future<Either<Failure, List<Product>>>> {
  final ProductRepository repository;

  ViewAllProductsUseCase(this.repository);

  @override
  Future<Either<Failure, List<Product>>> call() async {
    return await repository.getProducts();
  }
}
