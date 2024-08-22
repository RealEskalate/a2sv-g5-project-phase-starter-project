import 'package:dartz/dartz.dart';
import '../../../../core/errors/failure/failures.dart';
import '../../../../core/usecases/usecases.dart';
import '../repositories/product_repository.dart';
import 'package:equatable/equatable.dart';
import '../entities/product_entity.dart';

class DeleteUsecase implements UseCase<void, String> {
  final ProductRepository abstractProductRepository;

  DeleteUsecase(this.abstractProductRepository);

  @override
  Future<Either<Failure, String>> call(String id) async {
    return await abstractProductRepository.deleteProduct(id);
  }
}