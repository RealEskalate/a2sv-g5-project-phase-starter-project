import 'package:dartz/dartz.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/core/usecases/usecase.dart';
import 'package:product_6/features/product/domain/repository/product_repository.dart';

class DeleteProductUseCase implements UseCase<void, String> {
  final ProductRepository repository;

  DeleteProductUseCase(this.repository);

  @override
  Future<Either<Failure, void>> call(String id) async {
    return await repository.deleteProduct(id);
  }
}
