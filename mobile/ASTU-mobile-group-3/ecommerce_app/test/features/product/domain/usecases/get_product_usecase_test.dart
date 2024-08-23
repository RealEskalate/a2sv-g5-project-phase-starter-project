import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/features/product/domain/usecases/get_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late GetProductUseCase getProductUseCase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getProductUseCase = GetProductUseCase(mockProductRepository);
  });


  test('Testing the data flow inside the Repositrory of product list return', () async {
    /// Rearranging the functionality
    when(
        mockProductRepository.getProduct(TestingDatas.id)
    ).thenAnswer((_) async => const Right(TestingDatas.testDataEntity));

    /// action

    final result = await getProductUseCase.execute(TestingDatas.id);

    /// assertion
    ///
    expect(result, const Right(TestingDatas.testDataEntity));
  });
}