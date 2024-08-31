import 'package:get_it/get_it.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'core/network/http.dart';
import 'core/network/network_info.dart';
import 'features/authentication/data/data_sources/local/local_data_source.dart';
import 'features/authentication/data/data_sources/local/local_data_source_impl.dart';
import 'features/authentication/data/data_sources/remote/auth_remote_data_source.dart';
import 'features/authentication/data/data_sources/remote/auth_remote_datasource_impl.dart';
import 'features/authentication/data/repositories/auth_repo_impl.dart';
import 'features/authentication/domain/repositories/auth_repo.dart';
import 'features/authentication/domain/usecases/get_current_user_usecase.dart';
import 'features/authentication/domain/usecases/log_in_usecase.dart';
import 'features/authentication/domain/usecases/log_out_usecase.dart';
import 'features/authentication/domain/usecases/sign_up_usecase.dart';
import 'features/authentication/presentation/bloc/auth_bloc.dart';
import 'features/chat/data/data_sources/chat_data_source.dart';
import 'features/chat/data/repositories/chat_repository_impl.dart';
import 'features/chat/domain/repositories/chat_repository.dart';
import 'features/chat/domain/usecases/create_chat_usecase.dart';
import 'features/chat/domain/usecases/delete_chat_usecase.dart';
import 'features/chat/domain/usecases/get_chat_usecase.dart';
import 'features/chat/domain/usecases/get_chats_usecase.dart';
import 'features/chat/domain/usecases/send_message_usecase.dart';
import 'features/chat/presentation/blocs/bloc/chat_bloc.dart';
import 'features/product/data/data_sources/local/local_data_source.dart';
import 'features/product/data/data_sources/local/local_data_source_impl.dart';
import 'features/product/data/data_sources/remote/remote_data_source.dart';
import 'features/product/data/data_sources/remote/remote_data_source_impl.dart';
import 'features/product/data/repositories/product_repository_impl.dart';
import 'features/product/domain/repository/product_repository.dart';
import 'features/product/domain/usecases/add_product_usecase.dart';
import 'features/product/domain/usecases/delete_product_usecase.dart';
import 'features/product/domain/usecases/get_product_usecase.dart';
import 'features/product/domain/usecases/get_products_usecase.dart';
import 'features/product/domain/usecases/update_product_usecase.dart';
import 'features/product/presentation/bloc/product_bloc.dart';

final sl = GetIt.instance;

Future<void> init() async {
  //feature: Crud operations on product
  //bloc

  sl.registerFactory(
    () => ProductBloc(
      addProductUsecase: sl(),
      updateProductUsecase: sl(),
      deleteProductUsecase: sl(),
      getProductUsecase: sl(),
      getProductsUsecase: sl(),
    ),
  );
  // usecases
  sl.registerLazySingleton(() => AddProductUsecase(sl()));
  sl.registerLazySingleton(() => GetProductUsecase(sl()));
  sl.registerLazySingleton(() => GetProductsUsecase(sl()));
  sl.registerLazySingleton(() => UpdateProductUsecase(sl()));
  sl.registerLazySingleton(() => DeleteProductUsecase(sl()));
  // repository

  sl.registerLazySingleton<ProductRepository>(() => ProductRepositoryImpl(
        remoteDataSource: sl(),
        localDataSource: sl(),
        networkInfo: sl(),
      ));

  //data sources

  sl.registerLazySingleton<ProductRemoteDataSource>(
    () => RemoteDataSourceImpl(client: sl(), authLocalDataSource: sl()),
  );

  sl.registerLazySingleton<ProductLocalDataSource>(
    () => ProductLocalDataSourceImpl(sharedPreferences: sl()),
  );

  //!Core
  sl.registerLazySingleton<NetworkInfo>(
      () => NetworkInfoImpl(internetConnectionChecker: sl()));

  //external

  final sharedPreferences = await SharedPreferences.getInstance();
  sl.registerLazySingleton(() => sharedPreferences);
  sl.registerLazySingleton(() => http.Client());
  sl.registerLazySingleton(() => CustomHttp(client: sl()));
  sl.registerLazySingleton(() => InternetConnectionChecker());

  //feature: Authentication
  //bloc
  sl.registerFactory(
    () => AuthBloc(
      getCurrentUserUsecase: sl(),
      logInUsecase: sl(),
      logOutUsecase: sl(),
      signUpUsecase: sl(),
    ),
  );
  //usecases
  sl.registerLazySingleton(() => GetCurrentUserUsecase(authRepository: sl()));
  sl.registerLazySingleton(() => LogInUsecase(authRepository: sl()));
  sl.registerLazySingleton(() => SignUpUsecase(authRepository: sl()));
  sl.registerLazySingleton(() => LogOutUsecase(authRepository: sl()));
  //repositories

  sl.registerLazySingleton<AuthRepository>(
      () => AuthRepositoryImpl(authRemoteDataSource: sl()));

  //data sources
  sl.registerLazySingleton<AuthRemoteDataSource>(() => AuthRemoteDatasourceImpl(
        client: sl(),
        authLocalDataSource: sl(),
      ));
  sl.registerLazySingleton<AuthLocalDataSource>(() => AuthLocalDataSourceImpl(
        sharedPreferences: sl(),
      ));
  
  // Feature: chat

  // presentation
  sl.registerFactory(() => ChatBloc(
    getChatUsecase: sl(), 
    getChatsUsecase: sl(), 
    createChatUsecase: sl(), 
    sendMessageUsecase: sl(),
  ),);

  // usecase
  sl.registerLazySingleton(() => GetChatUsecase(sl()));
  sl.registerLazySingleton(() => GetChatsUsecase(sl()));
  sl.registerLazySingleton(() => CreateChatUsecase(sl()));
  sl.registerLazySingleton(() => DeleteChatUsecase(sl()));
  sl.registerLazySingleton(() => SendMessageUsecase(sl()));

  // repositores
  sl.registerLazySingleton<ChatRepository>(() => ChatRepositoryImpl(dataSource: sl()));

  // data source
  sl.registerLazySingleton<ChatDataSource>(() => ChatDataSourceImpl(client: sl()));
}
