import 'dart:convert';

import 'package:flutter_application_5/features/product/data/models/product_models.dart';
import 'package:flutter_application_5/features/product/domain/entities/product_entity.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../helpers/read_json.dart';


void main(){
  const testProductModel = ProductModel(
    id: '66b0b23928f63adda72ab38a',
    name: 'PC', 
    description: 'long description', 
    price: 123, 
    imagePath: null,);

  

  test(
      'should be a subclass of product model',
      () async {
      //assert
      expect(testProductModel, isA<ProductEntity>());
    }
  );
  test(
    'should return a valid model from json',
    ()async{
      //arrange
      final Map<String,dynamic> sonMap= jsonDecode(
        jsonDec('helpers/dummy_data/response.json'),
        
      );

      //act
      final result = ProductModel.fromJson(sonMap["data"][0]);

      //assert
      expect(result, equals(testProductModel));
    }

  );

  //To return json

  test(
    'should convert the product model to a valid JSON',
    () async {
      // arrange
      final expectedJson = {
        "id": "66b0b23928f63adda72ab38a",
        "name": "PC",
        "description": "long description",
        "price": 123,
        "imagePath": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1722855993/images/soyhb68osjiemyy2btte.png"
      };

      // act
      final result = testProductModel.toJson();

      // assert
      expect(result, equals(expectedJson));
    }

    
  );



}


