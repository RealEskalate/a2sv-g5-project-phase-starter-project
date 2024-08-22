import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/delete_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
 late  DeleteProduct deleteProduct;
  ProductEntity productEntity = ProductEntity(
      description: "description",
      id: "id",
      imageUrl: "imageUrl",
      name: "name",
      price: 400);

  setUp(() {
    mockProductRepository = MockProductRepository();

    deleteProduct = DeleteProduct(mockProductRepository, productEntity.id);
  });

  void testVal;
  Failure testFail = Failure("Erro occurred while deleting");
  test('trying to delete ', () async {
    when(mockProductRepository.deleteProduct(productEntity.id))
        .thenAnswer((_) async => Right("Success"));

    final result = await mockProductRepository.deleteProduct(productEntity.id);
    expect(result, Right(testVal));
  });

  test('Filure while trying to delete ', () async {
    when(mockProductRepository.deleteProduct(productEntity.id))
        .thenAnswer((_) async => Left(testFail));

    final result = await deleteProduct.execute();
    expect(result, Left(testFail));
  });
}
