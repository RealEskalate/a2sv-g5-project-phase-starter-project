import 'package:ecommerce_app/core/network/network_info.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../test_helper/test_helper_generation.mocks.dart';

void main(){
  late NetworkInfoImpl networkInfoImpl;
  late MockInternetConnectionChecker mockInternetConnectionChecker;

  setUp((){
    mockInternetConnectionChecker = MockInternetConnectionChecker();
    networkInfoImpl = NetworkInfoImpl(mockInternetConnectionChecker);
  });

  group('Testing the network info isConnected', (){

    test('Should return when the internet checker returns true', () async {
      /// arrange
      when(
          mockInternetConnectionChecker.hasConnection
      ).thenAnswer((_) async => true);
      ///action
      final result = await networkInfoImpl.isConnected;
      ///assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, true);

    });

    test('Should return false when there is no internet', () async {
      /// arrange
      when(
          mockInternetConnectionChecker.hasConnection
      ).thenAnswer((_) async => false);
      ///action
      final result = await networkInfoImpl.isConnected;
      ///assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, false);
    });
  });
}