import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repository/product_repository.dart';

class GetProductUsecase implements UseCase<ProductEntity, GetParams> {
  final ProductRepository repository;

  GetProductUsecase(this.repository);

  @override
  Future<Either<Failure, ProductEntity>> call(GetParams params) async {
    return await repository.getProduct(params.id);
  }
}

class GetParams extends Equatable {
  final String id;

  const GetParams({required this.id});
  @override
  List<Object?> get props => [id];
}
