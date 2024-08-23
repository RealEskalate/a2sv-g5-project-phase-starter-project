import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';


class GetProductUseCase {
  final ProductRepository getProductRepository;
  GetProductUseCase(this.getProductRepository);

  Future<Either<Failure, ProductEntity>> execute(String id) async {
    return await getProductRepository.getProduct(id);
  }
}