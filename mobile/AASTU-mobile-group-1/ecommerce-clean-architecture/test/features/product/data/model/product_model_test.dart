import 'dart:convert';

import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../../../helpers/dummy_data/dummy.dart';


void main(){

  const testproductmodel = ProductModel(
     id:'6672752cbd218790438efdb0', 
     image:'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg', 
     name: 'Anime website', 
     description: 'Explore anime characters.',
      price: 123
      );
    test('should be a subclass of product entity',
      () async  {
        expect(testproductmodel,isA<Productentity>());
      }
    
     );

     test('read from json and convert to dart object',
      () async  {
        final Map<String,dynamic> jsonMap = jsonDecode(
          data           
        );
        final result = ProductModel.fromJson(jsonMap['data'][0]); 
        // print(result);
        const expectedproduct =  ProductModel(
        id: '6672752cbd218790438efdb0',
        name: 'Anime website',
        description: 'Explore anime characters.',
        price: 123,
        image: 'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg'
        );

        expect(result,expectedproduct);
      }
    
      );

    test(
      'should return a json map containing proper data',
      () async{

        final result = testproductmodel.toJson();
        final expectedJsonMap={
          'id' : '6672752cbd218790438efdb0',
          'name': 'Anime website', 
          'description':'Explore anime characters.',
          'price':123,
          'imageUrl': 'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',

        };
        expect(result, equals(expectedJsonMap));      
      }


    );

}