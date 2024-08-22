import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/exception/exception.dart';
import 'package:product_8/features/auth/data/data_source/auth_local_data_source.dart';

import '../../../../helpers/test_helper.mocks.dart';


void main() {
  late AuthLocalDataSourceImpl localDataSource;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    localDataSource = AuthLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('cacheToken', () {
    const token = 'someToken123';

    test('should call SharedPreferences to cache the token', () async {
      // arrange
      when(mockSharedPreferences.setString(any, any))
          .thenAnswer((_) async => true);

      // act
      await localDataSource.cacheToken(token);

      // assert
      verify(mockSharedPreferences.setString('token', token));
    });
  });

  group('getToken', () {
    test('should return token from SharedPreferences when there is one in the cache', () async {
      // arrange
      when(mockSharedPreferences.getString(any)).thenReturn('someToken123');

      // act
      final result = await localDataSource.getToken();

      // assert
      expect(result, equals('someToken123'));
    });

    test('should throw CacheException when there is no cached token', () async {
      // arrange
      when(mockSharedPreferences.getString(any)).thenReturn(null);

      // act
      final call = localDataSource.getToken;

      // assert
      expect(() => call(), throwsA(isA<CacheException>()));
    });
  });

  group('deleteToken', () {
    test('should call SharedPreferences to remove the token', () async {
      // arrange
      when(mockSharedPreferences.remove(any))
          .thenAnswer((_) async => true);

      // act
      await localDataSource.deleteToken();

      // assert
      verify(mockSharedPreferences.remove('token'));
    });
  });
}
