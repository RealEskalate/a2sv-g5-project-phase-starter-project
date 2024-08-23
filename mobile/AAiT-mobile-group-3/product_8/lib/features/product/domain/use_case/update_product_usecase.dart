import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';

import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class UpdateProductUsecase implements UseCase<Product, UpdateParams> { 
  final ProductRepositories productRepository;

  UpdateProductUsecase({required this.productRepository});

  @override
  Future<Either<Failure , Product>> call(UpdateParams params) async {
    return await productRepository.updateProduct(params.product);
  }
}
class UpdateParams extends Equatable {
  final Product product;

  const UpdateParams({required this.product});
  
  @override
  List<Object?> get props => [product];
}