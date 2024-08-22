import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../../data/models/product_model.dart';
import '../repository/product_repository.dart';

class AddProductParams {
  final ProductModel product;
  final String imagePath;

  AddProductParams(this.product, this.imagePath);
}

class AddProduct extends UseCase<ProductModel, AddProductParams> {
  final ProductRepository repository;

  AddProduct(this.repository);

  @override
  Future<Either<Failure, ProductModel>> call(AddProductParams params) async {
    return repository.addProduct(params.product, params.imagePath);
  }
}
