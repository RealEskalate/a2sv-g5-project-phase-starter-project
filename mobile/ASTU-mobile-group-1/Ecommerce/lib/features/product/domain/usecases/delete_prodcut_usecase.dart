import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repositories/product_repository.dart';

class DeleteProductUsecase {
  final ProductRepository productRepository;

  DeleteProductUsecase({required this.productRepository});

  Future<Either<Failure, bool>> call(String id) async {
    return await productRepository.deleteProduct(id);
  }
}
