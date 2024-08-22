import 'package:dartz/dartz.dart';

import '../../../../core/error/failure.dart';
import '../entity/product.dart';
import '../repository/productRepository.dart';

class UpdateProductUseCase {
  final ProductRepository productRepository;
  UpdateProductUseCase(this.productRepository);

  Future <Either<Failure, Product>> call({required String id, required String name, required int price, required String description}){
    return productRepository.updateProduct(id:id, name:name, price:price, description:description);
  }

}
