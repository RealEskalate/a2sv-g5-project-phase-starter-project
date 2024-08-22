import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/domain/usecases/create_product.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late CreateProductUsecase createProductUsecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    createProductUsecase = CreateProductUsecase(mockProductRepository);
  });

  const testProductDetail = Product(
      id: '1',
      name: 'Apple',
      description: 'A testing apple.',
      imageUrl: 'imageUrl',
      price: 10.99);

  group('createProduct Usecase', () {
    test('should create a product', () async {
      //arrange
      when(mockProductRepository.createProduct(testProductDetail))
          .thenAnswer((_) async => const Right(testProductDetail));

      //act
      final result =
          await createProductUsecase(CreateParams(product: testProductDetail));

      //assert
      expect(result, const Right(testProductDetail));
    });

    test('should return a failure when failure occurs', () async {
      //arrange
      when(mockProductRepository.createProduct(testProductDetail)).thenAnswer(
          (_) async => const Left(ServerFailure('test error message')));

      //act
      final result =
          await createProductUsecase(CreateParams(product: testProductDetail));

      //expect
      expect(result, const Left(ServerFailure('test error message')));
    });
  });
}
