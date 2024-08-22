import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repository/product_repository.dart';

class GetProductsUsecase implements UseCase<List<ProductEntity>,NoParams>{
  final ProductRepository repository;

  GetProductsUsecase(this.repository);
  
  @override
  Future<Either<Failure, List<ProductEntity>>> call(NoParams params) async {
    return await repository.getProducts();
  }
}
