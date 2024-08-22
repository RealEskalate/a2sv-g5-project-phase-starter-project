import 'package:flutter_application_5/features/product/domain/entities/product_entity.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_application_5/features/product/domain/usecases/get_detail_usecase.dart';

import '../../../helpers/test_helper.mocks.dart';
import 'package:mockito/mockito.dart';
import 'package:dartz/dartz.dart';

void main(){
  GetDetailUseCase getDetailUseCase = GetDetailUseCase(MockProductRepository());
  MockProductRepository mockProductRepository = MockProductRepository();

  setUp(() {
    mockProductRepository = MockProductRepository();
    getDetailUseCase = GetDetailUseCase(mockProductRepository);
  });

  const tProductEntity = ProductEntity(
    id: '1',
    name: 'Test Product',
    description: 'Test Description',
    price: 10.0,
    imagePath: null,
  );
  const tId = '1';

  test(
    'should get product from the repository',
    () async {
      // Arrange
      when(mockProductRepository.getProduct(tId))
          .thenAnswer((_) async => const Right(tProductEntity));

      // Act
      final result = await getDetailUseCase.call(tId);

      // Assert
      expect(result, const Right(tProductEntity));
      verify(mockProductRepository.getProduct(tId));
      verifyNoMoreInteractions(mockProductRepository);
    },
  );

}