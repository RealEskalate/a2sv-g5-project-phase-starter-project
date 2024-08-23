abstract class AuthLocalDataSource {
  Future<bool> cacheToken(String token);
  Future<String> getToken();
  Future<void> removeToken();
}