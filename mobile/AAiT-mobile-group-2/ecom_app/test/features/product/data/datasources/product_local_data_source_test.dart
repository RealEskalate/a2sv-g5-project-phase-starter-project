import 'dart:convert';

import 'package:ecom_app/core/error/exception.dart';
import 'package:ecom_app/features/product/data/datasources/product_local_data_source.dart';
import 'package:ecom_app/features/product/data/models/product_model.dart';
import 'package:ecom_app/features/product/data/models/seller_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late ProductLocalDataSourceImpl localDataSourceImpl;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    localDataSourceImpl =
        ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('getAllProducts', () {
    final jsonList =
        json.decode(readJson('helpers/fixtures/dummy_products_cached.json'));
    final testProductModelList = ProductModel.fromJsonList(jsonList);
    const CACHED_PRODUCTS = 'CACHED_PRODUCTS';
    test(
        'should return cached products from shared reference when there is one in the cache',
        () async {
      //arrange
      when(mockSharedPreferences.getString(any))
          .thenReturn(json.encode(jsonList));

      //act
      final result = await localDataSourceImpl.getAllProducts();

      //assert
      verify(mockSharedPreferences.getString(CACHED_PRODUCTS));
      expect(result, equals(testProductModelList));
    });
    test('should throw cache exception when there is not a cached value',
        () async {
      //arrange
      when(mockSharedPreferences.getString(any)).thenReturn(null);

      //act
      final call = localDataSourceImpl.getAllProducts;

      //assert
      expect(() => call(), throwsA(const TypeMatcher<CacheException>()));
    });
  });

  group('cacheAllProducts', () {
    const CACHED_PRODUCTS = 'CACHED_PRODUCTS';
    const testProductModelList = [
      ProductModel(
          id: '1',
          name: 'Test Pineapple',
          description: 'A yellow pineapple for the summer',
          imageUrl: 'pineapple.jpg',
          price: 5.33,
          seller: SellerModel(id: '1', name: 'John', email: 'john@gmail.com'))
    ];

    test('should call SharedPreference to cache the data', () {
      //arrange
      final expectedJson =
          json.encode(ProductModel.toJsonList(testProductModelList));
      when(mockSharedPreferences.setString(any, expectedJson))
          .thenAnswer((_) async => true);

      //act
      localDataSourceImpl.cacheAllProducts(testProductModelList);

      //assert

      verify(mockSharedPreferences.setString(CACHED_PRODUCTS, expectedJson));
    });
  });
}
