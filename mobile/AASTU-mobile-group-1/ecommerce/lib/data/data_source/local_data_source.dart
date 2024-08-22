import 'package:ecommerce/core/import/import_file.dart';
class LocalDataSource extends LocalSource {
  final SharedPreferences sharedPreferences;

  LocalDataSource({required this.sharedPreferences});

  @override
  Stream<List<ProductModel>> getSavedProducts() async* {
  try {
    // Fetching the saved products from shared preferences
    String? productsJson = sharedPreferences.getString('productslist');
    
    // If no products are saved, yield an empty list
    if (productsJson == null) {
      yield [];
      return;
    }

    // Decoding the JSON string into a list of maps
    List<Map<String, dynamic>> productMap = List<Map<String, dynamic>>.from(jsonDecode(productsJson));
    
    // Mapping each map to a ProductModel object
    List<ProductModel> products = productMap.map((product) => ProductModel.fromjson(product)).toList();
    
    // Yielding the list of products
    yield products;
  } catch (e) {
    print("Failed to get saved products: $e");
    
    // If an error occurs, yield an empty list
    yield [];
  }
}


  @override
  Future<Either<Failure, void>> saveData(Stream<List<ProductModel>> productsStream) async {
  try {
    // Listen to the stream
    await for (var products in productsStream) {
      // Convert the list of ProductModel objects to a JSON string
      final List<Map<String, dynamic>> productsJson = products.map((product) => product.toJson()).toList();
      String jsonString = jsonEncode(productsJson);
      
      // Save the JSON string to SharedPreferences
      await sharedPreferences.setString('productslist', jsonString);
    }
    
    return const Right(null);
  } catch (e) {
    return Left(Failure(message: "Failed to cache data"));
  }
}

}
