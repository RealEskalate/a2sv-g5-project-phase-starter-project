import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';
import 'package:product_8/core/failure/failure.dart';
import 'package:product_8/core/usecase/usecase.dart';
import 'package:product_8/features/product/domain/entities/product_entity.dart';
import 'package:product_8/features/product/domain/use_case/delete_product_usecase.dart';
import 'package:product_8/features/product/domain/use_case/get_product_by_id_usecase.dart';
import 'package:product_8/features/product/domain/use_case/insert_product_usecase.dart';
import 'package:product_8/features/product/domain/use_case/update_product_usecase.dart';
import 'package:product_8/features/product/presentation/bloc/product_bloc.dart';
import 'package:product_8/features/product/presentation/bloc/product_event.dart';
import 'package:product_8/features/product/presentation/bloc/product_state.dart';

import '../../../../helpers/test_helper.mocks.dart';
void main() {
  late MockGetProductByIdUsecase mockGetProductByIdUsecase;
  late MockGetProductsUsecase mockGetProductsUsecase;
  late MockInsertProductUsecase mockInsertProductUsecase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late MockDeleteProductUsecase mockDeleteProductUsecase;
  late ProductBloc productBloc;

  setUp(() {
    mockGetProductByIdUsecase = MockGetProductByIdUsecase();
    mockGetProductsUsecase = MockGetProductsUsecase();
    mockInsertProductUsecase = MockInsertProductUsecase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockDeleteProductUsecase = MockDeleteProductUsecase();
    productBloc = ProductBloc(
        getProductByIdUsecase: mockGetProductByIdUsecase,
        getProductsUsecase: mockGetProductsUsecase,
        insertProductUsecase: mockInsertProductUsecase,
        updateProductUsecase: mockUpdateProductUsecase,
        deleteProductUsecase: mockDeleteProductUsecase);
  });

  const testProductDetails = Product(
      id: '1',
      name: 'Nike',
      description: 'Nike is the Best',
      price: 344,
      imageUrl: 'imageUrl');

  const testProductId = '1';
  test('initial state should be InitialState', () {
    //assert
    expect(productBloc.state, ProductInitial());
  });



  group('GetProductsbyId', () {
  
  blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductLoadedSingle] when GetProductbyIdEvent is added.',
        build: () {
         

          return productBloc;}
         
        ,
        setUp: () {
           when(mockGetProductByIdUsecase( const GetParams(id: testProductId)))
              .thenAnswer((_) async =>  const Right(testProductDetails));
        },
       
        
        act: (bloc) => bloc.add(const GetProductByIdEvent(id: testProductId)),
        wait:const Duration(milliseconds: 500),
        expect: () => [
              ProductLoading(),
              const ProductLoadedSingle(product: testProductDetails)
            ]);

              blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductErrorState] when GetProductbyIdEvent is added.',
        build: () {
           when(mockGetProductByIdUsecase.call( const GetParams(id: testProductId)))
              .thenAnswer((_) async =>  const Left(ServerFailure(message: 'server failure')));
          return productBloc;
         
        },
       
        
        act: (bloc) => bloc.add(const GetProductByIdEvent(id: testProductId)),
        wait:const Duration(milliseconds: 500),
        expect: () => [
              ProductLoading(),
               ProductError()
            ]);



  });


  group('CreateProductEvent', () {
    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductLoadedSingle] when CreateProductEvent is added.',
        build: () {
          when(mockInsertProductUsecase.call(const CreateParams(product: testProductDetails)))
              .thenAnswer((_) async => const Right(testProductDetails));
          return productBloc;
        },
        act: (bloc) => bloc.add(const CreateProductEvent(product: testProductDetails)),
       
        expect: () => [
              ProductLoading(),
              const ProductCreatedState(product: testProductDetails)
            ]);

    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductErrorState] when CreateProductEvent is added.',
        build: () {
          when(mockInsertProductUsecase.call(const CreateParams(product: testProductDetails)))
              .thenAnswer((_) async => const Left(ServerFailure(message: 'server failure')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const CreateProductEvent(product: testProductDetails)),
       
        expect: () => [
              ProductLoading(),
              const ProductCreatedErrorState(message: 'Error creating product')
            ]);
  });


  group('UpdateProductEvent', () {
    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductLoadedSingle] when UpdateProductEvent is added.',
        build: () {
          when(mockUpdateProductUsecase.call(const UpdateParams(product: testProductDetails)))
              .thenAnswer((_) async => const Right(testProductDetails));
          return productBloc;
        },
        act: (bloc) => bloc.add(const UpdateProductEvent(product: testProductDetails)),
        wait: const Duration(milliseconds: 500),
        expect: () => [
              ProductLoading(),
              const ProductUpdatedState(product: testProductDetails)
            ]);

    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductErrorState] when UpdateProductEvent is added.',
        build: () {
          when(mockUpdateProductUsecase.call(const UpdateParams(product: testProductDetails)))
              .thenAnswer((_) async => const Left(ServerFailure(message: 'server failure')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const UpdateProductEvent(product: testProductDetails)),
      
        expect: () => [
              ProductLoading(),
              const ProductUpdatedErrorState(message: 'Error updating product')
            ]);
  });


  group('DeleteProductEvent', () {
    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductDeleteState] when DeleteProductEvent is added.',
        build: () {
          when(mockDeleteProductUsecase.call(const DeleteParams(id: testProductId)))
              .thenAnswer((_) async => const Right(null));
          return productBloc;
        },
        act: (bloc) => bloc.add(const DeleteProductEvent(id: testProductId)),
       
        expect: () => [
              ProductLoading(),
              ProductDeleteState()
            ]);

    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductErrorState] when DeleteProductEvent is added.',
        build: () {
          when(mockDeleteProductUsecase.call(const DeleteParams(id: testProductId)))
              .thenAnswer((_) async => const Left(ServerFailure(message: 'server failure')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const DeleteProductEvent(id: testProductId)),
       
        expect: () => [
              ProductLoading(),
              ProductError()
            ]);
  });

  group('LoadProduct', () {
    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductLoaded] when LoadProduct is added.',
        build: () {
          when(mockGetProductsUsecase.call(NoParams()))
              .thenAnswer((_) async => const Right([testProductDetails]));
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadProduct()),
       
        expect: () => [
              ProductLoading(),
              const ProductLoaded(products: [testProductDetails])
            ]);

    blocTest<ProductBloc, ProductState>(
        'emits [ProductLoading, ProductErrorState] when LoadProduct is added.',
        build: () {
          when(mockGetProductsUsecase.call(NoParams()))
              .thenAnswer((_) async => const Left(ServerFailure(message: 'server failure')));
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadProduct()),
       
        expect: () => [
              ProductLoading(),
              ProductError()
            ]);
  });
}
