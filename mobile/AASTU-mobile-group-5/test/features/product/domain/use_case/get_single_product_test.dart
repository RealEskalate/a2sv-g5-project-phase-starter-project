import 'dart:io';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import 'package:task_9/features/product/domain/entities/product.dart';
import 'package:task_9/features/product/domain/repository/product_repository.dart';
import 'package:task_9/features/product/domain/use_case/get_product.dart';
import 'get_single_product_test.mocks.dart';

class MockFile extends Mock implements File {}

@GenerateMocks([ProductRepository])
void main() {
  late MockProductRepository productRepository;
  late GetProduct usecase;

  setUp(() {
    productRepository = MockProductRepository();
    usecase = GetProduct(productRepository);
  });
  const product = Product(
    name: 'Boots',
    id: '1',
    price: 200,
    imageUrl: 'assets/images/boots.jpg',
    description:
        'These boots are made for walking and that\'s just what they\'ll do one of these days these boots are gonna walk all over you',
  );
  test('should display a single product from the repository', () async {
    // Arrange
    when(productRepository.getProductById(product.id))
        .thenAnswer((_) async => const Right(product));

    // Act
    final result = await usecase(GetProductParams(product.id));

    // Assert
    expect(result, equals(const Right(product)));
    verify(productRepository.getProductById(product.id));
    verifyNoMoreInteractions(productRepository);
  });
}
