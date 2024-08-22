// ignore_for_file: avoid_renaming_method_parameters

import 'package:dartz/dartz.dart';
import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../../../product/domain/entitity/product.dart';
import '../../../product/domain/repositories/product_repository.dart';

class DeleteProductUsecase extends UseCase<Product, String> {
  final ProductRepository repository;

  DeleteProductUsecase(this.repository);

  @override
  Future<Either<Failure, Product>> execute(String id) async {
    return await repository.deleteProduct(id);
  }
}
