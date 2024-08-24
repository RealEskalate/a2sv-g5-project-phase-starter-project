import 'package:get_it/get_it.dart';

import '../../features/product/data/datasources/local_data_resource.dart';
import '../../features/product/data/datasources/remote_data_source.dart';
import '../../features/product/data/repositories/product_repository_impl.dart';
import '../../features/product/domain/repositories/product_repository.dart';
import '../../features/product/domain/usecase/delete_product.dart';
import '../../features/product/domain/usecase/get_all_product.dart';
import '../../features/product/domain/usecase/get_product.dart';
import '../../features/product/domain/usecase/insert_product.dart';
import '../../features/product/domain/usecase/update_product.dart';
import '../../features/product/presentation/bloc/product_bloc.dart';
import 'injection.dart';

class ProductInjection {
  Future<void> init() async {
    //! Features - Number Trivia
    // Bloc
    sl.registerFactory(
      () => ProductBloc(
        getProductUsecase: sl(),
        getAllProductUsecase: sl(),
        updateProductUsecase: sl(),
        deleteProductUsecase: sl(),
        insertProductUsecase: sl(),
      ),
    );

    // Use cases
    sl.registerLazySingleton(() => GetProductUsecase(sl()));
    sl.registerLazySingleton(() => GetAllProductUsecase(sl()));
    sl.registerLazySingleton(() => UpdateProductUsecase(sl()));
    sl.registerLazySingleton(() => DeleteProductUsecase(sl()));
    sl.registerLazySingleton(() => InsertProductUsecase(sl()));

    // Repository
    sl.registerLazySingleton<ProductRepository>(
      () => ProductRepositoryImpl(
        localDataSource: sl(),
        networkInfo: sl(),
        remoteDataSource: sl(),
      ),
    );

    // Data sources
    sl.registerLazySingleton<ProductRemoteDataSource>(
      () => ProductRemoteDataSourceImpl(client: sl()),
    );

    sl.registerLazySingleton<ProductLocalDataSource>(
      () => ProductLocalDataSourceImpl(sharedPreferences: sl()),
    );
  }
}
