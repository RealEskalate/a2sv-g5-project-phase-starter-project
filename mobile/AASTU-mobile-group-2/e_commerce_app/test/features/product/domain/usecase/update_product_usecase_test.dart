import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/update_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late UpdateProduct updateProduct;
  ProductEntity updatedProduct = ProductEntity(
      description: "updated shoe description",
      id: "id00013",
      imageUrl: "/assets/image.png",
      name: "Derby shoe",
      price: 27);

  setUp(() {
    mockProductRepository = MockProductRepository();
    updateProduct =
        UpdateProduct(mockProductRepository, updatedProduct, updatedProduct.id);
  });

  void testVal;
  Failure testFail = Failure("updating operation failed");

  test("success updating the product", () async {
    when(mockProductRepository.updateProduct(updatedProduct.id, updatedProduct))
        .thenAnswer((_) async => Right(updatedProduct));

    final result = await updateProduct.execute();
    expect(result, Right(testVal));
  });
  test("failure updating the product", () async {
    when(mockProductRepository.updateProduct(updatedProduct.id, updatedProduct))
        .thenAnswer((_) async => Left(testFail));

    final result = await updateProduct.execute();
    expect(result, Left(testFail));
  });
}
