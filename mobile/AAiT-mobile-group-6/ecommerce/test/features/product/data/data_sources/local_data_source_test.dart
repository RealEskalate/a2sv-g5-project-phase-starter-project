import 'dart:convert';

import 'package:ecommerce/core/error/exceptions.dart';
import 'package:ecommerce/features/product/data/datasources/local_data_resource.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late ProductLocalDataSourceImpl dataSource;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    dataSource =
        ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('get last product', () {
    final testProductModel = ProductModel.fromJson(
        json.decode(readJson('dummy_product_cached.json')));

    test('should return product when there is one in shared preferences',
        () async {
      when(mockSharedPreferences.getString(any))
          .thenReturn(readJson('dummy_product_cached.json'));

      final result = await dataSource.getLastProduct();
      verify(mockSharedPreferences.getString('cachedProduct'));
      expect(result, testProductModel);
    });

    test('should throw a CacheException when there is not a cached value', () {
      // arrange
      when(mockSharedPreferences.getString(any)).thenReturn(null);
      // act
      final call = dataSource.getLastProduct;
      // assert

      expect(() => call(), throwsA(TypeMatcher<CacheException>()));
    });
  });

  group('cache product', (){
    final testProductModel = const ProductModel(
      id: '1',
      name: 'name',
      description: 'description',
      price: 1.0,
      imageUrl: 'imageUrl',

    );

    test ('should call sharedPreferences to cache the data', () async {
      // arrange
      when(mockSharedPreferences.setString(any, any)).thenAnswer((_) async => true);
      // act
      dataSource.cacheProduct(testProductModel);
      // assert
      final expectedJsonString = json.encode(testProductModel.toJson());
      verify(mockSharedPreferences.setString('cachedProduct', expectedJsonString));
    });
  });
}
