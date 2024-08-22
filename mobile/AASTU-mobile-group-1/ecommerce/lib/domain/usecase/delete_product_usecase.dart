import 'package:ecommerce/core/import/import_file.dart';

class DeleteProductUsecase {
  ProductRepository productRepository;
  DeleteProductUsecase({required this.productRepository});
  Future<Either<Failure, void>> execute(String productId) {
    return productRepository.deleteProduct(productId);
  }

}
