// ignore_for_file: avoid_renaming_method_parameters

import 'package:dartz/dartz.dart';
import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../entitity/product.dart';
import '../repositories/product_repository.dart';

class InsertProductUsecase extends UseCase<Product, Product> {
  final ProductRepository repository;

  InsertProductUsecase(this.repository);

  @override
  Future<Either<Failure, Product>> execute(Product product) async {
    return await repository.insertProduct(product);
  }
}
