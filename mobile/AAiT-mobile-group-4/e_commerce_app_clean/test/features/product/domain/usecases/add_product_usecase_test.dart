import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/domain/usecases/add_product_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';
void main() {
  late MockProductRepository mockProductRepository;
  late AddProductUsecase addProductUsecase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    addProductUsecase = AddProductUsecase(mockProductRepository);
  });

  const tProductEntity = ProductEntity(
    id: '3',
    name: 'Product 1',
    price: 100,
    description: 'Description 1',
    imageUrl: 'image1.jpg',
  );
  test(
    'should add Product to ProductRepository and return ProductEntity',
    () async {
      // Arrange
      when(mockProductRepository.addProduct(tProductEntity))
          .thenAnswer((_) async => const Right(tProductEntity));
      
      // Act
      final result = await addProductUsecase(const CreateParams(product: tProductEntity));
      
      // Assert
      expect(result, const Right(tProductEntity));
    },
  );
}