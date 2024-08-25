import 'dart:convert';

import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/exceptions/product_exceptions.dart';
import 'package:ecommerce_app/features/product/data/data_resources/local_product_data_source.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockSharedPreferences mockSharedPreferences;
  late LocalProductDataSourceImpl localProductDataSourceImpl;
  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    localProductDataSourceImpl =
        LocalProductDataSourceImpl(mockSharedPreferences);
  });

  group('getProduct test', () {
    test(
        'should return ProductModel if the key of the product exist on shared preferences',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any))
          .thenAnswer((_) => TestingDatas.readJson());

      /// action
      final result =
          await localProductDataSourceImpl.getProduct(TestingDatas.id);

      /// assert
      verify(mockSharedPreferences.getString(TestingDatas.id));

      expect(result, TestingDatas.testDataModel);
    });

    test('Should throw cache exception if no data is found', () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenThrow(CacheException());

      /// action
      final result = localProductDataSourceImpl.getProduct;

      /// assert

      /// verify(mockSharedPreferences.getString(TestingDatas.id));
      expect(() async => await result(TestingDatas.id),
          throwsA(isA<CacheException>()));
    });
  });

  group('getAllProduc test', () {
    test(
        'getProduct should return list of available products on shared preferences',
        () async {
      /// arrange

      when(mockSharedPreferences.getString(any)).thenAnswer((invocation) {
        if (invocation.positionalArguments[0] == AppData.sharedProduct) {
          return TestingDatas.sharedPrefTest;
        } else {
          return TestingDatas.readJson();
        }
      });

      /// action

      final result = await localProductDataSourceImpl.getAllProducts();

      /// assert
      verify(mockSharedPreferences.getString(AppData.sharedProduct));
      // verify(mockSharedPreferences.getString(TestingDatas.id));
      expect(result, TestingDatas.productModelList);
    });

    test('Should throw cache exception if there is no data found', () {
      when(mockSharedPreferences.getString(any)).thenAnswer((_) => null);

      /// action

      final result = localProductDataSourceImpl.getAllProducts();

      verify(mockSharedPreferences.getString(AppData.sharedProduct));

      expect(result, throwsA(isA<CacheException>()));
    });
  });

  group('removeProduct test', () {
    test('Should return true when data is removed', () async {
      /// arrange
      when(mockSharedPreferences.remove(any)).thenAnswer((_) async => true);

      /// action
      final result =
          await localProductDataSourceImpl.removeProduct(TestingDatas.id);

      /// assert

      verify(mockSharedPreferences.remove(TestingDatas.id));
      expect(result, true);
    });

    test('Should return false when the data is not found locally', () async {
      /// arrange
      when(mockSharedPreferences.remove(any)).thenAnswer((_) async => false);

      /// action
      final result = localProductDataSourceImpl.removeProduct;

      /// assert

      expect(
          () async => result(TestingDatas.id), throwsA(isA<CacheException>()));
    });
  });

  group('updateProduct Test', () {
    test(
        'Should return true if the value is updated, and the list update should not be visited',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenAnswer((invocation) {
        if (invocation.positionalArguments[0] == AppData.sharedProduct) {
          return TestingDatas.sharedPrefTest;
        } else {
          return TestingDatas.readJson();
        }
      });
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      /// action
      final result = await localProductDataSourceImpl
          .updateProduct(TestingDatas.testDataModel);

      /// verify
      expect(result, true);
      verify(mockSharedPreferences.getString(any));
      verify(mockSharedPreferences.setString(any, any));
    });

    test('Should also act accordingly when the list doesn\'t present',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenAnswer((invocation) {
        if (invocation.positionalArguments[0] == AppData.sharedProduct) {
          return TestingDatas.sharedPrefTest;
        } else {
          return null;
        }
      });
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      /// action
      final result = await localProductDataSourceImpl
          .updateProduct(TestingDatas.testDataModel);

      /// assert
      expect(result, true);
      verify(mockSharedPreferences.getString(any));
      verify(
        mockSharedPreferences.setString(any, any),
      );
    });

    test('Should return False if one of Them returned False', () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenAnswer((invocation) {
        if (invocation.positionalArguments[0] == AppData.sharedProduct) {
          return TestingDatas.sharedPrefTest;
        } else {
          return null;
        }
      });
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => false);

      /// action
      final result = localProductDataSourceImpl.updateProduct;

      /// assert
      expect(() async => result(TestingDatas.testDataModel),
          throwsA(isA<CacheException>()));
      verify(mockSharedPreferences.getString(TestingDatas.testDataModel.id));
      verify(mockSharedPreferences.getString(AppData.sharedProduct));
    });
  });

  group('AddProduct should work properly', () {
    test('Should return true if product is added ', () async {
      /// arrange
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);
      when(mockSharedPreferences.getString(any)).thenAnswer((_) => null);

      /// action

      final result = await localProductDataSourceImpl
          .addProduct(TestingDatas.testDataModel);

      ///assert
      verify(mockSharedPreferences.setString(TestingDatas.testDataModel.id,
          json.encode(TestingDatas.testDataModel.toJson())));
      verify(mockSharedPreferences.getString(AppData.sharedProduct));
      verify(mockSharedPreferences.setString(AppData.sharedProduct,
          json.encode({TestingDatas.testDataModel.id: 1})));
      expect(result, true);
    });

    test('Should return true if product is added on existing data', () async {
      /// arrange
      when(mockSharedPreferences.setString(TestingDatas.id, any))
          .thenAnswer((_) async => true);

      when(mockSharedPreferences.getString(any))
          .thenAnswer((_) => TestingDatas.sharedPrefTest);

      when(mockSharedPreferences.setString(AppData.sharedProduct, any))
          .thenAnswer((_) async => true);

      /// action
      ///

      final result = await localProductDataSourceImpl
          .addProduct(TestingDatas.testDataModel);

      ///assert
      verify(mockSharedPreferences.setString(TestingDatas.testDataModel.id,
          json.encode(TestingDatas.testDataModel.toJson())));
      verify(mockSharedPreferences.getString(AppData.sharedProduct));
      var res = json.decode(TestingDatas.sharedPrefTest);
      res[TestingDatas.id] = 1;
      verify(mockSharedPreferences.setString(
          AppData.sharedProduct, json.encode(res)));
      expect(result, true);
    });

    test('Should return false if one of the operation failed', () async {
      /// arrange
      when(mockSharedPreferences.setString(TestingDatas.id, any))
          .thenAnswer((_) async => false);

      when(mockSharedPreferences.getString(any))
          .thenAnswer((_) => TestingDatas.sharedPrefTest);

      when(mockSharedPreferences.setString(AppData.sharedProduct, any))
          .thenAnswer((_) async => true);

      /// action
      ///

      final result = localProductDataSourceImpl.addProduct;

      ///assert
      expect(() async => result(TestingDatas.testDataModel),
          throwsA(isA<CacheException>()));
      verify(mockSharedPreferences.setString(TestingDatas.testDataModel.id,
          json.encode(TestingDatas.testDataModel.toJson())));
    });

    /// Error to be fixed later
    // test('should throw exception when error happends', () async {
    //   /// arrange
    //   when(
    //       mockSharedPreferences.setString(TestingDatas.id, any)
    //   ).thenAnswer((_) async =>  true);
    //
    //   when(
    //       mockSharedPreferences.getString(any)
    //   ).thenAnswer((_) => TestingDatas.sharedPrefTest);
    //
    //   when(
    //       mockSharedPreferences.setString(AppData.SHARED_PRODUCTS, any)
    //   ).thenAnswer((_) async => false);
    //   /// action
    //   ///
    //
    //   final result =  localProductDataSourceImpl.addProduct(TestingDatas.testDataModel);
    //
    //
    //   ///assert
    //
    //   expect(result, throwsA(isA<CacheException>()));
    //   verify(mockSharedPreferences.setString(TestingDatas.id, json.encode(TestingDatas.testDataModel.toJson())));
    //   verify(mockSharedPreferences.getString(AppData.SHARED_PRODUCTS));
    //
    // });
  });

  group('addAll product', () {
    test('Should add the models to the sharedpref if there is no data locally',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenAnswer((_) => null);
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      /// action
      final result = await localProductDataSourceImpl
          .addListOfProduct(TestingDatas.productModelList);
      expect(result, true);

      /// assert
      verify(mockSharedPreferences.setString(
          TestingDatas.id, json.encode(TestingDatas.testDataModel.toJson())));
      verify(mockSharedPreferences.setString(
          AppData.sharedProduct, json.encode({TestingDatas.id: 1})));
    });

    test(
        'Should not add to the sharepref when product list already exist though should return True',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any))
          .thenAnswer((_) => TestingDatas.sharedPrefTest);
      when(mockSharedPreferences.setString(TestingDatas.id, any))
          .thenAnswer((_) async => true);

      when(mockSharedPreferences.setString(AppData.sharedProduct, any))
          .thenAnswer((_) async => true);

      /// action
      final result = await localProductDataSourceImpl
          .addListOfProduct(TestingDatas.productModelList);

      /// assert
      expect(result, true);
      //verify(mockSharedPreferences.getString(AppData.sharedProduct));
    });

    test('Should not add all of the data when failures happends in between',
        () async {
      /// arrange
      when(mockSharedPreferences.getString(any)).thenAnswer((_) => null);
      when(mockSharedPreferences.setString(TestingDatas.id, any))
          .thenAnswer((_) async => false);

      when(mockSharedPreferences.setString(AppData.sharedProduct, any))
          .thenAnswer((_) async => true);

      /// action
      final result = localProductDataSourceImpl.addListOfProduct;

      /// assert
      expect(() async => result(TestingDatas.productModelList),
          throwsA(isA<CacheException>()));
      //verify(mockSharedPreferences.getString(AppData.sharedProduct));
      verify(mockSharedPreferences.setString(
          TestingDatas.id, json.encode(TestingDatas.testDataModel.toJson())));
    });
  });
}
