
import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/features/product/domain/usecases/update_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main(){
  late MockProductRepository mockProductRepository;
  late UpdateProductUsecase updateProductUsecase;

  setUp((){
    mockProductRepository = MockProductRepository();
    updateProductUsecase = UpdateProductUsecase(mockProductRepository);
  });

  test('Testing the data flow inside the Repositrory of updating product', () async {
    /// Rearranging the functionality
    when(
        mockProductRepository.updateProduct(TestingDatas.testDataEntity)
    ).thenAnswer((_) async =>  const Right(1));

    /// action

    final result = await updateProductUsecase.execute(TestingDatas.testDataEntity);

    /// assertion
    ///
    expect(result, const Right(1));


  });
}