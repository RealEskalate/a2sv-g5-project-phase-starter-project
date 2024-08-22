import 'package:ecommerce/core/import/import_file.dart';
final locator = GetIt.instance;
Future<void> setUp() async {
  final sharedPreferences = await SharedPreferences.getInstance();
  locator.registerLazySingleton<Dio>(() => Dio());
  locator.registerLazySingleton<InternetConnectionChecker>(
      () => InternetConnectionChecker());

  // Register NetworkInfo with NetworkInfoImpl
  locator.registerLazySingleton<NetworkInfo>(
    () => NetworkInfoImple(
        connectionChecker: locator<InternetConnectionChecker>()),
  );
  locator.registerLazySingleton<SharedPreferences>(() => sharedPreferences);

  locator.registerLazySingleton<LocalSource>(
    () => LocalDataSource(sharedPreferences: locator<SharedPreferences>()),
  );
  locator
      .registerLazySingleton<Api>(() => RemoteDataSource(dio: locator<Dio>()));
  locator.registerLazySingleton<ProductRepository>(() => ProductRepositoryImp(
      api: locator<Api>(),
      networkInfo: locator<NetworkInfo>(),
      localSource: locator<LocalSource>()));
  locator.registerLazySingleton<AddProductUsecase>(
      () => AddProductUsecase(productRepository: locator<ProductRepository>()));
  locator.registerLazySingleton<GetAllProductUsecase>(() =>
      GetAllProductUsecase(productRepository: locator<ProductRepository>()));
  locator.registerLazySingleton<GetProductUsecase>(
      () => GetProductUsecase(productRepository: locator<ProductRepository>()));
  locator.registerLazySingleton<DeleteProductUsecase>(() =>
      DeleteProductUsecase(productRepository: locator<ProductRepository>()));
  locator.registerLazySingleton<UpdateProductUsecase>(() =>
      UpdateProductUsecase(productRepository: locator<ProductRepository>()));

  // locator.registerLazySingleton<HomePageBloc>(()=>HomePageBloc(
  //   updateProductUsecase: locator<UpdateProductUsecase>(),
  //   addProductUsecase: locator<AddProductUsecase>(),
  //   deleteProductUsecase: locator<DeleteProductUsecase>(),
  //   getAllProductUsecase: locator<GetAllProductUsecase>(),
  // ));
}
