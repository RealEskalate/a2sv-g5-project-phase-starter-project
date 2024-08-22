import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';

import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class InsertProductUsecase implements UseCase<Product, CreateParams> {
  final ProductRepositories productRepository;

  InsertProductUsecase({required this.productRepository});

  @override
  Future<Either<Failure, Product>> call(CreateParams params) async {
    return await productRepository.createProduct(params.product);
  }
}
class CreateParams extends Equatable {
  final Product product;

  const CreateParams({required this.product});

  @override
  List<Object?> get props => [product];
}