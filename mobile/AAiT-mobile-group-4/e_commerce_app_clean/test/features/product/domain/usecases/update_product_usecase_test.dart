import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/domain/usecases/update_product_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late UpdateProductUsecase updateProductUsecase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    updateProductUsecase = UpdateProductUsecase(mockProductRepository);
  });

  const tProductEntity =  ProductEntity(
    id: '3',
    name: 'airjordan',
    description: 'something you wear',
    price: 566,
    imageUrl: 'https://www.google.com',
  );
  test(
    'should update Product from ProductRepository and return a bool',
    () async {
      // Arrange
      when(mockProductRepository.updateProduct(tProductEntity))
          .thenAnswer((_) async => const Right(tProductEntity));
      // Act
      final result = await updateProductUsecase(const UpdateParams(product: tProductEntity));
      // Assert
      expect(result, const Right(tProductEntity));
    },
  );
}
