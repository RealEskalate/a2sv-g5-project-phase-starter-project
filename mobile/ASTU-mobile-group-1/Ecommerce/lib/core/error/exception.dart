class ServerException implements Exception {
  final String message;
  ServerException(
      {this.message = 'An unknown error occurred. Please try again later.'});
}

class CacheException implements Exception {
  // final String message;
  // CacheException({required this.message});
}

class WrongPasswordOrEmailException implements Exception {}

class InvalidEmailException implements Exception {}

class AlreadyExistEmailException implements Exception {}

class UnknownException implements Exception {
  final String? message;

  UnknownException(
      {this.message = 'An unknown error occurred. Please try again later.'});
}
