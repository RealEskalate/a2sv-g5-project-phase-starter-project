import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:product_8/core/platform/network_info.dart';
import 'package:product_8/core/usecase/usecase.dart';
import 'package:product_8/features/auth/data/data_source/auth_local_data_source.dart';
import 'package:product_8/features/auth/data/data_source/auth_remote_data_source.dart';
import 'package:product_8/features/auth/domain/repositories/auth_repository.dart';
import 'package:product_8/features/auth/domain/use_case/sign_in_use_case.dart';
import 'package:product_8/features/auth/domain/use_case/sign_up_use_case.dart';
import 'package:product_8/features/product/data/data_source/local_data_source/product_local_data_source.dart';
import 'package:product_8/features/product/data/data_source/remote_data_source/product_remote_data_source.dart';

import 'package:product_8/features/product/domain/repositories/product_repository.dart';
import 'package:product_8/features/product/domain/use_case/delete_product_usecase.dart';
import 'package:product_8/features/product/domain/use_case/get_product_by_id_usecase.dart';
import 'package:product_8/features/product/domain/use_case/get_products_usecase.dart';
import 'package:product_8/features/product/domain/use_case/insert_product_usecase.dart';
import 'package:product_8/features/product/domain/use_case/update_product_usecase.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks([
  ProductRepositories,
  ProductRemoteDataSource,
  ProductLocalDataSource,
  NetworkInfo,
  InternetConnectionChecker,
  SharedPreferences ,
  UseCase, 
  InsertProductUsecase,
  GetProductsUsecase,
  GetProductByIdUsecase,
  DeleteProductUsecase,
  UpdateProductUsecase,
  AuthRepository,
  AuthRemoteDataSource,
  AuthLocalDataSource,
  SignInUseCase,
  SignUpUseCase
], customMocks: [
  MockSpec<http.Client>(as: #MockHttpClient)
])
void main() {}
