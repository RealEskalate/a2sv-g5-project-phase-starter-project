import 'package:flutter/material.dart';
import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'core/connectivity/network_info.dart';
import 'features/product/data/data_sources/product_local_data_source.dart';
import 'features/product/data/data_sources/product_remote_data_source.dart';
import 'features/product/data/repositories/product_repository_impl.dart';
import 'features/product/domain/repository/product_repository.dart';
import 'features/product/domain/use_case/add_product.dart';
import 'features/product/domain/use_case/delete_product.dart';
import 'features/product/domain/use_case/get_all_products.dart';
import 'features/product/domain/use_case/get_product.dart';
import 'features/product/domain/use_case/update_product.dart';
import 'features/product/presentation/bloc/add_page/add_page_bloc.dart';
import 'features/product/presentation/bloc/details_page/details_page_bloc.dart';
import 'features/product/presentation/bloc/home_page/home_page_bloc.dart';
import 'features/product/presentation/bloc/search_page/search_page_bloc.dart';
import 'features/product/presentation/bloc/update_page/update_page_bloc.dart';
import 'features/user/data/data_sources/user_local_data_source.dart';
import 'features/user/data/data_sources/user_name.dart';
import 'features/user/data/data_sources/user_remote_data_source.dart';
import 'features/user/data/repositories/user_respository_impl.dart';
import 'features/user/domain/repositories/user_repository.dart';
import 'features/user/domain/use_case/loggedin_user.dart';
import 'features/user/domain/use_case/login_user.dart';
import 'features/user/domain/use_case/logout_user.dart';
import 'features/user/domain/use_case/register_user.dart';
import 'features/user/presentation/bloc/authentication/authentication_bloc.dart';
import 'features/user/presentation/bloc/sign_in_page/sign_in_page_bloc.dart';
import 'features/user/presentation/bloc/sign_up_page/sign_up_page_bloc.dart';

final getIt = GetIt.instance;

Future<void> setupLocator() async {
  WidgetsFlutterBinding.ensureInitialized(); // Ensure binding is initialized

  final sharedPreferences = await SharedPreferences.getInstance();
  final client = http.Client();
  final internetConnectionChecker = InternetConnectionChecker();

  getIt.registerSingleton<SharedPreferences>(sharedPreferences);
  getIt.registerSingleton<http.Client>(client);
  getIt.registerSingleton<InternetConnectionChecker>(internetConnectionChecker);

  getIt.registerLazySingleton<UserLocalDataSource>(
    () => UserLocalDataSourceImpl(sharedPreferences: getIt()),
  );

  getIt.registerLazySingleton<ProductRemoteDataSource>(
    () => ProductRemoteDataSourceImpl(client: getIt(), userLocalDataSource: getIt(),),
  );
  getIt.registerLazySingleton<ProductLocalDataSource>(
    () => ProductLocalDataSourceImpl(sharedPreferences: getIt()),
  );
  getIt.registerLazySingleton<NetworkInfo>(
    () => NetworkInfoImpl(internetConnectionChecker),
  );

  getIt.registerSingleton<ProductRepository>(ProductRepositoryImpl(
    remoteDataSource: getIt(),
    localDataSource: getIt(),
    networkInfo: getIt(),
  ));
  getIt.registerFactory(() => AddProduct(getIt()));
  getIt.registerFactory(() => DeleteProduct(getIt()));
  getIt.registerFactory(() => UpdateProduct(getIt()));
  getIt.registerFactory(() => GetProduct(getIt()));
  getIt.registerFactory(() => GetAllProducts(getIt()));

  getIt.registerFactory<HomePageBloc>(
    () => HomePageBloc(getAllProducts: getIt<GetAllProducts>()),
  );
  getIt.registerFactory<SearchPageBloc>(
    () => SearchPageBloc(getAllProducts: getIt<GetAllProducts>()),
  );

  getIt.registerFactory(
      () => DetailsPageBloc(GetProduct(getIt()), DeleteProduct(getIt())));
  getIt.registerFactory(() => AddPageBloc(AddProduct(getIt())));
  getIt.registerFactory(
      () => UpdatePageBloc(UpdateProduct(getIt()), DeleteProduct(getIt())));

  getIt.registerLazySingleton<UserRemoteDataSource>(
    () => UserRemoteDataSourceImpl(client: getIt()),
  );

  
  
  getIt.registerSingleton<UserRepository>(
      UserRepositoryImpl(remoteDataSource: getIt(), localDataSource: getIt()));
    

  getIt.registerFactory(() => LoginUser(getIt()));
  getIt.registerFactory(() => RegisterUser(getIt()));

  getIt.registerFactory(() => SignInPageBloc(userRepository: getIt()));
  getIt.registerFactory(() => SignUpPageBloc(userRepository: getIt()));

  getIt.registerFactory(() => IsLoggedIn(getIt()));
  getIt.registerFactory(() => LogOut(getIt()));

  getIt.registerFactory(() => AuthenticationBloc(getIt()));

  getIt.registerLazySingleton<Future<String> Function()>(() => fetchUserName);

}
