import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repository/product_repository.dart';

class UpdateProductUsecase implements UseCase<ProductEntity, UpdateParams> {
  final ProductRepository repository;

  UpdateProductUsecase(this.repository);

  @override
  Future<Either<Failure, ProductEntity>> call(UpdateParams params) async {
   
    return await repository.updateProduct(params.product);
  }
}

class UpdateParams extends Equatable {
  final ProductEntity product;

  const UpdateParams({required this.product});

  @override
  List<Object?> get props => [product];
}
