


import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecommers/core/Error/failure.dart';
import 'package:ecommers/features/ecommerce/Domain/entity/ecommerce_entity.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/bottum_state.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/button_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/input_button_activation/button_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_bloc.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_event.dart';
import 'package:ecommers/features/ecommerce/presentation/state/product_bloc/product_state.dart';
import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helper/test_hlper.mocks.dart';

void main() {
  late MockEcommerceUsecase mockEcommerceUsecase;
  late ProductBloc productBloc;
  late ButtonBloc buttonBloc;

  setUp(() {
    mockEcommerceUsecase = MockEcommerceUsecase();
    productBloc = ProductBloc(ecommerceUsecase: mockEcommerceUsecase);
    buttonBloc = ButtonBloc(ecommerceUsecase: mockEcommerceUsecase);
  });


  test(
    'test the inisial state of the app',
    () {
      expect(productBloc.state, ProductIntialState());
    }
    );
 

  const EcommerceEntity ecommerceEntity = EcommerceEntity(
    id: '1', 
    name: 'pc', 
    description: 'description', 
    imageUrl: 'imageUrl', 
    price: 123.2);
  const  List<EcommerceEntity> listProduct = [
    EcommerceEntity(
    id: '1', 
    name: 'pc', 
    description: 'description', 
    imageUrl: 'imageUrl', 
    price: 123.2),
    EcommerceEntity(
    id: '1', 
    name: 'pc', 
    description: 'description', 
    imageUrl: 'imageUrl', 
    price: 123.2)
  ];
  group(
    'test the state of the app',
    () {
      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when get product by id success fully.',
        build: (){
          when(
            mockEcommerceUsecase.dataById(any)
          ).thenAnswer((_) async=> const Right(ecommerceEntity));
          return productBloc;
        },
        act: (bloc) => bloc.add(const GetSingleProductEvent(id: '1')),
          expect: () => [
            LoadingState(),
            const LoadedSingleProductState(product: ecommerceEntity)
          ],
      );

      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when get product by id return failure.',
        build: (){
          when(
            mockEcommerceUsecase.dataById(any)
          ).thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const GetSingleProductEvent(id: '1')),
          expect: () => [
            LoadingState(),
            const ProductErrorState(messages: 'try again'),
            
          ],
      );


      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when get all product return  success fully.',
        build: (){
          when(
            mockEcommerceUsecase.dataForAll()
          ).thenAnswer((_) async => const Right(listProduct));
          return productBloc;
        },
        act: (bloc) => bloc.add(const LoadAllProductEvent()),
          expect: () => [
            LoadingState(),
            const LoadedAllProductState(products: listProduct),
            
          ],
      );

      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when get  all product must return failure.',
        build: (){
          when(
            mockEcommerceUsecase.dataForAll()
          ).thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const LoadAllProductEvent()),
          expect: () => [
            LoadingState(),
            const ProductErrorState(messages: 'try again'),
            
          ],
      );

    

    


      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when delete  a product must return  success.',
        build: (){
          when(
            mockEcommerceUsecase.deleteProduct(any)
          ).thenAnswer((_) async => const Right(true));
          return productBloc;
        },
        act: (bloc) => bloc.add(const DeleteProductEvent(id: '1')),
          expect: () => [
            LoadingState(),
            const SuccessDelete(deleted:true),
            
          ],
      );

      blocTest<ProductBloc, ProductState>(
        'emits [MyState] when delete  a product must return failure.',
        build: (){
          when(
            mockEcommerceUsecase.deleteProduct(any)
          ).thenAnswer((_) async => const Left(ConnectionFailur(message: 'try again')));
          return productBloc;
        },
        act: (bloc) => bloc.add(const DeleteProductEvent(id: '1')),
          expect: () => [
            LoadingState(),
            const ProductErrorState(messages: 'try again'),
            
          ],
      );


      blocTest<ButtonBloc, BottumState>(
        'emits [MyState] when edit  a product must return  success.',
        build: (){
          when(
            mockEcommerceUsecase.editProduct(any,any),
          ).thenAnswer((_) async => const Right(true));
       
          return buttonBloc;
        },
        act: (bloc) => bloc.add( UpdateProductEvent()),
        
          expect: () => [
            AddLoadingState(),
            SuccessAddProduct(add:true),
            
          ],
      );

      blocTest<ButtonBloc, BottumState>(
        'emits [MyState] when edit  a product must return  failure.',
        build: (){
          when(
            mockEcommerceUsecase.editProduct(any,any)
          ).thenAnswer((_) async => const Left(ServerFailure(message: 'try again')));
          return buttonBloc;
        },
        act: (bloc) => bloc.add( UpdateProductEvent()),
          expect: () => [
            AddLoadingState(),
            AddErrorState(messages: 'try again'),
            
          ],
      );

      blocTest<ButtonBloc, BottumState>(
        'emits [MyState] when add  a product must return  success.',
        build: (){
          when(
            mockEcommerceUsecase.addProducts(any),
          ).thenAnswer((_) async => const Right(true));
       
          return buttonBloc;
        },
        act: (bloc) => bloc.add( AddProductEvent()),
        
          expect: () => [
            AddLoadingState(),
            SuccessAddProduct(add:true),
            
          ],
      );

      blocTest<ButtonBloc, BottumState>(
        'emits [MyState] when edit  a product must return  failure.',
        build: (){
          when(
            mockEcommerceUsecase.addProducts(any)
          ).thenAnswer((_) async => const Left(ServerFailure(message: 'try again')));
          return buttonBloc;
        },
        act: (bloc) => bloc.add( AddProductEvent()),
          expect: () => [
            AddLoadingState(),
            
            AddErrorState(messages: 'try again'),
            
          ],
      );
    }
    ); 
}