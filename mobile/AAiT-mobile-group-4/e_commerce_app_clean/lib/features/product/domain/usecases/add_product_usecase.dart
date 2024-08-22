import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../entities/product_entity.dart';
import '../repository/product_repository.dart';

class AddProductUsecase implements UseCase<ProductEntity,CreateParams> {
  final ProductRepository productRepository;

  AddProductUsecase(this.productRepository);

  @override
  Future<Either<Failure, ProductEntity>> call(CreateParams params) async {
    return await productRepository.addProduct(params.product);
  }
}

class CreateParams extends Equatable {
  final ProductEntity product;

  const CreateParams({required this.product});
  
  @override
  List<Object?> get props => [product];
}
