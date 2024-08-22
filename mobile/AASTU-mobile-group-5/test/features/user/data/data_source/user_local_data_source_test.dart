import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:shared_preferences/shared_preferences.dart';
import 'package:task_9/features/user/data/data_sources/user_local_data_source.dart';

import '../../../product/data/datasources/product_local_data_source_test.mocks.dart';



@GenerateMocks([SharedPreferences])
void main() {
  late UserLocalDataSourceImpl dataSource;
  late MockSharedPreferences mockSharedPreferences;

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    dataSource = UserLocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  group('saveAccessToken', () {
    const tAccessToken = 'test_token';

    test('should call SharedPreferences to save the token', () async {
      // Arrange
      when(mockSharedPreferences.setString(any, any)).thenAnswer((_) async => true);

      // Act
      await dataSource.saveAccessToken(tAccessToken);

      // Assert
      verify(mockSharedPreferences.setString('ACCESS_TOKEN', tAccessToken));
    });
  });

  group('getAccessToken', () {
    const tAccessToken = 'test_token';

    test('should return the saved token from SharedPreferences', () async {
      // Arrange
      when(mockSharedPreferences.getString(any)).thenReturn(tAccessToken);

      // Act
      final result = await dataSource.getAccessToken();

      // Assert
      verify(mockSharedPreferences.getString('ACCESS_TOKEN'));
      expect(result, tAccessToken);
    });

    test('should return null when there is no token saved', () async {
      // Arrange
      when(mockSharedPreferences.getString(any)).thenReturn(null);

      // Act
      final result = await dataSource.getAccessToken();

      // Assert
      verify(mockSharedPreferences.getString('ACCESS_TOKEN'));
      expect(result, null);
    });
  });

  group('deleteAccessToken', () {
    test('should call SharedPreferences to remove the token', () async {
      // Arrange
      when(mockSharedPreferences.remove(any)).thenAnswer((_) async => true);

      // Act
      await dataSource.deleteAccessToken();

      // Assert
      verify(mockSharedPreferences.remove('ACCESS_TOKEN'));
    });
  });
}
