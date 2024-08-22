import 'package:ecommerce/core/platform/network_info.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late NetworkInfoImpl networkInfo;
  late MockInternetConnectionChecker mockDataConnectionChecker;

  setUp(() {
    mockDataConnectionChecker = MockInternetConnectionChecker();
    networkInfo = NetworkInfoImpl(mockDataConnectionChecker);
  });

  group('isConnected', () {
    test(
      'should forward the call to DataConnectionChecker.hasConnection',
      () async {
        // arrange
        final tHasConnectionFuture = Future.value(true);

        when(mockDataConnectionChecker.hasConnection)
            .thenAnswer((_) => tHasConnectionFuture);
        // act

        final result = networkInfo.isConnected;
        // assert
        verify(mockDataConnectionChecker.hasConnection);

        expect(result, tHasConnectionFuture);
      },
    );

    test(
      'should forward the call to DataConnectionChecker.hasConnection when not connected',
      () async {
        // arrange
        final tHasConnectionFuture = Future.value(false);

        when(mockDataConnectionChecker.hasConnection)
            .thenAnswer((_) => tHasConnectionFuture);
        // act

        final result = networkInfo.isConnected;
        // assert
        verify(mockDataConnectionChecker.hasConnection);

        expect(result, tHasConnectionFuture);
      },
    );
  });
}
