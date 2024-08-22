import 'package:internet_connection_checker/internet_connection_checker.dart';

abstract class NetworkInfo {
  Future<bool> get isConnected;
}

class NetworkInfoImple extends NetworkInfo {
  final InternetConnectionChecker connectionChecker;

  NetworkInfoImple({required this.connectionChecker});

  @override
  Future<bool> get isConnected async => await connectionChecker.hasConnection;
}
