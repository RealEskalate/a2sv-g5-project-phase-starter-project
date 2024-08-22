import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
import 'package:ecommerce/features/product/domain/usecases/getproduct.dart';
// import 'package:ecommerce/features/product/domain/usecases/getallproduct.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late GetProductUsecase getProductUsecase;
  late MockProductRepository mockProductRepository;
  

  setUp((){
    mockProductRepository = MockProductRepository();
    getProductUsecase = GetProductUsecase(mockProductRepository);


  });

  const testproductdetail =  Productentity(id: '4', image: 'http/img', name: 'adidas', description: 'description', price: 300);
  const id = '4';
   
  test(
    'should get a specific product from the repository',

  
   () async {
      when(
        mockProductRepository.getproduct(id)

      ).thenAnswer((_) async=>  Right(testproductdetail));

      final result = await getProductUsecase.getprod(id);

    expect(result, Right(testproductdetail));
  });

  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message',
  () async {
    //arrange
    when(
      mockProductRepository.getproduct(id)
    ).thenAnswer((_) async =>  Left(failure));

    //act
    final result = await getProductUsecase.getprod(id);

    //assert
    expect(result, Left(failure));

  }
  );

}

