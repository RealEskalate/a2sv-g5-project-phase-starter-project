import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'core/network/custom_client.dart';
import 'core/platform/network_info.dart';
import 'features/auth/data/datasources/auth_local_data_source.dart';
import 'features/auth/data/datasources/auth_remote_data_source.dart';
import 'features/auth/data/repositories/auth_repository_impl.dart';
import 'features/auth/domain/repositories/auth_repository.dart';
import 'features/auth/domain/usecases/get_user.dart';
import 'features/auth/domain/usecases/login.dart';
import 'features/auth/domain/usecases/logout.dart';
import 'features/auth/domain/usecases/register.dart';
import 'features/auth/presentation/bloc/auth_bloc.dart';
import 'features/product/data/datasources/product_local_data_source.dart';
import 'features/product/data/datasources/product_remote_data_source.dart';
import 'features/product/data/repositories/product_repository_impl.dart';
import 'features/product/domain/repositories/product_repository.dart';
import 'features/product/domain/usecases/create_product.dart';
import 'features/product/domain/usecases/delete_product.dart';
import 'features/product/domain/usecases/get_all_products.dart';
import 'features/product/domain/usecases/get_current_product.dart';
import 'features/product/domain/usecases/update_product.dart';
import 'features/product/presentation/bloc/product_bloc.dart';

final sl = GetIt.instance;

Future<void> init() async {
  //Bloc
  sl.registerFactory(() => ProductBloc(
        sl(),
        sl(),
        sl(),
        sl(),
        sl(),
      ));
  sl.registerFactory(() => AuthBloc(
        sl(),
        sl(),
        sl(),
        sl(),
      ));
  //Use cases
  sl.registerLazySingleton(() => GetCurrentProductUsecase(sl()));
  sl.registerLazySingleton(() => GetAllProductsUsecase(sl()));
  sl.registerLazySingleton(() => CreateProductUsecase(sl()));
  sl.registerLazySingleton(() => UpdateProductUsecase(sl()));
  sl.registerLazySingleton(() => DeleteProductUsecase(sl()));

  //Repository
  sl.registerLazySingleton<ProductRepository>(() => ProductRepositoryImpl(
        remoteDataSource: sl(),
        localDataSource: sl(),
        networkInfo: sl(),
      ));

  //Product Data sources
  sl.registerLazySingleton<ProductRemoteDataSource>(
      () => ProductRemoteDataSourceImpl(client: sl()));
  sl.registerLazySingleton<ProductLocalDataSource>(
      () => ProductLocalDataSourceImpl(sharedPreferences: sl()));

  // Auth data source
  sl.registerLazySingleton<AuthLocalDataSource>(
      () => AuthLocalDataSourceImpl(sharedPreferences: sl()));
  sl.registerLazySingleton<AuthRemoteDataSource>(() => AuthRemoteDataSourceImpl(client: sl()));


  // Auth Use cases
  sl.registerLazySingleton(() => LoginUsecase(sl()));
  sl.registerLazySingleton(() => LogoutUsecase(sl()));
  sl.registerLazySingleton(() => RegisterUsecase(sl()));
  sl.registerLazySingleton(() => GetUserUsecase(sl()));


  //Auth Repository
  sl.registerLazySingleton<AuthRepository>(() => AuthRepositoryImpl(
        remoteDataSource: sl(),
        localDataSource: sl(),
        networkInfo: sl(),
      ));

  // Custom Http Client
  sl.registerLazySingleton(() => CustomHttpClient(
        client: sl(),
        authLocalDataSource: sl(),
      ));

  //Core
  sl.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(sl()));

  //External
  final sharedPreferences = await SharedPreferences.getInstance();
  sl.registerLazySingleton(() => sharedPreferences);
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton(() => InternetConnectionChecker());
}
