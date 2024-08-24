class Urls {
  static const String baseUrl3 =
      'https://g5-flutter-learning-path-be.onrender.com/api/v3';
  static String currentProductById(String id) => '$baseUrl3/products/$id';
}

class ErrorMessages {
  static const String noInternet = 'Failed to connect to the internet';
  static const String somethingWentWrong = 'Something went wrong';
  static const String serverError = 'An error has occurred';
  static const String cacheError = 'Failed to load cache';
  static const String socketError =
      'No Internet connection or server unreachable';
  static const String forbiddenError = 'Invalid Credentials! Please try again';
  static const String userAlreadyExists = 'User Already Exists';
}
