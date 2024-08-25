

import 'dart:convert';

import 'package:ecommers/features/ecommerce/Data/model/ecommerce_model.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helper/dummy_data/read_json.dart';


void main() {
  const EcommerceModel model = EcommerceModel(
    id: '6672776eb905525c145fe0bb',
    name: 'Anime website',
    description: 'Explore anime characters.',
    imageUrl: 'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg',
    price: 123
  );

  test(
    'the model must similar to the entity of domain',
    () async{
      

        expect(model,isA<EcommerceEntity>());
    
    });
  test(
    'the model from json test must similar to the midel',
    () async{

      final dynamic jsonData = json.decode(
        readJson('helper/dummy_data/json_respond_data.json')
      );
      final result = EcommerceModel.fromJson(jsonData['data']);
      expect(result, equals(model));
    
    });
  
  test(
    'convert object to json format',
     () async {
      final result = model.toJson();
      
      expect(result, isA<Map<String,dynamic>>());
     }
     );
}