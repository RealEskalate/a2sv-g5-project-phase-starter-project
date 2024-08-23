class Urls {
  static const String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/products';
  static const String autUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2';
  static const String baseUrl2 =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2';

  // static const String apiKey = 'api_key=1f54bd990f1cdfb230adb312546d765d';
  static String getProdutbyId(String id) => '$baseUrl/$id';
}
