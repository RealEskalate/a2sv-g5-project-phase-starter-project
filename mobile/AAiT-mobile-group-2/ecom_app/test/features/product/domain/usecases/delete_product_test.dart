import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/product/domain/usecases/delete_product.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late DeleteProductUsecase deleteProductUsecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    deleteProductUsecase = DeleteProductUsecase(mockProductRepository);
  });

  const testProductId = '1';

  group('deleteProduct Usecase', () {
    test('should delete a product given the product ID', () async {
      //arrange
      when(mockProductRepository.deleteProduct(testProductId))
          .thenAnswer((_) async => const Right(null));

      //act
      final result = await deleteProductUsecase(DeleteParams(id: testProductId));

      //expect
      expect(result, const Right(null));
    });

    test('should return a failure when failure occurs', () async {
      //arrange
      when(mockProductRepository.deleteProduct(testProductId)).thenAnswer(
          (_) async => const Left(ServerFailure('test error message')));

      //act
      final result = await deleteProductUsecase(DeleteParams(id: testProductId));

      //expect
      expect(result, const Left(ServerFailure('test error message')));
    });
  });
}
