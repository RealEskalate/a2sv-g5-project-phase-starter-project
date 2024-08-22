import 'dart:convert';
import 'dart:io';

import 'package:e_commerce_app/core/constants/constants.dart';
import 'package:e_commerce_app/core/failure/exception.dart';
import 'package:e_commerce_app/features/product/data/data_sources/product_remote_data_source.dart';
import 'package:e_commerce_app/features/product/data/models/product_model.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:http/http.dart' as http;

import '../../../../helpers/json_reader.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockHttpClient mockHttpClient;
  late ProductRemoteDataSource productRemoteDataSource;

  setUp(() {
    mockHttpClient = MockHttpClient();
    productRemoteDataSource = ProductRemoteDataSource(client: mockHttpClient);
  });
  //1 Get all product test
  group("get all product api request test", () {
    test("get all product data success ", () async {
      when(mockHttpClient.get(Uri.parse(Urls.baseUrl))).thenAnswer((_) async =>
          http.Response(
              readJson("helpers/dummy/dummy_product_list.json"), 200));

      //act
      final result = await productRemoteDataSource.getAllProduct();
//e
      expect(result, isA<List<ProductModel>>());
    });
    test("failure to get all prodcut data ", () async {
      when(mockHttpClient.get(Uri.parse(Urls.baseUrl)))
          .thenAnswer((_) async => http.Response('not found', 404));

      //act
      final result = productRemoteDataSource.getAllProduct();
//e
      expect(result, throwsA(isA<ServerException>()));
    });
  });

// 2 get one product test

  String id = "6672752cbd218790438efdb0";
  group("get one product api request test", () {
    test("success to get one product data ", () async {
      //assert
      when(mockHttpClient.get(Uri.parse(Urls.getProductById(id)))).thenAnswer(
          (_) async =>
              http.Response(readJson("helpers/dummy/dummy_product.json"), 200));

      //act
      final result = await productRemoteDataSource.getOneProduct(id);
//expect
      expect(result, isA<ProductModel>());
    });
    test("not found to get product model ", () async {
      //assert
      when(mockHttpClient.get(Uri.parse(Urls.getProductById(id))))
          .thenAnswer((_) async => http.Response('not found', 404));

      //act
      final result = productRemoteDataSource.getOneProduct(id);
//expect
      expect(result, throwsA(isA<ServerException>()));
    });
  });

  ProductModel testModel = ProductModel(
      description: "Explore anime characters.",
      id: "6672752cbd218790438efdb0",
      imageUrl:
          "https://res.cloudinary.com/g5-mobile-track/image/upload/v1718777132/images/zxjhzrflkvsjutgbmr0f.jpg",
      name: "Anime website",
      price: 123);
  ProductModel testNewProduct = ProductModel(
      description: "Explore anime characters addj.",
      id: "6672752cbd218790438efdb0",
      imageUrl:
          "D:/abd/A2SV/2024-internship-mobile-tasks/e_commerce_app/assets/download.jpeg",
      name: "Anime website",
      price: 123);

// 3 add new roduct test case
  group("test case to insert new product ", () {
      File image = File(testNewProduct.imageUrl);
      //assert
      var imageBytes = image.readAsBytesSync();
    test("success to insert new product ", () async {
    

      when(mockHttpClient.post(
        Uri.parse(Urls.baseUrl),
        headers: {"Content-Type": "multipart/form-data"},
        body: {
          "image": imageBytes,
          "name": testNewProduct.name,
          "description": testNewProduct.description,
          "price": testNewProduct.price,

        },
        encoding: Encoding.getByName("utf-8")
      )).thenAnswer((_) async =>
          http.Response(readJson("helpers/dummy/dummy_product.json"), 201));

      //act
      final result =
          await productRemoteDataSource.insertProduct(testNewProduct);
//expect
      expect(result, isA<ProductModel>());
    });
    test("Failure to insert new product ", () async {
      //assert
      when(mockHttpClient.post(
        Uri.parse(Urls.baseUrl),
        headers: {"Content-Type": "multipart/form-data"},
        body: {
          "image": imageBytes,
          "name": testNewProduct.name,
          "description": testNewProduct.description,
          "price": testNewProduct.price,

        },
        encoding: Encoding.getByName("utf-8")
      ))
          .thenAnswer((_) async => http.Response("not inserted", 404));

      //act
      final result = productRemoteDataSource.insertProduct(testNewProduct);
//expect
      expect(result, throwsA(isA<ServerException>()));
    });
  });

// 4 add new roduct test case
  group("test case to update new product ", () {
    test("success to update new product ", () async {
      //assert

      when(mockHttpClient.put(Uri.parse(Urls.getProductById(testModel.id)),
              body: testModel.toJson()))
          .thenAnswer((_) async =>
              http.Response(readJson("helpers/dummy/dummy_product.json"), 200));

      //act
      final result = await productRemoteDataSource.updateProduct(testModel);
//expect
      expect(result, isA<ProductModel>());
    });
    test("Failure to update new product ", () async {
      //assert

      when(mockHttpClient.put(Uri.parse(Urls.getProductById(testModel.id)),
              body: testModel.toJson()))
          .thenAnswer((_) async => http.Response('product not updated', 404));

      //act
      final result = productRemoteDataSource.updateProduct(testModel);
//expect
      expect(result, throwsA(isA<ServerException>()));
    });
  });

// 5 add new roduct test case
  group("test case to delete new product ", () {
    test("success to delete new product ", () async {
      //assert

      when(mockHttpClient.delete(Uri.parse(Urls.getProductById(testModel.id))))
          .thenAnswer((_) async => http.Response('deleted successfully', 200));

      //act
      final result = await productRemoteDataSource.deleteProduct(testModel.id);
//expect
      expect(result, isA<String>());
    });
    test("Failure to delete new product ", () async {
      //assert

      when(mockHttpClient.delete(Uri.parse(Urls.getProductById(testModel.id))))
          .thenAnswer((_) async => http.Response('product not deleted', 404));

      //act
      final result = productRemoteDataSource.deleteProduct(testModel.id);
//expect
      expect(result, throwsA(isA<ServerException>()));
    });
  });
}
