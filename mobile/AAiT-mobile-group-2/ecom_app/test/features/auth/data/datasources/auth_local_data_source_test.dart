import 'dart:convert';

import 'package:dartz/dartz.dart';
import 'package:ecom_app/features/auth/data/datasources/auth_local_data_source.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late AuthLocalDataSourceImpl authLocalDataSourceImpl;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    authLocalDataSourceImpl =
        AuthLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('AuthLocalDataSource', () {
    final tToken = '1234';
    test('getToken', () async {
      //arrange
      when(mockSharedPreferences.getString(any))
          .thenReturn(json.encode(tToken));

      //act
      final result = await authLocalDataSourceImpl.getToken();

      //assert
      expect(result, '1234');
    });

    test('cacheToken', () async {
      //arrange
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      //act
      final result = await authLocalDataSourceImpl.cacheToken(tToken);

      //assert
      expect(result, unit);
    });

    test('logout', () async {
      //arrange
      when(mockSharedPreferences.remove(any)).thenAnswer((_) async => true);

      //act
      final result = await authLocalDataSourceImpl.logout();

      //assert
      expect(result, unit);
    });
  });
}
