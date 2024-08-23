// import 'dart:convert';
// import 'dart:io';

// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/http.dart' as http;
// import 'package:mockito/mockito.dart';
// import 'package:product_8/core/constants/constants.dart';
// import 'package:product_8/core/exception/exception.dart';
// import 'package:product_8/features/product/data/data_source/remote_data_source/product_remote_data_source.dart';
// import 'package:product_8/features/product/data/models/product_model.dart';

// import '../../../../helpers/jason_reader.dart';
// import '../../../../helpers/test_helper.mocks.dart';

// void main() {
//   late MockHttpClient mockHttpCliant;
//   late ProductRemoteDataSourceImpl productRemoteDateSourceImpl;

//   setUp(() {
//     mockHttpCliant = MockHttpClient();
//     productRemoteDateSourceImpl =
//         ProductRemoteDataSourceImpl(client: mockHttpCliant);
//   });

//   const productId = '6672776eb905525c145fe0bb';
//   const jsonCurrent = 'helpers/dummy_data/mock_data.json';
//   const jsonAll = 'helpers/dummy_data/mock_data_list.json';
//   const testProductModel = ProductModel(
//       id: '6672776eb905525c145fe0bb',
//       name: 'Anime website',
//       description: 'Explore anime characters.',
//       imageUrl:
//           'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg',
//       price: 123.8);

// // get current product
//   group('get current product', () {
//     test('should return product model when the response code is 200', () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.getProdutbyId(productId))))
//           .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 200));
//       //act
//       final result =
//           await productRemoteDateSourceImpl.getProductById(productId);
//       //assert
//       expect(result, isA<ProductModel>());
//     });

//     test('should return product model when the response code is 404', () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.getProdutbyId(productId))))
//           .thenThrow(const SocketException('No Internet Connection'));
//       //act
//       final call = productRemoteDateSourceImpl.getProductById;
//       // assert
//       expect(() => call(productId), throwsA(isA<SocketException>()));
//     });

//     test('should throw a  socket  exception if it happens', () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.getProdutbyId(productId))))
//           .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 400));
//       //act  and assert
//       expect(() => productRemoteDateSourceImpl.getProductById(productId),
//           throwsA(isA<ServerException>()));
//     });
//   });

// // get all product
//   group('get all product', () {
//     test('should return list of product model when the response code is 200',
//         () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.baseUrl)))
//           .thenAnswer((_) async => http.Response(readJson(jsonAll), 200));
//       //act
//       final result = await productRemoteDateSourceImpl.getProducts();
//       //assert
//       expect(result, isA<List<ProductModel>>());
//     });

//     test('should throw a  socket  exception if it happens', () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.baseUrl)))
//           .thenThrow(const SocketException('No Internet Connection'));
//       //act  and assert
//       expect(() => productRemoteDateSourceImpl.getProducts(),
//           throwsA(isA<SocketException>()));
//     });

//     test('should throw a  server  exception if it happens', () async {
//       //arrange
//       when(mockHttpCliant.get(Uri.parse(Urls.baseUrl)))
//           .thenAnswer((_) async => http.Response(readJson(jsonAll), 400));
//       //act  and assert
//       expect(() => productRemoteDateSourceImpl.getProducts(),
//           throwsA(isA<ServerException>()));
//     });
//   });

// // delete product
//   group('deleteProduct', () {
//     test('should delete successfully', () async {
//       //arrange
//       when(mockHttpCliant.delete(Uri.parse(Urls.getProdutbyId(productId))))
//           .thenAnswer((_) async => http.Response('', 200));
//       //act

//       await productRemoteDateSourceImpl.deleteProduct(productId);
//       //assert
//       verify(mockHttpCliant.delete(Uri.parse(Urls.getProdutbyId(productId))));
//     });

//     test('should throw a ServerException when the response code is not 200',
//         () async {
//       // arrange
//       when(mockHttpCliant.delete(any))
//           .thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // act
//       final call = productRemoteDateSourceImpl.deleteProduct;

//       // assert
//       expect(() => call(productId), throwsA(isA<ServerException>()));
//     });

//     test('should throw a socket exception if it happens', () {
//       //arrange
//       when(mockHttpCliant.delete(any))
//           .thenThrow(const SocketException('No Internet Connection'));

//       //act
//       final call = productRemoteDateSourceImpl.deleteProduct;

//       //assert
//       expect(() => call(productId), throwsA(isA<SocketException>()));
//     });
//   });

// //update product
//   group('update product', () {
//     test('should update successfully', () async {
//       //arrange
//       final jsonBody = jsonEncode( {
//       'name': testProductModel.name,
//       'description': testProductModel.description,
//       'price': testProductModel.price,
//     });
//       when(mockHttpCliant.put(Uri.parse(Urls.getProdutbyId(productId)),
//               body: jsonBody ,headers: {'Content-Type': 'application/json'}))
//           .thenAnswer((_) async => http.Response(readJson(jsonCurrent), 200));
//       //act

//       final result =
//           await productRemoteDateSourceImpl.updateProduct(testProductModel);
//       //assert
//       expect(result, testProductModel);
//     });

//     test('should throw a ServerException when the response code is not 200',
//         () async {
//       // arrange
//       final jsonBody = jsonEncode( {
//       'name': testProductModel.name,
//       'description': testProductModel.description,
//       'price': testProductModel.price,
//     });
//       when(mockHttpCliant.put(Uri.parse(Urls.getProdutbyId(productId)),
//               body: jsonBody ,headers: {'Content-Type': 'application/json'}))
//           .thenAnswer((_) async =>  http.Response('Something went wrong', 404));

//       // act
//       final call = productRemoteDateSourceImpl.updateProduct;

//       // assert
//       expect(() => call(testProductModel), throwsA(isA<ServerException>()));
//     });

//     test('should throw a socket exception if it happens', () {
//       //arrange
//       final jsonBody = jsonEncode( {
//       'name': testProductModel.name,
//       'description': testProductModel.description,
//       'price': testProductModel.price,
//     });
//       when(mockHttpCliant.put(Uri.parse(Urls.getProdutbyId(productId)),
//               body: jsonBody ,headers: {'Content-Type': 'application/json'}))
//           .thenThrow(const SocketException('No Internet Connection'));

//       //act
//       final call = productRemoteDateSourceImpl.updateProduct;

//       //assert
//       expect(() => call(testProductModel), throwsA(isA<SocketException>()));
//     });
//   });

// //create product

//   // group('create product', () {
//   //   test('should create successfully', () async {
//   //     //arrange

//   //     final productJson = {
//   //       'name': testProductModel.name,
//   //       'description': testProductModel.description,
//   //       'imageUrl': testProductModel.imageUrl,
//   //       'price': testProductModel.price
//   //     };

//   //     when(mockHttpCliant.post(
//   //       Uri.parse(Urls.baseUrl),
//   //       body: productJson,
//   //     )).thenAnswer((_) async => http.Response(readJson(jsonCurrent), 200));
//   //     //act

//   //     final result =
//   //         await productRemoteDateSourceImpl.createProduct(testProductModel);
//   //     //assert
//   //     expect(result, testProductModel);
//   //   });

//   //   test('should throw a ServerException when the response code is not 200',
//   //       () async {
//   //     // arrange
//   //     final productJson = {
//   //       'name': testProductModel.name,
//   //       'description': testProductModel.description,
//   //       'imageUrl': testProductModel.imageUrl,
//   //       'price': testProductModel.price
//   //     };

//   //     when(mockHttpCliant.post(
//   //       Uri.parse(Urls.baseUrl),
//   //       body: productJson,
//   //     )).thenAnswer((_) async => http.Response('Something went wrong', 404));

//   //     // act
//   //     final call = productRemoteDateSourceImpl.createProduct;

//   //     // assert
//   //     expect(() => call(testProductModel), throwsA(isA<ServerException>()));
//   //   });

//   //   test('should throw a socket exception if it happens', () {
//   //     //arrange
//   //      final productJson = {
//   //       'name': testProductModel.name,
//   //       'description': testProductModel.description,
//   //       'imageUrl': testProductModel.imageUrl,
//   //       'price': testProductModel.price
//   //     };
//   //     when(mockHttpCliant.post(Uri.parse(Urls.baseUrl),
//   //             body: productJson,
//   //            ))
//   //         .thenThrow(const SocketException('No Internet Connection'));

//   //     //act
//   //     final call = productRemoteDateSourceImpl.createProduct;

//   //     //assert
//   //     expect(() => call(testProductModel), throwsA(isA<SocketException>()));
//   //   });
//   // });
// }
