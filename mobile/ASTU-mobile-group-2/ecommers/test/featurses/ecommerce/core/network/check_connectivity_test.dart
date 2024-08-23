import 'package:ecommers/core/network/check_connectivity.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import '../../../../helper/test_hlper.mocks.dart';
// Import the generated mock file

// @GenerateMocks([InternetConnectionChecker])
void main() {
  late NetworkInfoImpl networkInfo;
  late MockInternetConnectionChecker mockInternetConnectionChecker;

  setUp(() {
    mockInternetConnectionChecker = MockInternetConnectionChecker();
    networkInfo = NetworkInfoImpl(connectionChecker: mockInternetConnectionChecker);
  });

  test(
    'should return true when the device is connected to the internet',
    () async {
      // Arrange
      when(mockInternetConnectionChecker.hasConnection)
          .thenAnswer((_) async => true);

      // Act
      final result = await networkInfo.isConnected;

      // Assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, true);
    },
  );

  test(
    'should return false when the device is not connected to the internet',
    () async {
      // Arrange
      when(mockInternetConnectionChecker.hasConnection)
          .thenAnswer((_) async => false);

      // Act
      final result = await networkInfo.isConnected;

      // Assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, false);
    },

  );

  
}
