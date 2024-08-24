import 'dart:convert';
import 'dart:io';

import 'package:ecom_app/core/constants/constants.dart';
import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/product/data/datasources/product_remote_data_source.dart';
import 'package:ecom_app/features/product/data/models/product_model.dart';
import 'package:ecom_app/features/product/data/models/seller_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockCustomHttpClient mockHttpClient;
  late ProductRemoteDataSourceImpl productRemoteDataSourceImpl;

  setUp(() {
    mockHttpClient = MockCustomHttpClient();
    productRemoteDataSourceImpl =
        ProductRemoteDataSourceImpl(client: mockHttpClient);
  });

  const productId = '6672776eb905525c145fe0bb';
  const jsonCurrent = 'helpers/fixtures/mock_product_api.json';

  const jsonAll = 'helpers/fixtures/mock_products_list_api.json';
  const testProductModel = ProductModel(
      id: '6672776eb905525c145fe0bb',
      name: 'Anime website',
      description: 'Explore anime characters.',
      imageUrl:
          'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg',
      price: 123,
      seller: SellerModel(id: '1', name: 'John', email: 'john@gmail.com'));

  group('get current product', () {
    test('should return product model when the response code is 200', () async {
      //arrange
      when(mockHttpClient.get((Urls.currentProductById(productId))))
          .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 200));

      //act
      final result =
          await productRemoteDataSourceImpl.getCurrentProduct(productId);

      //assert
      expect(result, isA<ProductModel>());
    });

    test('should throw server exception if status code is other than 200',
        () async {
      //arrange
      when(mockHttpClient.get((Urls.currentProductById(productId))))
          .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 400));

      //act and assert
      // verify(mockHttpClient.get(any));
      expect(() => productRemoteDataSourceImpl.getCurrentProduct(productId),
          throwsA(isA<ServerException>()));
    });

    test('should throw a socket exception if it happens', () {
      //arrange
      when(mockHttpClient.get((Urls.currentProductById(productId)))).thenThrow(
          const SocketException(
              'No Internet connection or server unreachable'));

      //act
      final call = productRemoteDataSourceImpl.getCurrentProduct;

      //assert
      expect(() => call(productId), throwsA(isA<SocketException>()));
    });
  });

  group('getAllProducts', () {
    // test('should return a list of product models if status code is 200',
    //     () async {
    //   //arrange
    //   when(mockHttpClient.get(("${Urls.baseUrl3}/products")))
    //       .thenAnswer((_) async => http.Response(readJson(jsonAll), 200));

    //   //act

    //   final result = await productRemoteDataSourceImpl.getAllProducts();

    //   //assert
    //   expect(result, isA<List<ProductModel>>());
    // });

    test('should throw a server exception if status code is different from 200',
        () async {
      //arrange
      when(mockHttpClient.get(('${Urls.baseUrl3}/products')))
          .thenAnswer((_) async => http.Response(readJson(jsonAll), 400));

      //act and assert
      expect(() => productRemoteDataSourceImpl.getAllProducts(),
          throwsA(isA<ServerException>()));
    });

    test('should throw a socket exception if it happens', () {
      //arrange
      when(mockHttpClient.get(('${Urls.baseUrl3}/products'))).thenThrow(
          const SocketException(
              'No Internet connection or server unreachable'));

      //act
      final call = productRemoteDataSourceImpl.getAllProducts;

      //assert
      expect(() => call(), throwsA(isA<SocketException>()));
    });
  });

  group('deleteProduct', () {
    test('should delete successfully', () async {
      //arrange
      final url = (Urls.currentProductById(productId));
      when(mockHttpClient.delete(url))
          .thenAnswer((_) async => http.Response('', 200));

      //act
      await productRemoteDataSourceImpl.deleteProduct(productId);

      //assert
      verify(mockHttpClient.delete(url));
    });

    test('should throw a ServerException when the response code is not 200',
        () async {
      // arrange
      when(mockHttpClient.delete(any))
          .thenAnswer((_) async => http.Response('Something went wrong', 404));

      // act
      final call = productRemoteDataSourceImpl.deleteProduct;

      // assert
      expect(() => call(productId), throwsA(isA<ServerException>()));
    });
    test('should throw a socket exception if it happens', () {
      //arrange
      when(mockHttpClient.delete(any)).thenThrow(const SocketException(
          'No Internet connection or server unreachable'));

      //act
      final call = productRemoteDataSourceImpl.deleteProduct;

      //assert
      expect(() => call(productId), throwsA(isA<SocketException>()));
    });
  });

  group('updateProduct', () {
    test('should return an updated product model if status code is 200',
        () async {
      //arrange
      final jsonBody = jsonEncode({
        'name': testProductModel.name,
        'description': testProductModel.description,
        'price': testProductModel.price,
      });
      when(mockHttpClient.put(
        (Urls.currentProductById(productId)),
        body: jsonBody,
      )).thenAnswer((_) async => http.Response(readJson(jsonCurrent), 200));

      //act
      final result =
          await productRemoteDataSourceImpl.updateProduct(testProductModel);

      //assert
      expect(result, testProductModel);
    });

    test('should throw a ServerException when the response code is not 200',
        () async {
      // arrange
      final jsonBody = jsonEncode({
        'name': testProductModel.name,
        'description': testProductModel.description,
        'price': testProductModel.price,
      });
      when(mockHttpClient.put(
        (Urls.currentProductById(productId)),
        body: jsonBody,
      )).thenAnswer((_) async => http.Response('Something went wrong', 500));

      // act
      final call = await productRemoteDataSourceImpl.updateProduct;

      // assert
      expect(() => call(testProductModel), throwsA(isA<ServerException>()));
    });

    test('should throw a socket exception if it happens', () {
      //arrange
      final jsonBody = jsonEncode({
        'name': testProductModel.name,
        'description': testProductModel.description,
        'price': testProductModel.price,
      });
      when(mockHttpClient.put(
        (Urls.currentProductById(productId)),
        body: jsonBody,
      )).thenThrow(const SocketException(
          'No Internet connection or server unreachable'));

      //act
      final call = productRemoteDataSourceImpl.updateProduct;

      //assert
      expect(() => call(testProductModel), throwsA(isA<SocketException>()));
    });
  });

  // group('createProduct', () {
  //   test('should return a created product model if status code is 201',
  //       () async {
  //     //arrange
  //     final productJson = {
  //       'name': testProductModel.name,
  //       'description': testProductModel.description,
  //       'imageUrl': testProductModel.imageUrl,
  //       'price': testProductModel.price
  //     };
  //     when(mockHttpClient.post((Urls.baseUrl), body: productJson))
  //         .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 201));

  //     //act
  //     final result =
  //         await productRemoteDataSourceImpl.createProduct(testProductModel);

  //     //assert
  //     expect(result, testProductModel);
  //   });
  //   test('should throw a server exception if status code is different from 200',
  //       () async {
  //     //arrange
  //     final productJson = {
  //       'name': testProductModel.name,
  //       'description': testProductModel.description,
  //       'imageUrl': testProductModel.imageUrl,
  //       'price': testProductModel.price
  //     };
  //     when(mockHttpClient.post(Uri.parse(Urls.baseUrl), body: productJson))
  //         .thenAnswer((_) async => http.Response('something went wrong', 500));

  //     //act
  //     final call = await productRemoteDataSourceImpl.createProduct;

  //     //assert
  //     expect(() => call(testProductModel), throwsA(isA<ServerException>()));
  //   });

  //   test('should throw a socket exception if it happens', () {
  //     //arrange
  //     final productJson = {
  //       'name': testProductModel.name,
  //       'description': testProductModel.description,
  //       'imageUrl': testProductModel.imageUrl,
  //       'price': testProductModel.price
  //     };
  //     when(mockHttpClient.post(Uri.parse(Urls.baseUrl), body: productJson))
  //         .thenThrow(const SocketException(
  //             'No Internet connection or server unreachable'));

  //     //act
  //     final call = productRemoteDataSourceImpl.createProduct;

  //     //assert
  //     expect(() => call(testProductModel), throwsA(isA<SocketException>()));
  //   });
  // });

  // group('createProduct', () {
  //   final tProduct = const ProductModel(
  //     id: '1',
  //     name: 'Test Product',
  //     description: 'Test Description',
  //     imageUrl: 'C:/Users/VICTUS 15/Desktop/a2sv-project-phase/clean-architecture/2024-project-phase-mobile-tasks/mobile/ecommerce-tdd/ecommerce_a2sv/test/helpers/images/burger.jpg',
  //     price: 100,
  //   );

  //   test('should send a multipart request with image when image file exists',
  //       () async {

  //     final result = await productRemoteDataSourceImpl.createProduct(tProduct);

  //     // Assert
  //     expect(result, equals(tProduct));
  //   });

  //   test('getting products', () async{
  //     final result = await productRemoteDataSourceImpl.getAllProducts();
  //     expect(result, isA<List<ProductModel>>());  //   });

  //   // test('should throw ImageException when the image file does not exist',
  //   //     () async {
  //   //   // Arrange
  //   //   var fileExists = File(tProduct.imageUrl).existsSync();
  //   //   when(fileExists).thenReturn(false);

  //   //   // Act & Assert
  //   //   expect(
  //   //     () async => await productRemoteDataSourceImpl.createProduct(tProduct),
  //   //     throwsA(isA<ImageException>()),
  //   //   );
  //   // });

  //   // test('should throw ServerException when the response code is not 201',
  //   //     () async {
  //   //   // Arrange
  //   //   var fileExists = File(tProduct.imageUrl).existsSync();
  //   //   when(fileExists).thenReturn(true);

  //   //   when(mockHttpClient.send(any))
  //   //       .thenAnswer((_) async => http.StreamedResponse(
  //   //             Stream.value(utf8.encode('Error')),
  //   //             400,
  //   //           ));

  //   //   // Act & Assert
  //   //   expect(
  //   //     () async => await productRemoteDataSourceImpl.createProduct(tProduct),
  //   //     throwsA(isA<ServerException>()),
  //   //   );
  //   // });

  //   // test('should throw SocketException when there is no internet connection',
  //   //     () async {
  //   //   // Arrange
  //   //   var fileExists = File(tProduct.imageUrl).existsSync();
  //   //   when(fileExists).thenReturn(true);

  //   //   when(mockHttpClient.send(any))
  //   //       .thenThrow(const SocketException('No Internet'));

  //   //   // Act & Assert
  //   //   expect(
  //   //     () async => await productRemoteDataSourceImpl.createProduct(tProduct),
  //   //     throwsA(isA<SocketException>()),
  //   //   );
  //   // });
  // });
}
