import 'package:dartz/dartz.dart';
import 'package:ecommerse2/core/error/failure.dart';
import 'package:ecommerse2/features/product/domain/entity/product.dart';
import 'package:ecommerse2/features/product/domain/usecase/get_all_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

void main() {
  late GetAllProductUseCase getAllProductUseCase;
  late MockProductRepository mockProductRepository;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getAllProductUseCase = GetAllProductUseCase(mockProductRepository);
  });

  List<Product> products = [
    const Product(
        id: '1',
        name: 'Nike',
        category: 'Shoe',
        description: 'A great Shoe',
        image: 'The Nike',
        price: 99),
    const Product(
        id : '2',
        name: 'Puma',
        category: 'Shoe',
        description: 'The best Shoe',
        image: 'The Puma',
        price: 999),
  ];

  test('should get all the product listed', () async {
    //arrange
    when(mockProductRepository.getAllProduct())
        .thenAnswer((_) async => Right(products));

    //act
    final result = await getAllProductUseCase.execute();

    //assert
    expect(result, Right(products));
  });
  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message', () async {
    //arrange
    when(mockProductRepository.getAllProduct())
        .thenAnswer((_) async => Left(failure));

    //act
    final result = await getAllProductUseCase.execute();

    //assert
    expect(result, Left(failure));
  });
}
