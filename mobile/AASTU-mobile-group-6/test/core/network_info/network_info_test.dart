import 'package:ecommerce_app_ca_tdd/core/network/network_info.dart';
import 'package:internet_connection_checker/internet_connection_checker.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

// class MockDataConnectionChecker extends Mock implements InternetConnectionChecker {}

void main() {
  late MockInternetConnectionChecker mockDataConnectionChecker;
  late NetworkInfoImpl networkInfo;

  setUp(() {
    mockDataConnectionChecker = MockInternetConnectionChecker();
    networkInfo = NetworkInfoImpl(mockDataConnectionChecker);
  });


test(
  'should return true if there is an internet connection',
  () async {
    // arrange
    final tHasConnectionFuture = Future.value(true);
    when(mockDataConnectionChecker.hasConnection)
        .thenAnswer((_) async => await tHasConnectionFuture);
    // act
    final result = await networkInfo.isConnected;
    // assert
    // verify(mockDataConnectionChecker.hasConnection);
    expect(result, true);
  },
);

test(
  'should return false if there is no internet connection',
  () async {
    // arrange
    final tHasConnectionFuture = Future.value(false);
    when(mockDataConnectionChecker.hasConnection)
        .thenAnswer((_) => tHasConnectionFuture);
    // act
    final result = await networkInfo.isConnected;
    // assert
    verify(mockDataConnectionChecker.hasConnection);
    expect(result, false);
  },
);

}