import 'package:dartz/dartz.dart';
import 'package:ecommerse2/core/error/failure.dart';
import 'package:ecommerse2/features/product/domain/entity/product.dart';
import 'package:ecommerse2/features/product/domain/usecase/get_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

void main(){

  late GetProductUseCase getProductUseCase;
  late MockProductRepository mockProductRepository;

  setUp((){

    mockProductRepository = MockProductRepository();
    getProductUseCase = GetProductUseCase(mockProductRepository);

  });

  String id = '1';
  Product product = const Product(id: '1', name: 'Nike', category: 'Shoe', description: 'A great Shoe', image: 'The Nike', price: 99);

  test('Product Found', () async{

    //arrange
    when(mockProductRepository.getProduct(id)).thenAnswer((_) async => Right(product));   

    //act
    final result = await getProductUseCase.call(id);

    //assert
    expect(result, Right(product));
  });

  //testing failure
  Failure failure = const ConnectionFailure('Connection Error');

  test('Failure Get Product', () async {
    //arrange
    when(mockProductRepository.getProduct(id)).thenAnswer((_) async => Left(failure));

    //act
    final result = await getProductUseCase.call(id);

    //assert
    expect(result, Left(failure));

  });





}