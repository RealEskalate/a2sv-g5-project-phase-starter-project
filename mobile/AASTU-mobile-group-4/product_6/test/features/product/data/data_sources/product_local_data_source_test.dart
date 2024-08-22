import 'dart:convert';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/data/data_sources/product_local_data_source.dart';
import 'package:product_6/features/product/data/models/product_model.dart';

import '../../../../mock.mocks.dart';

// Mocking SharedPreferences using Mockito
// class MockSharedPreferences extends Mock implements SharedPreferences {}

void main() {
  late ProductLocalDataSourceImpl dataSource;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    dataSource = ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  const tProductModel = ProductModel(id: '1', name: 'Test Product', price: 9.99, description: '', imageUrl: '');
  final tProductList = [tProductModel];

  void setUpMockSharedPreferencesGetStringSuccess() {
    when(mockSharedPreferences.getString(any))
        .thenReturn(json.encode(tProductList.map((product) => product.toJson()).toList()));
  }

  group('getProducts', () {
    test(
      'should return List<ProductModel> when there is data in SharedPreferences',
      () async {
        // arrange
        setUpMockSharedPreferencesGetStringSuccess();
        // act
        final result = await dataSource.getProducts();
        // assert
        verify(mockSharedPreferences.getString(productDataKey));
        expect(result, equals(tProductList));
      },
    );

    test(
      'should throw CacheException when there is no data in SharedPreferences',
      () async {
        // arrange
        when(mockSharedPreferences.getString(any)).thenReturn(null);
        // act
        final call = dataSource.getProducts;
        // assert
        expect(() => call(), throwsA(const TypeMatcher<CacheException>()));
      },
    );
  });

  group('getProductById', () {
    test(
      'should return ProductModel when the product is found',
      () async {
        // arrange
        setUpMockSharedPreferencesGetStringSuccess();
        // act
        final result = await dataSource.getProductById('1');
        // assert
        verify(mockSharedPreferences.getString(productDataKey));
        expect(result, equals(tProductModel));
      },
    );

    test(
      'should throw CacheException when the product is not found',
      () async {
        // arrange
        setUpMockSharedPreferencesGetStringSuccess();
        // act
        final call = dataSource.getProductById;
        // assert
        expect(() => call('2'), throwsA(const TypeMatcher<CacheException>()));
      },
    );
  });

  group('createProduct', () {
    test(
  'should add a new product to SharedPreferences',
  () async {
    // arrange
    when(mockSharedPreferences.getString(any)).thenReturn(null);  // Initially, no products exist
    when(mockSharedPreferences.setString(any, any)).thenAnswer((_) async => true);

    // act
    final result = await dataSource.createProduct(tProductModel);

    // assert
    final expectedJsonString = json.encode([tProductModel.toJson()]);
    verify(mockSharedPreferences.setString(productDataKey, expectedJsonString)).called(1);
    expect(result, true);
  },
);

  });

  group('updateProduct', () {
    test(
      'should update an existing product in SharedPreferences',
      () async {
        // arrange
        setUpMockSharedPreferencesGetStringSuccess();
        when(mockSharedPreferences.setString(any, any)).thenAnswer((_) async => true);
        // act
        final result = await dataSource.updateProduct(tProductModel);
        // assert
        final expectedJsonString = json.encode(tProductList.map((product) => product.toJson()).toList());
        verify(mockSharedPreferences.setString(productDataKey, expectedJsonString));
        expect(result, true);
      },
    );

    test(
      'should throw CacheException when the product is not found',
      () async {
        // arrange
        when(mockSharedPreferences.getString(any)).thenReturn(json.encode([]));
        // act
        final call = dataSource.updateProduct;
        // assert
        expect(() => call(tProductModel), throwsA(const TypeMatcher<CacheException>()));
      },
    );
  });

  group('deleteProduct', () {
    test(
      'should remove a product from SharedPreferences',
      () async {
        // arrange
        setUpMockSharedPreferencesGetStringSuccess();
        when(mockSharedPreferences.setString(any, any)).thenAnswer((_) async => true);
        // act
        final result = await dataSource.deleteProduct('1');
        // assert
        final expectedJsonString = json.encode([]);
        verify(mockSharedPreferences.setString(productDataKey, expectedJsonString));
        expect(result, true);
      },
    );

    test(
      'should throw CacheException when no products are found',
      () async {
        // arrange
        when(mockSharedPreferences.getString(any)).thenReturn(null);
        // act
        final call = dataSource.deleteProduct;
        // assert
        expect(() => call('1'), throwsA(const TypeMatcher<CacheException>()));
      },
    );
  });
}








