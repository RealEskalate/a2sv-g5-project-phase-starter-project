import 'dart:convert';
import 'dart:io';

import 'package:flutter_test/flutter_test.dart';
import 'package:http/http.dart' as http;
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/core/constants/api_url.dart';
import 'package:task_9/core/error/exceptions.dart';
import 'package:task_9/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:task_9/features/product/data/models/product_model.dart';
import 'package:task_9/features/user/data/data_sources/user_local_data_source.dart';

import 'product_remote_data_source_test.mocks.dart';

@GenerateMocks([http.Client, UserLocalDataSource])
void main() {
  late ProductRemoteDataSourceImpl dataSource;
  late MockClient mockHttpClient;
  late MockUserLocalDataSource mockUserLocalDataSource;
  

  setUp(() {
    mockHttpClient = MockClient();
    mockUserLocalDataSource = MockUserLocalDataSource();
    dataSource = ProductRemoteDataSourceImpl(
      client: mockHttpClient,
      userLocalDataSource: mockUserLocalDataSource,
    );
  });

  const String accessToken = 'test_access_token';

  setUp(() {
    when(mockUserLocalDataSource.getAccessToken())
        .thenAnswer((_) async => accessToken);
  });

  final jsonResponse = jsonDecode(File(
          'D:/2024-internship-mobile-tasks/mobile/aryam_ezra/task_9/test/dummy_test/dummy_json.json')
      .readAsStringSync());

  const tProductModel = ProductModel(
    id: '6672776eb905525c145fe0bb',
    name: 'Anime website',
    description: 'Explore anime characters.',
    price: 123,
    imageUrl:
        'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg',
  );
  const tImagePath = 'assets/images/boots.jpg';

  group('getAllProducts', () {
    test('should return a list of ProductModels when the response code is 200',
      () async {
    // Arrange
    when(mockHttpClient.get(
      Uri.parse(Urls.getAllProducts()),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
    )).thenAnswer(
        (_) async => http.Response(jsonEncode(jsonResponse), 200));

    // Act
    final result = await dataSource.getAllProducts();

    // Assert
    expect(result, isA<List<ProductModel>>());
  });

    test(
        'should throw a ServerException when the response code is 404 or other',
        () async {
      // Arrange
      when(mockHttpClient.get(
        Uri.parse(Urls.getAllProducts()),
        headers: {'Authorization': 'Bearer $accessToken'},
      )).thenAnswer((_) async => http.Response('Not Found', 404));

      // Act
      final call = dataSource.getAllProducts();

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });
  });

  const String productId = '6672776eb905525c145fe0bb';

  group('getProductById', () {
final resp = jsonDecode(File(
          'D:/2024-internship-mobile-tasks/mobile/aryam_ezra/task_9/test/dummy_test/dummy_single_product.json')
      .readAsStringSync());
  test('should return a product when the response code is 200', () async {
    // Arrange
    when(mockHttpClient.get(
      Uri.parse(Urls.getProduct(productId)),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
    )).thenAnswer((_) async => http.Response(jsonEncode(resp), 200));

    // Act
    final result = await dataSource.getProductById(productId);

    // Assert
    expect(result, isA<ProductModel>());
  });



  test('should throw ProductNotFoundException when the response code is 404', () async {
  // Arrange
  when(mockHttpClient.get(
    Uri.parse(Urls.getProduct(productId)),
    headers: {
      'Content-Type': 'application/json',
      'Authorization': 'Bearer $accessToken',
    },
  )).thenAnswer((_) async => http.Response('Not Found', 404));

  // Act
  final call = dataSource.getProductById(productId);

  // Assert
  expect(() => call, throwsA(isA<ProductNotFoundException>()));
});


    test('should throw ServerException when an unexpected error occurs',
        () async {
      // Arrange
      when(mockHttpClient.get(
        Uri.parse(Urls.getProduct(productId)),
        headers: {'Authorization': 'Bearer $accessToken'},
      )).thenAnswer((_) async => throw ServerException());

      // Act
      final call = dataSource.getProductById('non-existing-id');

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });

  });

  group('deleteProduct', () {
    test('should delete successfully when the response code is 200',
        () async {
      // Arrange
      when(mockHttpClient.delete(
      Uri.parse(Urls.deleteProduct(productId)),
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
    )).thenAnswer(
        (_) async => http.Response(jsonEncode(jsonResponse), 200));

      // Act
      await dataSource.deleteProduct(productId);

      // Assert
      // Since there's no return value for successful deletion, no assertion is needed
    });

    test('should throw ProductNotFoundException when the response code is 404',
        () async {
      // Arrange
      when(mockHttpClient.delete(
        Uri.parse(Urls.deleteProduct('non-existing-id')),
       headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
      )).thenAnswer((_) async => http.Response('Not Found', 404));

      // Act
      final call = dataSource.deleteProduct('non-existing-id');

      // Assert
      expect(() => call, throwsA(isA<ProductNotFoundException>()));
    });

    test('should throw ServerException for other error codes', () async {
      // Arrange
      when(mockHttpClient.delete(
        Uri.parse(Urls.deleteProduct(productId)),
        headers: {'Authorization': 'Bearer $accessToken'},
      )).thenAnswer((_) async => http.Response('Server Error', 500));

      // Act
      final call = dataSource.deleteProduct(productId);

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });
  });

  group('addProduct', () {
  /* */
  // test('should send a POST request to add the product', () async {
  //   final expectedUrl = Uri.parse(Urls.addProduct());
  //   when(mockHttpClient.post(
  //     expectedUrl,
  //     headers: {
  //       'Content-Type': 'application/json',
  //       'Authorization': 'Bearer $accessToken',
  //     },
  //     body: jsonEncode(tProductModel.toJson()),
  //   )).thenAnswer(
  //     (_) async => http.Response(jsonEncode({'data': tProductModel.toJson()}), 201),
  //   );

  //   // Act
  //   ProductModel result;
  //   try {
  //     result = await dataSource.addProduct(tProductModel, tImagePath);
  //   } catch (e) {
  //     fail('Exception thrown: $e');
  //   }

  //   // Assert
  //   verify(mockHttpClient.post(
  //     expectedUrl,
  //     headers: {
  //       'Content-Type': 'application/json',
  //       'Authorization': 'Bearer $accessToken',
  //     },
  //     body: jsonEncode(tProductModel.toJson()),
  //   ));
  //   expect(result, tProductModel);
  // });

  test('should throw a ServerException when the response code is not 201', () async {
    // Arrange
    final expectedUrl = Uri.parse(Urls.addProduct());
    when(mockHttpClient.post(
      expectedUrl,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
      body: jsonEncode(tProductModel.toJson()), // Ensure this is the expected body
    )).thenAnswer(
        (_) async => http.Response('Something went wrong', 500));

    // Act
    final call = dataSource.addProduct(tProductModel, tImagePath);

    // Assert
    expect(() => call, throwsA(isA<ServerException>()));
  });
});

  group('updateProduct', () {

    test('should send a PUT request to update the product', () async {
    // Arrange
    final expectedUrl = Uri.parse(Urls.updateProduct(tProductModel.id));
    when(mockHttpClient.put(
      expectedUrl,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
      body: jsonEncode(tProductModel.toJson()), // Ensure this is the expected body
    )).thenAnswer(
        (_) async => http.Response(jsonEncode({'data': tProductModel.toJson()}), 200));

    // Act
    ProductModel result;
    try {
      result = await dataSource.updateProduct(tProductModel.id, tProductModel);
    } catch (e) {
      fail('Exception thrown: $e');
    }

    // Assert
    verify(mockHttpClient.put(
      expectedUrl,
      headers: {
        'Content-Type': 'application/json',
        'Authorization': 'Bearer $accessToken',
      },
      body: jsonEncode(tProductModel.toJson()),
    ));
    expect(result, tProductModel);
});


    test('should throw a ServerException when the response code is not 201',
        () async {
      // Arrange
      when(mockHttpClient.put(
        any,
        headers: {'Authorization': 'Bearer $accessToken'},
        body: anyNamed('body'),
      )).thenAnswer((_) async => http.Response('Something went wrong', 404));

      // Act
      final call = dataSource.updateProduct(tProductModel.id, tProductModel);

      // Assert
      expect(() => call, throwsA(isA<ServerException>()));
    });
  });
}

     


// verion 1 tests
// import 'dart:convert';
// import 'dart:io';

// import 'package:flutter_test/flutter_test.dart';
// import 'package:http/http.dart' as http;
// import 'package:mockito/annotations.dart';
// import 'package:mockito/mockito.dart';
// import 'package:task_9/core/error/exceptions.dart';
// import 'package:task_9/features/product/data/data_sources/product_remote_data_source.dart';
// import 'package:task_9/features/product/data/models/product_model.dart';

// import 'product_remote_data_source_test.mocks.dart';

// @GenerateMocks([http.Client])
// void main() {
//   late ProductRemoteDataSourceImpl dataSource;
//   late MockClient mockHttpClient;
  

//   setUp(() {
//     mockHttpClient = MockClient();
//     dataSource = ProductRemoteDataSourceImpl(client: mockHttpClient);
//   });
//   final jsonResponse = jsonDecode(File(
//           'D:/2024-internship-mobile-tasks/mobile/aryam_ezra/task_9/test/dummy_test/dummy_json.json')
//       .readAsStringSync());

//   const tProductModel = ProductModel(
//       id: '6672776eb905525c145fe0bb',
//       name: 'Anime website',
//       description: 'Explore anime characters.',
//       price: 123,
//       imageUrl:
//           'https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777711/images/clmxnecvavxfvrz9by4w.jpg');
//   const tImagePath = 'assets/images/boots.jpg';
//   group('getAllProducts', () {
//     test('should return a list of ProductModels when the response code is 200',
//         () async {
//       // Arrange

//       when(mockHttpClient.get(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products')))
//           .thenAnswer(
//               (_) async => http.Response(jsonEncode(jsonResponse), 200));

//       // Act
//       final result = await dataSource.getAllProducts();

//       // Assert
//       expect(result, isA<List<ProductModel>>());
//     });

//     test(
//         'should throw a ServerException when the response code is 404 or other',
//         () async {
//       // Arrange
//       when(mockHttpClient.get(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products')))
//           .thenAnswer((_) async => http.Response('Not Found', 404));

//       // Act
//       final call = dataSource.getAllProducts();

//       // Assert
//       expect(() => call, throwsA(isA<ServerException>()));
//     });
//   });

//   const String productId = '6672776eb905525c145fe0bb';

//   group('get Products by ID', () {
//     final resp = jsonDecode(File(
//             'D:/2024-internship-mobile-tasks/mobile/aryam_ezra/task_9/test/dummy_test/dummy_single_product.json')
//         .readAsStringSync());
//     test('should return a product when the response code is 200', () async {
//       // Arrange

//       when(mockHttpClient.get(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products/$productId')))
//           .thenAnswer((_) async => http.Response(jsonEncode(resp), 200));

//       // Act
//       final result = await dataSource.getProductById(productId);

//       // Assert
//       expect(result, isA<ProductModel>());
//     });

//     test('something expected occur', () async {
//       // Arrange
//       when(mockHttpClient.get(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products')))
//           .thenAnswer((_) async => throw ServerException());

//       // Act
//       final call = dataSource.getProductById('non-existing-id');

//       // Assert
//       expect(() => call, throwsA(isA<ServerException>()));
//     });

//     test('should throw ProductNotFoundException when the product is not found',
//         () async {
//       // Arrange
//       when(mockHttpClient.get(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products')))
//           .thenAnswer((_) async => http.Response('Not Found', 404));

//       // Act
//       final call = dataSource.getProductById('productId');

//       // Assert
//       expect(() => call, throwsA(isA<ServerException>()));
//     });
//   });

//   group('Delete Products by ID', () {
//     test('should complete successfully when the response code is 200',
//         () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products/productId')))
//           .thenAnswer((_) async =>
//               http.Response('', 200)); // Success response with empty body

//       // Act
//       dataSource.deleteProduct('productId');

//       // Assert
//       // Since there's no return value for successful deletion, no assertion is needed
//     });

//     test('should throw ProductNotFoundException when the response code is 404',
//         () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products/non-existing-id')))
//           .thenAnswer((_) async => http.Response('Not Found', 404));

//       // Act
//       final call = dataSource.deleteProduct('non-existing-id');

//       // Assert
//       expect(() => call, throwsA(isA<ProductNotFoundException>()));
//     });

//     test('should throw ServerException for other error codes', () async {
//       // Arrange
//       when(mockHttpClient.delete(Uri.parse(
//               'https://g5-flutter-learning-path-be.onrender.com/api/v1/products/productId')))
//           .thenAnswer((_) async => http.Response('Server Error', 500));

//       // Act
//       final call = dataSource.deleteProduct('productId');

//       // Assert
//       expect(() => call, throwsA(isA<ServerException>()));
//     });
//   });

//   group('addProduct', () {
//     test('should throw a ServerException when the response code is not 201',
//         () async {
//       // Arrange
//       when(mockHttpClient.send(any))
//           .thenAnswer((_) async => http.StreamedResponse(
//                 Stream.fromIterable([utf8.encode('Something went wrong')]),
//                 400,
//               ));

//       // Act
//       dataSource.addProduct;

//       // Assert
//       expect(() => dataSource.addProduct(tProductModel, tImagePath), throwsA(isA<ServerException>()));
  
//     });


    
//   });
//   group('updateProduct', () {
//     test('should send a PUT request to update the product', () async {
//       // Arrange
//       final expectedUrl = Uri.parse(
//           'https://g5-flutter-learning-path-be.onrender.com/api/v1/products/${tProductModel.id}');
//       when(mockHttpClient.put(
//         expectedUrl,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer(
//           (_) async => http.Response(jsonEncode(tProductModel.toJson()), 200));

//       // Act
//       final result =
//           await dataSource.updateProduct(tProductModel.id, tProductModel);

//       // Assert
//       verify(mockHttpClient.put(
//         expectedUrl,
//         headers: {'Content-Type': 'application/json'},
//         body: jsonEncode(tProductModel.toJson()),
//       ));
//       expect(result, tProductModel);
//     });

//     test('should throw a ServerException when the response code is not 200',
//         () async {
//       // Arrange
//       when(mockHttpClient.put(
//         any,
//         headers: anyNamed('headers'),
//         body: anyNamed('body'),
//       )).thenAnswer((_) async => http.Response('Something went wrong', 404));

//       // Act
//       final call = dataSource.updateProduct;

//       // Assert
//       expect(() => call(tProductModel.id, tProductModel),
//           throwsA(isA<ServerException>()));
//     });
//   });
// }
