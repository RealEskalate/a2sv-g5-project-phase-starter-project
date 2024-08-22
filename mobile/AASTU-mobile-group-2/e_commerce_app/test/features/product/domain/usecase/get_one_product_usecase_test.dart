import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/get_one_product_usecase..dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late GetOneProduct getOneProduct;
  ProductEntity productEntity = ProductEntity(
      description: "description",
      id: "id12",
      imageUrl: "assets/image.png",
      name: "Red shoes",
      price: 35);
  Failure testFail = Failure("failed to get one product");
  setUp(() {
    mockProductRepository = MockProductRepository();
    getOneProduct = GetOneProduct(mockProductRepository, productEntity.id);
  });

  test("success to get one product", () async {
    when(mockProductRepository.getOneProduct(productEntity.id))
        .thenAnswer((_) async => Right(productEntity));

    final result = await getOneProduct.execute();
    expect(result, Right(productEntity));
  });
  test("Failure to get one product", () async {
    when(mockProductRepository.getOneProduct(productEntity.id))
        .thenAnswer((_) async => Left(testFail));

    final result = await getOneProduct.execute();
    expect(result, Left(testFail));
  });
}
