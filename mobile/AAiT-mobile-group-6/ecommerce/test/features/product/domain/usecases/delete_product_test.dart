// ignore_for_file: prefer_const_declarations

import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/domain/usecase/delete_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late DeleteProductUsecase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = DeleteProductUsecase(mockProductRepository);
  });

  final tProductId = '1';
  final tProduct = const Product(
      id: '1',
      name: 'name',
      description: 'description',
      price: 1.0,
      imageUrl: 'imageUrl',
  );

  test('should delete product from the repository', () async {
    // Arrange
    when(mockProductRepository.deleteProduct(any))
        .thenAnswer((_) async => Right(tProduct));

    // Act
    final result = await usecase.execute( tProductId);

    // Assert
    expect(result, Right(tProduct));
    verify(mockProductRepository.deleteProduct(tProductId));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
