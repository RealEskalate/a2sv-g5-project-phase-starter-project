import 'package:data_connection_checker_tv/data_connection_checker.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';

import 'package:product_6/core/connections/network_info.dart';
import 'package:product_6/features/product/data/data_sources/product_remote_datasource.dart';
import 'package:product_6/features/product/domain/repository/product_repository.dart';
import 'package:product_6/features/product/domain/usecases/create_product.dart';
import 'package:product_6/features/product/domain/usecases/delete_product.dart';
import 'package:product_6/features/product/domain/usecases/update_product.dart';
import 'package:product_6/features/product/domain/usecases/view_all_products.dart';
import 'package:product_6/features/product/domain/usecases/view_product.dart';
import 'package:shared_preferences/shared_preferences.dart';

@GenerateMocks([
  ProductRepository,
  ProductRemoteDatasource,
  DataConnectionChecker,
  NetworkInfo,
  http.Client,
  SharedPreferences,
  CreateProductUseCase,
  DeleteProductUseCase,
  UpdateProductUseCase,
  ViewAllProductsUseCase,
  ViewProductUseCase,

])
void main() {}

