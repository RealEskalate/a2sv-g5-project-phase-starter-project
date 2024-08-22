
import 'package:ecommerce/core/import/import_file.dart';

class AddProductUsecase {
  ProductRepository productRepository;
  AddProductUsecase({required this.productRepository});

  Future<Either<Failure, ProductEntity>> execute(ProductEntity product) {
    return productRepository.addProduct(product);
  }
}
