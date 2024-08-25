import 'dart:convert';
import 'package:ecommers/core/Error/failure.dart';
import 'package:ecommers/features/ecommerce/Data/data_source/local_data_source.dart';
import 'package:ecommers/features/ecommerce/Data/model/ecommerce_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/dummy_data/read_json.dart';
import '../../../../helper/test_hlper.mocks.dart';

void main() {
  late MockSharedPreferences mockSharedPreferences;
  late LocalDataSourceImpl localDataSourceImpl;
  

  setUp(() {
    mockSharedPreferences = MockSharedPreferences();
    localDataSourceImpl = LocalDataSourceImpl(sharedPreferences: mockSharedPreferences);
  });

  final listOfModel = readJson('helper/dummy_data/all_data.json');
  final singleModel = readJson('helper/dummy_data/json_respond_data.json');
  final datas = json.decode(listOfModel);
  EcommerceModel model = const EcommerceModel(
    id: '1', 
    name: 'name', description: 'description', imageUrl: 'imageUrl', price: 2);
 
  String id = '3';
  group(
    'test the local data source',
    (){

      test(
        'test get all data localdata feaching process',
        () async{
          when(
            mockSharedPreferences.getString(any)
          ).thenReturn(listOfModel);
          final result = await localDataSourceImpl.getAllFromLocal();

          verify(mockSharedPreferences.getString(any));
          expect(result, equals(EcommerceModel.getAllProduct(datas)));
        });


        test(
          'test single data',
          () async{
            
            when(
              mockSharedPreferences.getString(any)
            ).thenReturn(singleModel);
  
            final result =  await localDataSourceImpl.getSingleProduct(id);
            
            verify(mockSharedPreferences.getString(any));

            final last = EcommerceModel.fromJson(json.decode(singleModel)['data']);

            expect(result, equals(last));
          });


          test(
          'test must return error when the data is null',
          () async{
            
            when(
              mockSharedPreferences.getString(any)
            ).thenThrow(const CachException(message: 'No data found'));
            expect(() async => await localDataSourceImpl.getAllFromLocal(), throwsA(isA<CachException>()));
          });

          test(
            'test must return error when the data is null for specific id',
            () async {
              // Arrange
              when(mockSharedPreferences.getString(any))
                  .thenThrow(const CachException(message: 'No data found'));

              // Act & Assert
              expect(() async => await localDataSourceImpl.getSingleProduct(id), throwsA(isA<CachException>()));
            }
          );

          test(
          'test must return true when the data is add correctly or fals if not added',

            () async{ 
            Future<bool> value = Future.value(false);
            const idOfData = 'local_ecommer_data';
            final mockData = json.encode([model.toJson()]);
            when(mockSharedPreferences.getString(idOfData)).thenReturn(mockData);
            when(mockSharedPreferences.setString(any, any)).thenAnswer((_) => value);
            
            final result =  localDataSourceImpl.addCach(model);
            
            expect(result, equals(value));
            }
          
          );


          
          
    }
  );
}