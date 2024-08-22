import 'package:bloc_test/bloc_test.dart';
import 'package:dartz/dartz.dart';
import 'package:ecom_app/core/constants/constants.dart';
import 'package:ecom_app/core/error/failure.dart';
import 'package:ecom_app/core/usecase/usecase.dart';
import 'package:ecom_app/features/product/domain/entities/product.dart';
import 'package:ecom_app/features/product/domain/usecases/create_product.dart';
import 'package:ecom_app/features/product/domain/usecases/delete_product.dart';
import 'package:ecom_app/features/product/domain/usecases/get_current_product.dart';
import 'package:ecom_app/features/product/domain/usecases/update_product.dart';
import 'package:ecom_app/features/product/presentation/bloc/product_bloc.dart';

import 'package:flutter_test/flutter_test.dart';
import 'package:mockito/mockito.dart';

import '../../../../helpers/test_helper.mocks.dart';

void main() {
  late MockGetCurrentProductUsecase mockGetCurrentProductUsecase;
  late MockGetAllProductsUsecase mockGetAllProductsUsecase;
  late MockCreateProductUsecase mockCreateProductUsecase;
  late MockUpdateProductUsecase mockUpdateProductUsecase;
  late MockDeleteProductUsecase mockDeleteProductUsecase;
  late ProductBloc productBloc;

  setUp(() {
    mockGetCurrentProductUsecase = MockGetCurrentProductUsecase();
    mockGetAllProductsUsecase = MockGetAllProductsUsecase();
    mockCreateProductUsecase = MockCreateProductUsecase();
    mockUpdateProductUsecase = MockUpdateProductUsecase();
    mockDeleteProductUsecase = MockDeleteProductUsecase();
    productBloc = ProductBloc(
        mockGetAllProductsUsecase,
        mockGetCurrentProductUsecase,
        mockCreateProductUsecase,
        mockUpdateProductUsecase,
        mockDeleteProductUsecase);
  });

  const testProductEntitiy = Product(
      id: '1',
      name: 'Product 1',
      description: 'Product 1 description',
      imageUrl: 'product1.jpg',
      price: 100);

  const testId = '1';

  test('initial state should be InitialState', () {
    expect(productBloc.state, ProductInitialState());
  });

  group('GetSingleProduct Event', () {
    blocTest<ProductBloc, ProductState>(
        'emits [LoadingState, LoadSingleProductState] when GetSingleProductEvent is added.',
        build: () {
          return productBloc;
        },
        setUp: () {
          when(mockGetCurrentProductUsecase(const GetParams(id: testId)))
              .thenAnswer((_) async => const Right(testProductEntitiy));
          return productBloc;
        },
        act: (bloc) => bloc.add(GetSingleProductEvent(id: testId)),
        expect: () => [
              ProductLoading(),
              LoadSingleProductState(product: testProductEntitiy)
            ]);
    blocTest<ProductBloc, ProductState>(
        'emits [LoadingState, ErrorState] when GetSingleProductEvent is unsuccessful.',
        build: () {
          return productBloc;
        },
        setUp: () {
          when(mockGetCurrentProductUsecase(const GetParams(id: testId)))
              .thenAnswer((_) async =>
                  const Left(ServerFailure(ErrorMessages.serverError)));
          return productBloc;
        },
        act: (bloc) => bloc.add(GetSingleProductEvent(id: testId)),
        expect: () => [
              ProductLoading(),
              ProductErrorState(message: ErrorMessages.serverError)
            ]);
  });

  group('LoadAllProduct Event', () {
    blocTest(
        'emits [LoadingState, LoadAllProductState] when LoadAllProduct Event is added',
        build: () {
          when(mockGetAllProductsUsecase(NoParams()))
              .thenAnswer((_) async => const Right([testProductEntitiy]));
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadAllProductEvent()),
        expect: () => [
              ProductLoading(),
              LoadAllProductState(products: const [testProductEntitiy])
            ]);
    blocTest(
        'emits [LoadingState, ErrorState] when LoadAllProductEvent is unsuccessful.',
        build: () {
          when(mockGetAllProductsUsecase(NoParams())).thenAnswer((_) async =>
              const Left(ServerFailure(ErrorMessages.serverError)));
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadAllProductEvent()),
        expect: () => [
              ProductLoading(),
              ProductErrorState(message: ErrorMessages.serverError)
            ]);
    blocTest(
        'emits [LoadingState, ErrorState] when LoadAllProductEvent is unsuccessful.',
        build: () {
          when(mockGetAllProductsUsecase(NoParams())).thenAnswer(
              (_) async => const Left(CacheFailure(ErrorMessages.cacheError)));
          return productBloc;
        },
        act: (bloc) => bloc.add(LoadAllProductEvent()),
        expect: () => [
              ProductLoading(),
              ProductErrorState(message: ErrorMessages.cacheError)
            ]);
  });

  group('CreateProductEvent', () {
    blocTest(
        'emits [LoadingState, LoadSingleProductState] when CreateProductEvent is added',
        build: () {
          when(mockCreateProductUsecase(
                  CreateParams(product: testProductEntitiy)))
              .thenAnswer((_) async => const Right(testProductEntitiy));
          return productBloc;
        },
        act: (bloc) =>
            bloc.add(CreateProductEvent(product: testProductEntitiy)),
        expect: () => [
              ProductLoading(),
              ProductCreatedState(product: testProductEntitiy)
            ]);

    blocTest(
        'emits [LoadingState, ErrorState] when CreateProductEvent is unsuccessful',
        build: () {
          when(mockCreateProductUsecase(
                  CreateParams(product: testProductEntitiy)))
              .thenAnswer((_) async =>
                  const Left(ServerFailure(ErrorMessages.serverError)));
          return productBloc;
        },
        act: (bloc) =>
            bloc.add(CreateProductEvent(product: testProductEntitiy)),
        expect: () => [
              ProductLoading(),
              ProductCreatedErrorState(message: ErrorMessages.serverError)
            ]);
  });

  group('UpdateProductEvent', () {
    blocTest(
        'emits [LoadingState, LoadSingleProductState] when UpdateProductEvent is added',
        build: () {
          when(mockUpdateProductUsecase(
                  UpdateParams(product: testProductEntitiy)))
              .thenAnswer((_) async => const Right(testProductEntitiy));
          return productBloc;
        },
        act: (bloc) =>
            bloc.add(UpdateProductEvent(product: testProductEntitiy)),
        expect: () => [
              ProductLoading(),
              ProductUpdatedState(product: testProductEntitiy)
            ]);

    blocTest(
        'emits [LoadingState, ErrorState] when UpdateProductEvent is unsuccessful',
        build: () {
          when(mockUpdateProductUsecase(
                  UpdateParams(product: testProductEntitiy)))
              .thenAnswer((_) async =>
                  const Left(ServerFailure(ErrorMessages.serverError)));
          return productBloc;
        },
        act: (bloc) =>
            bloc.add(UpdateProductEvent(product: testProductEntitiy)),
        expect: () => [
              ProductLoading(),
              ProductUpdatedErrorState(message: ErrorMessages.serverError)
            ]);
  });

  group('DeleteProductEvent', () {
    blocTest(
        'emits [LoadingState, ProductDeletedState] when DeleteProductEvent is added',
        build: () {
          when(mockDeleteProductUsecase(DeleteParams(id: testId)))
              .thenAnswer((_) async => const Right(null));
          return productBloc;
        },
        act: (bloc) => bloc.add(DeleteProductEvent(id: testId)),
        expect: () => [ProductLoading(), ProductDeletedState()]);

    blocTest(
        'emits [LoadingState, ErrorState] when DeleteProductEvent is unsuccessful',
        build: () {
          when(mockDeleteProductUsecase(DeleteParams(id: testId))).thenAnswer(
              (_) async =>
                  const Left(ServerFailure(ErrorMessages.serverError)));
          return productBloc;
        },
        act: (bloc) => bloc.add(DeleteProductEvent(id: testId)),
        expect: () => [
              ProductLoading(),
              ProductErrorState(message: ErrorMessages.serverError)
            ]);
  });
}
