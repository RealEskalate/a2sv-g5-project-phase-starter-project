// ignore_for_file: void_checks
import 'package:flutter_application_5/core/usecases/usecases.dart';
import 'package:flutter_application_5/features/product/domain/entities/product_entity.dart';
import 'package:flutter_application_5/features/product/domain/usecases/delete_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:flutter_application_5/features/product/domain/usecases/get_detail_usecase.dart';

import '../../../helpers/test_helper.mocks.dart';
import 'package:mockito/mockito.dart';
import 'package:dartz/dartz.dart';

void main(){
  DeleteUsecase deleteUsecase = DeleteUsecase(MockProductRepository());
  MockProductRepository mockProductRepository = MockProductRepository();

  setUp(() {
    mockProductRepository = MockProductRepository();
    deleteUsecase = DeleteUsecase(mockProductRepository);
  });

  const tProduct = '1';
  test(
    'should Delete product from the repository',
    () async {
      // Arrange
      when(mockProductRepository.deleteProduct(tProduct))
          .thenAnswer((_) async => const Right('Success'));

      // Act
      final result = await deleteUsecase.call(tProduct);

      // Assert
      expect(result, const Right('Success'));
      verify(mockProductRepository.deleteProduct(tProduct));
      verifyNoMoreInteractions(mockProductRepository);
    },
    );
}