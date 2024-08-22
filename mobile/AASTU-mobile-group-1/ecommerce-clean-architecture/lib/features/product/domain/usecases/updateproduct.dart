import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class UpdateProductUsecase{
  UpdateProductUsecase(this.productrepository);
  final ProductRepository productrepository;
  Future<Either<Failure, Productentity>> update(Productentity newproduct){
    return productrepository.updateProduct(newproduct);
  }
}

