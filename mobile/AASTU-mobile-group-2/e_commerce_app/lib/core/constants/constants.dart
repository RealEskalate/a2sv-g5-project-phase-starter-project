class Urls {
  static const String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/products';
  static const String loginUrl =
      "https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/login";
  static const String signUpUrl =
      "https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/register";
  static getProductById(String id) => "$baseUrl/$id";
  static String currentUserUrl = "https://g5-flutter-learning-path-be.onrender.com/api/v2/users/me";
  static String socketUrl = "https://g5-flutter-learning-path-be.onrender.com/";
}
 