import 'dart:async';

import 'package:data_connection_checker_tv/data_connection_checker.dart';
import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:shared_preferences/shared_preferences.dart';

import 'core/connections/network_info.dart';
import 'features/auth/data/data_source/auth_local_datasource.dart';
import 'features/auth/data/data_source/auth_remote_datasource.dart';
import 'features/auth/data/repository/auth_repository_imp.dart';
import 'features/auth/domain/repository/auth_repository.dart';
import 'features/auth/domain/usecases/get_user_profile.dart';
import 'features/auth/domain/usecases/login_usecase.dart';
import 'features/auth/domain/usecases/logout_usecase.dart';
import 'features/auth/domain/usecases/register_usecase.dart';
import 'features/auth/presentation/bloc/auth_bloc.dart';
import 'features/product/data/data_sources/product_local_data_source.dart';
import 'features/product/data/data_sources/product_remote_datasource.dart';
import 'features/product/data/repository/product_repository_imp.dart';
import 'features/product/domain/repository/product_repository.dart';
import 'features/product/domain/usecases/create_product.dart';
import 'features/product/domain/usecases/delete_product.dart';
import 'features/product/domain/usecases/update_product.dart';
import 'features/product/domain/usecases/view_all_products.dart';
import 'features/product/domain/usecases/view_product.dart';
import 'features/product/presentation/bloc/product_bloc.dart';

final GetIt locator = GetIt.instance;

Future<void> setupLocator() async {
  // Product External
  final sharedPreferences = await SharedPreferences.getInstance();
  locator.registerLazySingleton<SharedPreferences>(() => sharedPreferences);
  locator.registerLazySingleton<DataConnectionChecker>(
      () => DataConnectionChecker());
  locator.registerLazySingleton<http.Client>(() => http.Client());

  // Product Core
  locator.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(locator()));

  // Product Data sources
  locator.registerLazySingleton<ProductLocalDataSource>(
      () => ProductLocalDataSourceImpl(sharedPreferences: locator()));
  locator.registerLazySingleton<ProductRemoteDatasource>(
      () => ProductRemoteDatasourceImp(client: locator(), localDataSource: locator()));

  //Auth Data sources
  locator.registerLazySingleton<AuthRemoteDataSource>(
      () => AuthRemoteDataSourceImpl(client: locator()));
  locator.registerLazySingleton<AuthLocalDataSource>(
      () => AuthLocalDataSourceImpl(sharedPreferences: locator()));

  // Product Repository
  locator.registerLazySingleton<ProductRepository>(() => ProductRepositoryImp(
        productRemoteDatasource: locator(),
        networkInfo: locator(),
        productLocalDataSource: locator(),
      ));

  //Auth Repository
  locator.registerLazySingleton<AuthRepository>(() => AuthRepositoryImpl(
        remoteDataSource: locator(),
        localDataSource: locator(),
        networkInfo: locator(),
      ));

  // Product Use cases
  locator.registerLazySingleton<CreateProductUseCase>(
      () => CreateProductUseCase(locator()));
  locator.registerLazySingleton<DeleteProductUseCase>(
      () => DeleteProductUseCase(locator()));
  locator.registerLazySingleton<UpdateProductUseCase>(
      () => UpdateProductUseCase(locator()));
  locator.registerLazySingleton<ViewAllProductsUseCase>(
      () => ViewAllProductsUseCase(locator()));
  locator.registerLazySingleton<ViewProductUseCase>(
      () => ViewProductUseCase(locator()));
  // Auth Use cases
  locator.registerLazySingleton<Login>(() => Login(locator()));
  locator.registerLazySingleton<Register>(() => Register(locator()));
  locator
      .registerLazySingleton<GetUserProfile>(() => GetUserProfile(locator()));
  locator.registerLazySingleton<LogoutUseCase>(() => LogoutUseCase(locator()));

  // Product Bloc
  locator.registerFactory<ProductBloc>(() => ProductBloc(
        createProductUseCase: locator(),
        deleteProductUseCase: locator(),
        updateProductUseCase: locator(),
        viewAllProductsUseCase: locator(),
        viewProductUseCase: locator(),
      ));
  // Auth Bloc
  locator.registerFactory<AuthBloc>(() => AuthBloc(
        loginUseCase: locator(),
        registerUseCase: locator(),
        getUserProfileUseCase: locator(),
        logoutUseCase: locator(),
      ));
}















// import 'dart:async';
// import 'dart:io';

// import 'package:data_connection_checker_tv/data_connection_checker.dart';
// import 'package:get_it/get_it.dart';
// import 'package:shared_preferences/shared_preferences.dart';

// import 'core/connections/network_info.dart';
// import 'features/product/data/data_sources/product_local_data_source.dart';
// import 'features/product/data/data_sources/product_remote_datasource.dart';
// import 'features/product/data/repository/product_repository_imp.dart';
// import 'features/product/domain/repository/product_repository.dart';
// import 'features/product/domain/usecases/create_product.dart';
// import 'features/product/domain/usecases/delete_product.dart';
// import 'features/product/domain/usecases/update_product.dart';
// import 'features/product/domain/usecases/view_all_products.dart';
// import 'features/product/domain/usecases/view_product.dart';
// import 'features/product/presentation/bloc/product_bloc.dart';

// final GetIt locator = GetIt.instance;

// Future<void> setupLocator() async {
//   //Bloc
//   locator.registerFactory<ProductBloc>(() => ProductBloc(
//         createProductUseCase: locator(),
//         deleteProductUseCase: locator(),
//         updateProductUseCase: locator(),
//         viewAllProductsUseCase: locator(),
//         viewProductUseCase: locator(),
//       ));

//   //usecases
//   locator.registerLazySingleton<CreateProductUseCase>(
//       () => CreateProductUseCase(locator()));
//   locator.registerLazySingleton<DeleteProductUseCase>(
//       () => DeleteProductUseCase(locator()));
//   locator.registerLazySingleton<UpdateProductUseCase>(
//       () => UpdateProductUseCase(locator()));
//   locator.registerLazySingleton<ViewAllProductsUseCase>(
//       () => ViewAllProductsUseCase(locator()));
//   locator.registerLazySingleton<ViewProductUseCase>(
//       () => ViewProductUseCase(locator()));

//   //repository
//   locator.registerLazySingleton<ProductRepository>(() => ProductRepositoryImp(
//         productRemoteDatasource: locator(),
//         networkInfo: locator(),
//         productLocalDataSource: locator(),
//       ));

//   //data
//   locator.registerLazySingleton<ProductLocalDataSource>(
//       () => ProductLocalDataSourceImpl(sharedPreferences: locator()));
//   locator.registerLazySingleton<ProductRemoteDatasource>(
//       () => ProductRemoteDatasourceImp(client: locator()));

//   //core
//   locator.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(locator()));

//   //external
//   final sharedpref = await SharedPreferences.getInstance();
//   locator.registerLazySingleton<SharedPreferences>(() => sharedpref);
//   locator.registerLazySingleton<DataConnectionChecker>(
//       () => DataConnectionChecker());
//   locator.registerLazySingleton<HttpClient>(() => HttpClient());
// }
