// import '../../helpers/test_helper.mocks.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/network/network_info.dart';

import '../../helpers/test_helper.mocks.dart';

void main() {
  late MockInternetConnectionChecker mockInternetConnectionChecker;
  late NetworkInfoImpl networkInfoImpl;

  setUp(() {
    mockInternetConnectionChecker = MockInternetConnectionChecker();
    networkInfoImpl =
        NetworkInfoImpl(connectionChecker: mockInternetConnectionChecker);
  });

  group('isConnected', () {
    test('should forward the call to InternetConnectionChecker.hasConnection',
        () async {
      // arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer(
        (_) async => true,
        // act
      );
      final result = await networkInfoImpl.isConnected;
      // assert
      expect(result, true);
    });
  });
}
