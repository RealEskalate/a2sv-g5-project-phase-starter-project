import 'package:ecommerce_app/core/constants/constants.dart';
import 'package:ecommerce_app/core/errors/exceptions/product_exceptions.dart';
import 'package:ecommerce_app/features/auth/data/data_source/auth_local_data_source.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/auth_test_data/testing_data.dart';
import '../../../../test_helper/test_helper_generation.mocks.dart';

void main() {
  late AuthLocalDataSourceImpl authLocalDataSource;
  late MockSharedPreferences mockSharedPreferences;
  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    authLocalDataSource =
        AuthLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('Add Token', () {
    test('Should return true when saved', () async {
      when(mockSharedPreferences.setString(AppData.tokenPlacement, any))
          .thenAnswer((_) async => true);

      final result = await authLocalDataSource.saveToken(AuthData.tokenModel);

      expect(result, true);
    });

    test('Should throw exception if not saved', () async {
      when(mockSharedPreferences.setString(AppData.tokenPlacement, any))
          .thenThrow(CacheException());

      final result = authLocalDataSource.saveToken;

      expect(() async => result(AuthData.tokenModel),
          throwsA(isA<CacheException>()));
    });
  });
  group('Get token', () {
    test('Should return true when saved', () async {
      when(mockSharedPreferences.getString(AppData.tokenPlacement))
          .thenReturn(AuthData.tokenModel.token);

      final result = await authLocalDataSource.getToken();

      expect(result, AuthData.tokenModel);
    });

    test('Should throw exception', () async {
      when(mockSharedPreferences.getString(AppData.tokenPlacement))
          .thenThrow(CacheException());

      final result = authLocalDataSource.getToken;

      expect(() async => result(), throwsA(isA<CacheException>()));
    });
  });

  group('log out ', () {
    test('Should return true when the data is removed succesfully', () async {
      /// arrange
      when(mockSharedPreferences.remove(any)).thenAnswer((_) async => true);

      /// action
      final result = await authLocalDataSource.clearToken();

      /// assert
      expect(result, true);
    });

    test('Should return cache exception when exception is thrown', () async {
      /// arrange
      when(mockSharedPreferences.remove(any)).thenThrow(CacheException());

      /// action
      final result = authLocalDataSource.clearToken;

      /// assert
      expect(() async => result(), throwsA(isA<CacheException>()));
    });
  });
}
