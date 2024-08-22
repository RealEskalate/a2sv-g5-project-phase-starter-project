import 'package:dartz/dartz.dart';
import '../../../../core/error/failure.dart';
import '../entities/product.dart';
import '../repository/product_repository.dart';

class AddProductUsecase{
  AddProductUsecase(this.productrepository);
  final ProductRepository productrepository;
  Future<Either<Failure, void>> add(Productentity newproduct){
    return productrepository.addProduct(newproduct);
  }
}

