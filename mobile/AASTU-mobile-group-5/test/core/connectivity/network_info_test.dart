import 'package:flutter_test/flutter_test.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/connectivity/network_info.dart';

import 'network_info_test.mocks.dart';


@GenerateNiceMocks([MockSpec<InternetConnectionChecker>()])
void main() {
  late NetworkInfoImpl networkInfo;
  late MockInternetConnectionChecker mockInternetConnectionChecker;

  setUp(() {
    mockInternetConnectionChecker = MockInternetConnectionChecker();
    networkInfo = NetworkInfoImpl(mockInternetConnectionChecker as InternetConnectionChecker);
  });

  group('isConnected', () {
    test('should forward the call to InternetConnectionChecker.hasConnection and checks if device is connected to the internet', () async {
      // arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer((_) async => true);
      // act
      final result = await networkInfo.isConnected;
      // assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, true);
    });

    test('should return true when there is internet connection', () async {
      // arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer((_) async => true);
      // act
      final result = await networkInfo.isConnected;
      // assert
      expect(result, true);
    });

    test('should return false when there is no internet connection', () async {
      // arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer((_) async => false);
      // act
      final result = await networkInfo.isConnected;
      // assert
      expect(result, false);
    });
  });
}
