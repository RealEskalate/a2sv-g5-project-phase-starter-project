import 'package:ecommerce/core/import/import_file.dart';

class UpdateProductUsecase {
  ProductRepository productRepository;
  UpdateProductUsecase({required this.productRepository});
  Future<Either<Failure, ProductEntity>> execute(ProductEntity product) {
    return productRepository.updateProduct(product);
  }
  
}
