import '../../models/product_model.dart';

abstract class ProductRemoteDataSource {
  Future<ProductModel> getProduct(String id);
  Future<List<ProductModel>> getProducts();
  Future<bool> deleteProduct(String id);
  Future<ProductModel> updateProduct(ProductModel product);
  Future<ProductModel> addProduct(ProductModel product);
}