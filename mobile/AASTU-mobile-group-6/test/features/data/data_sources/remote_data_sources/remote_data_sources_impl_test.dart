// import 'package:flutter_application_5/core/network/network_info.dart';
// import 'package:flutter_application_5/features/product/data/data_sources/remote_data_source/remote_data_source.dart';
// import 'package:flutter_application_5/features/product/data/models/product_models.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/http.dart' as http;
// import 'package:mockito/mockito.dart';

// class MockHttpClient extends Mock implements http.Client {}

// void main() {
//   late ProductRemoteDataSourceImpl dataSource;
//   late MockHttpClient mockHttpClient;

//   setUp(() {
//     mockHttpClient = MockHttpClient();
//     dataSource = ProductRemoteDataSourceImpl(client: mockHttpClient);
//   });

//   group('getProduct', () {
//     final id = '1';
//     final productModel = ProductModel.fromJson(
//       const {
//     "statusCode": 200,
//     "message": "",
//     "data": [
//       {
//         "id": "1",
//         "name": "PC",
//         "description": "long description",
//         "price": 123,
//         "imageUrl": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1722855993/images/soyhb68osjiemyy2btte.png"
//       }
//     ]
//   }
//       // Provide sample JSON data for the product model
//     );

//     test(
//       'should perform a GET request on the specified URL and return a product model',
//       () async {
//         // arrange
//         when(mockHttpClient.get(Uri.parse('https://example.com/products/$id')))
//             .thenAnswer((_) async => http.Response(
//                 // Provide sample response data for a successful request
//                 ));
//         // act
//         final result = await dataSource.getProduct(id);
//         // assert
//         expect(result, equals(productModel));
//       },
//     );

//     test(
//       'should throw an exception if the HTTP request fails',
//       () async {
//         // arrange
//         when(mockHttpClient.get(Uri.parse('https://example.com/products/$id')))
//             .thenAnswer((_) async => http.Response('Error', 404));
//         // act
//         final call = dataSource.getProduct(id);
//         // assert
//         expect(call, throwsException);
//       },
//     );
//   });

//   // New Group with out the getProducts method
//   group('addProduct', () {
//   final productModel = ProductModel(
//     // id: '2',
//     name: 'Laptop',
//     description: 'powerful laptop',
//     price: 999,
//     imagePath: 'https://example.com/laptop.png',
//   );

//   test(
//     'should perform a POST request with the product data and return the added product',
//     () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('https://example.com/products'),
//         body: productModel.toJson(),
//       )).thenAnswer((_) async => http.Response(
//           // Provide sample response data for a successful request
//           ));
//       // act
//       final result = await dataSource.addProduct(productModel);
//       // assert
//       expect(result, equals(productModel));
//     },
//   );

//   test(
//     'should throw an exception if the HTTP request fails',
//     () async {
//       // arrange
//       when(mockHttpClient.post(
//         Uri.parse('https://example.com/products'),
//         body: productModel.toJson(),
//       )).thenAnswer((_) async => http.Response('Error', 404));
//       // act
//       final call = dataSource.addProduct(productModel);
//       // assert
//       expect(call, throwsException);
//     },
//   );
// });

// group('updateProduct', () {
//   final id = '1';
//   final updatedProductModel = ProductModel(
    
//     name: 'Updated PC',
//     description: 'updated description',
//     price: 456,
//     imagePath: 'https://example.com/updated_pc.png',
//   );

//   test(
//     'should perform a PUT request with the updated product data and return the updated product',
//     () async {
//       // arrange
//       when(mockHttpClient.put(
//         Uri.parse('https://example.com/products/$id'),
//         body: updatedProductModel.toJson(),
//       )).thenAnswer((_) async => http.Response(
//           // Provide sample response data for a successful request
//           ));
//       // act
//       final result = await dataSource.updateProduct(updatedProductModel);
//       // assert
//       expect(result, equals(updatedProductModel));
//     },
//   );

//   test(
//     'should throw an exception if the HTTP request fails',
//     () async {
//       // arrange
//       when(mockHttpClient.put(
//         Uri.parse('https://example.com/products/$id'),
//         body: updatedProductModel.toJson(),
//       )).thenAnswer((_) async => http.Response('Error', 404));
//       // act
//       final call = dataSource.updateProduct(updatedProductModel);
//       // assert
//       expect(call, throwsException);
//     },
//   );
// });

// group('deleteProduct', () {
//   final id = '1';

//   test(
//     'should perform a DELETE request on the specified URL and return true if the product is deleted',
//     () async {
//       // arrange
//       when(mockHttpClient.delete(Uri.parse('https://example.com/products/$id')))
//           .thenAnswer((_) async => http.Response(
//               // Provide sample response data for a successful request
//               ));
//       // act
//       final result = await dataSource.deleteProduct(id);
//       // assert
//       expect(result, isTrue);
//     },
//   );

//   test(
//     'should throw an exception if the HTTP request fails',
//     () async {
//       // arrange
//       when(mockHttpClient.delete(Uri.parse('https://example.com/products/$id')))
//           .thenAnswer((_) async => http.Response('Error', 404));
//       // act
//       final call = dataSource.deleteProduct(id);
//       // assert
//       expect(call, throwsException);
//     },
//   );
// });

//   // Add more test cases for other methods in the ProductRemoteDataSourceImpl class
// }