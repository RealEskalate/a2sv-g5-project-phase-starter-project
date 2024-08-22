import 'package:ecommerse2/features/product/data/data_sources/remote_data_source.dart';
import 'package:ecommerse2/features/product/domain/repository/productRepository.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';

@GenerateMocks(

  [
    ProductRepository,
    ProductRemoteDataSource, http.MultipartRequest, http.MultipartFile, http.StreamedResponse
  ],
  customMocks: [MockSpec<http.Client>(as: #MockHttpClient)],

)

void main() {}