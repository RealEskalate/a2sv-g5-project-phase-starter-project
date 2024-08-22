import 'package:flutter_application_5/core/network/network_info.dart';
import 'package:flutter_application_5/features/product/data/data_sources/local_data_source/local_data_source.dart';
import 'package:flutter_application_5/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:flutter_application_5/features/product/domain/repositories/product_repository.dart';
import 'package:http/http.dart' as http;

@GenerateMocks(
  [
    ProductRepository,
    ProductRemoteDataSource,
    ProductLocalDataSource,
    NetworkInfo,
    InternetConnectionChecker,
  ], 
  customMocks: [
    MockSpec<http.Client>(as: #MockHttpClient),
  ],
  )

void main() {}