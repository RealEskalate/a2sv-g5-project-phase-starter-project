import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/domain/usecases/get_current_product.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late GetCurrentProductUsecase getCurrentProductUsecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getCurrentProductUsecase = GetCurrentProductUsecase(mockProductRepository);
  });

  const testProductDetail = Product(
      id: '1',
      name: 'Apple',
      description: 'A testing apple.',
      imageUrl: 'imageUrl',
      price: 10.99);
  const testProductId = '1';

  group('getCurrentProduct Usecase', () {
    test('should get current product from the repository', () async {
      //arrange
      when(mockProductRepository.getCurrentProduct(testProductId))
          .thenAnswer((_) async => const Right(testProductDetail));

      //act
      final result =
          await getCurrentProductUsecase(const GetParams(id: testProductId));

      //assert
      expect(result, const Right(testProductDetail));
    });

    test('should return a failure when failure occurs', () async {
      //arrange
      when(mockProductRepository.getCurrentProduct(testProductId)).thenAnswer(
          (_) async => const Left(ServerFailure('test error message')));

      //act
      final result =
          await getCurrentProductUsecase(const GetParams(id: testProductId));

      //expect
      expect(result, const Left(ServerFailure('test error message')));
    });
  });
}
