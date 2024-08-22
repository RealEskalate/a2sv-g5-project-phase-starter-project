// ignore_for_file: prefer_const_declarations

import 'package:dartz/dartz.dart';
import 'package:ecommerce/features/product/domain/entitity/product.dart';
import 'package:ecommerce/features/product/domain/usecase/insert_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.mocks.dart';

void main() {
  late InsertProductUsecase usecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    usecase = InsertProductUsecase(mockProductRepository);
  });


  final tProduct = const Product(
      id: '1',
      name: 'name',
      description: 'description',
      price: 1.0,
      imageUrl: 'imageUrl',
);

  test('should add product to the repository', () async {
    // Arrange
    when(mockProductRepository.insertProduct(any))
        .thenAnswer((_) async => Right(tProduct));

    // Act
    final result = await usecase.execute(tProduct);

    // Assert
    expect(result, Right(tProduct));
    verify(mockProductRepository.insertProduct(tProduct));
    verifyNoMoreInteractions(mockProductRepository);
  });
}
