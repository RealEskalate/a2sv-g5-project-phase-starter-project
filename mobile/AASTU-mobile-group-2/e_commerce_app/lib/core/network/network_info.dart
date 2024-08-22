import 'package:internet_connection_checker/internet_connection_checker.dart';

abstract class NetworkInfo {
  Future<bool> get isConnected;
}

class NetworkInfoImplimentation implements NetworkInfo {
  InternetConnectionChecker internetStatus;
  NetworkInfoImplimentation({required this.internetStatus});

  @override
  Future<bool> get isConnected async => internetStatus.hasConnection;
}
