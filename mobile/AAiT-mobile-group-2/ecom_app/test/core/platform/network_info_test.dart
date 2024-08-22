
import 'package:ecom_app/core/platform/network_info.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

void main(){
  late NetworkInfoImpl networkInfoImpl;
  late MockInternetConnectionChecker mockInternetConnectionChecker;

  setUp((){
    mockInternetConnectionChecker = MockInternetConnectionChecker();
    networkInfoImpl = NetworkInfoImpl(mockInternetConnectionChecker);
  });


  group('isConnected', (){

    test('should return true when network is connected', ()async {

      final testHasConnection = Future.value(true);
      //arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer((_) => testHasConnection);

      //act
      final result = networkInfoImpl.isConnected;

      //assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, testHasConnection);
    });

    test("should return false when network isn't connected", ()async {

      final testHasConnection = Future.value(false);
      //arrange
      when(mockInternetConnectionChecker.hasConnection).thenAnswer((_) => testHasConnection);

      //act
      final result = networkInfoImpl.isConnected;

      //assert
      verify(mockInternetConnectionChecker.hasConnection);
      expect(result, testHasConnection);
    });
  });


}