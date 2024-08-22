import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_6/core/errors/failure.dart';
import 'package:product_6/features/product/domain/usecases/delete_product.dart';

import '../../../../mock.mocks.dart';


void main() {
  late DeleteProductUseCase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = DeleteProductUseCase(mockProductRepository);
  });

  const String tProductId = '1';

  test('should delete a product from the repository', () async {
    // arrange
    when(mockProductRepository.deleteProduct(any))
        .thenAnswer((_) async => const Right(null));

    // act
    final result = await usecase(tProductId);

    // assert
    expect(result, const Right(null));
    verify(mockProductRepository.deleteProduct(tProductId));
    verifyNoMoreInteractions(mockProductRepository);
  });

  test('should return failure when repository call is unsuccessful', () async {
    // arrange
    const tFailure = ServerFailure('Failed to delete product');
    when(mockProductRepository.deleteProduct(any))
        .thenAnswer((_) async => const Left(tFailure));

    // act
    final result = await usecase(tProductId);

    // assert
    expect(result, const Left(tFailure));
    verify(mockProductRepository.deleteProduct(tProductId));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
