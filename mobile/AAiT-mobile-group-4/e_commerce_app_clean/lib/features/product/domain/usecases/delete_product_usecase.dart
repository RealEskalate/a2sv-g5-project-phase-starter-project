import 'package:dartz/dartz.dart';
import 'package:equatable/equatable.dart';

import '../../../../core/error/failure.dart';
import '../../../../core/usecase/usecase.dart';
import '../repository/product_repository.dart';

class DeleteProductUsecase implements UseCase<bool, DeleteParams> {
  final ProductRepository repository;

  DeleteProductUsecase(this.repository);

  @override
  Future<Either<Failure, bool>> call(DeleteParams params) async {
    return await repository.deleteProduct(params.id);
  }
}

class DeleteParams extends Equatable {
  final String id;

  const DeleteParams({required this.id});
  @override
  List<Object?> get props => [id];
}
