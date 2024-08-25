import 'package:dartz/dartz.dart';

import '../../../../core/errors/failures/failure.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';


class GetAllProductUseCase {
  final ProductRepository getAllProductRepository;
  GetAllProductUseCase(this.getAllProductRepository);

  Future<Either<Failure, List<ProductEntity>>> execute() async {
    return await getAllProductRepository.getAllProducts();
  }
}