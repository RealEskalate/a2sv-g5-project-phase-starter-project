import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
import 'package:ecommerce/features/product/domain/usecases/addproduct.dart';
import 'package:ecommerce/features/product/domain/usecases/getallproduct.dart';
import 'package:ecommerce/features/product/domain/usecases/updateproduct.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late UpdateProductUsecase updateProductUsecase;
  late MockProductRepository mockProductRepository;
  

  setUp((){
    mockProductRepository = MockProductRepository();
    updateProductUsecase = UpdateProductUsecase(mockProductRepository);  
  });

    const id = '3';
    const testproductdetail =  Productentity(id: '3', image: 'http/img', name: 'adidas',  description: 'description', price: 300);
    const updatedproduct = Productentity(id: '3', image: 'http/ig', name: 'addas', description: 'desciption', price: 300);

   
  test(
    'should update the product to the repository',

  
   () async {
      when(
        mockProductRepository.updateproduct(updatedproduct)

      ).thenAnswer((_) async=>  Right(testproductdetail));

      final result = await updateProductUsecase.update(updatedproduct);

    expect(result, Right(testproductdetail));
  });

  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message',
  () async {
    //arrange
    when(
      mockProductRepository.updateproduct(updatedproduct)
    ).thenAnswer((_) async =>  Left(failure));

    //act
    final result = await updateProductUsecase.update(updatedproduct);

    //assert
    expect(result, Left(failure));

  }
  );

}

