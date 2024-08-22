import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/domain/entities/product.dart';
import 'package:product_6/features/product/domain/usecases/view_product.dart';

import '../../../../mock.mocks.dart';

void main() {
  late ViewProductUseCase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = ViewProductUseCase(mockProductRepository);
  });

  const String tProductId = '1';
  const Product tProduct = Product(
    id: tProductId,
    name: 'Product 1',
    description: 'Description 1',
    imageUrl: 'http://example.com/image1.jpg',
    price: 99.99,
  );

  test('should return product from repository when successful', () async {
    // arrange
    when(mockProductRepository.getProductById(any))
        .thenAnswer((_) async => const Right(tProduct));

    // act
    final result = await usecase(tProductId);

    // assert
    expect(result, const Right(tProduct));
    verify(mockProductRepository.getProductById(tProductId));
    verifyNoMoreInteractions(mockProductRepository);
  });

  test('should return failure when repository call is unsuccessful', () async {
    // arrange
    const tFailure = ServerFailure('Failed to get product');
    when(mockProductRepository.getProductById(any))
        .thenAnswer((_) async => const Left(tFailure));

    // act
    final result = await usecase(tProductId);

    // assert
    expect(result, const Left(tFailure));
    verify(mockProductRepository.getProductById(tProductId));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
