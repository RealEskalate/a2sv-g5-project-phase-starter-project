import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

class InsertProductUseCase {
  final ProductRepository inserProductRepository;
  InsertProductUseCase(this.inserProductRepository);

  Future<Either<Failure, int>> execute(ProductEntity product) async {
    return await inserProductRepository.insertProduct(product);
  }
}