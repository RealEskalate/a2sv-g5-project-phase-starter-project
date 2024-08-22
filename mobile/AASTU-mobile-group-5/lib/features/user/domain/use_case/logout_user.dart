import '../../data/data_sources/user_local_data_source.dart';

class LogOut {
  final UserLocalDataSource localDataSource;

  LogOut(this.localDataSource);

  Future<void> call() async {
    await localDataSource.deleteAccessToken();
  }
}
