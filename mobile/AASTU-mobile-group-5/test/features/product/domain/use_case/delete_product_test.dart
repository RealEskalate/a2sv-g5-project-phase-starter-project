import 'dart:io';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';

import 'package:task_9/features/product/domain/entities/product.dart';
import 'package:task_9/features/product/domain/repository/product_repository.dart';
import 'package:task_9/features/product/domain/use_case/delete_product.dart';
import 'delete_product_test.mocks.dart';

class MockFile extends Mock implements File {}

@GenerateMocks([ProductRepository])
void main() {
  late MockProductRepository productRepository;
  late DeleteProduct usecase;

  setUp(() {
    productRepository = MockProductRepository();
    usecase = DeleteProduct(productRepository);
  });
  const product = Product(
    name: 'Boots',
    id: '1',
    price: 200,
    imageUrl: 'assets/images/boots.jpg',
    description:
        'These boots are made for walking and that\'s just what they\'ll do one of these days these boots are gonna walk all over you',
  );
  test('should delete a product from the repository', () async {
    //Arrange
    when(productRepository.deleteProduct(product.id))
        // ignore: void_checks
        .thenAnswer((_) async => const Right(null));

    // Act
    usecase(DeleteProductParams(product.id));

    // Assert
    
    verify(productRepository.deleteProduct(product.id));
    verifyNoMoreInteractions(productRepository);
  });
}
