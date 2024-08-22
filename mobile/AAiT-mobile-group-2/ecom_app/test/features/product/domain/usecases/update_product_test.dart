import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/domain/usecases/update_product.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late UpdateProductUsecase updateProductUsecase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    updateProductUsecase = UpdateProductUsecase(mockProductRepository);
  });

  const testProductDetail = Product(
      id: '1',
      name: 'Apple',
      description: 'A testing apple.',
      imageUrl: 'imageUrl',
      price: 10.99);

  const updatedTestProductDetail = Product(
      id: '1',
      name: 'Banana',
      description: 'A testing banana.',
      imageUrl: 'imageUrl',
      price: 23.99);

  group('updateProduct Usecase', (){
      test('should update a product given the updated product', () async {
    //arrange
    when(mockProductRepository.updateProduct(testProductDetail))
        .thenAnswer((_) async => const Right(updatedTestProductDetail));

    //act
    final result = await updateProductUsecase(UpdateParams(product: testProductDetail));

    //expect
    expect(result, const Right(updatedTestProductDetail));
  });

  test('should return a failure when failure occurs', () async {
    //arrange
    when(mockProductRepository.updateProduct(testProductDetail))
        .thenAnswer((_) async => const Left(ServerFailure('test error message')));

    //act
    final result = await updateProductUsecase(UpdateParams(product: testProductDetail));

    //expect
    expect(result, const Left(ServerFailure('test error message')));
  });
  });


}
