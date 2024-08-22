import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/product.dart';
import '../repository/productRepository.dart';

class GetAllProductUseCase{

  final ProductRepository productRepository;
  GetAllProductUseCase(this.productRepository);


  Future <Either<Failure, List<Product>>> execute(){
    return productRepository.getAllProduct();
  }
}