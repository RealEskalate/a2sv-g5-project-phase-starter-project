
import 'package:dartz/dartz.dart';
import 'package:ecommerce_app/features/product/domain/usecases/insert_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../test_helper/test_helper_generation.mocks.dart';
import '../../../../test_helper/testing_datas/product_testing_data.dart';

void main(){
  late MockProductRepository mockProductRepository;
  late InsertProductUseCase insertProductUseCase;

  setUp((){
    mockProductRepository = MockProductRepository();
    insertProductUseCase = InsertProductUseCase(mockProductRepository);
  });


  test('Testing the data flow inside the Repositrory of inserting product', () async {
    /// Rearranging the functionality
    when(
        mockProductRepository.insertProduct(TestingDatas.testDataEntity)
    ).thenAnswer((_) async =>  const Right(1));

    /// action

    final result = await insertProductUseCase.execute(TestingDatas.testDataEntity);

    /// assertion
    ///
    expect(result, const Right(1));


  });
}