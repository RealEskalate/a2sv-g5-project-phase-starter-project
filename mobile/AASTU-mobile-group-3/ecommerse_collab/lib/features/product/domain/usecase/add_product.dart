import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/product.dart';
import '../repository/productRepository.dart';

class AddProductUseCase{
  final ProductRepository productRepository;

  AddProductUseCase(this.productRepository);
 
  Future<Either<Failure, void>> call(Product product){
    return productRepository.addProduct(product);
  }
}