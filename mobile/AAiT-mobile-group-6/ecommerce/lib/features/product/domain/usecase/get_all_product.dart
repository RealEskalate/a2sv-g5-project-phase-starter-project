import 'package:dartz/dartz.dart';
import '../../../../core/error/failures.dart';
import '../entitity/product.dart';
import '../repositories/product_repository.dart';
import '../../../../core/base_usecase.dart';

class NoParams {}

class GetAllProductUsecase extends UseCase<List<Product>, NoParams> {
  final ProductRepository repository;

  GetAllProductUsecase(this.repository);

  @override
  Future<Either<Failure, List<Product>>> execute(NoParams params) async {
    return await repository.getAllProduct();
  }
}
