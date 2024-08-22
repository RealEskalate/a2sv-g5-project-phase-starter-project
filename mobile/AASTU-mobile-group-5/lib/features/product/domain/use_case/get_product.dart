import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class GetProductParams {
  final String id;

  GetProductParams(this.id);
}

class GetProduct extends UseCase<Product, GetProductParams> {
  final ProductRepository repository;

  GetProduct(this.repository);

  @override
  Future<Either<Failure, Product>> call(GetProductParams params) async {
    return await repository.getProductById(params.id);
  }
}
