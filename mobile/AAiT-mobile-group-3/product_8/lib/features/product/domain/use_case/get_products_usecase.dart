import 'package:dartz/dartz.dart';

import '../../../../core/failure/failure.dart';

import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class GetProductsUsecase implements UseCase<List<Product>, NoParams> {
  final ProductRepositories productRepository;

  GetProductsUsecase({required this.productRepository});

  @override
  Future<Either<Failure, List<Product>>> call(NoParams params) async {
    return await productRepository.getProducts();
  }
}