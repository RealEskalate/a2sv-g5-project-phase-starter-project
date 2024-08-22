import 'dart:io';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/annotations.dart';
import 'package:mockito/mockito.dart';
import 'package:task_9/features/product/data/models/product_model.dart';
import 'package:task_9/features/product/domain/repository/product_repository.dart';
import 'package:task_9/features/product/domain/use_case/add_product.dart';
import 'add_product_test.mocks.dart';


class MockFile extends Mock implements File {}

@GenerateMocks([ProductRepository])
void main() {
  late MockProductRepository productRepository;
  late AddProduct usecase;

  setUp(() {
    productRepository = MockProductRepository();
    usecase = AddProduct(productRepository);
  });
    const productModel = ProductModel(
    name: 'Boots',
    id: '1',
    price: 200,
    imageUrl: 'assets/images/boots.jpg',
    description:
        'These boots are made for walking and that\'s just what they\'ll do one of these days these boots are gonna walk all over you',
  );

  const imagePath = 'assets/images/boots.jpg';

  test('should add a product to the repository', () async {
    // Arrange
    when(productRepository.addProduct(productModel, imagePath))
        .thenAnswer((_) async => const Right(productModel));

    // Act
    final result = await usecase(AddProductParams(productModel, imagePath));

    // Assert
    expect(result, const Right(productModel));
    verify(productRepository.addProduct(productModel, imagePath));
    verifyNoMoreInteractions(productRepository);
  });
}
