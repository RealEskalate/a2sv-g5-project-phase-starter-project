
import 'package:dartz/dartz.dart';
import 'package:ecommerse2/core/error/failure.dart';
import 'package:ecommerse2/features/product/domain/entity/product.dart';
import 'package:ecommerse2/features/product/domain/usecase/delete_product.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../helpers/test_helper.mocks.dart';

void main(){

  late DeleteProductUseCase deleteProductUseCase;
  late MockProductRepository mockProductRepository;

  setUp((){
    mockProductRepository = MockProductRepository();
    deleteProductUseCase = DeleteProductUseCase(mockProductRepository);
  });

  String id = '1';
  Product product = const Product(id: '1', name: 'Nike', category: 'Shoe', description: 'A great Shoe', image: 'The Nike', price: 99);

  test('Deleted Successfully', () async{

    //arrange
    when(mockProductRepository.deleteProduct(id)).thenAnswer((_) async => Right(product));


    //act
    final result = await deleteProductUseCase.call(id);

    //assert
    expect(result, Right(product));


  });

  //testing failure
  Failure failure = const ServerFailure('Failure Delete Product');

  test('Failure Delete Product', () async {

      //arrange
      when(mockProductRepository.deleteProduct(id)).thenAnswer((_) async => Left(failure));

      //act
      final result = await deleteProductUseCase.call(id);

      //assert

      expect(result, Left(failure));

  });

}