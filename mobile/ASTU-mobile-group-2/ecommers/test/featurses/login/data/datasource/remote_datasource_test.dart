import 'package:dartz/dartz.dart';
import 'package:ecommers/core/Error/failure.dart';
import 'package:ecommers/core/const/const.dart';
import 'package:ecommers/features/login/data/datasource/remote_datasource.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/read_json.dart';
import '../../../../helper/test_hlper.mocks.dart';



void main(){

    late MockClient mockClient;
    late RemoteDatasourceImpl remoteDatasourceImpl;
    late MockNetworkInfoImpl mockNetworkInfoImpl; 
    late MockSharedPreferences mockSharedPreferences;


    setUp((){
        mockClient = MockClient();
        mockNetworkInfoImpl = MockNetworkInfoImpl();
        mockSharedPreferences = MockSharedPreferences();
        remoteDatasourceImpl = RemoteDatasourceImpl(
          client: mockClient,
          networkInfo: mockNetworkInfoImpl,
          sharedPreferences: mockSharedPreferences);
    });
    final data = readJson('helper/dummy_data/login_success.json');
  
 
  


    test(
      'get data from remote  data source must return error',
      () async {
        when(mockClient.post(
          Uri.parse(LoginApi.loginApi),
          body: {'email': 'tolossamuel1@gmail.com', 'password': '12341234'},
        )).thenAnswer((_) async => http.Response(data, 201));

        // Call the login method and await the result
        final result = await remoteDatasourceImpl.login('tolossamuel1@gmail.com', '12341234');
    
         expect(result, isA<Left>());
         expect(result.fold((l) => l, (r) => r), isA<ConnectionFailur>());
      }

      
    );
}