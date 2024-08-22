import 'package:dartz/dartz.dart';
import '../../../../core/failure/failure.dart';
import '../../../../core/use_cases/use_case.dart';
import '../repository/product_repository.dart';
// Define a class to represent the parameters needed for deletion
class DeleteProductParams {
  final String id;

  DeleteProductParams(this.id);
}

class DeleteProduct
    extends UseCase<void, DeleteProductParams> {
  ProductRepository repository;

  DeleteProduct(this.repository);

  @override
  Future<Either<Failure, void>> call(DeleteProductParams params) async {
    return await repository.deleteProduct(params.id);
  }
}
