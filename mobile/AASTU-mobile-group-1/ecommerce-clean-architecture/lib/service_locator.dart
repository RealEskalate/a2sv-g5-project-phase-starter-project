
import 'package:ecommerce/features/auth/data/repositories/user_repository_impl.dart';
import 'package:ecommerce/features/auth/domain/repository/user_repository.dart';
import 'package:ecommerce/features/auth/domain/usecases/loginUser.dart';
import 'package:ecommerce/features/auth/presentation/bloc/authbloc/auth_bloc.dart';
import 'package:get_it/get_it.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:http/http.dart' as http;

import 'core/networkinfo.dart';
import 'features/auth/data/data_sources/local_data_source.dart';
import 'features/auth/data/data_sources/remote_data_source.dart';
import 'features/auth/domain/usecases/registerUser.dart';
import 'features/chat_feature/chat/data_layer/data_source/remote_abstract.dart';
import 'features/chat_feature/chat/data_layer/data_source/remote_chat_source.dart';
import 'features/chat_feature/chat/data_layer/repository_imp/chat_repository_imp.dart';
import 'features/chat_feature/chat/domain/repository/chat_repository.dart';
import 'features/chat_feature/chat/domain/usecase/initialize_chat.dart';
import 'features/product/data/data_sources/local_data_source.dart';
import 'features/product/data/data_sources/remote_data_source.dart';
import 'features/product/data/repositories/product_repository_imp.dart';
import 'features/product/domain/repository/product_repository.dart';
import 'features/product/domain/usecases/addproduct.dart';
import 'features/product/domain/usecases/deleteproduct.dart';
import 'features/product/domain/usecases/getallproduct.dart';
import 'features/product/domain/usecases/getproduct.dart';
import 'features/product/domain/usecases/updateproduct.dart';
import 'features/product/presentation/bloc/getallproductbloc/bloc/product_bloc.dart';

final getIt = GetIt.instance;

Future<void> setup() async {
  var sharedPreference = await SharedPreferences.getInstance();
  var client = http.Client();
  var connectivity = InternetConnectionChecker();
  getIt.registerSingleton<InternetConnectionChecker>(connectivity);
  getIt.registerSingleton<NetworkInfo>(NetworkInfoImpl(getIt()));
  getIt.registerSingleton<SharedPreferences>(sharedPreference);
  getIt.registerSingleton<http.Client>(client);
  getIt.registerSingleton<ProductRemoteDataSource>(ProductRemoteDataSourceImpl(client: getIt()));
  getIt.registerSingleton<localDataSource>(localDataSourceImpl(sharedPreference: getIt()));
  getIt.registerSingleton<ProductRepository>(ProductRepositoryImpl(productLocalDataSource: getIt(), productRemoteDataSource: getIt(), networkInfo: getIt()));
  getIt.registerSingleton<GetAllProductUsecase>(GetAllProductUsecase(getIt()));
  getIt.registerSingleton<AddProductUsecase>(AddProductUsecase(getIt()));
  getIt.registerSingleton<UpdateProductUsecase>(UpdateProductUsecase(getIt()));
  getIt.registerSingleton<DeleteProductUsecase>(DeleteProductUsecase(getIt()));
  getIt.registerSingleton<GetProductUsecase>(GetProductUsecase(getIt()));
  getIt.registerSingleton<ProductBloc>(ProductBloc(getAllProductUsecase: getIt(), addProductUsecase: getIt(), updateProductUsecase: getIt(), deleteProductUsecase: getIt(), getProductUsecase: getIt()));
  getIt.registerSingleton<UserRemoteDataSource>(UserRemoteDataSourceImpl(client: getIt()));
  getIt.registerSingleton<UserLocalDataSource>(UserLocalDataSourceImpl(sharedPreferences: getIt()));
  getIt.registerSingleton<UserRepository>(UserRepositoryImpl(localDataSource: getIt(), remoteDataSource: getIt(), networkInfo: getIt()));
  getIt.registerSingleton<RegisterUserUseCase>(RegisterUserUseCase(getIt()));
  getIt.registerSingleton<LoginUserUsecase>(LoginUserUsecase(getIt()));
  getIt.registerSingleton<UserBloc>(UserBloc(loginUserUsecase: getIt(), registerUserUsecase: getIt())); 
  getIt.registerSingleton<RemoteAbstract>(RemoteChatSource(client: getIt()));
  getIt.registerSingleton<ChatRepository>(ChatRepositoryImp(remoteAbstract: getIt()));
  getIt.registerSingleton<InitializeChat>(InitializeChat(repository: getIt()));
}
