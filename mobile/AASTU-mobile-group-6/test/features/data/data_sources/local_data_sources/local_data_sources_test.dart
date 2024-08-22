import 'dart:convert';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/models/product_models.dart';
import 'package:ecommerce_app_ca_tdd/features/product/data/data_sources/local_data_source/local_data_source.dart';

class MockSharedPreferences extends Mock implements SharedPreferences {}
void main() {
  late ProductLocalDataSourceImpl dataSource;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    dataSource = ProductLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('getProduct', () {
    const productId = '1';
    final productModel = ProductModel(
      id: '1',
      name: 'Test Product',
      description: 'Test Description',
      price: 10,
      imagePath: "",
    );
    final jsonString = '{"id": "1", "name": "Test Product", "description": "Test Description", "price": 10}';

    test('should return ProductModel when there is one in the cache', () async {
      // Arrange
      when(mockSharedPreferences.getString('PRODUCT_$productId'))
          .thenReturn(jsonString);

      // Act
      final result = await dataSource.getProduct(productId);

      // Assert
      expect(result, equals(productModel));
    });

    test('should throw an exception when there is nothing in the cache', () async {
      // Arrange
      when(mockSharedPreferences.getString('PRODUCT_$productId')).thenReturn(null);

      // Act
      final call = dataSource.getProduct;

      // Assert
      expect(() => call(productId), throwsException);
    });
  });

  group('addProduct', () {
    final productModel = ProductModel(
      id: '1',
      name: 'Test Product',
      description: 'Test Description',
      price: 10,
      imagePath: "",
    );

    test('should call SharedPreferences to cache the data', () async {
      // Arrange
      final jsonString = jsonEncode(productModel.toJson());

      // Act
      await dataSource.addProduct(productModel);

      // Assert
      verify(mockSharedPreferences.setString('PRODUCT_${productModel.id}', jsonString));
    });
  });

  // Similar test cases for deleteProduct, updateProduct, and getProducts...
}
