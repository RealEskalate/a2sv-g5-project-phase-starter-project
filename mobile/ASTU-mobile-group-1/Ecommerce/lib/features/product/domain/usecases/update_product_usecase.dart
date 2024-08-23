import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class UpdateProductUsecase {
  final ProductRepository productRepository;

  UpdateProductUsecase({required this.productRepository});

  Future<Either<Failure, bool>> call({
    required String id,
    required String name,
    required String description,
    required double price,
    required String imageUrl,
  }) async {
    final product = ProductEntity(
        id: id,
        name: name,
        description: description,
        price: price,
        imageUrl: imageUrl);

    final res = await productRepository.updateProduct(product);
    return res;
  }
}
