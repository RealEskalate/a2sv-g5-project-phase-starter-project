import 'dart:convert';


import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/exception/exception.dart';
import 'package:product_8/features/product/data/data_source/local_data_source/product_local_data_source.dart';
import 'package:product_8/features/product/data/models/product_model.dart';

import '../../../../helpers/jason_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';



void main() {
  late ProductLocalDataSourceImpl productLocalDataSourceImpl;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    productLocalDataSourceImpl =
        ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('cacheProducts', () {
    // ignore: constant_identifier_names
    const CACHED_PRODUCTS = 'CACHED_PRODUCTS';
    const testProductModelList = [
      ProductModel(
          id: '1',
          name: 'Test Pineapple',
          description: 'A yellow pineapple for the summer',
          imageUrl: 'pineapple.jpg',
          price: 5.33)
    ];
    test('should call SharedPreference to cache the data', () {
      //arrange
      final expectedJson =
          json.encode(ProductModel.toJsonList(testProductModelList));
      when(mockSharedPreferences.setString(any, expectedJson))
          .thenAnswer((_) async => Future.value(true));

      //act
      productLocalDataSourceImpl.cacheProducts(testProductModelList);

      //assert

      verify(mockSharedPreferences.setString(CACHED_PRODUCTS, expectedJson));
    });
  });


  group('getProducts', () {
    
    // ignore: constant_identifier_names
    const CACHED_PRODUCTS = 'CACHED_PRODUCTS';
    final jsonList = json.decode(readJson('helpers/dummy_data/dummy_products_cached.json'));
    final testProductModelList = ProductModel.fromJsonList(jsonList);
    test('should return a list of ProductModel when there is one in the cache', () async {
      //arrange
   
      when(mockSharedPreferences.getString(any))
          .thenReturn(json.encode(jsonList));

      //act
      final result = await productLocalDataSourceImpl.getProducts();

      //assert
      verify(mockSharedPreferences.getString(CACHED_PRODUCTS));
      expect(result, equals(testProductModelList));
    });

    test('should throw a CacheException when there is no data in the cache', () async {
      //arrange
      when(mockSharedPreferences.getString(any)).thenReturn(null);

      //act
      final call = productLocalDataSourceImpl.getProducts;

      //assert
      expect(() => call(), throwsA(const TypeMatcher<CacheException>()));
    });
  });
}
