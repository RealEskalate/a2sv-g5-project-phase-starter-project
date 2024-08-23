import 'package:application1/features/product/domain/entities/product_entity.dart'; // Import the ProductEntity class
// Import the ProductRepository class
import 'package:application1/features/product/domain/usecases/get_product_usecase.dart'; // Import the GetProductUsecase class
import 'package:dartz/dartz.dart'; // Import the dartz package
import 'package:mockito/mockito.dart'; // Import the mockito package
import 'package:test/test.dart'; // Import the test package

import '../../../../helper/test_helper.mocks.dart'; // Import the Failure class


void main() {
  late MockProductRepository mockProductRepository;
  late GetProductUsecase getProductUsecase;

  setUp(() {
    mockProductRepository = MockProductRepository();
    getProductUsecase = GetProductUsecase(mockProductRepository);
  });

  const tProductEntity = ProductEntity(
    id: '3',
    name: 'Product 1',
    price: 100,
    description: 'Description 1',
    imageUrl: 'image1.jpg',
   
  );

  test(
    'should get Product from ProductRepository',
    () async {
      // Arrange
      when(mockProductRepository.getProduct(tProductEntity.id))
          .thenAnswer((_) async => const Right(tProductEntity));
      
      // Act
      final result = await getProductUsecase(GetParams(id: tProductEntity.id));
      
      // Assert
      expect(result, const Right(tProductEntity));
    },
  );
}