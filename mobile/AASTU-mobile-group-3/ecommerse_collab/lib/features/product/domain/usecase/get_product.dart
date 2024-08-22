import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/product.dart';
import '../repository/productRepository.dart';

class GetProductUseCase{
  final ProductRepository productRepository;
  GetProductUseCase(this.productRepository);

  Future <Either<Failure, Product>> call(String id){
    return productRepository.getProduct(id);
  }
}