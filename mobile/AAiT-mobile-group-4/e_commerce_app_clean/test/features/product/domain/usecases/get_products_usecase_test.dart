
import 'package:application1/core/usecase/usecase.dart';
import 'package:application1/features/product/domain/entities/product_entity.dart';
import 'package:application1/features/product/domain/usecases/get_products_usecase.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late GetProductsUsecase getProductsUsecase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getProductsUsecase = GetProductsUsecase(mockProductRepository);
  });

  const tProductEntity = [
    ProductEntity(
      id: '1',
      name: 'Product 1',
      price: 100,
      description: 'Description 1',
      imageUrl: 'image1.jpg',
    ),
    ProductEntity(
      id: '2',
      name: 'Product 1',
      price: 100,
      description: 'Description 1',
      imageUrl: 'image1.jpg'
    ),
  ];
  test(
    'should call getProducts from ProductRepository',
    () async {
      //arrange
      when(mockProductRepository.getProducts())
          .thenAnswer((getProductsUsecase) async =>const Right(tProductEntity));
      //act
      final result = await getProductsUsecase(NoParams());
      //assert
      expect(result, const Right(tProductEntity));
    },
  );
}
