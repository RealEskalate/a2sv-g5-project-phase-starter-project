
import 'dart:convert';

import 'package:ecommers/core/Error/failure.dart';
import 'package:ecommers/core/const/const.dart';
import 'package:ecommers/features/ecommerce/Data/data_source/remote_data_source.dart';
import 'package:ecommers/features/ecommerce/Data/model/ecommerce_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/read_json.dart';
import '../../../../helper/test_hlper.mocks.dart';
import 'token.dart';

@GenerateMocks([http.Client])
void main() {
  late MockHttpClient mockHttpClient;

  late EcommerceRemoteDataSourceImpl ecommerceRemoteDataSourceImpl;
  late MockSharedPreferences sharedPreferences;



  setUpAll(() {
    mockHttpClient = MockHttpClient();
    sharedPreferences = MockSharedPreferences();
    ecommerceRemoteDataSourceImpl = EcommerceRemoteDataSourceImpl(client: mockHttpClient, sharedPreferences: sharedPreferences);
  });
  

  Map<String,dynamic> data = {
    'name' : 'name',
    'description' : 'description',
    'imageUrl' : 'imageUrl',
    'price' : 1
  };
  
  String id = '1';
  group(
    'remote data source must return the model',
    () {
     test(
        'it must return the data by id',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(
              mockHttpClient.get(
                Uri.parse(Urls.getByUrl(id)),
                headers: {
                  'Authorization': 'Bearer ${Token.tokin}',
                },
              ),
            ).thenAnswer(
            (_) async => http.Response (
              readJson('helper/dummy_data/remote_single.json'),200
            )
          );

          final result = await ecommerceRemoteDataSourceImpl.getProduct(id);

          expect(result, isA<EcommerceModel>());
        });  

        test(
        'it must return server error',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(
            mockHttpClient.get(Uri.parse(Urls.getByUrl(id)),
            headers: {
                  'Authorization': 'Bearer ${Token.tokin}',
                },
            )
          ).thenAnswer(
            (_) async => http.Response (
              'server errro',404,
            )
          );

          final result = ecommerceRemoteDataSourceImpl.getProduct(id);

          expect(result, throwsA(isA<ConnectionFailur>()));
        });  


      test(
        'it must return the all the data of the product',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(
            mockHttpClient.get(Uri.parse(Urls.getAll()),
            headers: {
                  'Authorization': 'Bearer ${Token.tokin}',
                },
            )
          ).thenAnswer(
            (_) async => http.Response (
              readJson('helper/dummy_data/remote_json.json'),200
            )
          );

          final result = await ecommerceRemoteDataSourceImpl.getAllProducts();

          expect(result, isA<List<EcommerceModel>>());
        });   


        test(
        'it must return true if the data is deleted else false from remote data source',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(

            mockHttpClient.delete(Uri.parse(Urls.deleteProduct(id)),
            headers: {
                  'Authorization': 'Bearer ${Token.tokin}',
                },
            )
          ).thenAnswer(
            (_) async => http.Response (
              'true',200
            )
          );

          final result = await ecommerceRemoteDataSourceImpl.deleteProduct(id);

          expect(result, true);
        });
        // ..headers['Authorization'] = 'Bearer $token'
        test(
        'it must return true if the data is updated from remote data source else false',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(
            mockHttpClient.put(
              Uri.parse(Urls.updateProduct(id)),
                headers: {'Content-Type': 'application/json','Authorization': 'Bearer ${Token.tokin}'},
                body: jsonEncode({
                  'name': data['name'],
                  'description': data['description'],
                  'price': data['price'],
                }),
                
              )
          ).thenAnswer(
            (_) async => http.Response (
              'true',200
            )
          );

          final result = await ecommerceRemoteDataSourceImpl.editProduct(id, data);

          expect(result, true);
        });

        test(
        'it must return true if the data is add to remote data source else false',
        () async{
          when(sharedPreferences.getString('key')).thenReturn(Token.tokin);
          when(
            mockHttpClient.post(
              Uri.parse(Urls.addNewProduct()),
              body: data
              )
          ).thenAnswer(
            (_) async => http.Response (
              'true',201
            )
          );

          final result = await ecommerceRemoteDataSourceImpl.addProduct(data);

          expect(result, false);
        });
    }
    
    
    );
}