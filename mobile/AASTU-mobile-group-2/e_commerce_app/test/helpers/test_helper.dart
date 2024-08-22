import 'package:e_commerce_app/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:e_commerce_app/features/product/domain/repositories/product_repository.dart';
import 'package:http/http.dart' as http;
import 'package:internet_connection_checker_plus/internet_connection_checker_plus.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:shared_preferences/shared_preferences.dart';

import 'dummy/dummy.dart';
// import '';

@GenerateMocks(
  [
  ProductRepository,
  ProductRemoteDataSource,
  SharedPreferences,
  
  ],

  customMocks: [MockSpec<http.Client>(as: #MockHttpClient)]
)

void main(){

}