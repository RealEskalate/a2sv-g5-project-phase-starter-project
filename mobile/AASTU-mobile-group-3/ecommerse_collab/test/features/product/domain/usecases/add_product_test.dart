

import 'package:dartz/dartz.dart';
import 'package:ecommerse2/core/error/failure.dart';
import 'package:ecommerse2/features/product/domain/entity/product.dart';
import 'package:ecommerse2/features/product/domain/usecase/add_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

void main(){

  late AddProductUseCase addProductUseCase;
  late MockProductRepository mockProductRepository;

  setUp((){

    mockProductRepository = MockProductRepository();
    addProductUseCase = AddProductUseCase(mockProductRepository);

  });

  Product product = const Product(id: '1', name: 'Nike', category: 'Shoe', description: 'A great Shoe', image: 'The Nike', price: 99);

  test('Added', () async {
      //arrange
      when(mockProductRepository.addProduct(product)).thenAnswer((_) async => const Right(null));

      //act
      final result = await addProductUseCase.call(product);

      //assert
      expect(result, const Right(null));

  });

  Failure failure = const ServerFailure('Failure');
  test('Failure Add Product', () async {
    //arrange
    when(mockProductRepository.addProduct(product)).thenAnswer((_) async => Left(failure));

    //act
    final result = await addProductUseCase.call(product);

    //assert
    expect(result, Left(failure));
    

  });

}