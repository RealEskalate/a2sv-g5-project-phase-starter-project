import 'package:dartz/dartz.dart';
import 'package:ecommerce/core/error/failure.dart';
import 'package:ecommerce/features/product/domain/entities/product.dart';
import 'package:ecommerce/features/product/domain/usecases/getallproduct.dart';
import 'package:mockito/mockito.dart';
import 'package:flutter_test/flutter_test.dart';
import '../../../../helpers/test_helper.mocks.dart';

void main(){
  late GetAllProductUsecase getAllProductUsecase;
  late MockProductRepository mockProductRepository;
  

  setUp((){
    mockProductRepository = MockProductRepository();
    getAllProductUsecase = GetAllProductUsecase(mockProductRepository);


  });

    List<Productentity> products = [
   const Productentity(id: '12' ,name: 'Nike',description: 'A great Shoe', image: 'The Nike', price: 99),
   const Productentity(id : '23' ,name: 'Puma',  description: 'The best Shoe', image: 'The Puma', price: 999),
  ];

   
  test(
    'should get all the products from the repository',

  
   () async {
      when(
        mockProductRepository.getallproduct()

      ).thenAnswer((_) async=>  Right(products));

      final result = await getAllProductUsecase.getall();

    expect(result, Right(products));
  });

  Failure failure = const ServerFailure('Failure Server');
  test('should print failure message',
  () async {
    //arrange
    when(
      mockProductRepository.getallproduct()
    ).thenAnswer((_) async =>  Left(failure));

    //act
    final result = await getAllProductUsecase.getall();

    //assert
    expect(result, Left(failure));

  }
  );

}

