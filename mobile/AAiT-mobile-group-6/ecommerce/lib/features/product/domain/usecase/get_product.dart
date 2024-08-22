// ignore_for_file: avoid_renaming_method_parameters

import 'package:dartz/dartz.dart';
import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../entitity/product.dart';
import '../repositories/product_repository.dart';

class GetProductUsecase extends UseCase<Product, String> {
  final ProductRepository repository;

  GetProductUsecase(this.repository);

  @override
  Future<Either<Failure, Product>> execute(String id) async {
    return await repository.getProduct(id);
  }
}
