import 'package:ecommerce/core/import/import_file.dart';

class GetAllProductUsecase {
  final ProductRepository productRepository;
  GetAllProductUsecase({required this.productRepository});

  Stream<List<ProductEntity>> execute() {
    return productRepository.getAllProducts();
  }
}
