import 'package:ecommerce/core/platform/network_info.dart';
import 'package:ecommerce/features/product/data/datasources/local_data_resource.dart';
import 'package:ecommerce/features/product/data/datasources/remote_data_source.dart';
import 'package:ecommerce/features/product/domain/repositories/product_repository.dart';
import 'package:ecommerce/features/product/domain/usecase/delete_product.dart';
import 'package:ecommerce/features/product/domain/usecase/get_all_product.dart';
import 'package:ecommerce/features/product/domain/usecase/get_product.dart';
import 'package:ecommerce/features/product/domain/usecase/insert_product.dart';
import 'package:ecommerce/features/product/domain/usecase/update_product.dart';
import 'package:ecommerce/features/product/presentation/bloc/product_bloc.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks(
  [
    ProductRepository,
    ProductRemoteDataSource,
    NetworkInfo,
    ProductLocalDataSource,
    InternetConnectionChecker,
    SharedPreferences,
    GetProductUsecase,
    GetAllProductUsecase,
    DeleteProductUsecase,
    InsertProductUsecase,
    UpdateProductUsecase,
    ProductBloc
  ],
  customMocks: [MockSpec<http.Client>(as: #MockHttpClient)],
)
void main() {}
