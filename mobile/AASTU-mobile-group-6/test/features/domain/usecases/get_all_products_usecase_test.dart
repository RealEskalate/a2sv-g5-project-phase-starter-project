import 'package:flutter_application_5/core/usecases/usecases.dart';
import 'package:flutter_application_5/features/product/domain/entities/product_entity.dart';
import 'package:flutter_application_5/features/product/domain/usecases/get_all_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_application_5/features/product/domain/usecases/get_detail_usecase.dart';

import '../../../helpers/test_helper.mocks.dart';
import 'package:mockito/mockito.dart';
import 'package:dartz/dartz.dart';

void main(){
  GetAllUsecase getAllUsecase = GetAllUsecase(MockProductRepository());
  MockProductRepository mockProductRepository = MockProductRepository();

  setUp(() {
    mockProductRepository = MockProductRepository();
    getAllUsecase = GetAllUsecase(mockProductRepository);
  });

  const tProduct = <ProductEntity>[
    ProductEntity(
      id: '1',
      name: 'Test Product',
      description: 'Test Description',
      price: 10,
      imagePath: null
    )
  ];

  test(
    'should get all products from the repository',
    () async {
      // Arrange
      when(mockProductRepository.getProducts())
          .thenAnswer((_) async => const Right(tProduct));

      // Act
      final result = await getAllUsecase.call(NoParams());

      // Assert
      expect(result, const Right(tProduct));
      verify(mockProductRepository.getProducts());
      verifyNoMoreInteractions(mockProductRepository);
    },
    );
}