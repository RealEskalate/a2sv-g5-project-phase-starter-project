import 'package:ecommerce/core/import/import_file.dart';

class GetProductUsecase {
  ProductRepository productRepository;
  GetProductUsecase({required this.productRepository});

  Future<Either<Failure, ProductEntity>> execute(String id) {
    return productRepository.getProduct(id);
  }
}
