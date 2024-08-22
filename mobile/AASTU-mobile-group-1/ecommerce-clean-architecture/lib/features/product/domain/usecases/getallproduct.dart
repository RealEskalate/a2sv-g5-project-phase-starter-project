import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class GetAllProductUsecase{
  GetAllProductUsecase(this.productrepository);
  final ProductRepository productrepository;
  Future<Either<Failure,List<Productentity>>> getall(){
    return productrepository.getAllProducts();
  }
}


