import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'core/network/network_info.dart';
import 'features/authentication/data/data_source/local_data_source.dart';
import 'features/authentication/data/data_source/remote_data_source.dart';
import 'features/authentication/data/repository/authentication_repository_impl.dart';
import 'features/authentication/domain/repository/authentication_repository.dart';
import 'features/authentication/domain/usecases/log_out.dart';
import 'features/authentication/domain/usecases/login.dart';
import 'features/authentication/domain/usecases/sign_up.dart';
import 'features/authentication/presentation/bloc/blocs.dart';

import 'features/product/data/data_sources/local_data_source.dart';
import 'features/product/data/data_sources/remote_data_source.dart';
import 'features/product/data/repositories/product_repository_impl.dart';
import 'features/product/domain/repository/productRepository.dart';
import 'features/product/domain/usecase/add_product.dart';
import 'features/product/domain/usecase/delete_product.dart';
import 'features/product/domain/usecase/get_all_product.dart';
import 'features/product/domain/usecase/get_product.dart';
import 'features/product/domain/usecase/update_product.dart';
import 'features/product/presentation/bloc/blocs.dart';
   

final getIt = GetIt.instance;

Future<void> setUp() async {
  final client = http.Client();
  InternetConnectionChecker connectionChecker = InternetConnectionChecker();
  SharedPreferences sharedPreferences = await SharedPreferences.getInstance();

  getIt.registerLazySingleton<InternetConnectionChecker>(() => connectionChecker);
  getIt.registerLazySingleton<http.Client>(() => client);
  getIt.registerLazySingleton<SharedPreferences>(() => sharedPreferences);
  
  // Ensure NetworkInfo is registered after its dependencies
  getIt.registerLazySingleton<NetworkInfo>(() => NetworkInfoImpl(getIt<InternetConnectionChecker>()));

  getIt.registerLazySingleton<LocalDataSource>(
      () => LocalDataSourceImpl(sharedPreferences: getIt<SharedPreferences>()));
  getIt.registerLazySingleton<ProductRemoteDataSource>(
      () => ProductRemoteDataSourceImpl(client: getIt<http.Client>()));
  getIt.registerLazySingleton<ProductRepository>(
      () => ProductRepositoryImpl(localDataSource:  getIt<LocalDataSource>(), productRemoteDataSource:  getIt<ProductRemoteDataSource>(),networkInfo:  getIt<NetworkInfo>()));
  getIt.registerLazySingleton<GetAllProductUseCase>(() => GetAllProductUseCase(getIt<ProductRepository>()));
  getIt.registerLazySingleton(() => GetProductUseCase(getIt<ProductRepository>()));
  getIt.registerLazySingleton(() => UpdateProductUseCase(getIt<ProductRepository>()));
  getIt.registerLazySingleton(() => DeleteProductUseCase(getIt<ProductRepository>()));
  getIt.registerLazySingleton(() => AddProductUseCase(getIt<ProductRepository>()));
  getIt.registerLazySingleton(() => ProductBloc(
        updateProductUseCase: getIt<UpdateProductUseCase>(),
        addProductUseCase: getIt<AddProductUseCase>(),
        deleteProductUseCase: getIt<DeleteProductUseCase>(),
        getAllProductUseCase: getIt<GetAllProductUseCase>(),
        getProductUseCase: getIt<GetProductUseCase>(),
      ));
  
  getIt.registerLazySingleton<UserRemoteDataSource>(() => UserRemoteDataSourceImpl(getIt<http.Client>()));
  getIt.registerLazySingleton<UserLocalDataSource>(() => UserLocalDataSourceImpl(sharedPreferences: getIt<SharedPreferences>()));
  getIt.registerLazySingleton<AuthenticationRepository>(() => AuthenticationRepositoryImpl( localDataSource:getIt<UserLocalDataSource>(), remoteDataSource:  getIt<UserRemoteDataSource>(),  networkInfo:  getIt<NetworkInfo>()));
  getIt.registerLazySingleton(() => LoginUseCase(getIt<AuthenticationRepository>()));
  getIt.registerLazySingleton(() => SignUpUseCase(getIt<AuthenticationRepository>())); 
  getIt.registerLazySingleton(() => LogOutUseCase(getIt<AuthenticationRepository>())); 
  getIt.registerLazySingleton(() =>
    UserBloc(
      signUpUseCase: getIt<SignUpUseCase>(),
      loginUseCase: getIt<LoginUseCase>(),
      logOutUseCase: getIt<LogOutUseCase>()
    ),
  );
}
