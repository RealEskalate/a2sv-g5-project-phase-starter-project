import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product.dart';
import '../repositories/product_repository.dart';

// class GetCurrentProductUsecase {
//   final ProductRepository productRepository;

//   GetCurrentProductUsecase(this.productRepository);

//   Future<Either<Failure, Product>> execute(String productId) async {
//     return await productRepository.getCurrentProduct(productId);
//   }
// }


class GetCurrentProductUsecase implements UseCase<Product, GetParams> {
  final ProductRepository productRepository;

  GetCurrentProductUsecase(this.productRepository);

  @override
  Future<Either<Failure,Product>> call(GetParams params) async{
    return await productRepository.getCurrentProduct(params.id);
  }
}

class GetParams extends Equatable{
  final String id;

  const GetParams({required this.id});
  
  @override
  List<Object?> get props => [id];
}