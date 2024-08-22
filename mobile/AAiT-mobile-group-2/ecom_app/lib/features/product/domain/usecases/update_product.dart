import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

class UpdateProductUsecase implements UseCase<Product, UpdateParams> {
  final ProductRepository productRepository;
  UpdateProductUsecase(this.productRepository);

  @override
  Future<Either<Failure,Product>> call(UpdateParams params){
    return productRepository.updateProduct(params.product);
  }
}

class UpdateParams extends Equatable {
  final Product product;

  UpdateParams({required this.product});

  @override
  List<Object?> get props => [product];
}
