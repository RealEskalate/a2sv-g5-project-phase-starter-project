import 'package:ecommerse2/core/constants/constants.dart';
import 'package:ecommerse2/core/error/exception.dart';
import 'package:ecommerse2/features/product/data/data_sources/remote_data_source.dart';
import 'package:ecommerse2/features/product/data/models/product_model.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';
import '../../helpers/json_reader.dart';
import '../../helpers/test_helper.mocks.dart';

void main() {
  final product = ProductModel.fromJson(const {
    "id": "6672940692adcb386d593686",
    "name": "PC",
    "description": "long description",
    "price": 123,
    "imageUrl":
        "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718785031/images/zqfvuxrxhip7shikyyj4.png"
  });

  late MockHttpClient mockHttpClient;
  late MockMultipartRequest mockRequest;
  late MockStreamedResponse mockResponse;

  late ProductRemoteDataSourceImpl productRemoteDataSourceImpl;

  setUp(() {
    mockHttpClient = MockHttpClient();
    productRemoteDataSourceImpl =
        ProductRemoteDataSourceImpl(client: mockHttpClient);
      mockRequest = MockMultipartRequest();
    mockResponse = MockStreamedResponse();
  });

//We are going to run two tests
// 1. Returning a vaild model when getting data is successful
// 2. Returning a server exception when getting data is failed
  group('get all products', () {
    test('should return list of product model when the response code is 200 ',
        () async {
      //arrange

      when(mockHttpClient.get(Uri.parse(Urls.baseUrl))).thenAnswer((_) async =>
          http.Response(
              readJson('features/product/helpers/dummy_data/dummy_json.json'),
              200));

      //act
      final result = await productRemoteDataSourceImpl.getAllProduct();

      //assert
      expect(result, isA<List<ProductModel>>());
    });
    test(
        'should throw a server exception when the response code is 404 or other ',
        () async {
      //arrange

      when(mockHttpClient.get(Uri.parse(Urls.baseUrl)))
          .thenAnswer((_) async => http.Response('Not Found', 404));

      //act
      final result = productRemoteDataSourceImpl.getAllProduct;

      //assert
      expect(() async => await result(), throwsA(isA<ServerException>()));
    });
  });

  group('get product by id', () {
    test('should return list of product model when the response code is 200 ',
        () async {
      when(mockHttpClient.get(Uri.parse('${Urls.baseUrl}/${product.id}')))
          .thenAnswer((_) async => http.Response(
              readJson('features/product/helpers/dummy_data/dummy_by_id.json'),
              200));
      final result = await productRemoteDataSourceImpl.getProduct(product.id);
      expect(result, equals(product));
    });

    test(
        'should throw a server exception when the response code is 404 or other ',
        () async {
      //arrange

      when(mockHttpClient.get(Uri.parse('${Urls.baseUrl}/${product.id}')))
          .thenAnswer((_) async => http.Response('Not Found', 404));

      //act
      final result = productRemoteDataSourceImpl.getProduct;

      //assert
      expect(() async => await result(product.id),
          throwsA(isA<ServerException>()));
    });
  });

  group('delete product', () {
    test('should throw server exception when response status code is not 200',
        () async {
      //arrange
      when(mockHttpClient.delete(Uri.parse('${Urls.baseUrl}/${product.id}')))
          .thenAnswer((_) async => http.Response('Something went wrong', 404));

      //act
      final result = productRemoteDataSourceImpl.deleteProduct;

      //assert
      expect(() async => await result(product.id),
          throwsA(isA<ServerException>()));
    });

    test(
        'should successfully delete the product when response status code is 200',
        () async {
      when(mockHttpClient.delete(Uri.parse('${Urls.baseUrl}/${product.id}')))
          .thenAnswer((_) async => http.Response('Delete', 200));

      await productRemoteDataSourceImpl.deleteProduct(product.id);

      verify(mockHttpClient.delete(Uri.parse('${Urls.baseUrl}/${product.id}')));
      verifyNoMoreInteractions(mockHttpClient);
    });
  });

  group('update the product', (){

      test('should return the product if updates successfully', () async { 

        when(mockHttpClient.put(Uri.parse('${Urls.baseUrl}/${product.id}'), headers: null,)).thenAnswer((_) async => http.Response('''
            {
              "statusCode": 200,
              "message": "",
              "data": {
                "id": "${product.id}",
                "name": "${product.name}",
                "description": "${product.description}",
                "price": ${product.price},
                "imageUrl": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1/some_image.jpg"
              }
            }
            ''',200));

        final result = await productRemoteDataSourceImpl.updateProduct(id: product.id, name: product.name, price: product.price, description: product.description);

        expect(result, product);

      });

  });

  group('Add Product', (){

      test('should send product data and return without throwing an exception', () async {
    // Arrange   
    

    // Act
    

    // Assert

  });

  } );
  
}
