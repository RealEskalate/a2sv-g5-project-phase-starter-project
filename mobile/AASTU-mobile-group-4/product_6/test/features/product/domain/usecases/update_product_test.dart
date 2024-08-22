import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/domain/entities/product.dart';
import 'package:product_6/features/product/domain/usecases/update_product.dart';

import '../../../../mock.mocks.dart';


void main() {
  late UpdateProductUseCase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = UpdateProductUseCase(mockProductRepository);
  });

  const Product tProduct = Product(
    id: '1',
    name: 'Product 1',
    description: 'Description 1',
    imageUrl: 'http://example.com/image1.jpg',
    price: 99.99,
  );

  test('should update a product in the repository', () async {
    // arrange
    when(mockProductRepository.updateProduct(any))
        .thenAnswer((_) async => const Right(tProduct));

    // act
    final result = await usecase(tProduct);

    // assert
    expect(result, const Right(tProduct));
    verify(mockProductRepository.updateProduct(tProduct));
    verifyNoMoreInteractions(mockProductRepository);
  });

  test('should return failure when repository call is unsuccessful', () async {
    // arrange
    const tFailure = ServerFailure('Failed to update product');
    when(mockProductRepository.updateProduct(any))
        .thenAnswer((_) async => Left(tFailure));

    // act
    final result = await usecase(tProduct);

    // assert
    expect(result, Left(tFailure));
    verify(mockProductRepository.updateProduct(tProduct));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
