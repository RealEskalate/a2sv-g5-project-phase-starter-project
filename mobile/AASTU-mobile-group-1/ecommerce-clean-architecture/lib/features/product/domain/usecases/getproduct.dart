import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class GetProductUsecase{
  GetProductUsecase(this.productrepository);
  final ProductRepository productrepository;
  Future<Either<Failure,Productentity>> getprod(String id){
    return productrepository.getProduct(id);
  }
}