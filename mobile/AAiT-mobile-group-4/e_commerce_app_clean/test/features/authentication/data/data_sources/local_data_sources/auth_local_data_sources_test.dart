import 'dart:convert';

import 'package:application1/core/error/exception.dart';
import 'package:application1/features/authentication/data/data_sources/local/local_data_source_impl.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../../helper/dummy_data/json_reader.dart';
import '../../../../../helper/test_helper.mocks.dart';

void main() {
  late MockSharedPreferences mockSharedPreferences;
  late AuthLocalDataSourceImpl authLocalDataSourceImpl;
  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    authLocalDataSourceImpl =
        AuthLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  // ignore: constant_identifier_names
  const String KEY = 'CACHED_TOKEN';
  const String cacheProductsPath =
      'helper/dummy_data/auth_model/dummy_cache_token.json';
  group('get token from cache', () {
    test(
      'Should get token from shared preferences',
      () async {
        // arrange
        when(mockSharedPreferences.getString(KEY))
            .thenReturn(readJson(cacheProductsPath));
        // act
        final result = await authLocalDataSourceImpl.getToken();
        final token = json.decode(result)['token'];
        // assert

        expect(token, 'mytoken');
      },
    );

    test(
      'Should recieve a cache exception when there is no token',
      () async {
        // arrange
        when(mockSharedPreferences.getString(KEY)).thenReturn(null);
        // act
        // assert
        expect(()async => await authLocalDataSourceImpl.getToken(), throwsA(isA<CacheException>()));
      },
    );
  });

  group('cache token', () {
    test(
      'Should add token to shared preferences',
      () async {
        // arrange
        final jsonString = readJson(cacheProductsPath);
        when(mockSharedPreferences.setString(KEY, jsonString))
            .thenAnswer((_) async => true);
        // act
        final result =
            await authLocalDataSourceImpl.cacheToken(jsonString);
        // assert
        expect(result, equals(true));
      },
    );
  });
  group('remove token', () {
    test(
      'Should remove token from shared preferences',
      () async {
        // arrange
        when(mockSharedPreferences.remove(KEY)).thenAnswer((_) async => true);
            //act
        final result =
            await authLocalDataSourceImpl.removeToken();
        // assert
        expect(result, equals(true));
      },
    );
  });
}

