class Urls {
  static const String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v1/products';
  static String getProductById(String id) {
    return '$baseUrl/$id';
  }

  static String getAllProducts() {
    return '$baseUrl';
  }
}
