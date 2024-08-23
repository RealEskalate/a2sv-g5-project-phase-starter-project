import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/failure/failure.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/domain/use_case/get_product_by_id_usecase.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late GetProductByIdUsecase getProductByIdUsecase;
  late MockProductRepositories mockProductRepositories;
  setUp(() {
    mockProductRepositories = MockProductRepositories();
    getProductByIdUsecase =
        GetProductByIdUsecase(productRepository: mockProductRepositories);
  });

   const testProductDetails = Product(
        id: '1',
        name: 'Nike',
        description: 'Nike is the Best',
        price: 344,
        imageUrl: 'imageUrl');
    const testProductId = '1';
    test('should get product by id from the repository', () async {
      //arrange
      when(mockProductRepositories.getProduct(testProductId))
          .thenAnswer((_) async => const Right(testProductDetails));
      //act
      final result = await getProductByIdUsecase.call(const GetParams(id: testProductId));
      //assert
      expect(result, const Right(testProductDetails));
      
    });
    test('should return a failure when the product fetching fails', () async {
      //arrange
      when(mockProductRepositories.getProduct(testProductId)).thenAnswer(
          (_) async => const Left(ServerFailure(message: 'server failure')));
      //act
      final result = await getProductByIdUsecase.call(const GetParams(id: testProductId));
      //assert
      expect(result, const Left(ServerFailure(message: 'server failure')));
    });
}
