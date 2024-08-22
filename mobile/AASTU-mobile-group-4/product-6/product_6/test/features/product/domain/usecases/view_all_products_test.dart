import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/domain/entities/product.dart';
import 'package:product_6/features/product/domain/usecases/view_all_products.dart';

import '../../../../mock.mocks.dart';




void main() {
  late ViewAllProductsUseCase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = ViewAllProductsUseCase(mockProductRepository);
  });

  const List<Product> tProducts = [
    Product(
      id: '1',
      name: 'Product 1',
      description: 'Description 1',
      imageUrl: 'http://example.com/image1.jpg',
      price: 99.99,
    ),
    Product(
      id: '1',
      name: 'Product 2',
      description: 'Description 2',
      imageUrl: 'http://example.com/image2.jpg',
      price: 49.99,
    ),
  ];

  test('should get list of products from the repository', () async {
    // arrange
    when(mockProductRepository.getProducts())
        .thenAnswer((_) async => const Right(tProducts));

    // act
    final result = await usecase();

    // assert
    expect(result, const Right(tProducts));
    verify(mockProductRepository.getProducts());
    verifyNoMoreInteractions(mockProductRepository);
  });

  test('should return failure when repository call is unsuccessful', () async {
    // arrange
    const tFailure = ServerFailure('Failed to load products');
    when(mockProductRepository.getProducts())
        .thenAnswer((_) async => const Left(tFailure));

    // act
    final result = await usecase();

    // assert
    expect(result, const Left(tFailure));
    verify(mockProductRepository.getProducts());
    verifyNoMoreInteractions(mockProductRepository);
  });
}
