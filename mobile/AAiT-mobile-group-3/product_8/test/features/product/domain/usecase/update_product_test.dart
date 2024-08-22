import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/failure/failure.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/domain/use_case/update_product_usecase.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late UpdateProductUsecase updateProductUsecase;
  late MockProductRepositories mockProductRepositories;
  setUp((){
    mockProductRepositories = MockProductRepositories();
    updateProductUsecase = UpdateProductUsecase(productRepository: mockProductRepositories);
  });
  const testProductDetails = Product(
      id: '1',
      name: 'Nike',
      description: 'Nike is the Best',
      price: 344,
      imageUrl: 'imageUrl');
  test('should update product to the repository', () async {
    //arrange
    when(mockProductRepositories.updateProduct(testProductDetails))
        .thenAnswer((_) async => const Right(testProductDetails));
    //act
    final result = await updateProductUsecase.call(const UpdateParams(product: testProductDetails));
    //assert
    expect(result, const Right(testProductDetails));
  });

  test('should return a failure when the product update fails', () async {
    //arrange
    when(mockProductRepositories.updateProduct(testProductDetails))
        .thenAnswer((_) async =>  const Left(ServerFailure(message: 'server failure')));
    //act
    final result = await updateProductUsecase.call(const UpdateParams(product: testProductDetails));
    //assert
    expect(result,  const Left(ServerFailure(message: 'server failure')));
  });
}