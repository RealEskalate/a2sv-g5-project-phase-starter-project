import '../../models/product_model.dart';

abstract class ProductLocalDataSource {
  Future<bool> cacheProducts(List<ProductModel> products);
  Future<List<ProductModel>> getProducts();
  
}