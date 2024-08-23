import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class GetProductUsecase {
  final ProductRepository productRepository;

  GetProductUsecase({required this.productRepository});

  Future<Either<Failure, ProductEntity>> call(String id) async {
    return await productRepository.getProductById(id);
  }
}
