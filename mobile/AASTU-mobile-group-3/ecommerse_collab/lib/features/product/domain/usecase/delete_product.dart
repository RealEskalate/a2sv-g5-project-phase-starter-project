
import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../repository/productRepository.dart';

class DeleteProductUseCase{

  final ProductRepository productRepository;
  DeleteProductUseCase(this.productRepository);

  Future <Either<Failure, void>> call(String id){
    return productRepository.deleteProduct(id);
  }

}