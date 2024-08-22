import 'package:dartz/dartz.dart';
import 'package:e_commerce_app/core/failure/failure.dart';
import 'package:e_commerce_app/features/product/domain/entities/product.dart';
import 'package:e_commerce_app/features/product/domain/usecase/get_all_product_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockProductRepository mockProductRepository;
  late GetAllProductUsecase getAllProductUsecase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getAllProductUsecase = GetAllProductUsecase(mockProductRepository);
  });
  List<ProductEntity> testAllProduct = <ProductEntity>[
    ProductEntity(
        id: 'uuu3456789dfghjk',
        name: 'derby leather',
        description: 'old shoes',
        imageUrl: '/assets/image.jpg',
        price: 15.5),
    ProductEntity(
        id: '3456789dfghjk',
        name: 'derby leather',
        description: 'new shoes',
        imageUrl: '/assets/image.jpg',
        price: 15.5),
  ];

  Failure testFail = Failure("You failed ");
  test('getting all product success', () async {
    //arrange
    when(mockProductRepository.getAllProduct())
        .thenAnswer((_) async => Right(testAllProduct));
    //act
    final result = await getAllProductUsecase.execute();

    //assert
    expect(result, Right(testAllProduct));
  });
  test('failed to get all product ', () async {
    //arrange
    when(mockProductRepository.getAllProduct())
        .thenAnswer((_) async => Left(testFail));
    //act
    final result = await getAllProductUsecase.execute();

    //assert
    expect(result, Left(testFail));
  });
}
