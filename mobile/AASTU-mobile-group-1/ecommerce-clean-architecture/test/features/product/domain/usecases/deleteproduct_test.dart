import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
// import 'package:ecommerce/features/product/domain/usecases/addproduct.dart';
import 'package:ecommerce/features/product/domain/usecases/deleteproduct.dart';
// import 'package:ecommerce/features/product/domain/usecases/getallproduct.dart';
// import 'package:ecommerce/features/product/domain/usecases/updateproduct.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late DeleteProductUsecase deleteProductUsecase;
  late MockProductRepository mockProductRepository;
  

  setUp((){
    mockProductRepository = MockProductRepository();
    deleteProductUsecase = DeleteProductUsecase(mockProductRepository);  
  });

    const id = '3';
    const deletedproductdetail =  Productentity(
      id: '3',
       image: 'http/img',
      name: 'adidas', 
      description: 'description', 
      price: 300);

   
  test(
    'should delete the product from the repository',

  
   () async {
      when(
        mockProductRepository.deleteproduct(id)

      ).thenAnswer((_) async=> Right(deletedproductdetail));

      final result = await deleteProductUsecase.delete(id);

    expect(result, Right(deletedproductdetail));
  });

  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message',
  () async {
    //arrange
    when(
      mockProductRepository.deleteproduct(id)
    ).thenAnswer((_) async =>  Left(failure));

    //act
    final result = await deleteProductUsecase.delete(id);

    //assert
    expect(result, Left(failure));

  }
  );

}

