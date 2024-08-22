// import 'dart:io';

// import 'package:dartz/dartz.dart';
// import 'package:ecommerse2/core/error/exception.dart';
// import 'package:ecommerse2/core/error/failure.dart';
// import 'package:ecommerse2/features/product/data/models/product_model.dart';
// import 'package:ecommerse2/features/product/data/repositories/product_repository_impl.dart';
// import 'package:ecommerse2/features/product/domain/entity/product.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';

// import '../../helpers/test_helper.mocks.dart';

// void main() {
//   late ProductRepositoryImpl productRepositoryImpl;
//   late MockProductRemoteDataSource mockProductRemoteDataSource;

//   setUp(() {
//     mockProductRemoteDataSource = MockProductRemoteDataSource();
//     productRepositoryImpl = ProductRepositoryImpl(mockProductRemoteDataSource);
//   });

 
  
//   List<Product> products = [
//     const Product(
//         id: '1',
//         name: 'Nike',
//         category: 'Shoe',
//         description: 'A great Shoe',
//         image: 'The Nike',
//         price: 99),
//     const Product(
//         id: '1',
//         name: 'Nike',
//         category: 'Shoe',
//         description: 'A great Shoe',
//         image: 'The Nike',
//         price: 99)
//   ];
//   List<ProductModel> productModels = [
//     const ProductModel(
//         id: '1',
//         name: 'Nike',
//         category: 'Shoe',
//         description: 'A great Shoe',
//         image: 'The Nike',
//         price: 99),
//     const ProductModel(
//         id: '1',
//         name: 'Nike',
//         category: 'Shoe',
//         description: 'A great Shoe',
//         image: 'The Nike',
//         price: 99)
//   ];
//   final product =  ProductModel.fromJson(const {
//       "id": "6672940692adcb386d593686",
//       "name": "PC",
//       "description": "long description",
//       "price": 123,
//       "imageUrl": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718785031/images/zqfvuxrxhip7shikyyj4.png"
//     }).toEntity();
//   ProductModel productModel = ProductModel.fromJson(const {
//       "id": "6672940692adcb386d593686",
//       "name": "PC",
//       "description": "long description",
//       "price": 123,
//       "imageUrl": "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718785031/images/zqfvuxrxhip7shikyyj4.png"
//     });


//   //There are three scenarios to be tested
//   // 1. returns products on successful API request
//   // 2. returns to server failure when the request to the API fails
//   // 3. returns connection failure when not connected with the internet

//   group('get all products', () {
//     test('should return all products when a call to data source is successfull',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getAllProduct())
//           .thenAnswer((_) async => productModels);

//       //act
//       final result = await productRepositoryImpl.getAllProduct();

//       final productsFromRepository = result.fold(
//         (failure) => null, // Handle failure case
//         (products) => products,
//       );

//       //assert
//       expect(productsFromRepository, equals(products));
//     });

//     test(
//         'should return server failure when a call to connection is unsuccessfull',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getAllProduct())
//           .thenThrow(ServerException());

//       //act
//       final result = await productRepositoryImpl.getAllProduct();

//       //assert
//       expect(
//           result, equals(const Left(ServerFailure('An error has occurred'))));
//     });
//     test(
//         'should return connection failure when a call to connection is unsuccessfull',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getAllProduct())
//           .thenThrow(const SocketException('An error has occurred'));

//       //act
//       final result = await productRepositoryImpl.getAllProduct();

//       //assert
//       expect(
//           result, equals(const Left(ConnectionFailure('Failed to connect'))));
//     });
//   });

//   group('get product by id', () {
//     test('should return product when a call to data source is successful',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getProduct(productModel.id))
//           .thenAnswer((_) async => productModel);

//       //act
//       final result = await productRepositoryImpl.getProduct(productModel.id);

//       //assert
//       expect(result, equals(Right(product)));
//     });

//     test(
//         'should return server failure when a call to connection is unsuccessfull',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getProduct(productModel.id))
//           .thenThrow(ServerException());

//       //act
//       final result = await productRepositoryImpl.getProduct(productModel.id);

//       //assert
//       expect(
//           result, equals(const Left(ServerFailure('An error has occurred'))));
//     });
//     test(
//         'should return connection failure when a call to connection is unsuccessfull',
//         () async {
//       //arrange
//       when(mockProductRemoteDataSource.getProduct(productModel.id))
//           .thenThrow(const SocketException('An error has occurred'));

//       //act
//       final result = await productRepositoryImpl.getProduct(productModel.id);

//       //assert
//       expect(
//           result, equals(const Left(ConnectionFailure('Failed to connect'))));
//     });
//   });

//   group('delete the product', () {
//        test('should return the product', () async { 
//       when(mockProductRemoteDataSource.deleteProduct(productModel.id))
//           .thenAnswer((_) async {});

//       final result = await productRepositoryImpl.deleteProduct(productModel.id);

//       expect(result, equals(const Right(null)));
//     });

     
//   });

//   group('update the product', (){

//   });

//   group('add product', (){

//   });
// }
