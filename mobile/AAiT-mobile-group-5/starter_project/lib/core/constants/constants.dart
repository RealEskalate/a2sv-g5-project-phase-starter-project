class Urls {
  static const String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v1/products';
  static String getID(String id) => '$baseUrl/$id';
}

class Messages {
  static const String noInternet = 'Failed to connect to the internet';
  static const String somethingWentWrong = 'Something went wrong';
  static const String serverError = 'An error occurred';
  static const String cacheError = 'Failed to load cache';
  static const String socketError =
      'No Internet connection or server unreachable';
  static const String productStatetErrorMessage = 'Error In Product State';
}
