import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';

import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repositories/product_repository.dart';

class GetProductByIdUsecase implements UseCase<Product, GetParams> {
  final ProductRepositories productRepository;

  GetProductByIdUsecase({required this.productRepository});

  @override
  Future<Either<Failure, Product>> call(GetParams params) async {
    return await productRepository.getProduct(params.id);
  }
}

class GetParams extends Equatable {
  final String id;

  const GetParams({required this.id});
  
  @override
  List<Object?> get props => [id];
}