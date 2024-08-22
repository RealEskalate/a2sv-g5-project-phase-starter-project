import 'package:dartz/dartz.dart';
import '../../../../core/base_usecase.dart';
import '../../../../core/error/failures.dart';
import '../../../product/domain/entitity/product.dart';
import '../../../product/domain/repositories/product_repository.dart';

class UpdateProductUsecase extends UseCase<Product, Product> {
  final ProductRepository repository;

  UpdateProductUsecase(this.repository);

  @override
  Future<Either<Failure, Product>> execute(Product product) async {
    return await repository.updateProduct(product);
  }
}
