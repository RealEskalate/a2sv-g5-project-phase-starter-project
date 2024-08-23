



import 'package:dartz/dartz.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/Domain/usecase/ecommerce_usecase.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_hlper.mocks.dart';



void main() {
  late EcommerceUsecase usecase;
  late MockEcommerceRepositories repositories;

  setUp((){
    repositories = MockEcommerceRepositories();
    usecase = EcommerceUsecase(repositories: repositories);
  });
  


  const List<EcommerceEntity> productData = [
    EcommerceEntity(id: '1', name: 'pc', description: 'hp', imageUrl: 'http', price: 2222.2),
    EcommerceEntity(id: '1', name: 'pc', description: 'hp', imageUrl: 'http', price: 2222.2),
  ];
  

  Map<String,dynamic> datasJson = {
    'name' : 'name',
    'description' : 'description',
    'imageUrl' : 'imageUrl',
    'price' : 1
  };

  
  test(
    'should return detail repository of current data',
    () async {
      when(repositories.getAllProduct()).thenAnswer((_) async => const Right(productData));
      
      final result = await usecase.dataForAll();

      expect(result, const Right(productData));
    },
  );
  Map<String,dynamic>  datas = 
 {
      'id': '6672776eb905525c145fe0bb',
      'name': 'Anime website',
      'description': 'Explore anime characters.',
      'price': 123,
     
  };
  const EcommerceEntity singleProduct = EcommerceEntity(id: '1', name: 'pc', description: 'hp', imageUrl: 'http', price: 2222.2);
  String id = '1';
  test(
    'should return detail or product by id',
    () async {
      when(repositories.getProductById(id)).thenAnswer((_) async => const Right(singleProduct));
      
      final result = await usecase.dataById(id);

      expect(result, const Right(singleProduct));
    },
  );

  test(
    'should return edited product',
    () async {
      when(repositories.editeProduct(any,any)).thenAnswer((_) async => const Right(true));
      
      final result = await usecase.editProduct(id,datas);

      expect(result, const Right(true));
    },
  );

  test(
    'should return deleted by id',
    () async {
      when(repositories.deleteProduct(id)).thenAnswer((_) async => const Right(true));
      
      final result = await usecase.deleteProduct(id);

      expect(result, const Right(true));
    },
  );

  test(
    'should add new product to the data',
    () async {
      
      when(repositories.addProduct(any)).thenAnswer((_) async => const Right(true));
      
      final result = await usecase.addProducts(datasJson);

      expect(result, const Right(true));
    },
  );

}