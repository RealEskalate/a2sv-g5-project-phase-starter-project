import 'dart:io';

import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/exception.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:e_commerce_app/features/product/data/repositories/product_repository_implimentation.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRemoteDataSource mockProductRemoteDataSource;
  late ProductRepositoryImplimentation productRepositoryImplimentation;

  setUp(() {
    mockProductRemoteDataSource = MockProductRemoteDataSource();
    productRepositoryImplimentation = ProductRepositoryImplimentation(
        productRemoteDataSource: mockProductRemoteDataSource);
  });

  List<ProductModel> testProductModel = [
    ProductModel(
        description: "Explore anime characters addj.",
        id: "6672752cbd218790438efdb0",
        imageUrl:
            "D:/abd/A2SV/2024-internship-mobile-tasks/e_commerce_app/assets/download.jpeg",
        name: "Anime website",
        price: 123)
  ];
  List<ProductEntity> testProductEntity = [
    ProductEntity(
        description: "Explore anime characters addj.",
        id: "6672752cbd218790438efdb0",
        imageUrl:
            "D:/abd/A2SV/2024-internship-mobile-tasks/e_commerce_app/assets/download.jpeg",
        name: "Anime website",
        price: 123)
  ];
  
//1
  group("get all product scenarion", () {
    test("success to get all product", () async {
      //arrange
      when(mockProductRemoteDataSource.getAllProduct())
          .thenAnswer((_) async => testProductModel);

      //act
      final result = await productRepositoryImplimentation.getAllProduct();
      //assert
      result.fold((l) => expect(l, testProductEntity),
          (r) => expect(r, testProductEntity));
    });
    test(" failure get all product scenario", () async {
      //arrange
      when(mockProductRemoteDataSource.getAllProduct())
          .thenThrow(ServerException());

      //act
      final result = await productRepositoryImplimentation.getAllProduct();
      //assert
      result.fold((l) => expect(l, ServerFailure("message")),
          (r) => expect(r, testProductEntity));
    });
    test("connection error", () async {
      //arrange
      when(mockProductRemoteDataSource.getAllProduct())
          .thenThrow(SocketException("Failed to connect"));

      //act
      final result = await productRepositoryImplimentation.getAllProduct();
      //assert
    result.fold((l) => expect(l, ConnectionFailure("Failed to connect to internet")),
          (r) => expect(r, testProductEntity));
    });
  });

  //2 get one product test cases
  group("get all product scenarion", () {});
  group("get all product scenarion", () {
    test("get all product scenario", () async {
      //arrange

      //act

      //assert
    });
  });

  // 3 insert one product 
  group("get all product scenarion", () {
    test("get all product scenario", () async {
      //arrange

      //act

      //assert
    });
  });

  // 4  update product 
  group("get all product scenarion", () {
    test("get all product scenario", () async {
      //arrange

      //act

      //assert
    });
  });

  // 5  delete product 
  group("get all product scenarion", () {
    test("get all product scenario", () async {
      //arrange

      //act

      //assert
    });
  });




}
