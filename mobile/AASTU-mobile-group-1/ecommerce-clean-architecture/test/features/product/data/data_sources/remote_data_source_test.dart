

import 'dart:convert';

import 'package:ecommerce/core/constants/constants.dart';
import 'package:ecommerce/core/error/exception.dart';
import 'package:ecommerce/features/product/data/data_sources/remote_data_source.dart';
import 'package:ecommerce/features/product/data/model/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart';
import 'package:http/http.dart';
import 'package:http/http.dart';
import 'package:http/http.dart';
import 'package:http/http.dart';
import 'package:mockito/annotations.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/mockito.dart';
import '../../../../helpers/dummy_data/dummy.dart';
import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';
class MockMultipartFile extends Mock implements http.MultipartFile {}

void main(){


late MockClient mockHttpClient;
late MockMultipartFile mockMultipartFile;

late ProductRemoteDataSourceImpl productRemoteDataSourceImpl;
setUp((){
  mockHttpClient = MockClient();
  productRemoteDataSourceImpl = ProductRemoteDataSourceImpl(client: mockHttpClient);
  mockMultipartFile = MockMultipartFile();
});
  const tProductModel =  ProductModel(
    id: '123',
    name: 'Nike Shoes',
    description: 'High-quality sports shoes',
    price: 99,
    image: 'assets/wifi.jpg'
  );

  final productModel = ProductModel.fromJson( const {
        "id": "6672752cbd218790438efdb0",
        "name": "Anime website",
        "description": 'Explore anime characters.',
        "price": 123,
        "imageUrl": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg"
      });

  final tProductModelJson = {
    'id': '123',
    'name': 'Nike Shoes',
    'description': 'High-quality sports shoes',
    'price': 99,
    'imageUrl': 'path/to/image.jpg'
  };
group('get all product',(){

test('should return list of product model when the response code is 200', ()async{
    when(
      mockHttpClient.get(Uri.parse(Urls.baseUrl)
      )
      ).thenAnswer( 
        (_) async=>http.Response(
          data,200
        ) 
      ); 

      final result = await productRemoteDataSourceImpl.getallproduct();
      expect(result,isA<List<ProductModel>>());

});
test('should throw a server exception when the response code is 404 or other', ()async{
    when(
      mockHttpClient.get(Uri.parse(Urls.baseUrl)
      )
      ).thenAnswer( 
        (_) async=>http.Response(
          'Not found',404
        ) 
      ); 

      // final result = await productRemoteDataSourceImpl.getallproduct();
    expect(
    () async => await productRemoteDataSourceImpl.getallproduct(),
    throwsA(isA<ServerException>()),
  );

});
 test('should return a product model when the response code is 200', ()async{
    when(
      mockHttpClient.get(Uri.parse('${Urls.baseUrl}/6672752cbd218790438efdb0'))
      ).thenAnswer( 
        (_) async=>http.Response(
          readJson('helpers/dummy_data/single_product.json'),200
        ) 
      ); 
      //  'https://g5-flutter-learning-path-be.onrender.com/api/v1/products';
      //  print(productModel.id == '6672752cbd218790438efdb0');
      //  print('6672752cbd218790438efdb0');

      final result = await productRemoteDataSourceImpl.getproduct(productModel.id);
      // print(result);
      expect(result,productModel);
});
 test('should return a server exception when getting product failed', ()async{
    when(
      mockHttpClient.get(Uri.parse('${Urls.baseUrl}/6672752cbd218790438efdb0'))
      ).thenAnswer( 
        (_) async=>http.Response(
          "Not found",404
        ) 
      ); 
       'https://g5-flutter-learning-path-be.onrender.com/api/v1/products';
       print(productModel.id == '6672752cbd218790438efdb0');
       print('6672752cbd218790438efdb0');

      final result = await productRemoteDataSourceImpl.getproduct(productModel.id);
    expect(
    () async => await productRemoteDataSourceImpl.getproduct(productModel.id),
    throwsA(isA<ServerException>()),
  );
});






test('should add product when the response list is 200 ',()async{
   var uri = Uri.parse(Urls.postUrl);

   final responseData = jsonEncode(tProductModelJson);


  final byteStream = Stream.fromIterable([utf8.encode(responseData)]);

  final streamResponse = http.StreamedResponse(byteStream, 201);
   when(mockHttpClient.send(any)
   ).thenAnswer((_)async=>streamResponse);
   
  const newProductModel =  ProductModel(
        id: '123', 
        name: 'Nike Shoes',
        description: 'High-quality sports shoes', 
        price: 99, 
        image: 'assets/wifi.jpg'
        );
  expect(
   productRemoteDataSourceImpl.addproduct(newProductModel),
    completes,
  );}, 

);

test('should update product when the response list is 200',() async {
    // Arrange
    when(
      mockHttpClient.put(
        Uri.parse('${Urls.putUrl}/123'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'name': tProductModel.name,
          'description': tProductModel.description,
          'price': tProductModel.price,
          'imageUrl': tProductModel.image,
        }),
      ),
    ).thenAnswer(
      (_) async => http.Response(jsonEncode(tProductModelJson), 200),
    );
    final result = await productRemoteDataSourceImpl.updateproduct(tProductModel);
    expect(result, equals(tProductModel));

 });

 test('should throw error when failed to update product',() async {
    // Arrange
    when(
      mockHttpClient.put(
        Uri.parse('${Urls.putUrl}/123'),
        headers: {'Content-Type': 'application/json'},
        body: jsonEncode({
          'name': tProductModel.name,
          'description': tProductModel.description,
          'price': tProductModel.price,
          'imageUrl': tProductModel.image,
        }),
      ),
    ).thenAnswer(
      (_) async => http.Response('Not found', 404),
    );
    // final result = await productRemoteDataSourceImpl.updateproduct(tProductModel);
    // expect(result, equals(tProductModel));
  expect(
    () async => await productRemoteDataSourceImpl.updateproduct(tProductModel),
    throwsA(isA<ServerException>()),
  );

 });
test('should delete the product',() async {
    // Arrange
    when(
      mockHttpClient.delete(
        Uri.parse('${Urls.putUrl}/123'),
      ),
    ).thenAnswer(
      (_) async => http.Response(jsonEncode(tProductModelJson), 200),
    );
    final result = await productRemoteDataSourceImpl.deleteproduct('123');
    expect(result, true);
 });

test('should show error if deletion failed',() async {
    // Arrange
    when(
      mockHttpClient.delete(
        Uri.parse('${Urls.putUrl}/123'),
      ),
    ).thenAnswer(
      (_) async => http.Response("Not Found", 404),
    );
    // final result = await productRemoteDataSourceImpl.deleteproduct('123');
    // expect(result, true);
    expect(
    () async => await productRemoteDataSourceImpl.deleteproduct('123'),
    throwsA(isA<ServerException>()),
  );
 });
}
);

}