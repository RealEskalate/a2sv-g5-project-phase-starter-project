// import 'package:dartz/dartz.dart';
// import 'package:e_commerce_app/core/failure/failure.dart';
// import 'package:e_commerce_app/features/product/domain/enteties/product.dart';
// import 'package:e_commerce_app/features/product/domain/usecase/insert_product_usecase.dart';
// import 'package:flutter_test/flutter_test.dart';
// import 'package:mockito/mockito.dart';

// import '../../../../helpers/test_helper.mocks.dart';

// void main() {
//   late MockProductRepository mockProductRepository;
//   late InsertProduct insertProduct;

//   ProductEntity newProduct = ProductEntity(
//       description: "description",
//       id: "id",
//       imageUrl: "imageUrl",
//       name: "name",
//       price: 400);

//   setUp(() {
//     mockProductRepository = MockProductRepository();
//     insertProduct = InsertProduct(mockProductRepository, newProduct);
//   });
//   void testVal;
//   Failure testFail = Failure("Failed to insert produt");

//   test("Success insert operation", () async {
//     when(mockProductRepository.insertProduct(newProduct))
//         .thenAnswer((_) async => Right(newProduct));

//     final result = await insertProduct.execute();

//     expect(result, Right(testVal));
//   });
//   test("Failure insert operation", () async {
//     when(mockProductRepository.insertProduct(newProduct))
//         .thenAnswer((_) async => Left(testFail));

//     final result = await insertProduct.execute();

//     expect(result, Left(testFail));
//   });



// }
