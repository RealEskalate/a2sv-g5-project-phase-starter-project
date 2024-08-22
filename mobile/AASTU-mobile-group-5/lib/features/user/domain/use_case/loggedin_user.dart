import '../../data/data_sources/user_local_data_source.dart';

class IsLoggedIn {
  final UserLocalDataSource localDataSource;

  IsLoggedIn(this.localDataSource);

  Future<bool> call() async {
    final token = await localDataSource.getAccessToken();
    return token != null;
  }
}
