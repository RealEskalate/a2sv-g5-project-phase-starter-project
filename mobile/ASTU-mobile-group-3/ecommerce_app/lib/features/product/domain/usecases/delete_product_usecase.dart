import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../repositories/product_repository.dart';

class DeleteProductUseCase {
  final ProductRepository deleteProductRepository;
  DeleteProductUseCase(this.deleteProductRepository);

  Future<Either<Failure, int>> execute(String id) async {
    return await deleteProductRepository.deleteProduct(id);
  }
}