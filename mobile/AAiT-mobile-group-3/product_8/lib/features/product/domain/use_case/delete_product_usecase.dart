import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/failure/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repositories/product_repository.dart';

class DeleteProductUsecase implements UseCase<void, DeleteParams> {
  final ProductRepositories productRepository;

  DeleteProductUsecase({required this.productRepository});

   @override
     Future<Either<Failure , void>> call(DeleteParams params) async {
    return await productRepository.deleteProduct(params.id);
  }
}
class DeleteParams extends Equatable {
  final String id;

  const DeleteParams({required this.id});
  
  @override
  List<Object?> get props => [id];
}