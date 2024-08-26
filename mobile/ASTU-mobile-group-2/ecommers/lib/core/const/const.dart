class Urls {
  static String getByUrl(String id) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/products/$id';
  static String getAll() =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/products';
  static String addNewProduct() =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/products';
  static String updateProduct(String id) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/products/$id';
  static String deleteProduct(String id) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/products/$id';
}

class LoginApi {
  static const String registerApi =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/register';
  static const String loginApi =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/auth/login';

  static const String findMe =
      'https://g5-flutter-learning-path-be.onrender.com/api/v2/users/me';
}

class ChatApi {
  static String deleteChatApi(String chatId) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats/$chatId';
  static String startChatApi() =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
  static String getMessagesApi(String userId) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats/$userId/messages';
  static String chatByIdApi(String chatId) =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats/$chatId';
  static String getChatsApi() =>
      'https://g5-flutter-learning-path-be.onrender.com/api/v3/chats';
}
