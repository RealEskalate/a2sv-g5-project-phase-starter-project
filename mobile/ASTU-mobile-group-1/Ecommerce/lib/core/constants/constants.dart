class Urls {
  static String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2';

  static String getProductById(String id) => '$baseUrl/products/$id';
  static String getAllProducts() => '$baseUrl/products';

  static String getUser() => '$baseUrl/users/me';
  static String signIn() => '$baseUrl/auth/login';
  static String signUp() => '$baseUrl/auth/register';

  static String baseChat =
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
  static String getChatById(String chatId) => '$baseChat/$chatId';
  static String getChatMessages(String chatId) => '$baseChat/$chatId/messages';
}

class Urls {
  static String baseUrl =
      'https://g5-flutter-learning-path-be.onrender.com/api/v3';

  static String getProductById(String id) => '$baseUrl/products/$id';
  static String getAllProducts() => '$baseUrl/products';

  static String getUser() => '$baseUrl/users/me';
  static String signIn() => '$baseUrl/auth/login';
  static String signUp() => '$baseUrl/auth/register';

  static String baseChat =
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
  static String getChatById(String chatId) => '$baseChat/$chatId';
  static String getChatMessages(String chatId) => '$baseChat/$chatId/messages';
}
