import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
import 'package:ecommerce/features/product/domain/usecases/addproduct.dart';
import 'package:ecommerce/features/product/domain/usecases/getallproduct.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late AddProductUsecase addProductUsecase;
  late MockProductRepository mockProductRepository;
  

  setUp((){
    mockProductRepository = MockProductRepository();
    addProductUsecase = AddProductUsecase(mockProductRepository);  
  });

    const newproductdetail = Productentity(id:'12' ,name: 'Nike', description: 'A great Shoe', image: 'The Nike', price: 99);


   
  test(
    'should add the product to the repository',

  
   () async {
      when(
        mockProductRepository.addproduct(newproductdetail)

      ).thenAnswer((_) async=>  Right(null));

      final result = await addProductUsecase.add(newproductdetail);

    expect(result, Right(null));
  });

  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message',
  () async {
    //arrange
    when(
      mockProductRepository.addproduct(newproductdetail)
    ).thenAnswer((_) async =>  Left(failure));

    //act
    final result = await addProductUsecase.add(newproductdetail);

    //assert
    expect(result, Left(failure));

  }
  );

}

