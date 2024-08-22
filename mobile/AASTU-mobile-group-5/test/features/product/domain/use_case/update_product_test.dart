import 'dart:io';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/features/product/data/models/product_model.dart';
import 'package:task_9/features/product/domain/repository/product_repository.dart';
import 'package:task_9/features/product/domain/use_case/update_product.dart';
import 'add_product_test.mocks.dart';


class MockFile extends Mock implements File {}

@GenerateMocks([ProductRepository])
void main() {
  late MockProductRepository productRepository;
  late UpdateProduct usecase;

  setUp(() {
    productRepository = MockProductRepository();
    usecase = UpdateProduct(productRepository);
  });
    const productModel = ProductModel(
    name: 'Boots',
    id: '1',
    price: 200,
    imageUrl: 'assets/images/boots.jpg',
    description:
        'These boots are made for walking and that\'s just what they\'ll do one of these days these boots are gonna walk all over you',
  );


  test('should update a product to the repository', () async {
    // Arrange
    when(productRepository.updateProduct(productModel))
        .thenAnswer((_) async => const Right(productModel));

    // Act
    final result = await usecase(UpdateProductParams(productModel));

    // Assert
    expect(result, const Right(productModel));
    verify(productRepository.updateProduct(productModel));
    verifyNoMoreInteractions(productRepository);
  });
}
