// ignore_for_file: prefer_const_declarations

import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/domain/usecase/get_all_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late GetAllProductUsecase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = GetAllProductUsecase(mockProductRepository);
  });

  final tProducts = [
    const Product(
        id: '1',
        name: 'Nike Air Max 270',

        price: 300.00,
        imageUrl: 'images/nike.jpg',
        description:
            'footwear option characterized by its open lacing system, where the shoelace eyelets are sewn on top of the vamp (the upper part of the shoe). This design feature provides a more relaxed and casual look compared to the closed lacing system of oxford shoes. Derby shoes are typically made of high-quality leather, known for its durability and elegance, making them suitable for both formal and casual occasions. With their timeless style and comfortable fit, derby leather shoes are a staple in any well-rounded wardrobe.',
      ),
    const Product(
        id: '2',
        name: 'Nike Air Max 270',

        price: 300.00,
        imageUrl: 'images/nike.jpg',
        description:
            'footwear option characterized by its open lacing system, where the shoelace eyelets are sewn on top of the vamp (the upper part of the shoe). This design feature provides a more relaxed and casual look compared to the closed lacing system of oxford shoes. Derby shoes are typically made of high-quality leather, known for its durability and elegance, making them suitable for both formal and casual occasions. With their timeless style and comfortable fit, derby leather shoes are a staple in any well-rounded wardrobe.',
    ),
  ];

  test('should get all products from the repository', () async {
    // Arrange
    when(mockProductRepository.getAllProduct())
        .thenAnswer((_) async => Right(tProducts));

    // Act
    final result = await usecase.execute(NoParams());

    // Assert
    expect(result, Right(tProducts));
    verify(mockProductRepository.getAllProduct());
    verifyNoMoreInteractions(mockProductRepository);
  });
}
