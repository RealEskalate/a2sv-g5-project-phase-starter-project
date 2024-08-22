// ignore_for_file: prefer_const_declarations

import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/domain/usecase/Update_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late UpdateProductUsecase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = UpdateProductUsecase(mockProductRepository);
  });

  
  final tProduct = const Product(
      id: '1',
      name: 'name',
      description: 'description',
      price: 1.0,
      imageUrl: 'imageUrl',
);

  test('should update product from the repository', () async {
    // Arrange
    when(mockProductRepository.updateProduct(any))
        .thenAnswer((_) async => Right(tProduct));

    // Act
    final result = await usecase.execute(tProduct);

    // Assert
    expect(result, Right(tProduct));
    verify(mockProductRepository.updateProduct(tProduct));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
