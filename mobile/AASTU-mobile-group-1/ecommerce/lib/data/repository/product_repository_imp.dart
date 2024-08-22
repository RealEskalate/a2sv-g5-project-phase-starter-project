import 'package:ecommerce/core/import/import_file.dart';


class ProductRepositoryImp implements ProductRepository {
  Api api;
  NetworkInfo networkInfo;
  LocalSource localSource;
  ProductRepositoryImp(
      {required this.api,
      required this.networkInfo,
      required this.localSource});
  @override
  Stream<List<ProductModel>> getAllProducts() async* {
    if (await networkInfo.isConnected) {
      try {
        // Fetching products from the API
        var products = await api.getAllProducts();

        // Assuming 'products' is of type List<ProductModel>
        yield* products;

        // Save the data locally for offline access
        await localSource.saveData(products);
      } catch (e) {
        print("Failed to get products from API: $e");

        // Fallback to locally saved products in case of failure
        var localProducts = await localSource.getSavedProducts();
        yield* localProducts;
      }
    } else {
      try {
        // When offline, load products from local storage
        var localProducts = localSource.getSavedProducts();
        yield* localProducts;
      } catch (e) {
        print("Failed to get products from local storage: $e");

        // Yielding an empty list in case of failure
        yield [];
      }
    }
  }

  Future<Either<Failure, ProductModel>> getProduct(id) async {
    if (await networkInfo.isConnected) {
      var res = await api.getProduct(id);
      return res;
    }

    return Left(Failure(message: "No Internet Connection"));
  }

  Future<Either<Failure, ProductModel>> addProduct(product) async {
    if (await networkInfo.isConnected) {
      var convert = ProductModel.fromEntity(product);
      var res = await api.addProduct(convert);
      
      return res;
    }
    return Left(Failure(message: "No Internet Connection"));
  }

  Future<Either<Failure, ProductModel>> updateProduct(product) async {
    if (await networkInfo.isConnected) {
      var convert = ProductModel.fromEntity(product);
      var res = await api.updateProduct(convert);
      return res;
    }
    return Left(Failure(message: "No Internet Connection"));
  }

  Future<Either<Failure, void>> deleteProduct(productId) async {
    if (await networkInfo.isConnected) {
      var res = api.deleteProduct(productId);
      return res;
    }
    return Left(Failure(message: "No Internet Connection"));
  }
}
