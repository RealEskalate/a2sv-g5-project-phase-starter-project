
import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';

import 'package:shared_preferences/shared_preferences.dart';

import '../../features/product/domain/usecase/get_all_product.dart';
import '../platform/network_info.dart';
import 'auth_injection.dart';
import 'product_injection.dart';

final sl = GetIt.instance;
Future<void> init() async {
  AuthInjection().init();
  ProductInjection().init();

  // core
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(sl()));
  sl.registerLazySingleton(() => InternetConnectionChecker());
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton(() => NoParams());
  final sharedPreferences = await SharedPreferences.getInstance();
  sl.registerLazySingleton(() => sharedPreferences);
}
