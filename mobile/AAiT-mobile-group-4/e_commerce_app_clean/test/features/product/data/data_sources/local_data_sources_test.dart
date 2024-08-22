import 'dart:convert';

import 'package:application1/core/error/exception.dart';
import 'package:application1/features/product/data/data_sources/local/local_data_source_impl.dart';
import 'package:application1/features/product/data/models/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/json_reader.dart';
import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockSharedPreferences mockSharedPreferences;
  late ProductLocalDataSourceImpl productLocalDataSourceImpl;
  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    productLocalDataSourceImpl =
        ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  List<ProductModel> tProductList = [
    const ProductModel(
      id: '3',
      name: 'airjordan',
      description: 'something you wear',
      price: 566,
      imageUrl: 'https://www.google.com',
    ),
    const ProductModel(
      id: '4',
      name: 'mickey mouse',
      description: 'what you wear',
      price: 536,
      imageUrl: 'pininterest.com',
    )
  ];
  const String cacheProductsPath = 'helper/dummy_data/product_model/dummy_cached_products.json';
  group('get products from cache', () {
    test(
      'Should get products from shared preferences',
      () async {
        // arrange
        when(mockSharedPreferences.getString('PRODUCTS')).thenReturn(
            readJson(cacheProductsPath));
        // act
        final result = await productLocalDataSourceImpl.getProducts();
        // assert
        expect(result, tProductList);
      },
    );
    test(
      'Should recieve a cache exception when there is no data',
      () async {
        // arrange
        when(mockSharedPreferences.getString('PRODUCTS')).thenReturn(null);
        // act
        final result = productLocalDataSourceImpl.getProducts();
        // assert
        expect(result, throwsA(isA<CacheException>()));
      },
    );
  });

  group(
    'add products to cache',
    () {
    test(
      'Should add products to shared preferences',
      () async {
        // arrange
        final jsonString = jsonEncode(tProductList.map((product) {
          return product.toJson();
        }).toList());
        when(mockSharedPreferences.setString('PRODUCTS', jsonString))
            .thenAnswer((_) async => true);
        // act
        final result =
            await productLocalDataSourceImpl.cacheProducts(tProductList);
        // assert
        expect(result, equals(true));
      },
    );
  
  });
}
