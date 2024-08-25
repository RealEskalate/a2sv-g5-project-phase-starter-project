import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';
import 'package:product_6/core/constants/constants.dart';
import 'package:product_6/core/error/exception.dart';
import 'package:product_6/features/auth/domain/entities/user_entity.dart';
import 'package:product_6/features/product/data/data_sources/remote_data_source.dart';
import 'package:product_6/features/product/data/models/product_model.dart';
import 'package:product_6/features/product/domain/entities/product_entity.dart';

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late MockProductLocalDataSource mockProductLocalDataSource;
  late ProductRemoteDataSourceImpl productRemoteDataSourceImpl;

  setUp(() {
    mockHttpClient = MockHttpClient();

    mockProductLocalDataSource = MockProductLocalDataSource();
    productRemoteDataSourceImpl = ProductRemoteDataSourceImpl(
      client: mockHttpClient,
      productLocalDataSource: mockProductLocalDataSource,
    );
  });

  const testId = '1';
  const testToken = 'test_token';
  const testProductModel = ProductModel(
      id: '1',
      name: 'Test Product',
      description: 'Description',
      price: 100.0,
      imageUrl:
          'C:/Users/bb/pra/2024-internship-mobile-tasks/on-boarding/product_6/lib/assets/images/profile.jpg',
      seller: UserEntity.empty);

  final header = {
    'Content-Type': 'application/json',
    'Authorization': 'Bearer $testToken',
  };

  group(
    'getAllProducts',
    () {
      test(
          'should return a list of ProductModels when the response code is 200',
          () async {
        // arrange
        when(mockProductLocalDataSource.getToken())
            .thenAnswer((_) async => testToken);
        when(mockHttpClient.get(Uri.parse(Urls.getAllProducts()),
                headers: header))
            .thenAnswer(
          (_) async => http.Response(
            readJson('helpers/dummy_data/dummy_products_response.json'),
            200,
          ),
        );

        when(mockProductLocalDataSource.cacheProducts(any))
            .thenAnswer((_) async => true);
        // print(readJson('helpers/dummy_data/dummy_products_response.json'));
        // act
        final result = await productRemoteDataSourceImpl.getAllProducts();

        // assert
        expect(result, isA<List<ProductModel>>());
      });

      group('getProductById', () {
        test('should return a  ProductModel when the response code is 200',
            () async {
          // arrange
          when(mockProductLocalDataSource.getToken())
              .thenAnswer((_) async => testToken);
          when(mockHttpClient.get(Uri.parse(Urls.getProductById(testId)),
                  headers: header))
              .thenAnswer(
            (_) async => http.Response(
              readJson('helpers/dummy_data/dummy_product_response.json'),
              200,
            ),
          );

          // act
          final result =
              await productRemoteDataSourceImpl.getProductById(testId);

          // assert
          expect(result, isA<ProductModel>());
        });

        test('should throw a ServerException when the response code is not 200',
            () async {
          // arrange
          when(mockProductLocalDataSource.getToken())
              .thenAnswer((_) async => testToken);
          when(mockHttpClient.get(
            Uri.parse(Urls.getProductById(testId)),
            headers: header,
          )).thenAnswer(
            (_) async => http.Response(
              'Something went wrong',
              404,
            ),
          );

          // act
          final call = productRemoteDataSourceImpl.getProductById;

          // assert
          expect(() => call(testId), throwsA(isA<ServerException>()));
        });
      });
      group('updateProduct', () {
        final body = jsonEncode({
          'name': 'Test Product',
          'description': 'Description',
          'price': 100.0,
        });

        test('should return true when the product is successfully updated',
            () async {
          // arrange
          when(mockProductLocalDataSource.getToken())
              .thenAnswer((_) async => testToken);

          when(mockHttpClient.put(
            Uri.parse(Urls.getProductById(testId)),
            headers: header,
            body: anyNamed('body'),
          )).thenAnswer((_) async => http.Response('', 200));

          // act
          final result =
              await productRemoteDataSourceImpl.updateProduct(testProductModel);

          // assert
          expect(result, true);
        });

        test('should return false when the product update fails', () async {
          // arrange
          when(mockProductLocalDataSource.getToken())
              .thenAnswer((_) async => testToken);

          when(mockHttpClient.put(
            Uri.parse(Urls.getProductById(testId)),
            headers: anyNamed('headers'),
            body: body,
          )).thenAnswer((_) async => http.Response('', 400));

          // act
          final result =
              await productRemoteDataSourceImpl.updateProduct(testProductModel);

          // assert
          expect(result, false);
        });
      });

      group(
        'deleteProduct',
        () {
          test('should return true when the product is successfully deleted',
              () async {
            // arrange
            when(mockProductLocalDataSource.getToken())
                .thenAnswer((_) async => testToken);
            when(mockHttpClient.delete(
              Uri.parse(Urls.getProductById(testId)),
              headers: {
                'Content-Type': 'application/json',
                'Authorization': 'Bearer $testToken',
              },
            )).thenAnswer(
              (_) async => http.Response('', 200),
            );

            // act
            final result =
                await productRemoteDataSourceImpl.deleteProduct(testId);

            // assert
            expect(result, true);
          });

          test(
            'should throw a ServerException when the product deletion fails',
            () async {
              // arrange
              when(mockProductLocalDataSource.getToken())
                  .thenAnswer((_) async => testToken);
              when(mockHttpClient.delete(
                Uri.parse(Urls.getProductById(testId)),
                headers: {
                  'Content-Type': 'application/json',
                  'Authorization': 'Bearer $testToken',
                },
              )).thenAnswer(
                (_) async => http.Response('', 404),
              );

              // act
              final call = productRemoteDataSourceImpl.deleteProduct;

              // assert
              expect(() => call(testId), throwsA(isA<ServerException>()));
            },
          );
        },
      );
    },
  );
}
