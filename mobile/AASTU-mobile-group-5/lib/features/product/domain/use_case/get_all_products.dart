import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/no_para_use_case.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class GetAllProducts extends NoParamsUseCase<Future<Either<Failure, List<Product>>>> {
  final ProductRepository repository;

  GetAllProducts(this.repository);

  @override
  Future<Either<Failure, List<Product>>> call() async {
    return await repository.getAllProducts();
  }
}
