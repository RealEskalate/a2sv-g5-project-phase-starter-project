import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/local_data_source/local_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/repositories/product_repository_impl.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/repositories/product_repository.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/add_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/delete_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_all_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_detail_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/get_user_info.dart';
import 'package:ecommerce_app_ca_tdd/features/product/domain/usecases/update_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/add/add_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/detail/detail_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/home_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/product/presentation/bloc/update/bloc/update_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/data/repository/user_repository_impl.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/repository/users_repository.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/login_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/domain/usecases/signup_usecase.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/get_user/get_user_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/login/bloc/sign_in_bloc.dart';
import 'package:ecommerce_app_ca_tdd/features/user_auth/presentation/bloc/signup/bloc/sign_up_bloc.dart';
import 'package:get/get.dart';
import 'package:get_it/get_it.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

final sl = GetIt.instance;


Future<void> init() async {
  
  
  //! Externals
  var ht = await SharedPreferences.getInstance();
  sl.registerLazySingleton<SharedPreferences>(() => ht);
  sl.registerLazySingleton<http.Client>(()=> http.Client());
  sl.registerLazySingleton<InternetConnectionChecker>(()=> InternetConnectionChecker());
  sl.registerLazySingleton<NetworkInfo>(()=> NetworkInfoImpl(sl()));

  // Data Sources
  sl.registerLazySingleton<ProductRemoteDataSource>(()=> ProductRemoteDataSourceImpl(client: sl()));
  sl.registerLazySingleton<ProductLocalDataSource>(()=> ProductLocalDataSourceImpl(sharedPreferences: sl()));
  // Auth Data Sources
  sl.registerLazySingleton<UserRemoteDataSource>(()=> UserRemoteDataSourceImpl(client: sl()));

  //! Features
  // Repositories

  sl.registerLazySingleton<ProductRepository>(()=> ProductRepositoryImpl(ProductRemoteDataSourceImpl(client: sl()),ProductLocalDataSourceImpl(sharedPreferences:  sl()),NetworkInfoImpl(sl())));
  // Auth Repositories
  sl.registerLazySingleton<UsersRepository>(()=> UserRepositoryImpl(userRemoteDataSource: UserRemoteDataSourceImpl(client: sl())));



  // Use Cases
  sl.registerLazySingleton<GetAllUsecase>(()=> GetAllUsecase(sl()));
  sl.registerLazySingleton<GetDetailUseCase>(()=> GetDetailUseCase(sl()));
  sl.registerLazySingleton<AddProductUseCase>(()=> AddProductUseCase(sl()));
  sl.registerLazySingleton<DeleteUsecase>(()=> DeleteUsecase(sl()));

  // AUth Use Cases
  sl.registerLazySingleton<RegisterUser>(()=> RegisterUser(sl()));
  sl.registerLazySingleton<LoginUsecase>(()=> LoginUsecase(sl()));
  sl.registerLazySingleton<GetUserInfo>(()=> GetUserInfo(sl()));
  // sl.registerLazySingleton(()=> Up)

    //BLoc
  sl.registerFactory(()=> HomeBloc(
    GetAllUsecase(sl()), GetDetailUseCase(sl())));
  sl.registerFactory(()=> DetailBloc(
    DeleteUsecase(sl())));
    sl.registerFactory(()=> addBloc(
    AddProductUseCase(sl())));
    sl.registerFactory(()=> UpdateBloc(
    UpdateUsecase(sl())));
    
    // Auth Bloc
    sl.registerFactory(()=> SignUpBloc(RegisterUser(sl())));
    sl.registerFactory(()=> LoginBloc(LoginUsecase(sl())));
    sl.registerFactory(()=> GetUserBloc(GetUserInfo(sl())));
  
  //!COre



}