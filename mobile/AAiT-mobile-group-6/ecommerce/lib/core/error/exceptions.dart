class ServerException implements Exception {
  final String message;
  ServerException({this.message = 'Server Error'});
}

class CacheException implements Exception{
  final String message;
  CacheException({this.message = 'Cache Error'});
}
