import 'dart:io';

import 'package:http/http.dart' as http;
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:product_6/core/network/network_info.dart';
import 'package:product_6/core/util/input_converter.dart';
import 'package:product_6/features/auth/data/data_sources/auth_local_data_source.dart';
import 'package:product_6/features/auth/data/data_sources/auth_remote_data_source.dart';
import 'package:product_6/features/chat/data/data_source/remote_data_source/remote_data_source.dart';
import 'package:product_6/features/product/data/data_sources/local_data_source.dart';
import 'package:product_6/features/product/data/data_sources/remote_data_source.dart';
import 'package:product_6/features/product/domain/repositories/product_repository.dart';
import 'package:product_6/features/product/domain/usecases/delete_prodcut_usecase.dart';
import 'package:product_6/features/product/domain/usecases/get_all_prodcuts_usecase.dart';
import 'package:product_6/features/product/domain/usecases/get_product_usecase.dart';
import 'package:product_6/features/product/domain/usecases/insert_prodcut_usecase.dart';
import 'package:product_6/features/product/domain/usecases/update_product_usecase.dart';
import 'package:product_6/features/product/presentation/bloc/product_bloc.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks([
  ProductRepository,
  ProductRemoteDataSource,
  ProductLocalDataSource,
  InternetConnectionChecker,
  SharedPreferences,
  GetAllProductsUsecase,
  DeleteProductUsecase,
  InsertProductUsecase,
  GetProductUsecase,
  UpdateProductUsecase,
  InputConverter,
  NetworkInfo,
  AuthRemoteDataSource,
  AuthLocalDataSource,
  ProductBloc,
  
], customMocks: [
  MockSpec<http.Client>(as: #MockHttpClient)
])
void main() {}
