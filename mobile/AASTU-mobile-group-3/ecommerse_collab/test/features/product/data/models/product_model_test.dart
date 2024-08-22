import 'dart:convert';

import 'package:ecommerse2/features/product/data/models/product_model.dart';
import 'package:ecommerse2/features/product/domain/entity/product.dart';
import 'package:flutter_test/flutter_test.dart';

import '../../helpers/dummy_data/dummy_product_response.dart';

void main(){

  ProductModel productModel = const ProductModel(
    id: '1', name: 'Nike', category: '', description: 'A great Shoe', image: 'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg', price: 99
  );

  final json = {
    'id' : '1',
    'name' : 'Nike',
    'description' : 'A great Shoe',
    'category' : '',
    'image' : 'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg',
    'price' : 99,
  };

test(
  'should be a subclass of product entity', //What should be a subclass of the product entity? The product model, implys we need a product model object
  () async {

    //assert
    expect(productModel, isA<Product>());

  }
);

test('should correctly convert from JSON to ProductModel', () async {

  //act
  final Map<String, dynamic> jsonMap = jsonDecode(data);
  final result = ProductModel.fromJson(jsonMap['data'][0]); 

  //assert
  expect(result,isA<ProductModel>());

});

test('should correctly convert to JSON from ProductModel', () async {
  final result = productModel.toJson();

  expect(result, json);
});


}