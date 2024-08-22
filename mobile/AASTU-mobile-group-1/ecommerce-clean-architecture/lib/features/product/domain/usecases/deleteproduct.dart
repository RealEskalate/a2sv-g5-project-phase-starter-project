import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class DeleteProductUsecase{
  DeleteProductUsecase(this.productrepository);
  final ProductRepository productrepository;
  Future<Either<Failure,bool>> delete(String id){
    return productrepository.deleteProduct(id);
  }
}