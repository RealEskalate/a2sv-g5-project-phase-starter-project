
class Urls {
  static const String baseUrl = 'https://g5-flutter-learning-path-be.onrender.com/api/v1/products';
  static String getProductId(String id) => '$baseUrl/$id';
  static String deleteProductId(String id) => '$baseUrl/$id';
  static const String getProducts = baseUrl;
  static const String addProduct = baseUrl;
  static String updateProductId(String id) => '$baseUrl/$id';
}
class Urls2 {
  static const String baseUrl = 'https://g5-flutter-learning-path-be.onrender.com/api/v2';
  static String getCurrentUser() => '$baseUrl/users/me';
  static String login() => '$baseUrl/auth/login';
  static String signUp() => '$baseUrl/auth/register';
  static String getProductId(String id) => '$baseUrl/products/$id';
  static String deleteProductId(String id) => '$baseUrl/products/$id';
  static const String getProducts = '$baseUrl/products';
  static const String addProduct = '$baseUrl/products';
  static String updateProductId(String id) => '$baseUrl/products/$id';
}
