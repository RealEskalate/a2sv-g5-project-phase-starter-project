import 'dart:convert';

import 'package:application1/core/constants/constants.dart';
import 'package:application1/features/authentication/domain/entities/user_data.dart';
import 'package:application1/features/product/data/data_sources/remote/remote_data_source_impl.dart';
import 'package:application1/features/product/data/models/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/json_reader.dart';
import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late RemoteDataSourceImpl remoteDataSourceImpl;
  late MockAuthLocalDataSource mockAuthLocalDataSource;
  setUp(() {
    mockHttpClient = MockHttpClient();
    mockAuthLocalDataSource = MockAuthLocalDataSource();
    remoteDataSourceImpl = RemoteDataSourceImpl(
        client: mockHttpClient, authLocalDataSource: mockAuthLocalDataSource);
  });
const String productsResponsePath =
      '/helper/dummy_data/product_response/dummy_products_GET_response.json';
  const String id = '6672776eb905525c145fe0bb';
  const String getProductResponsePath =
      '/helper/dummy_data/product_response/dummy_product_GET_response.json';
  const String deleteProductResponsePath =
      '/helper/dummy_data/product_response/dummy_product_DELETE_response.json';
  const String updateProductResponsePath =
      '/helper/dummy_data/product_response/dummy_product_PUT_response.json';
  
  const tUser = UserEntity(
    email: 'user@gmail.com',
    id: '1',
    name: 'Mr. User',
  );
  const tProductModel = ProductModel(
    id: '1',
    name: 'PC',
    description: 'long description',
    price: 123,
    imageUrl:
        'image.png',
    seller: tUser,
  );
  List<ProductModel> tProductList = [
    const ProductModel(
      id: '1',
      name: 'name1',
      description: 'long description',
      price: 123,
      imageUrl:
          'image.png',
      seller: tUser,
    ),
    const ProductModel(
      id: '2',
      name: 'name2',
      description: 'long description',
      price: 123,
      imageUrl:
          'image.png',
      seller: tUser,
    ),
  ];

  group('get a product', () {
    test('should return a product model if the response is 200', () async {
      //arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls3.getProductId(id)), headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer mytoken',
      })).thenAnswer(
          (_) async => http.Response(readJson(getProductResponsePath), 200));
      //act
      final result = await remoteDataSourceImpl.getProduct(id);
      //assert

      expect(result, tProductModel);
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
      
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls3.getProductId(id)),headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        }))
          .thenAnswer((_) async => http.Response('Not Found', 404));
      //act and assert

      expect(() async {
        await remoteDataSourceImpl.getProduct(id);
      }, throwsA(isA<Exception>()));
    });
  });

  group('get All products', () {
    test('should return a List of products  if the response is 200', () async {
      //arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls3.getProducts),headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        })).thenAnswer(
          (_) async => http.Response(readJson(productsResponsePath), 200));
      //act
      final result = await remoteDataSourceImpl.getProducts();
      //assert

      expect(result, tProductList);
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
       when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.get(Uri.parse(Urls3.getProducts),headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        }))
          .thenAnswer((_) async => http.Response('Not Found', 404));
      //act and assert

      expect(() async {
        await remoteDataSourceImpl.getProducts();
      }, throwsA(isA<Exception>()));
    });
  });

  group('delete a product', () {
    test('should return true if the response is 200', () async {
      //arrange
      
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.delete(Uri.parse(Urls3.deleteProductId(id)),headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        }))
          .thenAnswer((_) async =>
              http.Response(readJson(deleteProductResponsePath), 200));
      //act
      final result = await remoteDataSourceImpl.deleteProduct(id);
      //assert
      expect(result, true);
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      
      when(mockHttpClient.delete(Uri.parse(Urls3.deleteProductId(id)),headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        }))
          .thenAnswer((_) async => http.Response('Operation Failed', 404));
      //act
      //assert
      expect(() async {
        await remoteDataSourceImpl.deleteProduct(id);
      }, throwsA(isA<Exception>()));
    });
  });

  group('update a product', () {
    final jsonBody = jsonEncode({
      'name': tProductModel.name,
      'description': tProductModel.description,
      'price': tProductModel.price,
    });
    test('should return the updated product if the response is 200', () async {
      //arrange
       when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.put(
        Uri.parse(Urls3.updateProductId(tProductModel.id)),
        body: jsonBody,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        },
      )).thenAnswer(
          (_) async => http.Response(readJson(updateProductResponsePath), 200));
      //act
      final result = await remoteDataSourceImpl.updateProduct(tProductModel);
      //assert
      expect(result, tProductModel);
    });

    test('Should throw a server exception if the response is 404', () async {
      //arrange
       when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => 'mytoken');
      when(mockHttpClient.put(
        Uri.parse(Urls3.updateProductId(tProductModel.id)),
        body: jsonBody,
        headers: {
          'Content-Type': 'application/json',
          'Authorization': 'Bearer mytoken',
        },
      )).thenAnswer((_) async => http.Response('Operation Failed', 404));
      //act
      //assert
      expect(() async {
        await remoteDataSourceImpl.updateProduct(tProductModel);
      }, throwsA(isA<Exception>()));
    });
  });
}
