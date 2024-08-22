import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import 'package:product_6/core/connections/network_info.dart';

import '../mock.mocks.dart';

void main() {
  late NetworkInfoImpl networkInfo;
  late MockDataConnectionChecker mockDataConnectionChecker;

  setUp(() {
    mockDataConnectionChecker = MockDataConnectionChecker();
    networkInfo = NetworkInfoImpl(mockDataConnectionChecker);
  });

  group('isConnected', () {
    test('should forward the call to DataConnectionChecker.hasConnection', () async {
      // Arrange
      final tHasConnectionFuture = Future.value(true);

      when(mockDataConnectionChecker.hasConnection).thenAnswer((_) => tHasConnectionFuture);

      // Act
      final result = networkInfo.isConnected;

      // Assert
      verify(mockDataConnectionChecker.hasConnection);
      expect(result, equals(tHasConnectionFuture));
    });

    test('should return false when DataConnectionChecker.hasConnection is false', () async {
      // Arrange
      when(mockDataConnectionChecker.hasConnection).thenAnswer((_) async => false);

      // Act
      final result = await networkInfo.isConnected;

      // Assert
      expect(result, equals(false));
    });

    test('should return true when DataConnectionChecker.hasConnection is true', () async {
      // Arrange
      when(mockDataConnectionChecker.hasConnection).thenAnswer((_) async => true);

      // Act
      final result = await networkInfo.isConnected;

      // Assert
      expect(result, equals(true));
    });
  });
}














