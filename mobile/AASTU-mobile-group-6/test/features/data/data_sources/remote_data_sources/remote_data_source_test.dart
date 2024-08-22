
import 'package:ecommerce_app_ca_tdd/core/constants/constants.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
import '../../../../helpers/test_helper.mocks.dart';

class MockHttpClient extends Mock implements http.Client {}
void main() {
  late ProductRemoteDataSourceImpl dataSource;
  late MockHttpClient mockHttpClient;

  setUp(() {
    mockHttpClient = MockHttpClient();
    dataSource = ProductRemoteDataSourceImpl(client: mockHttpClient);
  });

  group('getProduct', () {
    const productId = '1';
    final productJson = '{"id": "1", "name": "Test Product", "description": "Test Description", "price": 10.0}';
    final productModel = ProductModel.fromJson({
      'id': '1',
      'name': 'Test Product',
      'description': 'Test Description',
      'price': 10.0,
    });

    test('should return ProductModel when the response code is 200 (success)', () async {
      when(mockHttpClient.get(Uri.parse(Urls.getProduct(productId))))
          .thenAnswer((_) async => http.Response(productJson, 200));

      final result = await dataSource.getProduct(productId);

      expect(result, equals(productModel));
    });

    test('should throw an Exception when the response code is not 200', () async {
      when(mockHttpClient.get(Uri.parse(Urls.getProduct(productId))))
          .thenAnswer((_) async => http.Response('Something went wrong', 404));

      expect(() => dataSource.getProduct(productId), throwsException);
    });
  });

  // Similar test cases can be written for addProduct, deleteProduct, updateProduct, and getProducts.
}
