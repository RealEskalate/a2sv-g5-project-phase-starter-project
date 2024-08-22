import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

class CreateProductUsecase implements UseCase<Product, CreateParams> {
  final ProductRepository productRepository;
  CreateProductUsecase(this.productRepository);

  @override
  Future<Either<Failure, Product>> call(CreateParams params) {
    return productRepository.createProduct(params.product);
  }
}

class CreateParams extends Equatable {
  final Product product;

  CreateParams({required this.product});
  @override
  List<Object?> get props => [product];
}
