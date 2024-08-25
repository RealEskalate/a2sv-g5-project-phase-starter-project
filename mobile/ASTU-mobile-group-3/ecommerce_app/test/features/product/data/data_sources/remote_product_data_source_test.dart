import 'dart:convert';
import 'dart:io';

import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/exceptions/product_exceptions.dart';
import 'package:ecommerce_app/features/auth/data/model/token_model.dart';
import 'package:ecommerce_app/features/product/data/data_resources/remote_product_data_source.dart';
import 'package:ecommerce_app/features/product/data/models/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
// ignore: depend_on_referenced_packages
import 'package:http_parser/http_parser.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late RemoteProductDataSourceImp remoteProductDataSourceImp;
  late MockAuthLocalDataSource mockAuthLocalDataSource;

  List<ProductModel> expectedListOfModel = <ProductModel>[];
  Map<String, dynamic> finalResult = json.decode(TestingDatas.readProductV3());
  for (Map<String, dynamic> jsonModel in finalResult['data']) {
    expectedListOfModel.add(ProductModel.fromJson(jsonModel));
  }
  ProductModel expectedSingleProduct =
      ProductModel.fromJson(json.decode(TestingDatas.readJson()));
  setUp(() {
    mockHttpClient = MockHttpClient();
    mockAuthLocalDataSource = MockAuthLocalDataSource();
    remoteProductDataSourceImp =
        RemoteProductDataSourceImp(mockHttpClient, mockAuthLocalDataSource);
  });

  group('testing the getAllProducts', () {
    test('Should list prouct of model when success', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(Uri.parse(AppData.allProductUrl),
              headers: anyNamed('headers')))
          .thenAnswer((_) async =>
              http.Response(TestingDatas.getAllProductResponce(), 200));

      /// action
      final result = await remoteProductDataSourceImp.getAllProducts();
      verify(mockHttpClient.get(any, headers: anyNamed('headers')));

      expect(result, expectedListOfModel);
    });

    test('Should return ServerException  when socket exception is thrown', () {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(Uri.parse(AppData.allProductUrl),
              headers: anyNamed('headers')))
          .thenThrow(const SocketException('Failed'));

      /// action
      final result = remoteProductDataSourceImp.getAllProducts;

      /// assert
      expect(() async => result(), throwsA(isA<ServerException>()));
      //verify(mockHttpClient.get(any, headers: anyNamed('headers')));
    });

    test('Should return ServerException when status code is not 200', () {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(Uri.parse(AppData.allProductUrl),
              headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('Not Found', 404));

      /// action
      ///
      final result = remoteProductDataSourceImp.getAllProducts;

      expect(() async => result(), throwsA(isA<ServerException>()));

      //verify(mockHttpClient.get(any, headers: anyNamed('headers')));
    });
  });

  group('This is to test getProduct ', () {
    test('Should return product model when single data with id is required',
        () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(
              Uri.parse('${AppData.allProductUrl}/${TestingDatas.id}'),
              headers: anyNamed('headers')))
          .thenAnswer(
              (_) async => http.Response(TestingDatas.getSingleProduct(), 200));

      /// action
      ///
      final result =
          await remoteProductDataSourceImp.getProduct(TestingDatas.id);

      /// assert
      expect(result, expectedSingleProduct);
      verify(mockHttpClient.get(any, headers: anyNamed('headers')));
    });

    test('Should throw ServerExceptoiin when query failed', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(
              Uri.parse('${AppData.allProductUrl}/${TestingDatas.apiId}'),
              headers: anyNamed('headers')))
          .thenThrow(const SocketException('Failed'));

      /// action
      ///
      final result = remoteProductDataSourceImp.getProduct;

      /// assert
      expect(() async => result(TestingDatas.apiId),
          throwsA(isA<ServerException>()));
      //verify(mockHttpClient.get(any, headers: anyNamed('headers')));
    });

    test('Should return Server exception when status code is not 200',
        () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.get(
              Uri.parse('${AppData.allProductUrl}/${TestingDatas.apiId}'),
              headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('Not found', 404));

      /// action
      ///
      final result = remoteProductDataSourceImp.getProduct(TestingDatas.apiId);

      /// assert
      expect(result, throwsA(isA<ServerException>()));
      //verify(mockHttpClient.get(any, headers: anyNamed('headers')));
    });
  });

  group('insertList', () {
    test('Should success', () async {
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      const fakePath = '/home/chera/Downloads/certificate.jpg';

      final uri = Uri.parse(AppData.baseUrl);
      final request = http.MultipartRequest('POST', uri);
      final myTest = ProductModel(
          id: TestingDatas.id,
          name: TestingDatas.testDataModel.name,
          description: TestingDatas.testDataModel.description,
          price: TestingDatas.testDataModel.price,
          imageUrl: fakePath,
          seller: AuthData.userModel);
      request.files.add(
        await http.MultipartFile.fromPath(
          'image',
          fakePath,
          contentType: MediaType('image', 'png'),
        ),
      );
      request.fields['name'] = TestingDatas.testDataModel.name;
      request.fields['description'] = TestingDatas.testDataModel.description;
      request.fields['price'] = TestingDatas.testDataModel.price.toString();

      const fileContent = 'file content';

      final expecterStream =
          Stream<List<int>>.fromIterable([utf8.encode(fileContent)]);
      final streamedResponse = http.StreamedResponse(expecterStream, 201);

      /// arrange
      when(mockHttpClient.send(any)).thenAnswer((_) async => streamedResponse);

      /// action
      final result = await remoteProductDataSourceImp.insertProduct(myTest);

      /// assert
      expect(result, AppData.successInsert);
      verify(mockHttpClient.send(any));
    });

    test('Should success', () async {
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      const fakePath = '/home/chera/Downloads/certificate.jpg';

      final uri = Uri.parse(AppData.baseUrl);
      final request = http.MultipartRequest('POST', uri);
      final myTest = ProductModel(
          id: TestingDatas.id,
          name: TestingDatas.testDataModel.name,
          description: TestingDatas.testDataModel.description,
          price: TestingDatas.testDataModel.price,
          imageUrl: fakePath,
          seller: AuthData.sellerData);
      request.files.add(
        await http.MultipartFile.fromPath(
          'image',
          fakePath,
          contentType: MediaType('image', 'png'),
        ),
      );
      request.fields['name'] = TestingDatas.testDataModel.name;
      request.fields['description'] = TestingDatas.testDataModel.description;
      request.fields['price'] = TestingDatas.testDataModel.price.toString();

      /// arrange
      when(mockHttpClient.send(any)).thenThrow(ServerException());

      /// action
      final result = remoteProductDataSourceImp.insertProduct;

      /// assert
      expect(() async => result(myTest), throwsA(isA<ServerException>()));
      //verify(mockHttpClient.send(any));
    });
  });

  group('updateProducr', () {
    test('Should update the product if all value are valid', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.put(
              Uri.parse('${AppData.allProductUrl}/${TestingDatas.id}'),
              body: anyNamed('body'),
              headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('', 200));

      /// action
      final result = await remoteProductDataSourceImp
          .updateProduct(TestingDatas.testDataModel);

      /// assert
      expect(result, AppData.successUpdate);
      // verify(mockHttpClient.put(any,
      //     body: {
      //       'name': TestingDatas.testDataModel.name,
      //       'description': TestingDatas.testDataModel.description,
      //       'price': '${TestingDatas.testDataModel.price}',
      //     },
      //     headers: anyNamed('headers')));
    });

    test('Should throw server exception when request fails, socket exception',
        () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.put(any,
              body: anyNamed('body'), headers: anyNamed('headers')))
          .thenThrow(const SocketException('Failed'));

      /// action
      final result = remoteProductDataSourceImp.updateProduct;

      /// assert
      expect(() async => result(TestingDatas.testDataModel),
          throwsA(isA<ServerException>()));
      // verify(mockHttpClient.put(any,
      //     body: {
      //       'name': TestingDatas.testDataModel.name,
      //       'description': TestingDatas.testDataModel.description,
      //       'price': '${TestingDatas.testDataModel.price}',
      //     },
      //     headers: anyNamed('headers')));
    });

    test('Should throw server exception when status code is not 200', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.put(any,
              body: anyNamed('body'), headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('Not found', 404));

      /// action
      final result = remoteProductDataSourceImp.updateProduct;

      /// assert
      expect(() async => result(TestingDatas.testDataModel),
          throwsA(isA<ServerException>()));
      //   verify(mockHttpClient.put(any,
      //       body: {
      //         'name': TestingDatas.testDataModel.name,
      //         'description': TestingDatas.testDataModel.description,
      //         'price': '${TestingDatas.testDataModel.price}',
      //       },
      //       headers: anyNamed('headers')));
    });
  });

  group('deleteProduct', () {
    test('Should delete the product if all value is valid', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.delete(
              Uri.parse('${AppData.allProductUrl}/${TestingDatas.id}'),
              headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('', 200));

      /// action
      final result =
          await remoteProductDataSourceImp.deleteProduct(TestingDatas.id);

      /// assert
      expect(result, AppData.successDelete);
      verify(mockHttpClient.delete(any, headers: anyNamed('headers')));
    });

    test('Should throw server exception when request fails, socket exception',
        () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.delete(any, headers: anyNamed('headers')))
          .thenThrow(const SocketException('Failed'));

      /// action
      final result = remoteProductDataSourceImp.deleteProduct;

      /// assert
      expect(
          () async => result(TestingDatas.id), throwsA(isA<ServerException>()));
      //verify(mockHttpClient.delete(any, headers: anyNamed('headers')));
    });

    test('Should throw server exception when status code is not 200', () async {
      /// arrange
      when(mockAuthLocalDataSource.getToken())
          .thenAnswer((_) async => const TokenModel(token: AuthData.token));
      when(mockHttpClient.delete(any, headers: anyNamed('headers')))
          .thenAnswer((_) async => http.Response('Not found', 404));

      /// action
      final result = remoteProductDataSourceImp.deleteProduct;

      /// assert
      expect(
          () async => result(TestingDatas.id), throwsA(isA<ServerException>()));
      //verify(mockHttpClient.delete(any, headers: anyNamed('headers')));
    });
  });
}
